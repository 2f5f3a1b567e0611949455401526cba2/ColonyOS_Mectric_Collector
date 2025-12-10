// ColonyOS_Metric_Exporter/agent/main.go

// vår agent as of now, sole purpose is to send dummy data to backend

package main

import (
    "bytes"
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "time"
)

type Metric struct {
    Host      string    `json:"host"`
    CPU       float64   `json:"cpu"`
    Memory    float64   `json:"memory"`
    Timestamp time.Time `json:"timestamp"`
}

func main() {
    backendURL := "http://localhost:8080/metrics"

    for {
        m := Metric{
            Host:      "agent-1",
            CPU:       rand.Float64() * 100,
            Memory:    rand.Float64() * 100,
            Timestamp: time.Now(),
        }

        body, _ := json.Marshal(m)

        resp, err := http.Post(backendURL, "application/json", bytes.NewReader(body))
        if err != nil {
            log.Println("error sending metric:", err)
        } else {
            resp.Body.Close()
            log.Println("metric sent:", m.CPU)
        }

        time.Sleep(5 * time.Second)
    }
}
