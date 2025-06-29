package services

import (
	"fmt"
	// "time"
	// For connecting to Postgres DB
	// "github.com/jackc/pgx/v5"

	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/models"
)

func GetDehumidifiers() ([]models.Dehumidifier, error) {
	sql := `
       SELECT energy_star_unique_id, brand_name, model_name, model_number FROM dehumidifiers
	   ORDER BY energy_star_unique_id ASC
		 
    `

	rows, err := db.Pool.Query(db.Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying dehumidifiers table: %w", err)
	}
	defer rows.Close()

	var dehumidifiers []models.Dehumidifier
	for rows.Next() {
		var dehumidifier models.Dehumidifier
		err := rows.Scan(
			&dehumidifier.EnergyStarUniqueID,
			&dehumidifier.BrandName,
			&dehumidifier.ModelName,
			&dehumidifier.ModelNumber)

		if err != nil {
			return nil, fmt.Errorf("error scanning a row in the dehumidifers table: %w", err)
		}
		dehumidifiers = append(dehumidifiers, dehumidifier)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in dehumidifers table: %w", err)
	}

	return dehumidifiers, nil
}

func GetDehumidifierByID(id float64) (*models.Dehumidifier, error) {

	sql := `
       SELECT energy_star_unique_id, brand_name, model_name, model_number FROM dehumidifiers
	   WHERE energy_star_unique_id = $1
    `

	row := db.Pool.QueryRow(db.Ctx, sql, id)
	var dehumidifier models.Dehumidifier
	err := row.Scan(&dehumidifier.EnergyStarUniqueID, &dehumidifier.BrandName, &dehumidifier.ModelName, &dehumidifier.ModelNumber)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &dehumidifier, nil
}
