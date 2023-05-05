package dto

type (
	UpdateBalance struct {
		Account string  `json:"account"`
		Amount  float32 `json:"amount"`
	}

	Transfer struct {
		Tnx    string  `json:"tnx"`
		SrcAcc string  `json:"srcAcc"`
		DstAcc string  `json:"dstAcc"`
		Amount float32 `json:"amount"`
	}
)
