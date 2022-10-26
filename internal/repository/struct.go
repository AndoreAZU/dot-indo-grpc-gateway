package repository

import "time"

type User struct {
	Xid       string    `json:"xid" gorm:"primaryKey"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	TierId    string    `json:"tier_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Following struct {
	XidUser      string `json:"xid_user"`
	XidFollowing string `json:"xid_following"`
}
