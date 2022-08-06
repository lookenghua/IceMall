package user

type Role string

const (
	RoleADMIN Role = "ADMIN"
	RoleUSER  Role = "USER"
)

type Sex int

const (
	SexMAN     Sex = 1
	SexWOMAN   Sex = 2
	SexUNKNOWN Sex = 0
)
