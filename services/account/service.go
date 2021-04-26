package account

import "context"

type Service interface {
	CreateAccount(ctx context.Context, account Account) (string, error)
}