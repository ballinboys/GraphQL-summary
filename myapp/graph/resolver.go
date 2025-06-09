package graph

import (
	"database/sql"
	"os"
	"sync"
)

type Resolver struct {
	DB *sql.DB
}

var (
	mlServiceURL string
	once         sync.Once
)

// GetMLServiceURL baca environment variable ML_SERVICE_URL hanya sekali
func GetMLServiceURL() string {
	once.Do(func() {
		mlServiceURL = os.Getenv("ML_SERVICE_URL")
		if mlServiceURL == "" {
			mlServiceURL = "http://localhost:5000/generate" // default fallback
		}
	})
	return mlServiceURL
}
