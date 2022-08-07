package config

// the initialization of the configuration file
func Init() {
	// init viper
	InitViper()
	// init log
	InitLogrus()
	// init mysql connection
	InitMysqlCon()
	// init redis connection
	InitRedisCon()
}
