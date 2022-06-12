package service

type (
	Auth interface {
		SignUp() error
		SignIn() error
	}
	AuthRepository interface {
	}
)
