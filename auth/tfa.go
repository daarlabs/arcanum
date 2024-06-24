package auth

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"time"
	
	"github.com/dchest/uniuri"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	
	"github.com/daarlabs/arcanum/quirk"
	
	"github.com/daarlabs/arcanum/cache"
	"github.com/daarlabs/arcanum/cookie"
)

type TfaManager interface {
	GetPendingUserId() (int, error)
	GetPendingVerification() (bool, error)
	GetActive() (bool, error)
	Enable(id ...int) error
	Disable(id ...int) error
	Verify(otp string) (string, error)
	VerifyCodes(email, codes string) (bool, error)
	CreateQrImageBase64(id ...int) (string, error)
	
	MustGetPendingUserId() int
	MustGetPendingVerification() bool
	MustGetActive() bool
	MustEnable(id ...int)
	MustDisable(id ...int)
	MustVerify(otp string) string
	MustVerifyCodes(email, codes string) bool
	MustCreateQrImageBase64(id ...int) string
}

type tfaManager struct {
	manager *manager
	db      *quirk.DB
	cookie  cookie.Cookie
	cache   cache.Client
}

const (
	TfaCookieKey = "X-Tfa"
	TfaImageSize = 200
)

const (
	tfaSecretCodesLength = 160
)

func createTfaManager(
	manager *manager,
) TfaManager {
	return &tfaManager{
		manager: manager,
		db:      manager.db,
		cookie:  manager.cookie,
		cache:   manager.cache,
	}
}

func (m tfaManager) GetPendingUserId() (int, error) {
	var u User
	token := m.cookie.Get(TfaCookieKey)
	err := m.cache.Get(createTfaCacheKey(token), &u)
	return u.Id, err
}

func (m tfaManager) MustGetPendingUserId() int {
	userId, err := m.GetPendingUserId()
	if err != nil {
		panic(err)
	}
	return userId
}

func (m tfaManager) GetPendingVerification() (bool, error) {
	token := m.cookie.Get(TfaCookieKey)
	n := len(token)
	if n == 0 {
		return false, nil
	}
	return n > 0, nil
}

func (m tfaManager) MustGetPendingVerification() bool {
	pending, err := m.GetPendingVerification()
	if err != nil {
		panic(err)
	}
	return pending
}

func (m tfaManager) GetActive() (bool, error) {
	user, err := m.manager.User().Get()
	if err != nil {
		return false, err
	}
	return user.Tfa && len(user.TfaUrl.V) > 0 && len(user.TfaCodes.V) > 0 && len(user.TfaSecret.V) > 0, nil
}

func (m tfaManager) MustGetActive() bool {
	active, err := m.GetActive()
	if err != nil {
		panic(err)
	}
	return active
}

func (m tfaManager) Verify(otp string) (string, error) {
	token := m.cookie.Get(TfaCookieKey)
	if len(token) == 0 {
		return "", ErrorMissingTfaCookie
	}
	var u User
	if err := m.cache.Get(createTfaCacheKey(token), &u); err != nil {
		return "", err
	}
	if u.Id == 0 {
		return "", ErrorInvalidUser
	}
	err := quirk.New(m.db).
		Q(fmt.Sprintf(`SELECT id, email, roles, tfa_secret FROM %s`, usersTable)).
		Q("WHERE id = @id", quirk.Map{"id": u.Id}).
		Exec(&u)
	if err != nil {
		return "", err
	}
	if valid := totp.Validate(otp, u.TfaSecret.V); !valid {
		return "", ErrorInvalidOtp
	}
	if err := m.cache.Set(token, "", time.Millisecond); err != nil {
		return "", err
	}
	m.cookie.Set(TfaCookieKey, "", time.Millisecond)
	return m.manager.Session().New(u)
}

func (m tfaManager) MustVerify(otp string) string {
	token, err := m.Verify(otp)
	if err != nil {
		panic(err)
	}
	return token
}

func (m tfaManager) VerifyCodes(email, codes string) (bool, error) {
	u, err := m.manager.CustomUser(0, email).Get()
	if err != nil {
		return false, err
	}
	return u.TfaCodes.V == codes, nil
}

func (m tfaManager) MustVerifyCodes(email, codes string) bool {
	verified, err := m.VerifyCodes(email, codes)
	if err != nil {
		panic(err)
	}
	return verified
}

func (m tfaManager) Enable(id ...int) error {
	userId, err := m.getUserId(id...)
	if err != nil {
		return err
	}
	u, err := m.manager.User().Get()
	if err != nil {
		return err
	}
	key, err := totp.Generate(
		totp.GenerateOpts{
			Issuer:      m.getHost(),
			AccountName: u.Email,
		},
	)
	if err != nil {
		return err
	}
	codes := uniuri.NewLen(tfaSecretCodesLength)
	return quirk.New(m.db).
		Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(
			"SET tfa = @tfa, tfa_codes = @tfa-codes, tfa_secret = @tfa-secret, tfa_url = @tfa-url", quirk.Map{
				"tfa":        true,
				"tfa-codes":  codes,
				"tfa-secret": key.Secret(),
				"tfa-url":    key.String(),
			},
		).
		Q("WHERE id = @id", quirk.Map{"id": userId}).
		Exec()
}

func (m tfaManager) MustEnable(id ...int) {
	err := m.Enable(id...)
	if err != nil {
		panic(err)
	}
}

func (m tfaManager) Disable(id ...int) error {
	userId, err := m.getUserId(id...)
	if err != nil {
		return err
	}
	err = quirk.New(m.db).
		Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q("SET tfa = false, tfa_codes = NULL, tfa_secret = NULL, tfa_url = NULL").
		Q("WHERE id = @id", quirk.Map{"id": userId}).
		Exec()
	if err != nil {
		return err
	}
	m.cookie.Destroy(TfaCookieKey)
	return nil
}

func (m tfaManager) MustDisable(id ...int) {
	err := m.Disable(id...)
	if err != nil {
		panic(err)
	}
}

func (m tfaManager) CreateQrImageBase64(id ...int) (string, error) {
	userId, err := m.getUserId(id...)
	if err != nil {
		return "", err
	}
	var u User
	err = quirk.New(m.db).
		Q(fmt.Sprintf(`SELECT tfa_url FROM %s`, usersTable)).
		Q("WHERE id = @id", quirk.Map{"id": userId}).
		Exec(&u)
	if err != nil {
		return "", err
	}
	key, err := otp.NewKeyFromURL(u.TfaUrl.V)
	if err != nil {
		return "", err
	}
	img, err := key.Image(TfaImageSize, TfaImageSize)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	if err = png.Encode(&buffer, img); err != nil {
		return "", err
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func (m tfaManager) MustCreateQrImageBase64(id ...int) string {
	qrImageBase64, err := m.CreateQrImageBase64(id...)
	if err != nil {
		panic(err)
	}
	return qrImageBase64
}

func (m tfaManager) getHost() string {
	protocol := "http"
	if m.manager.req.TLS != nil {
		protocol = "https"
	}
	return protocol + "://" + m.manager.req.Host
}

func (m tfaManager) getUserId(id ...int) (int, error) {
	var userId int
	idn := len(id)
	if idn == 0 {
		user, err := m.manager.Session().Get()
		if err != nil {
			return userId, err
		}
		userId = user.Id
	}
	if idn > 0 {
		userId = id[0]
	}
	return userId, nil
}
