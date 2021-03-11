package service

import (
	"context"
)

type tinaService struct {
}

func NewTinaService() *tinaService {
	return &tinaService{}
}

func (s *tinaService) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: req.GetName() + " 你好,《山海情》真精彩！"}, nil
}

func (s *tinaService) LearnForeignLanguage(ctx context.Context, req *LanguageRequest) (*LanguageReply, error) {
	return &LanguageReply{Score: 100}, nil
}
