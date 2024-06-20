package config

type Config struct{

	AppName			string
	HttpPort		string
	DbUser			string
	DbUserPassword	string
	DbHost			string
	DbPort			int
	DbName			string

}

func Load()*Config{

	var Config Config

	Config.AppName="todo_app"
	Config.HttpPort=":8080"

	Config.DbUser="jasur"
	Config.DbUserPassword="1001"
	Config.DbHost="localhost"
	Config.DbPort=5432
	Config.DbName="todo"

	return &Config
}