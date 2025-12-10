// ColonyOS_Metric_Collector/backend/storage/metrics.go

package storage

import (
	"context"
	"time"

	"ColonyOS_Metric_Collector/backend/models"
)

func InsertMetric(m models.Metric) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := MetricsCollection.InsertOne(ctx, m)
	return err
}
