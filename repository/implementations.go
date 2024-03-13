package repository

import (
	"context"
	"errors"

	"github.com/SawitProRecruitment/UserService/utils"
	util "github.com/SawitProRecruitment/UserService/utils"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

const mainQuery = `
	select
		id,
		email,
		full_name,
		gender,
		occupation,
		password,
		phone_number,
		login_count
	from sp_users

`

func (r *Repository) GetUserById(ctx context.Context, input GetUserByIdInput) (output UserOutput, err error) {
	query := mainQuery + `where id = $1`
	err = r.Db.QueryRowContext(ctx, query, input).Scan(
		&output.Id,
		&output.Email,
		&output.FullName,
		&output.Gender,
		&output.Occupation,
		&output.Password,
		&output.PhoneNumber,
		&output.LoginCount,
	)
	if err != nil {
		return
	}
	return
}
func (r *Repository) GetUserByEmail(ctx context.Context, input GetUserByEmailInput) (output UserOutput, err error) {
	query := mainQuery + `where email = $1`
	err = r.Db.QueryRowContext(ctx, query, input).Scan(
		&output.Id,
		&output.Email,
		&output.FullName,
		&output.Gender,
		&output.Occupation,
		&output.Password,
		&output.PhoneNumber,
		&output.LoginCount,
	)
	if err != nil {
		return
	}
	return
}
func (r *Repository) GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (output UserOutput, err error) {
	query := mainQuery + `where phone_number = $1`
	err = r.Db.QueryRowContext(ctx, query, input).Scan(
		&output.Id,
		&output.Email,
		&output.FullName,
		&output.Gender,
		&output.Occupation,
		&output.Password,
		&output.PhoneNumber,
		&output.LoginCount,
	)
	if err != nil {
		return
	}
	return
}
func (r *Repository) CreateUser(ctx context.Context, input CreateUserInput) (err error) {
	query := `
		INSERT INTO sp_users 
			(
				email, 
				full_name, 
				gender, 
				occupation, 
				password, 
				phone_number
			)
		VALUES ($1, $2, $3, $4, $5, $6);
		`
	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return
	}
	result, err := stmt.Exec(
		input.Email,
		input.FullName,
		input.Gender,
		input.Occupation,
		input.Password,
		input.Password,
	)
	if err != nil {
		return
	}

	_, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}
func (r *Repository) UpdateUser(ctx context.Context, input UpdateUserInput) (err error) {
	query := `update sp_users 
		set 
			full_name = $1,
			phone_number = $2
		where 
			user_id = $3
		`

	stmt, err := r.Db.Prepare(query)
	if err != nil {
		return
	}
	result, err := stmt.Exec(
		input.FullName,
		input.PhoneNumber,
		input.UserId,
	)
	if err != nil {
		return
	}

	_, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}

// can be separated into multiple functions and add transactional for testing just create a simple one
func (r *Repository) Login(ctx context.Context, input LoginInput) (output LoginOutput, err error) {
	data, err := r.GetUserByPhoneNumber(ctx, GetUserByPhoneNumberInput(input.PhoneNumber))
	if err != nil {
		return
	}
	if !util.VerifyPassword(input.Password, data.Password) {
		return "", errors.New("bad password")
	}

	jwt, err := utils.GenerateJWT(data.Id)
	if err != nil {
		return "", err
	}

	//add login count
	loginCount := data.LoginCount + 1
	queryUpdate := `update sp_users
		set 
			login_count = $1
		where 
			user_id = $2
		`

	stmtUpdate, err := r.Db.Prepare(queryUpdate)
	if err != nil {
		return
	}
	_, err = stmtUpdate.Exec(
		loginCount,
		data.Id,
	)
	if err != nil {
		return
	}

	return LoginOutput(jwt), err
}
