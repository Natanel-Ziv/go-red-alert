package alerts

import "encoding/json"

// AlertType is the numeric code for each alert.
//
//go:generate stringer -type=AlertType -trimprefix=AlertType
type AlertType int

const (
	AlertTypeNone                          AlertType = 0
	AlertTypeMissiles                      AlertType = 1
	AlertTypeRadiologicalEvent             AlertType = 2
	AlertTypeEarthQuake                    AlertType = 3
	AlertTypeTsunami                       AlertType = 4
	AlertTypeHostileAircraftIntrusion      AlertType = 5
	AlertTypeHazardousMaterials            AlertType = 6
	AlertTypeTerroristInfiltration         AlertType = 7
	AlertTypeMissilesDrill                 AlertType = 11
	AlertTypeEarthQuakeDrill               AlertType = 12
	AlertTypeEnd                           AlertType = 13
	AlertTypeGetReady                      AlertType = 14
	AlertTypeHostileAircraftIntrusionDrill AlertType = 15
	AlertTypeHazardousMaterialsDrill       AlertType = 16
	AlertTypeTerroristInfiltrationDrill    AlertType = 17
	AlertTypeUnknown                       AlertType = -1
)

func (a AlertType) MarshalJSON() ([]byte, error) {
	s := a.String()
	if s == "" {
		s = AlertTypeUnknown.String()
	}
	return json.Marshal(s)
}

// ActiveAlert is the JSON schema returned by GET /current
type ActiveAlert struct {
	Type         AlertType `json:"type"`
	Cities       []string  `json:"cities"`
	Instructions string    `json:"instructions,omitempty"`
}

type CurrentAlert struct {
	ID    string   `json:"id"`
	Cat   string   `json:"cat"`
	Title string   `json:"title"`
	Data  []string `json:"data"`
	Desc  string   `json:"desc"`
}

type CurrentAlertResponse struct {
	Alert   bool          `json:"alert"`
	Current *CurrentAlert `json:"current,omitempty"`
}

type HistoryItem struct {
	AlertDate string `json:"alertDate"`
	Title     string `json:"title"`
	Data      string `json:"data"`
	Category  int    `json:"category"`
}

type HistoryResponse struct {
	History []HistoryItem `json:"history"`
}
