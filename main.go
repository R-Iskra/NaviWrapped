package main

import (
	"github.com/R-Iskra/NaviWrapped/db"
)

func main() {
	db.Init()
	db.Close()
}
