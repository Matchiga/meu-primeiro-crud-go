package service

import (
	"fmt"

	rest_err "github.com/Matchiga/meu-primeiro-crud-go/src"
	"github.com/Matchiga/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/Matchiga/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
