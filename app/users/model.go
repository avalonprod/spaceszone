package users

import "time"

type User struct {
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	IsBlocked   bool
	IsDeleted   bool
	LastVisitAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
