package dto

type (
	UpdateBalance struct {
		Account string  `json:"account"`
		Amount  float32 `json:"amount"`
	}
)
