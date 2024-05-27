package auth

import (
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"time"
	
	"github.com/dchest/uniuri"
	"github.com/matthewhartstonge/argon2"
	
	"github.com/daarlabs/arcanum/cache"
	
	"github.com/daarlabs/arcanum/quirk"
)

type UserManager interface {
	Get(id ...int) (User, error)
	Create(r User) (int, error)
	Update(r User, columns ...string) error
	ResetPassword(token ...string) (string, error)
	UpdatePassword(actualPassword, newPassword string) error
	ForceUpdatePassword(newPassword string) error
	Enable(id ...int) error
	Disable(id ...int) error
	
	MustGet(id ...int) User
	MustCreate(r User) int
	MustUpdate(r User, columns ...string)
	MustResetPassword(token ...string) string
	MustUpdatePassword(actualPassword, newPassword string)
	MustForceUpdatePassword(newPassword string)
	MustEnable(id ...int)
	MustDisable(id ...int)
}

type User struct {
	Id           int              `json:"id"`
	Active       bool             `json:"active"`
	Roles        []string         `json:"roles"`
	Email        string           `json:"email"`
	Password     string           `json:"password"`
	Tfa          bool             `json:"tfa"`
	TfaSecret    sql.Null[string] `json:"tfaSecret"`
	TfaCodes     sql.Null[string] `json:"tfaCodes"`
	TfaUrl       sql.Null[string] `json:"tfaUrl"`
	LastActivity time.Time        `json:"lastActivity"`
	CreatedAt    time.Time        `json:"createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt"`
}

type userManager struct {
	db         *quirk.DB
	cache      cache.Client
	id         int
	email      string
	driverName string
	data       quirk.Map
}

const (
	UserActive       = "active"
	UserRoles        = "roles"
	UserEmail        = "email"
	UserPassword     = "password"
	UserTfa          = "tfa"
	UserTfaSecret    = "tfa_secret"
	UserTfaCodes     = "tfa_codes"
	UserTfaUrl       = "tfa_url"
	UserLastActivity = "last_activity"
)

const (
	usersTable  = "users"
	paramPrefix = "@"
)

const (
	operationInsert = "insert"
	operationUpdate = "update"
)

var (
	argon = argon2.DefaultConfig()
)

func CreateUserManager(db *quirk.DB, cache cache.Client, id int, email string) UserManager {
	return &userManager{
		db:         db,
		cache:      cache,
		email:      email,
		id:         id,
		data:       make(map[string]any),
		driverName: db.DriverName(),
	}
}

func (u *userManager) Get(id ...int) (User, error) {
	if len(id) > 0 {
		u.id = id[0]
	}
	var r User
	if u.id == 0 && u.email == "" {
		return r, ErrorInvalidUser
	}
	err := quirk.New(u.db).Q(`SELECT *`).
		Q(fmt.Sprintf(`FROM %s`, usersTable)).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Q(`LIMIT 1`).
		Exec(&r)
	clear(u.data)
	return r, err
}

func (u *userManager) MustGet(id ...int) User {
	r, err := u.Get(id...)
	if err != nil {
		panic(err)
	}
	return r
}

func (u *userManager) Create(r User) (int, error) {
	if u.id != 0 {
		return u.id, ErrorUserAlreadyExists
	}
	if err := u.readData(operationInsert, r, []string{}); err != nil {
		return 0, err
	}
	columns, placeholders := u.insertValues()
	err := quirk.New(u.db).Q(fmt.Sprintf(`INSERT INTO %s`, usersTable)).
		Q(fmt.Sprintf(`(%s)`, columns)).
		Q(fmt.Sprintf(`VALUES (%s)`, placeholders), u.args()...).
		Q(`RETURNING id`).
		Exec(&u.id)
	u.email = r.Email
	clear(u.data)
	return u.id, err
}

func (u *userManager) MustCreate(r User) int {
	id, err := u.Create(r)
	if err != nil {
		panic(err)
	}
	return id
}

func (u *userManager) Update(r User, columns ...string) error {
	if u.id == 0 && u.email == "" {
		return ErrorInvalidUser
	}
	if err := u.readData(operationUpdate, r, columns); err != nil {
		return err
	}
	err := quirk.New(u.db).Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(fmt.Sprintf(`SET %s`, u.updateValues()), u.args()...).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Exec()
	clear(u.data)
	return err
}

func (u *userManager) MustUpdate(r User, columns ...string) {
	err := u.Update(r, columns...)
	if err != nil {
		panic(err)
	}
}

func (u *userManager) ResetPassword(token ...string) (string, error) {
	if len(token) > 0 {
		var r User
		err := u.cache.Get(u.createResetPasswordKey(token[0]), &r)
		u.email = r.Email
		return r.Email, err
	}
	t := uniuri.New()
	return t, u.cache.Set(
		u.createResetPasswordKey(t),
		User{Email: u.email},
		time.Hour,
	)
}

func (u *userManager) MustResetPassword(token ...string) string {
	t, err := u.ResetPassword(token...)
	if err != nil {
		panic(err)
	}
	return t
}

func (u *userManager) UpdatePassword(actualPassword, newPassword string) error {
	if u.id == 0 && u.email == "" {
		return ErrorMissingUser
	}
	user, err := u.Get()
	if err != nil {
		return err
	}
	if ok, err := argon2.VerifyEncoded([]byte(actualPassword), []byte(user.Password)); !ok || err != nil {
		return ErrorMismatchPassword
	}
	hash, err := u.hashPassword(newPassword)
	if err != nil {
		return err
	}
	err = quirk.New(u.db).Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(`SET password = @password`, quirk.Map{"password": hash}).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Exec()
	clear(u.data)
	return err
}

func (u *userManager) MustUpdatePassword(actualPassword, newPassword string) {
	err := u.UpdatePassword(actualPassword, newPassword)
	if err != nil {
		panic(err)
	}
}

func (u *userManager) ForceUpdatePassword(newPassword string) error {
	if u.id == 0 && u.email == "" {
		return ErrorMissingUser
	}
	hash, err := u.hashPassword(newPassword)
	if err != nil {
		return err
	}
	err = quirk.New(u.db).Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(`SET password = @password`, quirk.Map{"password": hash}).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Exec()
	clear(u.data)
	return err
}

func (u *userManager) MustForceUpdatePassword(newPassword string) {
	err := u.ForceUpdatePassword(newPassword)
	if err != nil {
		panic(err)
	}
}

func (u *userManager) Enable(id ...int) error {
	if len(id) > 0 {
		u.id = id[0]
	}
	if u.id == 0 && u.email == "" {
		return ErrorInvalidUser
	}
	err := quirk.New(u.db).Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(`SET active = true`).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Exec()
	clear(u.data)
	return err
}

func (u *userManager) MustEnable(id ...int) {
	err := u.Enable(id...)
	if err != nil {
		panic(err)
	}
}

func (u *userManager) Disable(id ...int) error {
	if len(id) > 0 {
		u.id = id[0]
	}
	if u.id == 0 && u.email == "" {
		return ErrorInvalidUser
	}
	err := quirk.New(u.db).Q(fmt.Sprintf(`UPDATE %s`, usersTable)).
		Q(`SET active = false`).
		If(u.id > 0, `WHERE id = @id`, quirk.Map{"id": u.id}).
		If(u.id == 0, `WHERE email = @email`, quirk.Map{"email": u.email}).
		Exec()
	clear(u.data)
	return err
}

func (u *userManager) MustDisable(id ...int) {
	err := u.Disable(id...)
	if err != nil {
		panic(err)
	}
}

func (u *userManager) readData(operation string, data User, columns []string) error {
	columnsExist := len(columns) > 0
	if operation == operationInsert && slices.Contains(columns, quirk.Id) {
		u.data[quirk.Id] = data.Id
	}
	if !columnsExist || slices.Contains(columns, UserActive) {
		u.data[UserActive] = data.Active
	}
	if !columnsExist || slices.Contains(columns, UserEmail) {
		u.data[UserEmail] = data.Email
	}
	if !columnsExist || slices.Contains(columns, UserPassword) {
		hash, err := u.hashPassword(data.Password)
		if err != nil {
			return err
		}
		u.data[UserPassword] = hash
	}
	if !columnsExist || slices.Contains(columns, UserRoles) {
		u.data[UserRoles] = data.Roles
	}
	if !columnsExist || slices.Contains(columns, UserTfa) {
		u.data[UserTfa] = data.Tfa
	}
	if !columnsExist || slices.Contains(columns, UserTfaSecret) {
		u.data[UserTfaSecret] = data.TfaSecret.V
	}
	if !columnsExist || slices.Contains(columns, UserTfaCodes) {
		u.data[UserTfaCodes] = data.TfaCodes.V
	}
	if !columnsExist || slices.Contains(columns, UserTfaUrl) {
		u.data[UserTfaUrl] = data.TfaUrl.V
	}
	return nil
}

func (u *userManager) hashPassword(password string) (string, error) {
	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u *userManager) args() []quirk.Map {
	if len(u.data) == 0 {
		return []quirk.Map{}
	}
	result := u.data
	vectors := make([]any, 0)
	for name, v := range u.data {
		if name == UserPassword {
			continue
		}
		vectors = append(vectors, v)
	}
	switch u.driverName {
	case quirk.Postgres:
		if len(vectors) > 0 {
			result[quirk.Vectors] = quirk.CreateTsVector(vectors...)
		}
	}
	return []quirk.Map{result}
}

func (u *userManager) insertValues() (string, string) {
	columns := []string{quirk.Id}
	placeholders := []string{quirk.Default}
	for name := range u.data {
		columns = append(columns, name)
		placeholders = append(placeholders, paramPrefix+name)
	}
	switch u.driverName {
	case quirk.Postgres:
		if len(u.data) > 0 {
			columns = append(columns, quirk.Vectors)
			placeholders = append(placeholders, fmt.Sprintf("to_tsvector(%s%s)", paramPrefix, quirk.Vectors))
		}
	}
	columns = append(columns, UserLastActivity)
	placeholders = append(placeholders, quirk.CurrentTimestamp)
	
	columns = append(columns, quirk.CreatedAt)
	placeholders = append(placeholders, quirk.CurrentTimestamp)
	
	columns = append(columns, quirk.UpdatedAt)
	placeholders = append(placeholders, quirk.CurrentTimestamp)
	return strings.Join(columns, ","), strings.Join(placeholders, ",")
}

func (u *userManager) updateValues() string {
	result := make([]string, 0)
	for column := range u.data {
		if column == quirk.Id {
			continue
		}
		result = append(result, fmt.Sprintf("%s = %s%s", column, paramPrefix, column))
	}
	result = append(result, fmt.Sprintf("%s = %s", UserLastActivity, quirk.CurrentTimestamp))
	result = append(result, fmt.Sprintf("%s = %s", quirk.UpdatedAt, quirk.CurrentTimestamp))
	switch u.driverName {
	case quirk.Postgres:
		vectors := make([]any, 0)
		for column, v := range u.data {
			if column == quirk.Id {
				continue
			}
			vectors = append(vectors, v)
		}
		if len(vectors) > 0 {
			result = append(result, fmt.Sprintf("%s = to_tsvector(%s%s)", quirk.Vectors, paramPrefix, quirk.Vectors))
		}
	}
	return strings.Join(result, ",")
}

func (u *userManager) createResetPasswordKey(token string) string {
	return "reset-password:" + token
}
