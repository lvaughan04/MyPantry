package main

import (
	"context"
	"fmt"
	"os"
	
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))

	if err != nil{
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "SELECT 'Hello, MyPantry!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}