package repository

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Call struct {
	Origin      string
	Destination string
	Duration    int
	TimeStamp   time.Time
}

type CallRepository struct {
	calls []Call
}

func NewCallRepository(path string) *CallRepository {
	return &CallRepository{
		calls: loadCalls(path),
	}
}

func (c *CallRepository) GetCalls(phoneNumber string) []Call {
	return c.calls
}

// loadCalls loads the calls from a CSV file
// It returns a map with the phone number as key and a slice of calls as value
// If there is an error loading the file, it returns an empty map
func loadCalls(path string) []Call {
	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	if err != nil {
		log.Printf("Error opening file: %v \n", err)
		return []Call{}
	}

	defer file.Close()

	reader := csv.NewReader(file)
	calls := []Call{}

	// Skip header
	reader.Read()
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Printf("Error reading record: %v", err)
			continue
		}

		if record[0] == record[1] {
			continue
		}

		date, err := time.Parse("2006-01-02T15:04:05Z", record[3])
		if err != nil {
			log.Printf("Error converting date: %v", err)
			continue
		}

		duration, err := strconv.Atoi(record[2])
		if err != nil {
			log.Printf("Error converting duration: %v", err)
			continue
		}

		calls = append(calls, Call{
			Origin:      record[0],
			Destination: record[1],
			Duration:    duration,
			TimeStamp:   date,
		})
	}

	return calls
}
