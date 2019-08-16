package mysql

import (
	"redcup/spider/package/storage/options"
	"redcup/spider/package/storage/realization/mysql/db"
	"redcup/spider/package/storage/realization/mysql/models"
	"redcup/spider/package/types"
)

type mysql struct {
	responseChan         chan types.Response
	fetchResponseFuncMap map[string]func(response types.Response) error
}

func (s *mysql) AddFetchResponseFunc(requestRule string, fn func(response types.Response) error) {
	s.fetchResponseFuncMap[requestRule] = fn
}

func NewMysql(options *options.StorageOptions) *mysql {
	db.AddTables(models.UserInfo{},models.Work{})
	db.InitEngine(options)
	return &mysql{
		fetchResponseFuncMap: make(map[string]func(response types.Response) error, 0),
	}
}

func (s *mysql) AddResponse(response types.Response) {
	s.responseChan <- response
}

func (s *mysql) Run() {

	s.responseChan = make(chan types.Response, 0)

	go func() {
		for {
			select {
			case response := <-s.responseChan:
				models.NewInsertResult().Insert(response)
			}
		}
	}()
}
