package logic

import (
	"context"

	"go_zero_test/inner/svc"
	"go_zero_test/inner/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Go_zero_testLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGo_zero_testLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Go_zero_testLogic {
	return &Go_zero_testLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Go_zero_testLogic) Go_zero_test(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
