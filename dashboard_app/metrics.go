
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "io/ioutil"
)

func setupRoutes(router *gin.Engine) {
    // Monthly Revenue Trend
    router.GET("/dashboard/monthly-revenue", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/monthly_sales_summary.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read monthly_sales_summary.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Top 10 Products
    router.GET("/dashboard/top-products", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/top_products.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read top_products.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Sales by Region
    router.GET("/dashboard/region-sales", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/region_wise_performance.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read region_wise_performance.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Anomaly Records
    router.GET("/dashboard/anomalies", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/anomaly_records.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read anomaly_records.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Category Discount Heatmap
    router.GET("/dashboard/category-discount", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/category_discount_map.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read category_discount_map.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Revenue per Order
    router.GET("/dashboard/revenue-per-order", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/revenue_per_order.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read revenue_per_order.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })

    // Repeat Customer Rate
    router.GET("/dashboard/repeat-customers", func(c *gin.Context) {
        data, err := ioutil.ReadFile("output/repeat_customer_stats.csv")
        if err != nil {
            c.String(http.StatusInternalServerError, "Cannot read repeat_customer_stats.csv")
            return
        }
        c.Data(http.StatusOK, "text/csv", data)
    })
}
