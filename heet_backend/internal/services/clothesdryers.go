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

func GetResidentialClothesDryerByID(id float64) (*models.ResidentialClothesDryer, error) {

	sql := `
       SELECT energy_star_unique_id, brand_name, model_name, model_number FROM residential_clothes_dryers
	   WHERE energy_star_unique_id = $1
    `

	row := db.Pool.QueryRow(db.Ctx, sql, id)
	var dryer models.ResidentialClothesDryer
	err := row.Scan(&dryer.EnergyStarUniqueID, &dryer.BrandName, &dryer.ModelName, &dryer.ModelNumber)

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &dryer, nil
}
