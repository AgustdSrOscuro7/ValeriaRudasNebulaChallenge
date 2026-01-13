# ValeriaRudasNebulaChallenge

# SSL Labs TLS Checker – Nebula Challenge

## 📌 Overview

This project is a simple Go program that uses the SSL Labs API to analyze the TLS/SSL security configuration of a given domain.

The application sends a request to the SSL Labs analysis endpoint, polls the API until the scan is completed, and then prints a summary of the results, including the TLS grade for each detected endpoint.

This project was developed as part of the Nebula Challenge.

## 🛠️ Technologies Used

- **Go (Golang)**
- **SSL Labs API v2**
- **Standard Go libraries:**
  - `net/http`
  - `encoding/json`
  - `time`
  - `os`

## ⚙️ How It Works

1. The program receives a domain name as a command-line argument.
2. It starts a new SSL Labs assessment for the domain.
3. Since the SSL Labs API is asynchronous, the program periodically checks the scan status.
4. While the scan is running, the status is shown as `IN_PROGRESS`.
5. Once the status is `READY`, the program prints:
   - The IP address of each endpoint
   - The endpoint status
   - The TLS grade (if available)
6. If the API returns an error (for example, a blacklisted hostname), the program reports it and exits gracefully.

## ▶️ How to Run

### 1️⃣ Prerequisites

- Go 1.20 or later installed
- Verify installation:
  ```bash
  go version
  ```

### 2️⃣ Run the program

```bash
go run main.go <domain>
```

**Example:**
```bash
go run main.go neverssl.com
```

## 🧪 Recommended Test Domains

Some domains are blacklisted or take a long time to scan due to high traffic or complex infrastructure.

Recommended domains for testing:
- `neverssl.com`
- `httpbin.org`
- `badssl.com`

## ⚠️ Notes About the SSL Labs API

The API is rate-limited and asynchronous:
- Large or popular domains (e.g., Google, Facebook) may take several minutes
- Some domains (such as `example.com`) are blacklisted and will return an error
- Polling too frequently may slow down the analysis

The program handles these cases by:
- Waiting between requests
- Checking the returned status
- Stopping execution when an error occurs

## 📌 Sample Output

```
Starting SSL Labs analysis for: neverssl.com
Current status: IN_PROGRESS
Current status: READY
TLS Analysis completed for: neverssl.com
--------------------------------------------------
IP Address: 34.223.124.45
Status: Ready
Grade: A
--------------------------------------------------
IP Address: 2600:1f13:37c:1400:ba21:7165:5fc7:736e
Status: Ready
Grade: A
--------------------------------------------------
```

## 📈 Possible Improvements

- Add a configurable timeout to avoid long waits
- Export results in JSON or CSV format
- Improve error handling for specific HTTP status codes
- Add unit tests for response parsing

## 👩‍💻 Author

#### Valeria Rudas
#### Nebula Challenge – Truora