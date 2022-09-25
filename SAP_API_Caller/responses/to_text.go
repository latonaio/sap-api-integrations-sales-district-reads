package responses

type ToText struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesDistrict     string `json:"SalesDistrict"`
			Language          string `json:"Language"`
			SalesDistrictName string `json:"SalesDistrictName"`
		} `json:"results"`
	} `json:"d"`
}
