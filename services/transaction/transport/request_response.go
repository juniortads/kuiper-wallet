package transport

import "github.com/juniortads/kuiper-wallet/services/transaction"

type (
	CreateTransactionRequest struct {
		Transaction transaction.Transaction
	}
	CreateTransactionResponse struct {
		Err error
		Id string
	}
)