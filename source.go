package libredis

// 一些预定义的数据源名称
const (
	SourceDefault = ""
	SourceMaster  = "master"
)

// Configuration 数据源配置
type Configuration struct {
	Name     string
	Enabled  bool
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

// Source 数据源
type Source interface {
	Client() Client
}

// SourceFactory 数据源工厂
type SourceFactory interface {
	Open(cfg *Configuration) (Source, error)
}

// SourceRegistration 数据源注册信息
type SourceRegistration struct {
	Enabled bool
	Name    string
	Config  *Configuration
	Factory SourceFactory
	Source  Source
}

// SourceRegistry 数据源注册器
type SourceRegistry interface {
	ListRegistrations() []*SourceRegistration
}

// SourceManager 数据源管理器
type SourceManager interface {
	GetSource(name string) (Source, error)
}
