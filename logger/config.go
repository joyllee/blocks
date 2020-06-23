package logger

var ServerConfig Config

type Config struct {
	Mode   string `default:"release"`
	Port   int32  `default:"62004"`
	Logger struct {
		LogLevel    string `default:"errors"`
		LogDir      string `default:"/opt/log"`
		LogFileName string `default:"demo.log"`
		LogFormat   string `default:"text"` // text or json
	}
}

