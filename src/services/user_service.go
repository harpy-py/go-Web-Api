package services

import (
	dto "github.com/harpy-py/go-Web-Api/api/DTO"
	"github.com/harpy-py/go-Web-Api/common"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/constants"
	"github.com/harpy-py/go-Web-Api/data/db"
	"github.com/harpy-py/go-Web-Api/data/models"
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

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUserName(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByMobileNuber(mobilenumber string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).Select("count(*) > 0").Where("mobile_number = ?", mobilenumber).Find(&exists).Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error){
	if err = s.database.Model(&models.Role{}).Select("id").Where("name = ?", constants.DefaultRoleName).First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}