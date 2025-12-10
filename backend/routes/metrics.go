// Package routes ColonyOS_Metric_Collector/backend/routes/metrics.go
package routes

import (
	"ColonyOS_Metric_Collector/backend/models"
	"ColonyOS_Metric_Collector/backend/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// register all metric routes
func RegisterMetricRoutes(r *gin.Engine) {
	r.GET("/metrics", MetricsGET)
	r.POST("/metrics", MetricsPOST)
}

func MetricsPOST(c *gin.Context) {
	var m models.Metric
	log.Printf("Received metric: %+v\n", m)

	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{"error: ": err.Error()})
		return
	}
	if err := storage.InsertMetric(m); err != nil {
		c.JSON(500, gin.H{"error on failed to insert metric: ": err.Error()})
		return
	}
	log.Printf("Received metric: %+v\n", m)

	c.Status(201)
}

// get metrics, ie query it
func MetricsGET(c *gin.Context) {
	host := c.Query("host")
	limitStr := c.DefaultQuery("limit", "20")

	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	filter := bson.M{}
	if host != "" {
		filter["host"] = host
	}

	opts := options.Find().
		SetSort(bson.M{"timestamp": -1}).
		SetLimit(limit)

	cursor, err := storage.MetricsCollection.Find(
		context.TODO(),
		filter,
		opts,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "db query failed"})
		return
	}

	var results []models.Metric
	if err := cursor.All(context.TODO(), &results); err != nil {
		c.JSON(500, gin.H{"error": "cursor decode failed"})
		return
	}
	c.JSON(200, results)
}
