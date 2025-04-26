package alerts

import (
    "bytes"
    "context"
    "encoding/json"
    "io"
    "net/http"
)

func (c *Client) FetchHistory(ctx context.Context) (*HistoryResponse, error) {
    if c.cfg.TestHistoryMode {
        return &HistoryResponse{History: []HistoryItem{
            {"2024-07-03 17:26:48", "ירי רקטות וטילים", "תל אביב", 1},
            {"2024-07-03 17:26:48", "ירי רקטות וטילים", "ניו יורק", 1},
            {"2024-07-03 17:26:48", "ירי רקטות וטילים", "מוסקבה", 1},
        }}, nil
    }

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cfg.HistoryURL, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "https://www.oref.org.il/")
    req.Header.Set("Referer", "https://www.oref.org.il//12481-he/Pakar.aspx")
    req.Header.Set("X-Requested-With", "XMLHttpRequest")

    resp, err := c.http.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    raw, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    raw = bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))
    raw = bytes.TrimSpace(raw)

    if len(raw) == 0 {
        return &HistoryResponse{History: nil}, nil
    }

    var items []HistoryItem
    if err := json.Unmarshal(raw, &items); err != nil {
        return nil, err
    }
    return &HistoryResponse{History: items}, nil
}

