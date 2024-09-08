package main

import (
	"github.com/open-auth/internal/initialize"
	"github.com/open-auth/pkg/utils"
)

func main() {
	utils.GenerateRSA(2048)
	initialize.Run()
}
