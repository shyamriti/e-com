package main

import (
	"e-com/pkg/database"
	"e-com/pkg/routes"
)

func main() {
	database.Connection()
	r := routes.Routers()
	r.Run(":8000")
}
