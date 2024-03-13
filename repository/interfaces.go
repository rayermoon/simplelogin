// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetUserById(ctx context.Context, input GetUserByIdInput) (output UserOutput, err error)
	GetUserByEmail(ctx context.Context, input GetUserByEmailInput) (output UserOutput, err error)
	GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (output UserOutput, err error)
	CreateUser(ctx context.Context, input CreateUserInput) (err error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (err error)
	Login(ctx context.Context, input LoginInput) (output LoginOutput, err error)
}
