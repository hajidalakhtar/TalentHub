package config

import (
	"TalentHub/exception"
	"github.com/joho/godotenv"
	"os"
)

type Config interface {
	Get(key string) string
	JwtKey() []byte
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}
func (config *configImpl) JwtKey() []byte {
	key := os.Getenv("JWT_KEY")
	return []byte(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
