package sap_api_output_formatter

type SalesDistrictReads struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	Product           string `json:"Product"`
	APISchema         string `json:"api_schema"`
	SalesDistrictCode string `json:"sales_district_code"`
	Deleted           string `json:"deleted"`
}

type SalesDistrict struct {
	SalesDistrict string `json:"SalesDistrict"`
	ToText        string `json:"to_Text"`
}

type Text struct {
	SalesDistrict     string `json:"SalesDistrict"`
	Language          string `json:"Language"`
	SalesDistrictName string `json:"SalesDistrictName"`
}

type ToText struct {
	SalesDistrict     string `json:"SalesDistrict"`
	Language          string `json:"Language"`
	SalesDistrictName string `json:"SalesDistrictName"`
}
