package config
import (
	"database/sql"
	// "time"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Type string
	Username string
	Password string
	Host string
	Port string
	Database string
	DB *sql.DB
}

func InitDB(conf Config) DatabaseConfig{
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	return DatabaseConfig{
		Type: conf.Database.Type,
		Username: conf.Database.User,
		Password: conf.Database.Password,
		Host: conf.Database.Host,
		Port: conf.Database.Port,
		Database: conf.Database.DBName,
		DB: db,
	}
}