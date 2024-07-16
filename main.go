package main

import (
	"github.com/PerryHuan9/learn_go/router"
	"github.com/PerryHuan9/learn_go/sql"
)

func main() {
	sql.Connect()
	router.StartServer()

}
