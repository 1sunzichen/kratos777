package service

import (
	"helloworld/api/opea"
	"helloworld/internal/biz"
)

type AppService struct {
	opea.UnimplementedAppServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewAppService(uc *biz.GreeterUsecase) *AppService {
	return &AppService{uc: uc}
}
