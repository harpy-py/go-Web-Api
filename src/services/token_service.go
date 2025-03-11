package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	dto "github.com/harpy-py/go-Web-Api/api/DTO"
	"github.com/harpy-py/go-Web-Api/config"
	"github.com/harpy-py/go-Web-Api/pkg/logging"
	"github.com/harpy-py/go-Web-Api/pkg/service_errors"
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
	tokenDetail := &dto.TokenDetail{}
	tokenDetail.AccessTokenExpireTime = time.Now().Add(s.conf.JWT.AccessTokenExpireDuration * time.Minute).Unix()
	tokenDetail.RefreshTokenExpireTime = time.Now().Add(s.conf.JWT.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaim := jwt.MapClaims{}

	accessTokenClaim["uesr_id"] = token.UserId
	accessTokenClaim["first_name"] = token.FirstName
	accessTokenClaim["last_name"] = token.LastName
	accessTokenClaim["username"] = token.Username
	accessTokenClaim["email"] = token.Email
	accessTokenClaim["mobile_number"] = token.PhoneNumber
	accessTokenClaim["roles"] = token.Roles
	accessTokenClaim["exp"] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)

	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(s.conf.JWT.Secret))

	if err != nil {
		return nil, err
	}

	refreshTokenClaim := jwt.MapClaims{}

	refreshTokenClaim["uesr_id"] = token.UserId
	refreshTokenClaim["exp"] = tokenDetail.AccessTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)

	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(s.conf.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}
	
	return tokenDetail, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error){
	aToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.conf.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return aToken, nil
}

func (s *TokenService) GetClaim(token string) (claimMap map[string]interface{}, err error){
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil{
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid{
		for k, v := range claims{
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}