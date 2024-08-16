package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	pb "finance_tracker/notification_service/genproto"
	"finance_tracker/notification_service/helper"

	"github.com/google/uuid"
)

type NotificationRepo struct {
	db    *sql.DB
	UserC pb.AuthServiceClient
}

func NewNotificationRepo(db *sql.DB, user_con pb.AuthServiceClient) *NotificationRepo {
	return &NotificationRepo{db: db, UserC: user_con}
}
func (r *NotificationRepo) CreateNotification(req *pb.NotificationCreate) (*pb.Void, error) {
	var sender_id string
	if req.SenderId == "" {
		query := "select id from users where role = 'admin' and deleted_at = 0 limit 1"
		row := r.db.QueryRow(query)
		err := row.Scan(&sender_id)
		if err == sql.ErrNoRows {
			return nil, errors.New("admin not found")
		} else if err != nil {
			return nil, err
		}
	} else {
		sender_id = req.SenderId
	}

	tr, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	//geting the email
	var user_email, user_name string
	email_query := "select email,username from users where id = $1 and deleted_at = 0"
	err = r.db.QueryRow(email_query, sender_id).Scan(&user_email, &user_name)
	if err == sql.ErrNoRows {
		return nil, errors.New("sender not found")
	}

	//sending email
	from := "farruhbekturdiqulov@gmail.com"
	password := "f0313y0513"
	err = helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       user_email,
		Message:  fmt.Sprintf("Hi %s", user_name),
		Code:     req.Message,
	})

	if err != nil {
		tr.Rollback()
		return nil, errors.New("failed to send notification email" + err.Error())
	}
	query := `insert into notifications(id,
										reciever_id,
										sender_id,
										message,
										status)
						values($1,$2,$3,$4,$5)`
	_, err = tr.Exec(query,
		uuid.NewString(),
		req.RecieverId,
		sender_id,
		req.Message,
		"pending")
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) NotifyAll(req *pb.NotificationMessage) (*pb.Void, error) {
	if req.SenderId == "" {
		query := "select id from users where role = 'admin' where deleted_at = 0 limit1"
		row := r.db.QueryRow(query)
		err := row.Scan(&req.SenderId)
		if err == sql.ErrNoRows {
			return nil, errors.New("admin not found")
		} else if err != nil {
			return nil, err
		}
	}

	fil := pb.Filter{}
	filter := pb.UserFilter{
		Role:   req.TargetGroup,
		Filter: &fil,
	}
	users, err := r.UserC.GetAllUsers(context.Background(), &filter)
	if err != nil {
		return nil, err
	}

	for _, user := range users.Users {
		sender_id := req.SenderId

		tr, err := r.db.Begin()
		if err != nil {
			return nil, err
		}

		//sending email
		from := "farruhbekturdiqulov@gmail.com"
		password := "f0313y0513"
		err = helper.SendVerificationCode(helper.Params{
			From:     from,
			Password: password,
			To:       user.Email,
			Message:  fmt.Sprintf("Hi %s", user.FullName),
			Code:     req.Message,
		})

		if err != nil {
			tr.Rollback()
			return nil, errors.New("failed to send notification email" + err.Error())
		}
		query := `insert into notifications(id,
											reciever_id,
											sender_id,
											message)
							values($1,$2,$3,$4)`
		_, err = tr.Exec(query,
			uuid.NewString(),
			user.Id,
			sender_id,
			req.Message)
		if err != nil {
			tr.Rollback()
			return nil, err
		}

		err = tr.Commit()
		if err != nil {
			tr.Rollback()
			return nil, err
		}
	}
	// if err != nil {
	// 	return nil, err
	// }
	return &pb.Void{}, nil
}
func (r *NotificationRepo) DeleteNotificcation(id *pb.ById) (*pb.Void, error) {
	query := `update notifications set deleted_at = EXTRACT(EPOCH FROM now()) 
				where id = $1 and deleted_at = 0`
	_, err := r.db.Exec(query, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (r *NotificationRepo) UpdateNotificcation(req *pb.NotificationUpdate) (*pb.Void, error) {
	query := "UPDATE notifications SET "
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.Body.Message != "" && req.Body.Message != "string" {
		cons = append(cons, fmt.Sprintf("message=$%d", len(args)+1))
		args = append(args, req.Body.Message)
	}
	if req.Body.Status != "" && req.Body.Status != "string" {
		cons = append(cons, fmt.Sprintf("status=$%d", len(args)+1))
		args = append(args, req.Body.Status)
	}

	// Ensure there's at least one field to update
	if len(cons) == 0 {
		return nil, fmt.Errorf("no fields to update")
	}

	query += strings.Join(cons, ", ")
	query += fmt.Sprintf(" WHERE deleted_at = 0 and id=$%d", len(args)+1)
	args = append(args, req.NotificationId)

	// Execute the query
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
func (r *NotificationRepo) GetNotifications(req *pb.NotifFilter) (*pb.NotificationList, error) {

	log.Println("----------------------------", req.RecieverId)
	query := `SELECT id, 
					reciever_id, 
					message, 
					status, 
					created_at,
					sender_id
		FROM notifications 
		WHERE deleted_at = 0`
	var cons []string
	var args []interface{}

	// Dynamically build the query
	if req.RecieverId != "" && req.RecieverId != "string" {
		cons = append(cons, fmt.Sprintf("reciever_id = $%d", len(args)+1))
		args = append(args, req.RecieverId)
	}

	if req.Status != "" && req.Status != "string" {
		cons = append(cons, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, req.Status)
	}
	if req.SenderId != "" && req.SenderId != "string" {
		cons = append(cons, fmt.Sprintf("sender_id = $%d", len(args)+1))
		args = append(args, req.SenderId)
	}

	// Append conditions to query if any exist
	if len(cons) > 0 {
		query += " AND " + strings.Join(cons, " AND ")
	}

	if req.Filter.Limit != 0 || req.Filter.Offset != 0 {
		if req.Filter.Limit != 0 {
			query += fmt.Sprintf("LIMIT $%d", len(args)+1)
			args = append(args, req.Filter.Limit)
		}
		if req.Filter.Offset != 0 {
			query += fmt.Sprintf(" OFFSET $%d", len(args)+1)
			args = append(args, req.Filter.Offset)
		}
	}

	// Execute the query
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	// Prepare the response
	var notificationList pb.NotificationList
	for rows.Next() {
		var notif pb.NotificationGet
		if err := rows.Scan(&notif.Id,
			&notif.UserId,
			&notif.Message,
			&notif.Status,
			&notif.CreatedAt,
			&notif.SenderId); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		notificationList.Notifications = append(notificationList.Notifications, &notif)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}
	for _, n := range notificationList.Notifications {
		query := `update notifications set status = 'read' 
					where deleted_at = 0 and id = $1`
		_, err := r.db.Exec(query, n.Id)
		if err != nil {
			return nil, err
		}
	}

	return &notificationList, nil
}
func (r *NotificationRepo) GetNotification(id *pb.ById) (*pb.NotificationGet, error) {
	query := `select id,
					reciever_id,
					message,
					status,
					created_at,
					sender_id
			from notifications where deleted_at = 0 and id = $1`
	row := r.db.QueryRow(query, id.Id)

	var notif pb.NotificationGet
	err := row.Scan(&notif.Id,
		&notif.UserId,
		&notif.Message,
		&notif.Status,
		&notif.CreatedAt,
		&notif.SenderId)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &notif, nil
}
