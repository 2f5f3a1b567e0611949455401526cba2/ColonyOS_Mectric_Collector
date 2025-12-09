// colony-metrics/backend/main.go
package main

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "colony-metrics/backend/models"
    "colony-metrics/backend/storage"
)

func main() {
    r := gin.Default()

    r.POST("/metrics", func(c *gin.Context) {
        var m models.Metric
        if err := c.ShouldBindJSON(&m); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        log.Printf("Received metric: %+v\n", m)
        if err := storage.InitMongo("mongodb://localhost:27017", "colony_metrics", "metrics"); err != nil {
            log.Fatal("mongo init failed:", err) } 
        c.Status(http.StatusCreated)
    })

    // temp health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    if err := r.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
