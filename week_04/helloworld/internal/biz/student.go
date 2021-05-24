package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	Id   int64
	Name string
	Age  int
}

type StudentRepo interface {
	ListStudent(ctx context.Context) ([]*Student, error)
}

type StudentUsecase struct {
	repo StudentRepo
}

func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo}
}

func (this *StudentUsecase) ListStudent(ctx context.Context) (ps []*Student, err error) {
	ps, err = this.repo.ListStudent(ctx)
	if err != nil {
		return
	}
	return
}
