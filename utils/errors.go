package utils

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewError(err error) Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = err.Error()
	return e
}

func PasswordsNotMatch() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "passwords don't match"
	return e
}

func AccessForbidden() Error {
	e := Error{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}
