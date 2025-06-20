package models

type ResidentialClothesWasher struct {
	EnergyStarUniqueID               float64 `json:"ENERGY_STAR_unique_id"`
	BrandName                        string  `json:"brand_name"`
	ModelNumber                      string  `json:"model_number"`
	AdditionalModelInfo              string  `json:"additional_model_information"`
	UPC                              string  `json:"upc"`
	LoadConfiguration                string  `json:"load_configuration"`
	SpecialType                      string  `json:"special_type"`
	AdditionalWasherFeatures         string  `json:"additional_washer_features"`
	IntendedMarket                   string  `json:"intended_market"`
	VolumeCuFt                       float64 `json:"volume_(cu. ft.)"`
	HeightInches                     float64 `json:"height_(inches)"`
	WidthInches                      float64 `json:"width_(inches)"`
	DepthInches                      float64 `json:"depth_(inches)"`
	IMEF                             float64 `json:"integrated_modified_energy_factor_(IMEF)"`
	FederalIMEFStandard              float64 `json:"US_federal_standard_(IMEF)"`
	AnnualEnergyUse                  float64 `json:"annual_energy_use_(kWh/yr)"`
	IWF                              float64 `json:"integrated_water_factor_(IWF)"`
	FederalIWFStandard               float64 `json:"US_federal_standard_(IWF)"`
	AnnualWaterUse                   float64 `json:"annual_water_use_(gallons/yr)"`
	DrumCapacity                     string  `json:"drum_capacity_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	HeatPumpTechnology               string  `json:"heat_pump_technology_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	CombinedEnergyFactor             string  `json:"combined_energy_factor_(CEF)_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	EstimatedAnnualDryerEnergyUse    string  `json:"estimated_annual_energy_use_(kWh/yr)_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	EstimatedEnergyTestCycleTime     string  `json:"estimated_energy_test_cycle_time_(min)_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	EnergyTestCycleInfo              string  `json:"energy_test_cycle_information_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	CombinedEnergyFactorMaxDryness   string  `json:"calculated_combined_energy_factor_max_dryness_setting_(lbs/kWh)_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	AdditionalDryerFeatures          string  `json:"additional_dryer_features_for_the_dryer_in_a_combination_all-in-one_washer-dryer"`
	Connected                        string  `json:"connected"`
	ConnectsUsing                    string  `json:"connects_using"`
	CommunicationStandard            string  `json:"communication_standard_application_layer"`
	OpenStandardBasedInterconnection string  `json:"direct_on-premises_open-standard_based_interconnection"`
	PairedDryerAvailable             string  `json:"paired_ENERGY_STAR_clothes_dryer_available"`
	PairedDryerModelIdentifier       string  `json:"paired_ENERGY_STAR_clothes_dryer_ENERGY_STAR_model_identifier"`
	DateAvailableOnMarket            string  `json:"date_available_on_market"`
	DateCertified                    string  `json:"date_certified"`
	Markets                          string  `json:"markets"`
	ModelIdentifier                  string  `json:"ENERGY_STAR_model_identifier"`
	MeetsMostEfficient2025Criteria   string  `json:"meets_ENERGY_STAR_most_efficient_2025_criteria"`
}
