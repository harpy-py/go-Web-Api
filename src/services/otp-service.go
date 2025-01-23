package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/constants"
	"github.com/harpy-py/go-Web-Api/data/cache"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
	"github.com/harpy-py/go-Web-Api/pkg/service_errors"
)

type OtpService struct {
	logger logging.Logger
	conf config.Config
	redis *redis.Client
}

type OtpDto struct{
	Value string
	Used bool
}

func NewOtpService(conf *config.Config) *OtpService{
	logger := logging.NewLogger(conf)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, conf: *conf, redis: redis}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error{
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used: false,
	}

	res, err := cache.Get[OtpDto](s.redis, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	}else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}
	err = cache.Set(s.redis, key, val, s.conf.Otp.ExpireTime * time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error{
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	res, err := cache.Get[OtpDto](s.redis, key)
	if err != nil {
		return nil
	}else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}else if err == nil && !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpInvalid}
	}else if err == nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.conf.Otp.ExpireTime * time.Second)
	}
	return nil
}