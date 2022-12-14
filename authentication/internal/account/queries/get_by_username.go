package queries

import (
	"context"
	"github.com/ce-final-project/backend_game_server/authentication/config"
	"github.com/ce-final-project/backend_game_server/authentication/internal/account/repository"
	"github.com/ce-final-project/backend_game_server/authentication/internal/models"
	"github.com/ce-final-project/backend_game_server/pkg/logger"
	"github.com/opentracing/opentracing-go"
	"strconv"
)

type GetAccountByUsernameHandler interface {
	Handle(ctx context.Context, query *GetAccountByUsernameQuery) (*models.Account, error)
}

type getAccountByUsernameHandler struct {
	log         logger.Logger
	cfg         *config.Config
	accountRepo repository.AccountRepository
	cacheRepo   repository.CacheRepository
}

func NewGetAccountByUsernameQueryHandler(log logger.Logger, cfg *config.Config, accountRepo repository.AccountRepository, cacheRepo repository.CacheRepository) GetAccountByUsernameHandler {
	return &getAccountByUsernameHandler{log: log, cfg: cfg, accountRepo: accountRepo, cacheRepo: cacheRepo}
}

func (q *getAccountByUsernameHandler) Handle(ctx context.Context, query *GetAccountByUsernameQuery) (*models.Account, error) {
	var span opentracing.Span
	span, ctx = opentracing.StartSpanFromContext(ctx, "getAccountByEmailHandler.Handle")
	defer span.Finish()

	if account, err := q.cacheRepo.GetAccountByKeyReference(ctx, query.Username); err == nil && account != nil {
		return account, nil
	}

	account, err := q.accountRepo.GetAccountByUsername(ctx, query.Username)
	if err != nil {
		return nil, err
	}

	q.cacheRepo.PutKeyReference(ctx, account.Username, strconv.FormatUint(account.ID, 10))
	return account, nil
}
