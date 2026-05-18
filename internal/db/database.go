package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	DB  *sql.DB
	RDB *redis.Client
	Ctx = context.Background()
)

func InitDatabases() {
	var err error

	// PostgreSQL Connection (password=userpass আপডেট করবেন)
	pgDSN := "host=localhost port=5432 user=dev_user password=Mr.Hridoy445! dbname=pothorokkhi_db sslmode=disable"
	DB, err = sql.Open("postgres", pgDSN)
	if err != nil {
		log.Fatalf("Postgres connection configuration failed: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Postgres database is unreachable: %v", err)
	}
	fmt.Println("✅ PostgreSQL (PostGIS) Connected successfully!")

	// Spatial extension অন করা
	_, _ = DB.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")

	// Redis Connection (Password: "userpass" আপডেট করবেন)
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "RedisHridoy445!",
		DB:       0,
	})

	if _, err = RDB.Ping(Ctx).Result(); err != nil {
		log.Fatalf("Redis Geospatial Cache connection failed: %v", err)
	}
	fmt.Println("✅ Redis Spatial Cache Connected successfully!")
}
