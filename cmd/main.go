package main

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/db" // Добавлен импорт db
	"rest-project/internal/routes"
)

func main() {
	db.InitDB() // Теперь корректно инициализирует БД

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
