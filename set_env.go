package main

import (
	"github.com/joho/godotenv"
)

func setenv() {
	godotenv.Load(".client")
}
