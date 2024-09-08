package main

import (
	"github.com/go-open-auth/internal/initialize"
	"github.com/go-open-auth/pkg/utils"
)

func main() {
	utils.GenerateRSA(2048)
	initialize.Run()
}
