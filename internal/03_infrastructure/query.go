// internal/03_infrastructure/query.go
package infrastructure

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Queries struct {
    GetRecentActiveCustomers string `json:"get_recent_active_customers"`
}

func LoadQueriesFromFile(path string) (*Queries, error) {
    f, err := os.Open(filepath.Clean(path))
    if err != nil {
        return nil, err
    }
    defer f.Close()
    var q Queries
    err = json.NewDecoder(f).Decode(&q)
    if err != nil {
        return nil, err
    }
    return &q, nil
}
