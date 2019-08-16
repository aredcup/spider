package options

import (
	"redcup/spider/package/scheduler"
	schedularoptions "redcup/spider/package/scheduler/options"
	"redcup/spider/package/storage"
	storageoptions "redcup/spider/package/storage/options"
	"redcup/spider/package/types"
)

type EngineOptions struct {
	MaxRequestCount int
	ListenAddr      string
	Scheduler       scheduler.Scheduler
	Storage         storage.Storage
}

func NewEngineOptions(baseOptions *types.BaseOptions) *EngineOptions {
	return &EngineOptions{
		MaxRequestCount: baseOptions.EngineOptions.MaxRequestCount,
		ListenAddr:      baseOptions.EngineOptions.ListenAddr,
		Scheduler:       scheduler.NewScheduler(schedularoptions.NewSchedulerOptions(baseOptions)),
		Storage:         storage.NewStorage(storageoptions.NewStorageOptions(baseOptions)),
	}
}
