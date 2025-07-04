package services

import (
	"fmt"
	// "time"
	// For connecting to Postgres DB
	// "github.com/jackc/pgx/v5"

	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/models"
)

func GetResidentialClothesWashers() ([]models.ResidentialClothesWasher, error) {
	sql := `
       SELECT energy_star_unique_id, brand_name, model_number FROM residential_clothes_washers
	   ORDER BY energy_star_unique_id ASC
		 
    `

	rows, err := db.Pool.Query(db.Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying residential_clothes_washers table: %w", err)
	}
	defer rows.Close()

	var residentialClothesWashers []models.ResidentialClothesWasher
	for rows.Next() {
		var residentialClothesWasher models.ResidentialClothesWasher
		err := rows.Scan(
			&residentialClothesWasher.EnergyStarUniqueID,
			&residentialClothesWasher.BrandName,
			&residentialClothesWasher.ModelNumber)

		if err != nil {
			return nil, fmt.Errorf("error scanning a row in the residential_clothes_washers table: %w", err)
		}
		residentialClothesWashers = append(residentialClothesWashers, residentialClothesWasher)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in residential_clothes_washers table: %w", err)
	}

	return residentialClothesWashers, nil
}

func GetResidentialClothesWasherByID(id float64) (*models.ResidentialClothesWasher, error) {

	sql := `
       SELECT energy_star_unique_id, brand_name, model_number FROM residential_clothes_washers
	   WHERE energy_star_unique_id = $1
    `

	row := db.Pool.QueryRow(db.Ctx, sql, id)
	var clotheswasher models.ResidentialClothesWasher
	err := row.Scan(&clotheswasher.EnergyStarUniqueID, &clotheswasher.BrandName, &clotheswasher.ModelNumber)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &clotheswasher, nil
}
