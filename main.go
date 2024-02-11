package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"go-public-image/controller"
	"go-public-image/repository"
	"go-public-image/service"
)

func main() {
	// dbHost := "localhost"
	// // dbHost := "mysql-db"
	// dbPort := "3306"
	// dbUser := "root"
	// dbPassword := "password"

	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	// Create the MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")

	r := gin.Default()

	svc := service.NewMovieService(repository.NewMovieRepository(db))
	h := controller.NewMovieHandler(svc)
	r.GET("/movies", h.GetMovies)

	err = r.Run()
	if err != nil {
		fmt.Println("Error to run api...")
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
