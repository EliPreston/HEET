package services

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

type LightFixture struct {
	EnergyStarUniqueID       float64 `json:"ENERGY_STAR_unique_id"`
	Partner                  string  `json:"ENERGY_STAR_partner"`
	BrandName                string  `json:"brand_name"`
	ModelName                string  `json:"model_name"`
	ModelNumber              string  `json:"model_number"`
	AdditionalModelInfo      string  `json:"additional_model_information"`
	UPC                      string  `json:"upc"`
	IndoorOutdoor            string  `json:"indoor_outdoor"`
	FixtureType              string  `json:"fixture_type"`
	TotalLightOutput         float64 `json:"total_light_output_(lumens)"`
	TotalInputPower          float64 `json:"total_input_power_(Watts)"`
	Efficiency               float64 `json:"energy_efficiency_(lumens/Watt)"`
	ColorTemperature         float64 `json:"appearance_correlated_color_temperature_(Kelvin)"`
	ColorQuality             float64 `json:"color_quality_(CRI)"`
	LightSourceLife          int     `json:"light_source_life_(hours)"`
	StandbyPower             float64 `json:"standby_power_consumption_(Watts)"`
	ConnectedCapable         string  `json:"connected_capable"`
	ConnectsUsing            string  `json:"connects_using"`
	NetworkSecurityStandards string  `json:"network_security_standards"`
	CanSizes                 string  `json:"can_sizes_(inches)"`
	CanRatings               string  `json:"can_ratings"`
	LightSourceConnection    string  `json:"light_source_connection_base_type"`
	PowerFactor              float64 `json:"power_factor"`
	SpecialFeatures          string  `json:"special_features"`
	Notes                    string  `json:"notes"`
	DateAvailableOnMarket    string  `json:"date_available_on_market"`
	DateCertified            string  `json:"date_certified"`
	Markets                  string  `json:"markets"`
	CBModelIdentifier        string  `json:"CB_model_identifier"`
}

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

type ClothesWasherDryerCombo struct {
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
