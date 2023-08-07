package main

import (
	"fish/internal/config"
	"fish/internal/web"
)

func main() {
	config.AddFile("db")
	web.NewWeb().Run(":8000")
}
