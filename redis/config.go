package redis

import (
	"github.com/go-redis/redis/v7"
	"time"
	"utils/logger"
)

const (
	//ClusterMode using clusterClient
	ClusterMode string = "cluster"
	//StubMode using reidsClient
	StubMode string = "stub"
)

// Config for redis, contains RedisStubConfig and RedisClusterConfig
type Config struct {
	// Addrs 实例配置地址
	Addrs []string `json:"addrs" toml:"addrs"`
	// Addr stubConfig 实例配置地址
	Addr string `json:"addr" toml:"addr"`
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
	// Debug开关
	Debug bool `json:"debug" toml:"debug"`
	// ReadOnly 集群模式 在从属节点上启用读模式
	ReadOnly bool `json:"readOnly" toml:"readOnly"`
	// 是否开启链路追踪，开启以后。使用DoCotext的请求会被trace
	EnableTrace bool `json:"enableTrace" toml:"enableTrace"`
	// 慢日志门限值，超过该门限值的请求，将被记录到慢日志中
	SlowThreshold time.Duration `json:"slowThreshold" toml:"slowThreshold"`
	// OnDialError panic|error
	OnDialError string `json:"level"  toml:"onDialError"`
	logger      logger.Logger
}

func (self *Config) SetLogger(logger logger.Logger) {
	self.logger = logger
}

// DefaultRedisConfig default config ...
func DefaultRedisConfig() Config {
	return Config{
		DB:            0,
		PoolSize:      10,
		MaxRetries:    3,
		MinIdleConns:  100,
		DialTimeout:   1 * time.Second,
		ReadTimeout:   1 * time.Second,
		WriteTimeout:  1 * time.Second,
		IdleTimeout:   60 * time.Second,
		ReadOnly:      false,
		Debug:         false,
		EnableTrace:   false,
		SlowThreshold: time.Duration(250),
		OnDialError:   "panic",
		logger:        logger.DefaultLogger(),
	}
}

// Build ...
func (config Config) Build() *Redis {
	count := len(config.Addrs)
	if count < 1 {
		config.logger.Panic("no address in redis config", logger.ArgAny("config", config))
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
			//config.logger.Warn("redis config has only 1 address but with cluster mode")
		}
		client = config.buildCluster()
	case StubMode:
		if count > 1 {
			//config.logger.Warn("redis config has more than 1 address but with stub mode")
		}
		client = config.buildStub()
	default:
		//config.logger.Panic("redis mode must be one of (stub, cluster)")
	}
	return &Redis{
		Config: &config,
		Client: client,
	}
}

func (config Config) buildStub() *redis.Client {
	stubClient := redis.NewClient(&redis.Options{
		Addr:         config.Addrs[0],
		Password:     config.Password,
		DB:           config.DB,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  config.DialTimeout * time.Second * 600,
		ReadTimeout:  config.ReadTimeout * time.Second * 600,
		WriteTimeout: config.WriteTimeout * time.Second * 600,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		IdleTimeout:  config.IdleTimeout * time.Second * 600,
	})

	if err := stubClient.Ping().Err(); err != nil {
		switch config.OnDialError {
		case "panic":
			//config.logger.Panic("dial redis fail", zap.Any("err", err), zap.Any("config", config))
		default:
			//config.logger.Error("dial redis fail", zap.Any("err", err), zap.Any("config", config))
		}
	}

	return stubClient

}

//
func (config Config) buildCluster() *redis.ClusterClient {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.Addrs,
		MaxRedirects: config.MaxRetries,
		ReadOnly:     config.ReadOnly,
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
		switch config.OnDialError {
		case "panic":
			//config.logger.Panic("start cluster redis", zap.Any("err", err))
		default:
			//config.logger.Error("start cluster redis", zap.Any("err", err))
		}
	}
	return clusterClient
}

// StdRedisStubConfig ...
func StdRedisStubConfig(name string) Config {
	return RawRedisStubConfig("maoti.redis." + name + ".stub")
}

// RawRedisStubConfig ...
func RawRedisStubConfig(key string) Config {
	var config = DefaultRedisConfig()
	config.Addrs = []string{config.Addr}
	config.Mode = StubMode
	return config
}

// StdRedisClusterConfig ...
func StdRedisClusterConfig(name string) Config {
	return RawRedisClusterConfig("maoti.redis." + name + ".cluster")
}

// RawRedisClusterConfig ...
func RawRedisClusterConfig(key string) Config {
	var config = DefaultRedisConfig()
	config.Mode = ClusterMode
	return config
}
