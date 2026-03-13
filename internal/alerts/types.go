package alerts

import "encoding/json"

// AlertType is the numeric code for each alert.
//
//go:generate stringer -type=AlertType -trimprefix=AlertType
type AlertType int

const (
	AlertTypeNone                 AlertType = 0
	AlertTypeMissileAlert         AlertType = 1
	AlertTypeUAV                  AlertType = 2
	AlertTypeNonConventional      AlertType = 3
	AlertTypeWarning              AlertType = 4
	AlertTypeMemorialDay1         AlertType = 5
	AlertTypeHostileAircraft      AlertType = 6
	AlertTypeEarthquakeAlert1     AlertType = 7
	AlertTypeEarthquakeAlert2     AlertType = 8
	AlertTypeCBRNE                AlertType = 9
	AlertTypeTerrorAttack         AlertType = 10
	AlertTypeTsunami              AlertType = 11
	AlertTypeHazmat               AlertType = 12
	AlertTypeUpdate               AlertType = 13
	AlertTypeFlash                AlertType = 14
	AlertTypeMissileAlertDrill    AlertType = 15
	AlertTypeUAVDrill             AlertType = 16
	AlertTypeNonConventionalDrill AlertType = 17
	AlertTypeWarningDrill         AlertType = 18
	AlertTypeMemorialDayDrill1    AlertType = 19
	AlertTypeMemorialDayDrill2    AlertType = 20
	AlertTypeEarthquakeDrill1     AlertType = 21
	AlertTypeEarthquakeDrill2     AlertType = 22
	AlertTypeCBRNEDrill           AlertType = 23
	AlertTypeTerrorAttackDrill    AlertType = 24
	AlertTypeTsunamiDrill         AlertType = 25
	AlertTypeHazmatDrill          AlertType = 26
	AlertTypeUpdateDrill          AlertType = 27
	AlertTypeFlashDrill           AlertType = 28
	AlertTypeUnknown              AlertType = -1
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
