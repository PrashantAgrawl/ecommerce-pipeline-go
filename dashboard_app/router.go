package main

import (
	"ecommerce_pipeline_go/config"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter1() *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Static endpoints to serve preprocessed CSVs
	endpoints := map[string]string{
		"/dashboard/top-products":      config.OutputTopProducts,
		"/dashboard/monthly-revenue":   config.OutputMonthlySummary,
		"/dashboard/region-sales":      config.OutputRegionPerformance,
		"/dashboard/anomalies":         config.OutputAnomalyRecords,
		"/dashboard/category-discount": config.OutputCategoryDiscountMap,
		"/dashboard/revenue-per-order": config.OutputRevenuePerOrder,
		"/dashboard/repeat-customers":  config.OutputRepeatCustomerStats,
	}

	for route, path := range endpoints {
		p := path
		router.GET(route, func(c *gin.Context) {
			data, err := ioutil.ReadFile(p)
			if err != nil {
				c.String(http.StatusInternalServerError, "Failed to read: %s", p)
				return
			}
			c.Data(http.StatusOK, "text/csv", data)
		})
	}

	return router
}
