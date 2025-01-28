package services

import (
	dto "github.com/harpy-py/go-Web-Api/api/DTO"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
)

type TokenService struct {
	logger logging.Logger
	conf *config.Config
}

type tokenDto struct{
	UserId int
	FirstName string
	LastName string
	Username string
	PhoneNumber string
	Email string
	Roles []string
}

func NewTokenService(conf *config.Config) *TokenService {
	logger := logging.NewLogger(conf)
	return &TokenService{
		conf: conf,
		logger: logger,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error){

}