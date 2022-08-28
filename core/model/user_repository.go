package model

import "time"

type UserRepository struct {
	Id                 int64 `xorm:"not null pk autoincr INTEGER"`
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (r *UserRepository) TabelName() string {
	return "user_repository"
}
