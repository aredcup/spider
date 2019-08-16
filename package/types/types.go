package types

type BaseOptions struct {
	MasterAddr  string         `yaml:"masterAddr"`
	BaseRoot    string         `yaml:"baseRoot"`
	Requests    []*RequestYaml `yaml:"requests"`
	MatchRules  []MatchRule    `yaml:"matchRules"`
	UrlPathRule UrlPathRule    `yaml:"urlPathRule"`

	SchedulerOptions SchedulerConfigOptions `yaml:"schedulerOptions"`
	EngineOptions    EngineConfigOptions    `yaml:"engineOptions"`
	ClientOptions    ClientConfigOptions    `yaml:"clientOptions"`
	StorageOptions   StorageConfigOptions   `yaml:"storageOptions"`
}

type EngineConfigOptions struct {
	MaxRequestCount int    `yaml:"maxRequestCount"`
	ListenAddr      string `yaml:"listenAddr"`
}

type StorageConfigOptions struct {
	DataSourceName []string `yaml:"dataSourceName"`
	ConnInterval   int64    `yaml:"connInterval"`
	IsSyncTables   bool     `yaml:"isSyncTables"`
	IsDropTables   bool     `yaml:"isDropTables"`
}

type ClientConfigOptions struct {
	MasterAddr string `yaml:"masterAddr"`
	LimitTime  uint64 `yaml:"limitTime"`
}

type SchedulerConfigOptions struct {
}

type UrlPathRule struct {
}
