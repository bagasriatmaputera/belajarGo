package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Activity struct {
	ID           int       `json: "id"`
	Title        string    `json: "title"`
	Description  string    `json: "description"`
	ActivityDate time.Time `json: "activity_date"`
	Status       string    `json: "status"`
	CreatedAt    time.Time `json: "created_at"`
}

func main() {
	_ = godotenv.Load()

	cfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// Wajib untuk port 6543 (transaction pooler):
	// matikan prepared statement
	cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	var version string
	if err := pool.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Terhubung:", version)

	app := fiber.New()

	// Fetch Activity
	app.Get("activities", func(c fiber.Ctx) error {
		rows, err := pool.Query(context.Background(), "SELECT id, title, description, activity_date, status, created_at FROM activities")
		if err != nil {
			log.Println("gagal koneksi ke database", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal mengabil data dari database",
			})
		}

		defer rows.Close()

		// Menggunakan manual iterasi

		// activities := []Activity{}

		// for rows.Next() {
		// 	// var a Activity
		// 	a := Activity{}

		// 	err := rows.Scan(
		// 		&a.ID,
		// 		&a.Title,
		// 		&a.Description,
		// 		&a.ActivityDate,
		// 		&a.Status,
		// 		&a.CreatedAt,
		// 	)
		// 	if err != nil {
		// 		log.Println("Gagal Scan Row:", err)
		// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 			"error": "Somthing went wrong! Gagal memproses data.",
		// 		})
		// 	}

		// 	activities = append(activities, a)
		// }

		// if err := rows.Err(); err != nil {
		// 	log.Println("Gagal dalam iterasi:", err)
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": err.Error(),
		// 	})
		// }

		// Menggunakan helper pgx v5
		activities, err := pgx.CollectRows(rows, pgx.RowToStructByName[Activity])

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(activities)

	})

	app.Listen(":8080")

}
