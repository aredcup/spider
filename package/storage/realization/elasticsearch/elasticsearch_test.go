package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"redcup/spider/package/storage/options"
	"redcup/spider/package/types"
	"strings"
	"testing"
)

func TestNewElastic(t *testing.T) {
	elastic := NewElastic(options.NewStorageOptions(&types.BaseOptions{
		StorageOptions: types.StorageConfigOptions{
			DataSourceName: []string{"http://192.168.99.100:9200/"},
		},
	}))

	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: "1",
		Body:       strings.NewReader(`{"title" : "` + "dsadas" + `"}`),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), elastic.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
