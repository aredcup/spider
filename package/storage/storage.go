package storage

import (
	"redcup/spider/package/storage/options"
	"redcup/spider/package/storage/realization/mysql"
	"redcup/spider/package/types"
)

type Storage interface {
	AddResponse(response types.Response)
	AddFetchResponseFunc(requestRule string, fn func(response types.Response) error)
	Run()
}

func NewStorage(options *options.StorageOptions) Storage {
	return mysql.NewMysql(options)
	//return elasticsearch.NewElastic(options)
}
