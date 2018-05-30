package base


var(
	Logger = LoggerSetting()
	Cdb = Lines()
	env = Environment()
	cdbinfo = map[string]map[string]string{
		"prod":{
				"user":"******",
				"password":"*******",
				},
		"dev":{
				"user":"*****",
				"password":"*******",
				},
	}
)
