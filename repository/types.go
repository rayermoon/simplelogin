// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type GetUserByEmail string
type GetUserByPhoneNumberInput string
type GetUserByEmailInput string
type GetUserByIdInput string
type CheckLogedIn string
type CreateUserInput struct {
	Email       string
	FullName    string
	Gender      string
	Occupation  string
	Password    string
	PhoneNumber string
}
type UpdateUserInput struct {
	UserId      string
	FullName    string
	PhoneNumber string
}
type LoginInput struct {
	UserEmail   string
	PhoneNumber string
	Password    string
}

type UserOutput struct {
	Id          string
	Email       string
	FullName    string
	Gender      string
	Occupation  string
	Password    string
	PhoneNumber string
	LoginCount  int
}

type LoginOutput string
