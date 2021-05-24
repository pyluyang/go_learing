package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

func NewStudentServiceService(student *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{
		log:     log.NewHelper("student", logger),
		student: student,
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return &pb.CreateStudentReply{}, nil
}
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	return &pb.UpdateStudentReply{}, nil
}
func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	return &pb.DeleteStudentReply{}, nil
}
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	return &pb.GetStudentReply{}, nil
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	ps, err := s.student.ListStudent(ctx)
	reply := &pb.ListStudentReply{}
	for i := range ps {
		reply.Students = append(reply.Students, &pb.Student{
			Id:   ps[i].Id,
			Name: ps[i].Name,
			Age:  int64(ps[i].Age),
		})
	}
	return reply, err
}
