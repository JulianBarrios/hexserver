package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/julianbarrios/hexserver/infrastructure/server"
	"github.com/julianbarrios/hexserver/infrastructure/storage/mysql"
)

const (
	host   = "localhost"
	port   = 8080
	dbUser = "root"
	dbPass = "root"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", mysqlUri)
	if err != nil {
		return fmt.Errorf("Hubo un error al inicializar la base de datos: %s", err)
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
