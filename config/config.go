package config

type Configuration interface {
	GetString(name string) (configVal string, found bool)
	GetInt(name string) (configVal int, found bool)
	GetBool(name string) (configVal bool, found bool)
	GetFloat(name string) (configVal float64, found bool)

	GetStringDefault(name string, defVal string) (configVal string)
	GetIntDefault(name string, defVal int) (configVal int)
	GetBoolDefault(name string, defVal bool) (configVal bool)
	GetFloatDefault(name string, defVal float64) (configVal float64)

	GetSection(sectionName string) (section Configuration, found bool)
}
