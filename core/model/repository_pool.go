package model

import "time"

type RepositoryPool struct {
	Id        int64 `xorm:"not null pk autoincr INTEGER"`
	Identity  string
	Hash      string
	Name      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (r *RepositoryPool) TabelName() string {
	return "repository_pool"
}
