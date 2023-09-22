package main

import (
	"companies/internal/config"
	"companies/internal/rest"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db := DBConnection(cfg)

	defer db.Close()

	ExposeRestServer(db)
}

func ExposeRestServer(db *gorm.DB) {
	engine := gin.Default()

	router := rest.NewRouter(engine.RouterGroup, db)
	router.RouteControllers()

	err := engine.Run()
	if err != nil {
		log.Fatalf("error starting server, %v", err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func DBConnection(cfg config.Config) *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error connecting db %v", err)
	}

	return db
}
