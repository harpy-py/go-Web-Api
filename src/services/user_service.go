package services

import (
	dto "github.com/harpy-py/go-Web-Api/api/DTO"
	"github.com/harpy-py/go-Web-Api/common"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/data/db"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger logging.Logger
	conf *config.Config
	OtpService *OtpService
	database *gorm.DB
}

func NewUserService(conf *config.Config) *UserService{
	database := db.GetDb()
	logger := logging.NewLogger(conf)
	return &UserService{
		conf: conf,
		database: database,
		logger: logger,
		OtpService: NewOtpService(conf),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOTP(*s.conf)
	err := s.OtpService.SetOtp(req.MobileNumber, otp)
	if err != nil{
		return err
	}
	return nil
}