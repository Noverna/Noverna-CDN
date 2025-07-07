package config

type Config struct {
	Server   Server   `toml:"server"`
	CDN      CDN      `toml:"cdn"`
	Cache    Cache    `toml:"cache"`
	Premium  Premium  `toml:"premium"`
	Advanced Advanced `toml:"advanced"`
	Debug    Debug    `toml:"debug"`
}

type Server struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	LogLevel string `toml:"log_level"`
	DataDir  string `toml:"data_dir"`
	TempDir  string `toml:"temp_dir"`
}

type CDN struct {
	Enabled     bool    `toml:"enabled"`
	StaticServe string  `toml:"static_serve"`
	StaticDir   string  `toml:"static_dir"`
	UrlPrefix   string  `toml:"url_prefix"`
	Cors        CDNCors `toml:"cdn.cors"`
}

type CDNCors struct {
	Enabled        bool     `toml:"enabled"`
	AllowedOrigin  string   `toml:"allowed_origin"`
	AllowedMethods []string `toml:"allowed_methods"`
}

type Cache struct {
	Enabled      bool   `toml:"enabled"`
	Strategy     string `toml:"strategy"`
	MaxMemory    int    `toml:"max_memory_mb"`
	ExpiresAfter int    `toml:"expires_after_seconds"`
}

type Premium struct {
	Enabled     bool     `toml:"enabled"`
	MultiRegion bool     `toml:"multi_region"`
	EdgeNodes   []string `toml:"edge_nodes"`
	Analytics   bool     `toml:"analytics_enabled"`
	ApiKey      string   `toml:"api_key"`
}

type Advanced struct {
	GZip  bool `toml:"gzip_compression"`
	ETag  bool `toml:"etag_support"`
	Http2 bool `toml:"http2"`
}

type Debug struct {
	Enabled bool `toml:"enabled"`
}