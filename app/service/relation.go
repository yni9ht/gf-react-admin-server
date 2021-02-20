package service

type relationService struct{}

var (
	// Relation 供外部调用
	Relation = &relationService{}
)
