package config

import (
	"gorm.io/gorm"
)

var(
	db *gorm.DB
	logger *Logger
)

func Init() error{
	//return errors.New("Um siri fazendo barra!!!!")
	return nil
}

func GetLoger(p string) *Logger{
	logger = NewLogger(p)
	return logger
}
