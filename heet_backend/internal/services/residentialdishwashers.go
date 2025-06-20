package services

import (
	"fmt"
	// "time"
	// For connecting to Postgres DB
	// "github.com/jackc/pgx/v5"

	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/models"
)

func GetResidentialDishwashers() ([]models.ResidentialDishwasher, error) {
	sql := `
       SELECT energy_star_unique_id, brand_name, model_number, depth_inches, width_inches FROM residential_dishwashers
	   ORDER BY energy_star_unique_id ASC
		 
    `

	rows, err := db.Pool.Query(db.Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying residential_dishwashers table: %w", err)
	}
	defer rows.Close()

	var residentialDishwashers []models.ResidentialDishwasher
	for rows.Next() {
		var residentialDishwasher models.ResidentialDishwasher
		err := rows.Scan(
			&residentialDishwasher.EnergyStarUniqueID,
			&residentialDishwasher.BrandName,
			&residentialDishwasher.ModelNumber,
			&residentialDishwasher.DepthInches,
			&residentialDishwasher.WidthInches)

		if err != nil {
			return nil, fmt.Errorf("error scanning a row in the residential_dishwashers table: %w", err)
		}
		residentialDishwashers = append(residentialDishwashers, residentialDishwasher)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in residential_dishwashers table: %w", err)
	}

	return residentialDishwashers, nil
}
