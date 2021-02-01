package db

type DbCofig struct {
	DbName   string
	Username string
	Password string
}

func GetDbConfig() *DbCofig {
	return &DbCofig{
		DbName:   "todoapp",
		Username: "adityaladwa",
		Password: "",
	}
}
