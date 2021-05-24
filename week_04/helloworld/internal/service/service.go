package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewStudentServiceService)

//todo 修改相关代码
type StudentService struct {
	pb.UnimplementedStudentServiceServer

	log *log.Helper

	student *biz.StudentUsecase
}
