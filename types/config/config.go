package config

type AppConfigure struct {
	Version   string     `yaml:"Version"`
	Server    *Server    `yaml:"Server"`
	Jwt       *Jwt       `yaml:Jwt`
	Redis     *Redis     `yaml:Redis`
	Database  *Database  `yaml:Database`
	Logger    *Logger    `yaml:Logger`
	Env       string     `yaml:Env`
	MinChat   *MinChat   `yaml:MinChat`
	RabbitMq  *RabbitMq  `yaml:RabbitMq`
	WechatPay *WechatPay `yaml:WechatPay`
}

type Jwt struct {
	JwtSecret           string `yaml:"jwtSecret"`
	TokenExpireDuration int    `yaml:"tokenExpireDuration"`
}

type Server struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Database struct {
	Debug    bool   `yaml:"debug"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Migrate  bool   `yaml:"migrate"`
}
type Logger struct {
	LogPath    string `yaml:"logPath"`
	LogErrPath string `yaml:"logErrPath"`
	Level      string `yaml:"level"`
}

type MinChat struct {
	AppId  string `yaml:"appId"`
	Secret string `yaml:"secret"`
}

type RabbitMq struct {
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	VHost    string `yaml:"vHost"`
}

type WechatPay struct {
	MchID                      string `yaml:"MchID"`
	MchCertificateSerialNumber string `yaml:"MchCertificateSerialNumber"`
	MchPrivateKey              string `yaml:"MchPrivateKey"`
	MchAPIv3Key                string `yaml:"MchAPIv3Key"`
}

const (
	EnvLOCAL  = "local"
	EnvDEV    = "dev"
	EnvONLINE = "online"
)
