package elasticsearch

import (
	"github.com/elastic/go-elasticsearch"
	"redcup/spider/package/storage/options"
	"redcup/spider/package/types"
)

type elasticSearch struct {
	client               *elasticsearch.Client
	responseChan         chan types.Response
	fetchResponseFuncMap map[string]func(response types.Response) error
}

func (s *elasticSearch) AddFetchResponseFunc(requestRule string, fn func(response types.Response) error) {
	s.fetchResponseFuncMap[requestRule] = fn
}

func NewElastic(options *options.StorageOptions) *elasticSearch {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: options.DataSourceName,
	})
	if err != nil {
		panic(err)
	}

	return &elasticSearch{
		client:               client,
		fetchResponseFuncMap: make(map[string]func(response types.Response) error, 0),
	}
}

func (s *elasticSearch) AddResponse(response types.Response) {
	s.responseChan <- response
}

func (s *elasticSearch) Run() {
	go func() {
		for {
			select {
			case response := <-s.responseChan:
				s.save(response)
			}
		}
	}()
}

func (s *elasticSearch) save(response types.Response) {
	//_, err := s.client.Index().Index("spider").Type("test").Do(context.Background())
	//if err != nil {
	//	glog.Error(err)
	//}
}
