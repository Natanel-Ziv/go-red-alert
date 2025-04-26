package alerts

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

