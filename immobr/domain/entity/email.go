package entity

type Email struct {
	Value string `json:"email"`
}

func (e *Email) Validate() error {
	return nil
}
