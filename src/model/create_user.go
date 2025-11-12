package model

import (
	"fmt"

	rest_err "github.com/Matchiga/meu-primeiro-crud-go/src"
	"github.com/Matchiga/meu-primeiro-crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
