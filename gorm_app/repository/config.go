package repository

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     int // TODO type
}
type Config struct {
	DB struct {
		Local      DBConfig
		Production DBConfig
	}
}

func NewConfig() *Config {
	c := new(Config)

	// local環境設定
	c.DB.Local.Host = "127.0.0.1"
	c.DB.Local.Username = "root"
	c.DB.Local.Password = "root"
	c.DB.Local.DBName = "test_database"
	c.DB.Local.Port = 3306
	return c
}
