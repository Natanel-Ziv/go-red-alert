# oref-alerts-go

A lightweight Go implementation of the Oref Alerts Proxy Service originally written in Java. This service fetches current and historical alert data from Oref (Israel Home Front Command) and exposes two HTTP endpoints.

---

## Features

- `/current`: Returns the latest alert status
- `/history`: Returns the alert history
- Configurable via environment variables
- Built-in test/mock modes
- BOM and empty-response handling
- Dockerized for easy deployment
- CI/CD GitHub Actions workflow for building & publishing Docker images

---

## Prerequisites

- Go 1.24+
- Docker (optional, for containerized deployment)
- `make` (optional, for local convenience)

---

## Installation & Local Development

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/oref-alerts-go.git
   cd oref-alerts-go
   ```

2. **Configure environment variables** Create a `.env` file or export directly:

   ```ini
   PORT=9001
   CURRENT_ALERT_MOCK_URL=https://www.oref.org.il/WarningMessages/alert/alerts.json
   HISTORY_MOCK_URL=https://www.oref.org.il/warningMessages/alert/History/AlertsHistory.json
   SSL_INSECURE_SKIP_VERIFY=true
   HTTP_TIMEOUT_SEC=10
   CURRENT_ALERT_TEST_MODE=false
   HISTORY_TEST_MODE=false
   CURRENT_ALERT_TEST_MODE_LOC=Tel+Aviv
   ```

3. **Build & run**

   ```bash
   make run-local
   ```

4. **Test endpoints**

   ```bash
   curl http://localhost:9001/current
   curl http://localhost:9001/history
   ```

---

## Docker

1. **Build**

   ```bash
   make docker-build
   ```

2. **Run**

   ```bash
   make docker-run
   ```

Your service will be available on `http://localhost:9001`.

---

## Credits

- Original Java implementation: [dmatik/oref-alerts-proxy-ms](https://github.com/dmatik/oref-alerts-proxy-ms)
- Go rewrite by Natanel Ziv

---

*Licensed under the MIT License.*

