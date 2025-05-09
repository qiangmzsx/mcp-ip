package ip

import (
	"encoding/json"
	"strings"
)

type Location struct {
	CountryCode string `json:"country_code,omitempty"`
	Country     string `json:"country,omitempty"`
	Province    string `json:"province,omitempty"`
	City        string `json:"city,omitempty"`
	CityCode    string `json:"city_code,omitempty"`
	ISP         string `json:"isp,omitempty"`
	Latitude    string `json:"latitude,omitempty"`
	Longitude   string `json:"longitude,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}

func (l *Location) String() string {
	data, _ := json.MarshalIndent(l, "", "  ")
	return string(data)
}

func XDB2Location(rawData string) *Location {
	fields := strings.Split(rawData, "|")
	if len(fields) < 5 {
		return &Location{}
	}

	return &Location{
		Country:  strings.TrimSpace(fields[0]),
		Province: strings.TrimSuffix(strings.TrimSpace(fields[2]), "省"),
		City:     strings.TrimSuffix(strings.TrimSpace(fields[3]), "市"),
		ISP:      strings.TrimSpace(fields[4]),
	}
}

type IPReq struct {
	IP string `json:"ip" form:"ip"`
}
