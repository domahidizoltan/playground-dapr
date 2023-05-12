package dto

import "errors"

const (
	LockBalanceCommandType         CommandType = "lockBalance"
	DebitSourceCommandType         CommandType = "debitSource"
	UpdateSourceBalanceCommandType CommandType = "updateSourceBalance"
	CreditDestCommandType          CommandType = "creditDest"
	UpdateDestBalanceCommandType   CommandType = "updateDestBalance"
)

var BusinessError error = errors.New("error")

type (
	CommandType string
	// UpdateBalance struct {
	// 	Account string  `json:"account"`
	// 	Amount  float32 `json:"amount"`
	// }

	TransferCommand struct {
		Command CommandType `json:"command"`
		Tnx     string      `json:"tnx"`
		SrcAcc  string      `json:"srcAcc"`
		DstAcc  string      `json:"dstAcc"`
		Amount  float64     `json:"amount"`
	}
)
