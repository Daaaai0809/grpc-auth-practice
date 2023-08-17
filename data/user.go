package data

const (
	Admin float64 = 1
	User float64 = 2
	Guest float64 = 3
)

type UserData struct {
	UserName string
	Password string
	Role float64
}

var Users = []UserData{
	{
		UserName: "admin",
		Password: "$2a$10$VI4jAbgSAk1fuWZSW1RZ2.N8v/ZDkPuw2jX8onE8hvYWQ6x4phwne",
		Role: Admin,
	},
	{
		UserName: "user",
		Password: "$2a$10$SbT7ZuDYTAFLpFAt/QEH..33YtTOuNx2g.eQUlZTS9Bhs0/ymg6ra",
		Role: User,
	},
}