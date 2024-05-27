package auth

import "errors"

var (
	ErrorMissingSessionCookie = errors.New("session cookie does not exist")
	ErrorMissingTfaCookie     = errors.New("tfa cookie does not exist")
	ErrorCredentialsMismatch  = errors.New("client is not equal with session")
	ErrorMissingUser          = errors.New("user doesn't exist")
	ErrorMismatchPassword     = errors.New("passwords aren't equal")
	ErrorUserAlreadyExists    = errors.New("user already exists")
	ErrorInvalidUser          = errors.New("invalid user")
	ErrorInvalidOtp           = errors.New("invalid otp")
	ErrorInvalidCredentials   = errors.New("invalid credentials")
)
