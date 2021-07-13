package repositories

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"search-bar-project/adapters/config"
	"search-bar-project/domain"
	"strings"
)

func Save(search domain.Search) (domain.Search, error) {

	elasticsearch := config.GetConnection()

	body, err := json.Marshal(search)
	if err != nil {
		log.Fatalf("Error parsing the registry: %s", err)
		return domain.Search{}, err
	}

	req := esapi.IndexRequest{
		Index:      "search_index",
		DocumentID: search.ID,
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), elasticsearch)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%v", res.Status(), search.ID)
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return search, nil

}
