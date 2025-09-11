package db

import (
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/infra/db/seeder"
	"github.com/rs/zerolog/log"
)

func (db *Database) Seed() {
	queryInsertProduct := `
		INSERT INTO products (
			name,
			price
		) VALUES (
			:name,
			:price
		);`

	successCount := 0
	errorCount := 0

	for _, p := range seeder.Products {
		_, err := db.Conn.NamedExec(queryInsertProduct, p)
		if err != nil {
			log.Error().
				Err(err).
				Str("product_name", p.Name).
				Msg("[SEEDER][PRODUCT] Failed to insert product")

			errorCount++
			continue
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("errors", errorCount).
		Msg("[SEEDER][PRODUCT] Seeding completed")
}
