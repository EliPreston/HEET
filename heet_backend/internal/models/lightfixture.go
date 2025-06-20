package models

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
