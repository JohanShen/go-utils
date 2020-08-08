package kredis

import (
	"github.com/JohanShen/go-utils/logger"
	"github.com/go-redis/redis/v7"
	"time"
)

const (
	//ClusterMode using clusterClient
	ClusterMode string = "cluster"
	//StubMode using reidsClient
	StubMode string = "stub"
)

// Config for kredis, contains RedisStubConfig and RedisClusterConfig
type Config struct {
	// Addrs 实例配置地址
	Addrs []string `json:"addrs" toml:"addrs"`
	// Mode Redis模式 cluster|stub
	Mode string `json:"mode" toml:"mode"`
	// Password 密码
	Password string `json:"password" toml:"password"`
	// DB，默认为0, 一般应用不推荐使用DB分片
	DB int `json:"db" toml:"db"`
	// PoolSize 集群内每个节点的最大连接池限制 默认每个CPU10个连接
	PoolSize int `json:"poolSize" toml:"poolSize"`
	// MaxRetries 网络相关的错误最大重试次数 默认8次
	MaxRetries int `json:"maxRetries" toml:"maxRetries"`
	// MinIdleConns 最小空闲连接数
	MinIdleConns int `json:"minIdleConns" toml:"minIdleConns"`
	// DialTimeout 拨超时时间
	DialTimeout time.Duration `json:"dialTimeout" toml:"dialTimeout"`
	// ReadTimeout 读超时 默认3s
	ReadTimeout time.Duration `json:"readTimeout" toml:"readTimeout"`
	// WriteTimeout 读超时 默认3s
	WriteTimeout time.Duration `json:"writeTimeout" toml:"writeTimeout"`
	// IdleTimeout 连接最大空闲时间，默认60s, 超过该时间，连接会被主动关闭
	IdleTimeout time.Duration `json:"idleTimeout" toml:"idleTimeout"`
	// 慢日志门限值，超过该门限值的请求，将被记录到慢日志中
	SlowThreshold time.Duration `json:"slowThreshold" toml:"slowThreshold"`
	logger        logger.Logger
}

func (self *Config) SetLogger(logger logger.Logger) *Config {
	self.logger = logger
	return self
}

// DefaultRedisConfig default config ...
func DefaultRedisConfig() *Config {
	return &Config{
		DB:            0,
		PoolSize:      10,
		MaxRetries:    3,
		MinIdleConns:  100,
		DialTimeout:   15 * time.Second,
		ReadTimeout:   15 * time.Second,
		WriteTimeout:  15 * time.Second,
		IdleTimeout:   60 * time.Second,
		SlowThreshold: time.Duration(250),
		logger:        logger.DefaultLogger(),
	}
}

// Build ...
func (config *Config) Build() *Redis {
	count := len(config.Addrs)
	if count < 1 {
		logger.Panic(config.logger, "no address in kredis config", logger.ArgAny("config", *config))
	}
	if len(config.Mode) == 0 {
		config.Mode = StubMode
		if count > 1 {
			config.Mode = ClusterMode
		}
	}
	var client redis.Cmdable
	switch config.Mode {
	case ClusterMode:
		if count == 1 {
			logger.Debug(config.logger, "kredis config has only 1 address but with cluster mode")
		}
		client = config.buildCluster()
	case StubMode:
		if count > 1 {
			logger.Debug(config.logger, "kredis config has more than 1 address but with stub mode")
		}
		client = config.buildStub()
	default:
		logger.Debug(config.logger, "kredis mode must be one of (stub, cluster)")
	}

	return &Redis{
		Config:  config,
		Cmdable: client,
	}
}

func (config *Config) buildStub() *redis.Client {
	options := &redis.Options{
		Addr:         config.Addrs[0],
		Password:     config.Password,
		DB:           config.DB,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  config.IdleTimeout,
	}
	stubClient := redis.NewClient(options)

	//fmt.Printf("%+v", options)

	if err := stubClient.Ping().Err(); err != nil {
		logger.ErrorWithArg(config.logger, "dial kredis fail", err, logger.ArgAny("config", *config))
	}

	return stubClient

}

//
func (config *Config) buildCluster() *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.Addrs,
		MaxRedirects: config.MaxRetries,
		Password:     config.Password,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  config.IdleTimeout,
	})
	if err := clusterClient.Ping().Err(); err != nil {
		logger.ErrorWithArg(config.logger, "dial kredis fail", err, logger.ArgAny("config", *config))
	}
	return clusterClient
}
