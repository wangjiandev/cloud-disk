package model

type UserBasic struct {
	Id       uint64
	Identity string
	Name     string
	Password string
	Email    string
}

func (u *UserBasic) TabelName() string {
	return "user_basic"
}
