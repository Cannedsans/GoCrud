package handler

import (
	"github.com/Cannedsans/GoCrud/config"
	"gorm.io/gorm"
)

var(
	logger *config.Logger
	db*gorm.DB
)

func InitializeHandler(){
	logger = config.GetLoger("handler")
	db = config.GetSQLite()
}
