
# E-Commerce Analytics Pipeline in Go

A scalable Go-based pipeline that ingests, cleans, transforms, and exposes sales analytics via APIs.

---

## Features
- Chunked CSV ingestion (100M+ rows supported)
- Row-wise cleaning and validation
- Business metric generation (monthly sales, top products, anomalies)
- REST API with Gin
- Streamlit frontend dashboard

---

## 🚀 Setup Instructions

### Install dependencies

```bash
go mod tidy
go mod vendor
pip install streamlit pandas plotly
```

### Run the backend server

```bash
go run dashboard_app/main.go
```

---

## Run the frontend dashboard

```bash
streamlit run dashboard_app/streamlit_dashboard.py
```

---

## File Upload API

```bash
curl -X POST http://localhost:8080/upload -F "file=@data/raw_sales.csv"
```

---

## Dashboard Endpoints (served from preprocessed files)

```bash
curl http://localhost:8080/dashboard/top-products
curl http://localhost:8080/dashboard/monthly-revenue
curl http://localhost:8080/dashboard/region-sales
curl http://localhost:8080/dashboard/anomalies
curl http://localhost:8080/dashboard/category-discount
curl http://localhost:8080/dashboard/revenue-per-order
curl http://localhost:8080/dashboard/repeat-customers
```

---

## 📁 Folder Structure

- `pipeline/` - ingestion, cleaning, transformation logic
- `dashboard_app/` - Gin server and Streamlit UI
- `config/constants.go` - all path constants
- `data/raw_sales.csv` - input file (replaceable) -- Replace with raw_sales_2.csv if needs to be tested fro 5k rows
- `output/` - analytics output used by dashboard

---

## 🧪 Example Output Files

- `output/top_products.csv`
- `output/monthly_sales_summary.csv`
- `output/region_wise_performance.csv`
- `output/anomaly_records.csv`
