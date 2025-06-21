package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/natanel-ziv/oref-alerts-go/internal/alerts"
    "github.com/natanel-ziv/oref-alerts-go/pkg/config"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("config: %v", err)
    }

    client := alerts.NewClient(cfg)
    mux := http.NewServeMux()
    mux.HandleFunc("/current", withJSON(currentHandler(client)))
    mux.HandleFunc("/history", withJSON(historyHandler(client)))

    srv := &http.Server{
        Addr:         ":" + cfg.Port,
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Printf("🚀 listening on %s", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}

func withJSON(h func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        h(w, r)
    }
}

func currentHandler(c *alerts.Client) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        alert, err := c.FetchCurrent(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadGateway)
            return
        }
        json.NewEncoder(w).Encode(alert)
    }
}

func historyHandler(c *alerts.Client) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        resp, err := c.FetchHistory(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadGateway)
            return
        }
        json.NewEncoder(w).Encode(resp)
    }
}
