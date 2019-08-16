package options

import (
	"redcup/spider/package/types"
)

type StorageOptions struct {
	DataSourceName []string
	ConnInterval   int64
	IsSyncTables   bool
	IsDropTables   bool
}

func NewStorageOptions(baseOptions *types.BaseOptions) *StorageOptions {
	return &StorageOptions{
		DataSourceName: baseOptions.StorageOptions.DataSourceName,
		ConnInterval:   baseOptions.StorageOptions.ConnInterval,
		IsSyncTables:   baseOptions.StorageOptions.IsSyncTables,
		IsDropTables:   baseOptions.StorageOptions.IsDropTables,
	}
}
