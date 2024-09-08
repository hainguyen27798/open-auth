package setting

type Config struct {
	Server ServerSettings `mapstructure:"server"`
	Mysql  MysqlSettings  `mapstructure:"mysql"`
	Logger LoggerSettings `mapstructure:"log"`
	Redis  RedisSettings  `mapstructure:"redis"`
	SMTP   SMTPSettings   `mapstructure:"smtp"`
}

type ServerSettings struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type SMTPSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type MysqlSettings struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Username        string `mapstructure:"user"`
	Password        string `mapstructure:"pass"`
	Database        string `mapstructure:"db_name"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	ConnMaxLifeTime int    `mapstructure:"conn_max_life_time"`
}

type LoggerSettings struct {
	FileName   string `mapstructure:"file_name"`
	Level      string `mapstructure:"log_level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisSettings struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"Password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
