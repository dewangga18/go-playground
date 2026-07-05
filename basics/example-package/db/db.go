package db

var connection string

func init() {
	connection = "MySql"
}

func GetDB() string {
	return connection
}