package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"search-bar-project/domain"
)

func GetConnection() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"elasticsearch:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}


	return es
}

type SearchResponse struct {
	Took	int64
	Hits	struct{
		Total struct{
			Value 	int64
		}
		Hits []*SearchHit
	}
}

type SearchHit struct {
	Score 	float64			`json:"_score"`
	Index	string			`json:"_index"`
	Type	string			`json:"_type"`
	Version	int64			`json:"_version,omitempty"`

	Source domain.Search	`json:"_source"`
}
