// colony-metrics/backend/routes/metrics.go
package routes

import (
    "context"
    "colony-metrics/backend/models"
    "colony-metrics/backend/storage"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "strconv"
)

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
    cursor.All(context.TODO(), &results)

    c.JSON(200, results)
}
