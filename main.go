package main

import (
	"github.com/Cannedsans/GoCrud/config"
	"github.com/Cannedsans/GoCrud/router"
)

var(
	logger config.Logger
)
func main() {
	logger = *config.GetLoger("main")
	// inicialização das configurações
	err := config.Init()
	if err != nil {
		logger.ErroF("Erro ao inicializar as configurações.: %v", err)
		return
	}

	router.Initialize()
}
