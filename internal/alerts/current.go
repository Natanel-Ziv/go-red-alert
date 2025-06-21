package alerts

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) FetchCurrent(ctx context.Context) (*ActiveAlert, error) {
	// test-mode shortcut
	if c.cfg.TestCurrentMode {
		return &ActiveAlert{
			Type:         AlertTypeMissiles,
			Cities:       []string{"Tel Aviv"},
			Instructions: "היכנסו למרחב מוגן",
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
		return &ActiveAlert{Type: AlertTypeNone}, nil
	}

	// otherwise unmarshal
	var data CurrentAlert
	if err = json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	// Parse the cat-code to int and cast to AlertType
	code, err := strconv.Atoi(data.Cat)
	at := AlertType(code)
	if err != nil {
		at = AlertTypeUnknown
	}

	// Build our public response
	return &ActiveAlert{
		Type:         at,
		Cities:       data.Data,
		Instructions: data.Title,
	}, nil
}
