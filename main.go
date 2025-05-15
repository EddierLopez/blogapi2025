package main

import (
	"database/sql"
	"log"
	"restapi/api"
	"restapi/dto"
	"restapi/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("No se puede cargar el archivo de configuración", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("No se puede establecer la conexión", err)
	}
	dbtx := dto.NewDbTransaction(conn)
	server, err := api.NewServer(dbtx)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
	err = server.Start(config.ServerURL)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
}
