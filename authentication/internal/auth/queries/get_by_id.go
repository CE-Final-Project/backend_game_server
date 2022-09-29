package queries

import (
	"context"
	accountService "github.com/ce-final-project/backend_game_server/account/proto"
	"github.com/ce-final-project/backend_game_server/authentication/config"
	"github.com/ce-final-project/backend_game_server/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

type GetAccountByIdHandler interface {
	Handle(ctx context.Context, query *GetAccountByIdQuery) (*accountService.Account, error)
}

type getAccountByIdHandler struct {
	log      logger.Logger
	cfg      *config.Config
	asClient accountService.AccountServiceClient
}

func NewGetAccountByIdHandler(log logger.Logger, cfg *config.Config, asClient accountService.AccountServiceClient) *getAccountByIdHandler {
	return &getAccountByIdHandler{
		log:      log,
		cfg:      cfg,
		asClient: asClient,
	}
}

func (l *getAccountByIdHandler) Handle(ctx context.Context, query *GetAccountByIdQuery) (*accountService.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAccountByIdHandler.Handle")
	defer span.Finish()

	res, err := l.asClient.GetAccountById(ctx, &accountService.GetAccountByIdReq{AccountID: query.AccountID})
	if err != nil {
		return nil, err
	}

	return res.GetAccount(), nil
}