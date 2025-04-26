package alerts

import (
    "bytes"
    "context"
    "encoding/json"
    "io"
    "net/http"
)

func (c *Client) FetchCurrent(ctx context.Context) (*CurrentAlertResponse, error) {
    // test-mode shortcut
    if c.cfg.TestCurrentMode {
        return &CurrentAlertResponse{
            Alert: true,
            Current: &CurrentAlert{
                ID:    "133043790410000000",
                Cat:   "1",
                Title: "ירי רקטות וטילים",
                Data:  []string{c.cfg.TestCurrentLocation},
                Desc:  "היכנסו למרחב המוגן ושהו בו 10 דקות",
            },
        }, nil
    }

    // doGet but reading raw so we can strip BOM/empty
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.cfg.CurrentURL, nil)
    if err != nil {
        return nil, err
    }
    // same headers
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
    // strip UTF-8 BOM if present
    raw = bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))
    // trim whitespace/newlines
    raw = bytes.TrimSpace(raw)

    // empty → no alert
    if len(raw) == 0 {
        return &CurrentAlertResponse{Alert: false}, nil
    }

    // otherwise unmarshal
    var data CurrentAlert
    if err := json.Unmarshal(raw, &data); err != nil {
        return nil, err
    }
    out := &CurrentAlertResponse{Alert: data.ID != ""}
    if out.Alert {
        out.Current = &data
    }
    return out, nil
}

