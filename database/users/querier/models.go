// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package users

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UsersGender string

const (
	UsersGenderM UsersGender = "M"
	UsersGenderF UsersGender = "F"
)

func (e *UsersGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersGender(s)
	case string:
		*e = UsersGender(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersGender: %T", src)
	}
	return nil
}

type NullUsersGender struct {
	UsersGender UsersGender `json:"users_gender"`
	Valid       bool        `json:"valid"` // Valid is true if UsersGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersGender) Scan(value interface{}) error {
	if value == nil {
		ns.UsersGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersGender), nil
}

type User struct {
	ID                 int32       `json:"id"`
	Username           string      `json:"username"`
	Password           string      `json:"password"`
	AuthTicket         string      `json:"auth_ticket"`
	Email              string      `json:"email"`
	RankID             int32       `json:"rank_id"`
	AccountCreatedDate time.Time   `json:"account_created_date"`
	LastOnlineDate     time.Time   `json:"last_online_date"`
	IsOnline           bool        `json:"is_online"`
	Motto              string      `json:"motto"`
	Look               string      `json:"look"`
	Gender             UsersGender `json:"gender"`
	IpRegister         string      `json:"ip_register"`
	IpCurrent          string      `json:"ip_current"`
	HomeRoom           int32       `json:"home_room"`
}
