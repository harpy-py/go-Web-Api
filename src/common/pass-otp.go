package common

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/harpy-py/go-Web-Api/config"
)

func GenerateOTP(conf config.Config) string {
	otpLength := conf.Otp.Digits
	otp := ""
	for i := 0; i < int(otpLength); i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(otpLength))
		if err != nil {
			fmt.Println("Error generating OTP:", err)
		}
		otp += fmt.Sprintf("%d", num.Int64())
	}

	fmt.Println("Generated OTP:", otp)
	return otp
}