package models

type Dehumidifier struct {
	EnergyStarUniqueID             float64 `json:"ENERGY_STAR_unique_id"`
	Partner                        string  `json:"ENERGY_STAR_partner"`
	BrandName                      string  `json:"brand_name"`
	ModelName                      string  `json:"model_name"`
	ModelNumber                    string  `json:"model_number"`
	AdditionalModelInfo            string  `json:"additional_model_information"`
	UPC                            string  `json:"upc"`
	Type                           string  `json:"dehumidifier_type"`
	AlternateConfigType            string  `json:"alternate_configuration_type"`
	Efficiency                     float64 `json:"dehumidifier_efficiency_(Integrated Energy Factor - L/kWh)"`
	WaterRemovalCapacity           float64 `json:"dehumidifier_water_removal_capacity_(pints/day)"`
	CaseVolume                     float64 `json:"dehumidifier_case_volume_(cu. ft.)"`
	AnnualEnergyConsumption        float64 `json:"annual_energy_consumption_(kWh/year)"`
	RefrigerantType                string  `json:"refrigerant_type"`
	RefrigerantWithGWP             string  `json:"refrigerant_with_GWP"`
	RefrigerantCharge              float64 `json:"refrigerant_charge_(ounces)"`
	ConnectedCapable               string  `json:"connected_capable"`
	ConnectsUsing                  string  `json:"connects_Using"`
	ConnectedFeatures              string  `json:"connected_features"`
	DateAvailableOnMarket          string  `json:"date_available_on_market"`
	DateCertified                  string  `json:"date_certified"`
	Markets                        string  `json:"markets"`
	CBModelIdentifier              string  `json:"CB_model_identifier"`
	MeetsMostEfficient2025Criteria string  `json:"meets_ENERGY_STAR_most_efficient_2025_criteria"`
}
