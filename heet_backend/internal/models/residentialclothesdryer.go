package models

type ResidentialClothesDryer struct {
	EnergyStarUniqueID               float64 `json:"ENERGY_STAR_unique_id"`
	BrandName                        string  `json:"brand_name"`
	ModelName                        string  `json:"model_name"`
	ModelNumber                      string  `json:"model_number"`
	AdditionalModelInfo              string  `json:"additional_model_information"`
	UPC                              string  `json:"upc"`
	ProductType                      string  `json:"product_type"`
	FuelType                         string  `json:"fuel_type"`
	SpecialType                      string  `json:"special_type"`
	HeatPumpTechnology               string  `json:"heat_pump_technology"`
	Voltage                          float64 `json:"voltage_(V)"`
	DrumCapacity                     float64 `json:"drum_Capacity_(cu-ft)"`
	HeightInches                     float64 `json:"height_(inches)"`
	WidthInches                      float64 `json:"width_(inches)"`
	DepthInches                      float64 `json:"depth_(inches)"`
	CombinedEnergyFactor             float64 `json:"combined_energy_factor_(CEF)"`
	EstimatedAnnualEnergyUse         float64 `json:"estimated_annual_energy_use_(kWh/yr)"`
	EstimatedTestCycleTime           string  `json:"estimated_energy_test_cycle_time_(min)"`
	EnergyTestCycleInformation       string  `json:"energy_test_cycle_information"`
	AdditionalDryerFeatures          string  `json:"additional_dryer_features"`
	VentedOrVentless                 string  `json:"vented_ventless"`
	RefrigerantType                  string  `json:"refrigerant_type"`
	RefrigerantWithGWP               string  `json:"refrigerant_with_GWP"`
	Connected                        string  `json:"connected"`
	ConnectsUsing                    string  `json:"connects_using"`
	CommunicationStandard            string  `json:"communication_standard_application_layer"`
	OpenStandardBasedInterconnection string  `json:"direct_on-premises_open-standard_based_interconnection"`
	CombinedEnergyFactorMaxDryness   float64 `json:"calculated_combined_energy_factor_max_dryness_setting_(lbs/kWh)"`
	PairedWasherAvailable            string  `json:"paired_ENERGY_STAR_clothes_washer_available"`
	PairedWasherModelIdentifier      string  `json:"paired_ENERGY_STAR_clothes_washer_ENERGY_STAR_model_identifier"`
	DateAvailableOnMarket            string  `json:"date_available_on_market"`
	DateCertified                    string  `json:"date_certified"`
	Markets                          string  `json:"markets"`
	ModelIdentifier                  string  `json:"CB_model_identifier"`
	MeetsMostEfficient2025Criteria   string  `json:"meets_ENERGY_STAR_most_efficient_2025_criteria"`
}
