// main.go

package main

import (
	"database/sql"
	"log"

	"example/usermanagementapi/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite database connection
	var err error
	db, err := sql.Open("sqlite3", "C:/Users/batuh/Desktop/user-management/UserInfo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Set up user router with the database connection
	routers.SetupUserRouter(db, router)

	// Run the server on port 8080
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
