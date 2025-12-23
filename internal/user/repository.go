package user

import (
	"context"
	"time"
)

//O que é: contrato da camada de persistência.
//Por que importa: desacopla handlers da tecnologia do banco; permite //trocar por fury_nosql-go-sdk depois sem mexer no HTTP.

type UserRepository interface {
	Create(ctx context.Context, user User) (string, error)
	FindAll(ctx context.Context) ([]User, error)
	FindByDpp(ctx context.Context, dpp time.Time) ([]User, error)
	FindByCity(ctx context.Context, city string) ([]User, error)
	FindByAge(ctx context.Context, age int) ([]User, error)
}
