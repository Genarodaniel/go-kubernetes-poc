package main

import (
	"address-crud-1/internal/server"
)

func main() {
	// if err := config.Load(); err != nil {
	// 	panic(err)
	// }

	// db := database.Connect()
	// defer db.Close()

	// database.Migrate(db)

	s := server.Init()
	s.Run()

}
