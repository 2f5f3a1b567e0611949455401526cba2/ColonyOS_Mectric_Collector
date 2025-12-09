// colony-metrics/backend/storage/metrics.go

package storage

import (
    "context"
    "time"

    "colony-metrics/backend/models"
)

func InsertMetric(m models.Metric) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := MetricsCollection.InsertOne(ctx, m)
    return err
}
