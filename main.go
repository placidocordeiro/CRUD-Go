package main;

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/placidocordeiro/CRUD-Go/src/controller/routes"
);

func main() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file");
	}

	fmt.Println(os.Getenv("TEST"));

	router := gin.Default();

	routes.InitRoutes(&router.RouterGroup);

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err);
	}
}
