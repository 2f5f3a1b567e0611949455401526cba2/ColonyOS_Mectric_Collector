// colony-metrics/backend/models/metric.go
package models

import "time"

type Metric struct {
    Host      string    `json:"host" bson:"host"`
    CPU       float64   `json:"cpu" bson:"cpu"`
    Memory    float64   `json:"memory" bson:"memory"`
    Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
