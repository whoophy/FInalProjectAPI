package main

import (
	"finalproject/handler"
	"finalproject/infra"
)

func main() {
	infra.InitDB()
	r := handler.StartApp()
	r.Run(":8080")
}
