package main

import (
	"github.com/open-auth/global"
	"github.com/open-auth/internal/initialize"
	"github.com/open-auth/pkg/utils"
)

func main() {
	utils.GenerateRSA(2048, global.AdminScope)
	utils.GenerateRSA(2048, global.UserScope)
	initialize.Run()
}
