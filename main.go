package main

import (
	"github.com/OctavianoRyan25/TODO-List/database"
	"github.com/OctavianoRyan25/TODO-List/router"
)
func main() {
	database.StartDB()
	r:= router.StartServer()
	r.Run(":8080")
}