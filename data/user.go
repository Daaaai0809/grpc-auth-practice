package data

type Role int

const (
	Admin Role = iota
	User
	Guest
)

type UserData struct {
	UserName string
	Password string
	Role     Role
}

var Users = []UserData{
	{
		UserName: "admin",
		Password: "$2a$10$VI4jAbgSAk1fuWZSW1RZ2.N8v/ZDkPuw2jX8onE8hvYWQ6x4phwne",
		Role:     Admin,
	},
	{
		UserName: "user",
		Password: "$2a$10$SbT7ZuDYTAFLpFAt/QEH..33YtTOuNx2g.eQUlZTS9Bhs0/ymg6ra",
		Role:     User,
	},
}
