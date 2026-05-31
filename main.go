package main

import (
	"github.com/R-Iskra/NaviWrapped/db"
	"github.com/R-Iskra/NaviWrapped/server"
)

func main() {
	db.Init()
	server.Start()
}
