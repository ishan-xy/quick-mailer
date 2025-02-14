package main

import (
	"backend/common"
	"backend/database"
	"backend/router"
)

func main() {
    common.Init()
    database.Init()
    router.Init()
}