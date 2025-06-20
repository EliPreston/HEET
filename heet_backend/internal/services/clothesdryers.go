package services

import (
	"fmt"
	// "time"
	// For connecting to Postgres DB
	// "github.com/jackc/pgx/v5"

	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/models"
)

func GetResidentialClothesDryers() ([]models.ResidentialClothesDryer, error) {
	sql := `
       SELECT energy_star_unique_id, brand_name, model_name, model_number FROM residential_clothes_dryers
	   ORDER BY energy_star_unique_id ASC
		 
    `

	rows, err := db.Pool.Query(db.Ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("error querying residential_clothes_dryers table: %w", err)
	}
	defer rows.Close()

	var residentialClothesDryers []models.ResidentialClothesDryer
	for rows.Next() {
		var residentialClothesDryer models.ResidentialClothesDryer
		err := rows.Scan(
			&residentialClothesDryer.EnergyStarUniqueID,
			&residentialClothesDryer.BrandName,
			&residentialClothesDryer.ModelName,
			&residentialClothesDryer.ModelNumber)

		if err != nil {
			return nil, fmt.Errorf("error scanning a row in the residential_clothes_dryers table: %w", err)
		}
		residentialClothesDryers = append(residentialClothesDryers, residentialClothesDryer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows in residential_clothes_dryers table: %w", err)
	}

	return residentialClothesDryers, nil
}
