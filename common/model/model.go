package model

type TransferStatus string

const (
	TransferStatusPending                 = "PENDING"
	TransferStatusSourceDebited           = "SOURCE_DEBITED"
	TransferStatusSourceDebitFailed       = "SOURCE_DEBIT_FAILED"
	TransferStatusDestinationCredited     = "DESTINATION_CREDITED"
	TransferStatusDestinationCreditFailed = "DESTINATION_CREDIT_FAILED"
	TransferStatusCompleted               = "COMPLETED"
)

const (
	BankAccount = "BANK"
	AccPrefix   = "ACC"
	AccPattern  = AccPrefix + "[0-9]{3}"
)

type Transfer struct {
	TNX    string
	SrcAcc string
	DstAcc string
	Amount float32
}
