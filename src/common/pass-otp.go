package passwords

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/harpy-py/go-Web-Api/config"
)

func generateOTP(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	otp := ""
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", num.Int64())
	}

	return otp, nil
}

func main(conf config.Config) {
	otpLength := conf.Otp.Digits
	otp, err := generateOTP(otpLength)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}

	fmt.Println("Generated OTP:", otp)
}