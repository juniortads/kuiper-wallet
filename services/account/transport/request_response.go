package transport

import "github.com/juniortads/kuiper-wallet/services/account"

type (
	CreateAccountRequest struct {
		Account account.Account
	}
	CreateAccountResponse struct {
		Err error
		Id string
	}
)
