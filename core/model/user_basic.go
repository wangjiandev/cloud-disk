package model

import "time"

type UserBasic struct {
	Id        uint64 `xorm:"not null pk autoincr INTEGER"`
	Identity  string
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (u *UserBasic) TabelName() string {
	return "user_basic"
}
