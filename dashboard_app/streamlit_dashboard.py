
import streamlit as st
import pandas as pd
import plotly.express as px

BASE_URL = "http://localhost:8080/dashboard"

st.title("E-Commerce Sales Dashboard")

def load_csv(endpoint):
    url = f"{BASE_URL}/{endpoint}"
    try:
        return pd.read_csv(url)
    except Exception as e:
        st.error(f"Failed to load {endpoint}: {e}")
        return pd.DataFrame()

# Monthly Revenue Trend
st.header("📈 Monthly Revenue Trend")
monthly_df = load_csv("monthly-revenue")
if not monthly_df.empty:
    fig = px.line(monthly_df, x="month", y="revenue", title="Revenue Over Time")
    st.plotly_chart(fig)

# Top Products
st.header("🏆 Top 10 Products")
top_df = load_csv("top-products")
if not top_df.empty:
    fig = px.bar(top_df.head(10), x="product_name", y="revenue", title="Top Products by Revenue")
    st.plotly_chart(fig)

# Region-wise Sales
st.header("🌍 Sales by Region")
region_df = load_csv("region-sales")
if not region_df.empty:
    fig = px.pie(region_df, names="region", values="revenue", title="Regional Sales Split")
    st.plotly_chart(fig)

# Anomaly Records
st.header("🚨 Anomaly Records")
anomaly_df = load_csv("anomalies")
if not anomaly_df.empty:
    st.dataframe(anomaly_df)

# Category Discount Heatmap
st.header("🔥 Category Discount Heatmap")
heatmap_df = load_csv("category-discount")
if not heatmap_df.empty:
    fig = px.density_heatmap(heatmap_df, x="category", y="avg_discount", z="revenue", histfunc="avg")
    st.plotly_chart(fig)

# Revenue Per Order
st.header("💰 Revenue Per Order")
rpo_df = load_csv("revenue-per-order")
if not rpo_df.empty:
    st.dataframe(rpo_df)

# Repeat Customer Stats
st.header("🔁 New vs. Repeat Customers")
repeat_df = load_csv("repeat-customers")
if not repeat_df.empty:
    fig = px.pie(repeat_df, names="customer_type", values="revenue", title="Revenue Contribution")
    st.plotly_chart(fig)
