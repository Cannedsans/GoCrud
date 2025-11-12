package config

import (
	"os"

	"github.com/Cannedsans/GoCrud/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLoger("sqlite")
	dbpath := "./db/main.db"
	// check if the database file exist
	_, err := os.Stat(dbpath)
	if os.IsNotExist(err){
		logger.Info("Criando arquivo de banco de dados...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil{
			return nil, err
		}

		file, err := os.Create(dbpath)

		if err != nil{
			return nil, err
		}

		file.Close()
	}

	//create database and connect
	db, err :=gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	if err != nil{
		logger.ErroF("Erro ao inicializar o banco de dados.: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil{
		logger.ErroF("Erro ao migrar o banco de dados.: %v", err)
		return nil, err
	}

	return db, nil
}
