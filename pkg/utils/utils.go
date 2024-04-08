package utils

import (
	"math/rand"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomString(length int) string {
	return stringWithCharset(length, charset)
}

func GenerateUUID4() string {
	return uuid.NewString()
}

func ValidateStruct(model interface{}) error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}
