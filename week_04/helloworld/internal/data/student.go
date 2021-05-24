package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper("data/student", logger),
	}
}

func (s studentRepo) ListStudent(ctx context.Context) ([]*biz.Student, error) {
	ps, err := s.data.db.Debug().Student.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Student, 0)
	for i := range ps {
		rv = append(rv, &biz.Student{
			Id:   ps[i].ID,
			Name: ps[i].Name,
			Age:  ps[i].Age,
		})
	}
	return rv, nil
}
