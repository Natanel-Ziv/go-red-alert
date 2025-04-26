package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
    Port                 string
    CurrentURL           string
    HistoryURL           string
    InsecureSkipVerify   bool
    Timeout              time.Duration
    TestCurrentMode      bool
    TestHistoryMode      bool
    TestCurrentLocation  string
}

func Load() (*Config, error) {
    c := &Config{
        Port:               getEnv("PORT", "9001"),
        CurrentURL:         getEnv("CURRENT_ALERT_MOCK_URL", "https://www.oref.org.il/WarningMessages/alert/alerts.json"),
        HistoryURL:         getEnv("HISTORY_MOCK_URL", "https://www.oref.org.il/warningMessages/alert/History/AlertsHistory.json"),
        InsecureSkipVerify: getEnvBool("SSL_INSECURE_SKIP_VERIFY", true),
        Timeout:            time.Duration(getEnvInt("HTTP_TIMEOUT_SEC", 10)) * time.Second,
        TestCurrentMode:    getEnvBool("CURRENT_ALERT_TEST_MODE", false),
        TestHistoryMode:    getEnvBool("HISTORY_TEST_MODE", false),
        TestCurrentLocation: os.Getenv("CURRENT_ALERT_TEST_MODE_LOC"),
    }
    return c, nil
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

func getEnvBool(key string, def bool) bool {
    if v := os.Getenv(key); v != "" {
        b, err := strconv.ParseBool(v)
        if err == nil {
            return b
        }
    }
    return def
}

func getEnvInt(key string, def int) int {
    if v := os.Getenv(key); v != "" {
        if i, err := strconv.Atoi(v); err == nil {
            return i
        }
    }
    return def
}

