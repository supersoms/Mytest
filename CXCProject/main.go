package main

import (
	orm "CXCProject/database"
	"CXCProject/router"
	"github.com/gin-gonic/gin"
)

func main() {
	defer orm.DB.Close()
	orm.DB.LogMode(true)
	gin.ForceConsoleColor()
	router := router.InitRouter()
	router.Run("127.0.0.1:8085")
}
