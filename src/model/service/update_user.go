package service

import (
	rest_err "github.com/Matchiga/meu-primeiro-crud-go/src"
	"github.com/Matchiga/meu-primeiro-crud-go/src/model"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
