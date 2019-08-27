package conf

type Mysql struct {
	DataSource string
	User       string
	Passwd     string
}

type Config struct {
	Mysql *Mysql
}

var (
	confPath string
	Conf     *Config
)
