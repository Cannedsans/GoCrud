package config

import (
	"fmt"

	"gorm.io/gorm"
)

var(
	db *gorm.DB
	logger *Logger
)

func Init() error{
	var err error

	// inicialize db

	db, err = InitializeSQLite()

	if err != nil{
		return fmt.Errorf("Erro ao inicializar o banco de dados", err)
	}
	//return errors.New("Um siri fazendo barra!!!!")
	return nil
}

func GetSQLite() *gorm.DB{
	return db
}

func GetLoger(p string) *Logger{
	logger = NewLogger(p)
	return logger
}
