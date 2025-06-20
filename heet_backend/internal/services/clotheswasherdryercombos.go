package services

import (
	"fmt"
	// "time"
	// For connecting to Postgres DB
	// "github.com/jackc/pgx/v5"

	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/models"
)

func GetWasherDryerCombos() ([]models.ClothesWasherDryerCombo, error) {
	sql := `
       SELECT energy_star_unique_id, brand_name, model_number FROM washer_dryer_combos
	   ORDER BY energy_star_unique_id ASC
		 
    `

	rows, err := db.Pool.Query(db.Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying washer_dryer_combos table: %w", err)
	}
	defer rows.Close()

	var washerDryerCombos []models.ClothesWasherDryerCombo
	for rows.Next() {
		var washerDryerCombo models.ClothesWasherDryerCombo
		err := rows.Scan(
			&washerDryerCombo.EnergyStarUniqueID,
			&washerDryerCombo.BrandName,
			&washerDryerCombo.ModelNumber)

		if err != nil {
			return nil, fmt.Errorf("error scanning a row in the washer_dryer_combos table: %w", err)
		}
		washerDryerCombos = append(washerDryerCombos, washerDryerCombo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in washer_dryer_combos table: %w", err)
	}

	return washerDryerCombos, nil
}
