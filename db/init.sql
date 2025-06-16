-- Create a sample table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT INTO users (name, email) VALUES ('Tester01', 'test@test.com');


-- -----------------------------------------------------------

-- Dehumidifiers
-- ENERGY STAR Unique ID,ENERGY STAR Partner,Brand Name,Model Name,Model Number,Additional Model Information,UPC,Dehumidifier Type,Alternate Configuration Type,Dehumidifier Efficiency (Integrated Energy Factor - L/kWh),Dehumidifier Water Removal Capacity (pints/day),Whole-home Dehumidifier Case Volume (cu. ft.),Annual Energy Consumption (kWh/year),Refrigerant Type,Refrigerant with GWP,Refrigerant Charge (ounces),Connected Capable,Connects Using,Connected Features,Date Available on Market,Date Certified,Markets,CB Model Identifier,Meets ENERGY STAR Most Efficient 2025 Criteria

CREATE TABLE dehumidifiers (
    energy_star_unique_id NUMERIC PRIMARY KEY,
    energy_star_partner TEXT,
    brand_name TEXT,
    model_name TEXT,
    model_number TEXT,
    additional_model_information TEXT,
    upc TEXT,
    dehumidifier_type TEXT,
    alternate_configuration_type TEXT,
    efficiency_l_per_kwh NUMERIC,
    water_removal_capacity_pints_per_day NUMERIC,
    whole_home_case_volume_cuft NUMERIC,
    annual_energy_consumption_kwh_per_year NUMERIC,
    refrigerant_type TEXT,
    refrigerant_with_gwp TEXT,
    refrigerant_charge_oz NUMERIC,
    connected_capable TEXT,
    connects_using TEXT,
    connected_features TEXT,
    date_available_on_market DATE,
    date_certified DATE,
    markets TEXT,
    cb_model_identifier TEXT,
    meets_most_efficient_2025 TEXT
);

COPY dehumidifiers FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Dehumidifiers_20250607.csv'
DELIMITER ','
CSV HEADER;


-- FOR SOME REASON THERE IS AN ISSUE WITH THIS CSV FILE, KEEP GETTING "invalid record length at 0/19139F0: expected at least 24, got 0"
    -- Attempted to remove any blank rows if there were, validate rows all had same length as header with small script, still getting error
-- Light fixtures - Downlights
-- ENERGY STAR Unique ID,ENERGY STAR Partner,Brand Name,Model Name,Model Number,Additional Model Information,UPC,Indoor/Outdoor,Fixture Type,Total Light Output (lumens),Total Input Power (Watts),Energy Efficiency (lumens/Watt),Appearance/Correlated Color Temperature (Kelvin),Color Quality (CRI),Light Source Life (hours),Standby Power Consumption (Watts),Connected Capable,Connects Using,Network Security Standards,Can Size(s) (inches),Can Rating(s),Light Source Connection/Base Type,Power Factor,Special Features,Notes,Date Available On Market,Date Certified,Markets,CB Model Identifier

-- CREATE TABLE lighting_fixtures_downlights (
--     energy_star_unique_id NUMERIC PRIMARY KEY,
--     energy_star_partner TEXT,
--     brand_name TEXT,
--     model_name TEXT,
--     model_number TEXT,
--     additional_model_information TEXT,
--     upc TEXT,
--     indoor_outdoor TEXT,
--     fixture_type TEXT,
--     total_light_output_lumens NUMERIC,
--     total_input_power_watts NUMERIC,
--     energy_efficiency_lumens_per_watt NUMERIC,
--     correlated_color_temp_kelvin NUMERIC,
--     color_quality_cri NUMERIC,
--     light_source_life_hours NUMERIC,
--     standby_power_watts NUMERIC,
--     connected_capable TEXT,
--     connects_using TEXT,
--     network_security_standards TEXT,
--     can_sizes_inches TEXT,
--     can_ratings TEXT,
--     light_source_connection_base_type TEXT,
--     power_factor NUMERIC,
--     special_features TEXT,
--     notes TEXT,
--     date_available_on_market TEXT,
--     date_certified TEXT,
--     markets TEXT,
--     cb_model_identifier TEXT
-- );

-- COPY lighting_fixtures_downlights FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Light_Fixtures_-_Downlights_20250607.csv'
-- DELIMITER ','
-- CSV HEADER;


-- Dishwashers - Residential
-- ENERGY STAR Unique ID,Brand Name,Model Number,Additional Model Information,UPC,Type,Width (inches),Depth (inches),Capacity - Maximum Number of Place Settings,Soil-Sensing Capability,Tub Material,Drying Method,Additional Product Features,Annual Energy Use (kWh/yr),US Federal Standard (kWh/yr),% Better than US Federal Standard (kWh/yr),Water Use (gallons/cycle),US Federal Standard (gallons/cycle),% Better than US Federal Standard (gallons/cycle),Connected Capable,Connects Using,Communication Hardware Architecture,DR Protocol,Direct on-premises Open-standard Based Interconnection,Date Available On Market, Date Certified,Markets,CB Model Identifier,Meets ENERGY STAR Most Efficient 2025 Criteria

CREATE TABLE residential_dishwashers (
    energy_star_unique_id TEXT PRIMARY KEY,
    brand_name TEXT,
    model_number TEXT,
    additional_model_information TEXT,
    upc TEXT,
    dishwasher_type TEXT,
    width_inches NUMERIC,
    depth_inches NUMERIC,
    capacity_max_place_settings NUMERIC,
    soil_sensing_capability TEXT,
    tub_material TEXT,
    drying_method TEXT,
    additional_product_features TEXT,
    annual_energy_use_kwh NUMERIC,
    us_federal_standard_kwh NUMERIC,
    percent_better_than_us_federal_standard_kwh NUMERIC,
    water_use_gallons_per_cycle NUMERIC,
    us_federal_standard_gallons_per_cycle NUMERIC,
    percent_better_than_us_federal_standard_gallons NUMERIC,
    connected_capable TEXT,
    connects_using TEXT,
    communication_hardware_architecture TEXT,
    dr_protocol TEXT,
    open_standard_direct_on_prem TEXT,
    date_available_on_market TEXT,
    date_certified TEXT,
    markets TEXT,
    cb_model_identifier TEXT,
    meets_most_efficient_2025 TEXT
);

COPY residential_dishwashers FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Residential_Dishwashers_20250607.csv'
DELIMITER ','
CSV HEADER;


-- Clothes Dryers - Residential
-- ENERGY STAR Unique ID,Brand Name,Model Name,Model Number,Additional Model Information,UPC,Product Type,Fuel Type,Special Type,Heat Pump Technology,Voltage (V),Drum Capacity (cu-ft),Height (inches),Width (inches),Depth (inches),Combined Energy Factor (CEF),Estimated Annual Energy Use (kWh/yr),Estimated Energy Test Cycle Time (min),Energy Test Cycle Information,Additional Dryer Features,Vented or Ventless,Refrigerant Type,Refrigerant with GWP,Connected,Connects Using,Communication Standard - Application Layer,Direct on-premises Open-standard Based Interconnection,Calculated Combined Energy Factor - Max Dryness Setting (lbs/kWh),Paired ENERGY STAR Clothes Washer Available,Paired ENERGY STAR Clothes Washer ENERGY STAR Model Identifier,Date Available on Market,Date Certified,Markets,CB Model Identifier,Meets ENERGY STAR Most Efficient 2025 Criteria

CREATE TABLE residential_clothes_dryers (
    energy_star_unique_id NUMERIC PRIMARY KEY,
    brand_name TEXT,
    model_name TEXT,
    model_number TEXT,
    additional_model_information TEXT,
    upc TEXT,
    product_type TEXT,
    fuel_type TEXT,
    special_type TEXT,
    heat_pump_technology TEXT,
    voltage_v NUMERIC,
    drum_capacity_cu_ft NUMERIC,
    height_inches NUMERIC,
    width_inches NUMERIC,
    depth_inches NUMERIC,
    combined_energy_factor_cef NUMERIC,
    estimated_annual_energy_use_kwh NUMERIC,
    estimated_test_cycle_time_min NUMERIC,
    energy_test_cycle_info TEXT,
    additional_dryer_features TEXT,
    vented_or_ventless TEXT,
    refrigerant_type TEXT,
    refrigerant_with_gwp TEXT,
    connected TEXT,
    connects_using TEXT,
    comm_standard_app_layer TEXT,
    open_standard_direct_on_prem TEXT,
    calculated_cef_max_dryness NUMERIC,
    paired_energy_star_washer_available TEXT,
    paired_washer_model_identifier TEXT,
    date_available_on_market TEXT,
    date_certified TEXT,
    markets TEXT,
    cb_model_identifier TEXT,
    meets_most_efficient_2025 TEXT
);

COPY residential_clothes_dryers FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Residential_Clothes_Dryers_20250607.csv'
DELIMITER ','
CSV HEADER;


-- Washer-Dryer Combos
-- ENERGY STAR Unique ID,Brand Name,Model Number,Additional Model Information,UPC,Load Configuration,Special Type,Additional Washer Features,Intended Market,Volume (cu. ft.),Height (inches),Width (inches),Depth (inches),Integrated Modified Energy Factor (IMEF),US Federal Standard (IMEF),Annual Energy Use (kWh/yr),Integrated Water Factor (IWF),US Federal Standard (IWF),Annual Water Use (gallons/yr),Drum Capacity for the dryer in a Combination All-in-One Washer-Dryer,Heat Pump Technology for the dryer in a Combination All-in-One Washer-Dryer,Combined Energy Factor (CEF) for the dryer in a Combination All-in-One Washer-Dryer,Estimated Annual Energy Use (kWh/yr) for the dryer in a Combination All-in-One Washer-Dryer,Estimated Energy Test Cycle Time (min) for the dryer in a Combination All-in-One Washer-Dryer,Energy Test Cycle Information for the dryer in a Combination All-in-One Washer-Dryer,Calculated Combined Energy Factor - Max Dryness Setting (lbs/kWh) for the dryer in a Combination All-in-One Washer-Dryer,Additional Dryer Features for the dryer in a Combination All-in-One Washer-Dryer,Connected,Connects Using,Communication Standard Application Layer,Direct on-premises Open-standard Based Interconnection,Paired ENERGY STAR Clothes Dryer Available,Paired ENERGY STAR Clothes Dryer ENERGY STAR Model Identifier,Date Available On Market,Date Certified,Markets,ENERGY STAR Model Identifier,Meets ENERGY STAR Most Efficient 2025 Criteria

CREATE TABLE washer_dryer_combos (
    energy_star_unique_id NUMERIC PRIMARY KEY,
    brand_name TEXT,
    model_number TEXT,
    additional_model_information TEXT,
    upc TEXT,
    load_configuration TEXT,
    special_type TEXT,
    additional_washer_features TEXT,
    intended_market TEXT,
    volume_cu_ft NUMERIC,
    height_inches NUMERIC,
    width_inches NUMERIC,
    depth_inches NUMERIC,
    imef NUMERIC,
    imef_us_federal_standard NUMERIC,
    annual_energy_use_kwh NUMERIC,
    iwf NUMERIC,
    iwf_us_federal_standard NUMERIC,
    annual_water_use_gallons NUMERIC,
    drum_capacity_dryer_combo NUMERIC,
    heat_pump_tech_combo_dryer TEXT,
    combined_energy_factor_cef NUMERIC,
    estimated_annual_energy_use_dryer NUMERIC,
    estimated_test_cycle_time_mins NUMERIC,
    energy_test_cycle_info TEXT,
    calculated_cef_max_dryness NUMERIC,
    additional_dryer_features TEXT,
    connected TEXT,
    connects_using TEXT,
    comm_standard_app_layer TEXT,
    open_standard_direct_on_prem TEXT,
    paired_energy_star_dryer_available TEXT,
    paired_dryer_model_identifier TEXT,
    date_available_on_market TEXT,
    date_certified TEXT,
    markets TEXT,
    energy_star_model_identifier TEXT,
    meets_most_efficient_2025 TEXT
);

COPY washer_dryer_combos FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Res-Combo-Washer-Dryer_20250607.csv'
DELIMITER ','
CSV HEADER;


-- Clothes Washers - Residential
-- ENERGY STAR Unique ID,Brand Name,Model Number,Additional Model Information,UPC,Load Configuration,Special Type,Additional Washer Features,Intended Market,Volume (cu. ft.),Height (inches),Width (inches),Depth (inches),Integrated Modified Energy Factor (IMEF),US Federal Standard (IMEF),Annual Energy Use (kWh/yr),Integrated Water Factor (IWF),US Federal Standard (IWF),Annual Water Use (gallons/yr),Drum Capacity for the dryer in a Combination All-in-One Washer-Dryer,Heat Pump Technology for the dryer in a Combination All-in-One Washer-Dryer,Combined Energy Factor (CEF) for the dryer in a Combination All-in-One Washer-Dryer,Estimated Annual Energy Use (kWh/yr) for the dryer in a Combination All-in-One Washer-Dryer,Estimated Energy Test Cycle Time (min) for the dryer in a Combination All-in-One Washer-Dryer,Energy Test Cycle Information for the dryer in a Combination All-in-One Washer-Dryer,Calculated Combined Energy Factor - Max Dryness Setting (lbs/kWh) for the dryer in a Combination All-in-One Washer-Dryer,Additional Dryer Features for the dryer in a Combination All-in-One Washer-Dryer,Connected,Connects Using,Communication Standard Application Layer,Direct on-premises Open-standard Based Interconnection,Paired ENERGY STAR Clothes Dryer Available,Paired ENERGY STAR Clothes Dryer ENERGY STAR Model Identifier,Date Available On Market,Date Certified,Markets,ENERGY STAR Model Identifier,Meets ENERGY STAR Most Efficient 2025 Criteria


CREATE TABLE residential_clothes_washers (
    energy_star_unique_id NUMERIC PRIMARY KEY,
    brand_name TEXT,
    model_number TEXT,
    additional_model_information TEXT,
    upc TEXT,
    load_configuration TEXT,
    special_type TEXT,
    additional_washer_features TEXT,
    intended_market TEXT,
    volume_cu_ft NUMERIC,
    height_inches NUMERIC,
    width_inches NUMERIC,
    depth_inches NUMERIC,
    imef NUMERIC,
    imef_us_federal_standard NUMERIC,
    annual_energy_use_kwh NUMERIC,
    iwf NUMERIC,
    iwf_us_federal_standard NUMERIC,
    annual_water_use_gallons NUMERIC,
    drum_capacity_dryer_combo NUMERIC,
    heat_pump_tech_combo_dryer TEXT,
    combined_energy_factor_cef NUMERIC,
    estimated_annual_energy_use_dryer NUMERIC,
    estimated_test_cycle_time_mins NUMERIC,
    energy_test_cycle_info TEXT,
    calculated_cef_max_dryness NUMERIC,
    additional_dryer_features TEXT,
    connected TEXT,
    connects_using TEXT,
    comm_standard_app_layer TEXT,
    open_standard_direct_on_prem TEXT,
    paired_energy_star_dryer_available TEXT,
    paired_dryer_model_identifier TEXT,
    date_available_on_market TEXT,
    date_certified TEXT,
    markets TEXT,
    energy_star_model_identifier TEXT,
    meets_most_efficient_2025 TEXT
);

COPY residential_clothes_washers FROM '/docker-entrypoint-initdb.d/data_csv/ENERGY_STAR_Certified_Residential_Clothes_Washers_20250607.csv'
DELIMITER ','
CSV HEADER;