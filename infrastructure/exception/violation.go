package exception

type FieldViolation struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}
