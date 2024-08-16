package postgres

import (
	// "context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	redis "github.com/go-redis/redis"
	"github.com/google/uuid"

	pb "finance_tracker/auth_service/genproto"
	"finance_tracker/auth_service/helper"
	t "finance_tracker/auth_service/token"
)

type UserRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewUserRepo(db *sql.DB, rdb *redis.Client) *UserRepo {
	return &UserRepo{db: db, rdb: rdb}
}

func (ur *UserRepo) RegisterUser(req *pb.UserCreateReq) (*pb.Void, error) {
	query := `insert into users(id,
	                            username, 
								email, 
								password, 
								full_name, 
								dob,
								role)
				values($1,$2,$3,$4,$5,$6,$7)`
	_, err := ur.db.Exec(query,
		uuid.NewString(),
		req.UserName,
		req.Email,
		req.Password,
		req.FullName,
		req.Dob,
		"user")
	if err != nil {
		return nil, err
	}

	code, err := helper.GenerateRandomCode(6)
	if err != nil {
		return nil, errors.New("failed to generate code for verification" + err.Error())
	}

	ur.rdb.Set(code, req.Email, time.Minute*3)

	from := "shamsioqilov@gmail.com"
	password := "rdpo ehng abtm fuzy"
	err = helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification:%s", req.Email, code),
		Code:     code,
	})

	if err != nil {
		return nil, errors.New("failed to send email verification code" + err.Error())
	}

	return &pb.Void{}, nil
}

func (r *UserRepo) Login(req *pb.LoginReq) (*pb.Token, error) {
	query := `select id, username, role from users where username=$1 and password=$2 and confirmed = true`
	var id, username, role string
	err := r.db.QueryRow(query, req.UserName, req.Password).Scan(&id, &username, &role)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found: invalid username or password")
	} else if err != nil {
		return nil, err
	}

	access, _ := t.GenerateJWTToken(id, username, role)

	t := time.Now().Add(time.Minute * 60).String()
	return &pb.Token{AccessToken: access, ExpiresAt: t}, nil
}
func (r *UserRepo) RefreshToken(req *pb.Token) (*pb.Token, error) {
	access := req.AccessToken

	claims, err := t.JustExtractClaim(access)
	if err != nil {
		return nil, err
	}

	id := claims["id"]
	username := claims["username"]
	role := claims["role"]

	new_token, _ := t.GenerateJWTToken(id, username, role)

	t := time.Now().Add(time.Minute * 60).String()
	return &pb.Token{AccessToken: new_token, ExpiresAt: t}, nil
}
func (r *UserRepo) UpdateProfile(req *pb.UserUpdate) (*pb.Void, error) {
	query := "UPDATE users SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.FullName != "" && req.Body.FullName != "string" {
		cons = append(cons, fmt.Sprintf("full_name=$%d", len(args)+1))
		args = append(args, req.Body.FullName)
	}
	if req.Body.Dob != "" && req.Body.Dob != "string" {
		cons = append(cons, fmt.Sprintf("dob=$%d", len(args)+1))
		args = append(args, req.Body.Dob)
	}
	if req.Body.Language != "" && req.Body.Language != "string" {
		cons = append(cons, fmt.Sprintf("language=$%d", len(args)+1))
		args = append(args, req.Body.Language)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d and confirmed = true", len(args)+1)
	args = append(args, req.Id)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *UserRepo) GetUserProfile(id *pb.ById) (*pb.UserCreateRes, error) {
	query := `SELECT id, 
					username, 
					email, 
					full_name, 
					dob, 
					created_at
                FROM users
                WHERE id=$1 AND deleted_at = 0 and confirmed = true`
	var user pb.UserCreateRes
	err := r.db.QueryRow(query, id.Id).Scan(&user.Id,
		&user.UserName,
		&user.Email,
		&user.FullName,
		&user.Dob,
		&user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepo) ChangePassword(req *pb.PasswordChangeReq) (*pb.Void, error) {
	// Start a transaction
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var curPassword string
	query := `SELECT password FROM users WHERE id = $1 AND deleted_at = 0 and confirmed = true`

	// Get the current password
	err = tx.QueryRow(query, req.UserId).Scan(&curPassword)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found with these credentials")
	} else if err != nil {
		return nil, fmt.Errorf("failed to query current password: %v", err)
	}

	// Check if the old password matches
	if curPassword != req.Body.OldPassword {
		return nil, fmt.Errorf("your password does not match your current password")
	}

	// Update the password
	updateQuery := `UPDATE users SET password = $1 WHERE id = $2 AND deleted_at = 0`
	_, err = tx.Exec(updateQuery, req.Body.NewPassword, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to update password: %v", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &pb.Void{}, nil
}
func (r *UserRepo) ForgotPassword(req *pb.ForgotPasswordReq) (*pb.Void, error) {

	code, err := helper.GenerateRandomCode(6)
	if err != nil {
		return nil, errors.New("failed to generate code for verification" + err.Error())
	}

	// r.rdb.Set(context.Background(),req.Body.Email,code,time.Minute * 3)
	r.rdb.Set(code, req.Body.Email, time.Minute*7)

	from := "shamsioqilov@gmail.com"
	password := "rdpo ehng abtm fuzy"
	err = helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       req.Body.Email,
		Message:  fmt.Sprintf("Hi %s, your verification:%s", req.Body.Email, code),
		Code:     code,
	})

	if err != nil {
		return nil, errors.New("failed to send verification email" + err.Error())
	}
	return &pb.Void{}, nil
}
func (r *UserRepo) ResetPassword(req *pb.PasswordResetReq) (*pb.Void, error) {
	em, err := r.rdb.Get(req.Body.ResetCode).Result()
	log.Println(req.Body.ResetCode, err)
	if err != nil {
		return nil, errors.New("invalid code or code expired")
	}
	log.Println(em)

	query := `update users set password = $1 where email = $2`
	_, err = r.db.Exec(query, req.Body.NewPassword, em)
	log.Println(req.Body.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to reset password: %v", err)
	}
	return &pb.Void{}, nil
}
func (r *UserRepo) ConfirmEmail(req *pb.EmailConfirm) (*pb.Void, error) {
	em, err := r.rdb.Get(req.Code).Result()
	// log.Println(req.Body.ResetCode, err)
	if err != nil {
		return nil, errors.New("invalid code or code expired")
	}
	// log.Println(em)

	query := `update users set confirmed = true  where email = $1 and deleted_at = 0`
	_, err = r.db.Exec(query, em)
	// log.Println(req.Body.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to update usrs to confirm the email: %v", err)
	}
	return &pb.Void{}, nil
}
func (r *UserRepo) ResendCode(req *pb.ResendReq) (*pb.Void, error) {
	code, err := helper.GenerateRandomCode(6)
	if err != nil {
		return nil, errors.New("failed to generate code for verification" + err.Error())
	}

	// r.rdb.Set(context.Background(),req.Body.Email,code,time.Minute * 3)
	r.rdb.Set(code, req.Email, time.Minute*5)

	from := "shamsioqilov@gmail.com"
	password := "rdpo ehng abtm fuzy"
	err = helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi %s, your verification:%s", req.Email, code),
		Code:     code,
	})

	if err != nil {
		return nil, errors.New("failed to send verification email" + err.Error())
	}
	return &pb.Void{}, nil
}
func (r *UserRepo) CreateManger(req *pb.UserCreateReq) (*pb.Void, error) {
	query := `insert into users(id,
	username, 
	email, 
	password, 
	full_name, 
	dob,
	role,
	confirmed)
	values($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := r.db.Exec(query,
		uuid.NewString(),
		req.UserName,
		req.Email,
		req.Password,
		req.FullName,
		req.Dob,
		"kitchen_manager",
		"true",
	)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *UserRepo) GetAllUsers(req *pb.UserFilter) (*pb.UserList, error) {
	query := `select id,
					username, 
					email,
					full_name,
					dob,
					created_at 
					from users where deleted_at = 0`
	var cons []string
	var args []interface{}

	if req.Role != "" && req.Role != "string" {
		cons = append(cons, fmt.Sprintf("role=$%d", len(args)+1))
		args = append(args, req.Role)
	}
	if req.IsWorking != "" && req.IsWorking != "string" {
		cons = append(cons, fmt.Sprintf("dob=$%d", len(args)+1))
		args = append(args, req.IsWorking)
	}
	// Append conditions to query if any exist
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Filter.Limit, req.Filter.Offset)

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	var users pb.UserList
	defer rows.Close()
	for rows.Next() {
		var user pb.UserCreateRes
		if err := rows.Scan(&user.Id,
			user.UserName,
			user.Email,
			user.FullName,
			user.Dob,
			user.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users.Users = append(users.Users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}
	return &users, nil
}
