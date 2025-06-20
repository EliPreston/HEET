package models

type ResidentialDishwasher struct {
	EnergyStarUniqueID                        float64 `json:"ENERGY_STAR_unique_id"`
	BrandName                                 string  `json:"brand_name"`
	ModelNumber                               string  `json:"model_number"`
	AdditionalModelInformation                string  `json:"additional_model_information"`
	UPC                                       string  `json:"upc"`
	DishwasherType                            string  `json:"dishwasher_type"`
	WidthInches                               float64 `json:"width_inches"`
	DepthInches                               float64 `json:"depth_inches"`
	CapacityMaxPlaceSettings                  float64 `json:"capacity_max_place_settings"`
	SoilSensingCapability                     string  `json:"soil_sensing_capability"`
	TubMaterial                               string  `json:"tub_material"`
	DryingMethod                              string  `json:"drying_method"`
	AdditionalProductFeatures                 string  `json:"additional_product_features"`
	AnnualEnergyUseKWh                        float64 `json:"annual_energy_use_kwh"`
	USFederalStandardKWh                      float64 `json:"us_federal_standard_kwh"`
	PercentBetterThanUSFederalStandardKWh     float64 `json:"percent_better_than_us_federal_standard_kwh"`
	WaterUseGallonsPerCycle                   float64 `json:"water_use_gallons_per_cycle"`
	USFederalStandardGallonsPerCycle          float64 `json:"us_federal_standard_gallons_per_cycle"`
	PercentBetterThanUSFederalStandardGallons float64 `json:"percent_better_than_us_federal_standard_gallons"`
	ConnectedCapable                          string  `json:"connected_capable"`
	ConnectsUsing                             string  `json:"connects_using"`
	CommunicationHardwareArchitecture         string  `json:"communication_hardware_architecture"`
	DRProtocol                                string  `json:"DR_protocol"`
	OpenStandardDirectOnPrem                  string  `json:"open_standard_direct_on_prem"`
	DateAvailableOnMarket                     string  `json:"date_available_on_market"`
	DateCertified                             string  `json:"date_certified"`
	Markets                                   string  `json:"markets"`
	CBModelIdentifier                         string  `json:"CB_model_identifier"`
	MeetsMostEfficient2025Criteria            string  `json:"meets_ENERGY_STAR_most_efficient_2025_criteria"`
}
