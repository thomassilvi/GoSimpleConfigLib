package simple_config

import (
	"testing"
)

type DatabaseConfig struct {
	Hostname string
	Port     uint16
	Name     string
	Login    string
	Password string
}

type ConfigType1 struct {
	AppName		 string
	Web struct {
		Hostname string
		Port     uint16
	}
	Database DatabaseConfig
}

//-------------------------------------------------------------------------------------------------

func Test_1_ConfigRead(t *testing.T) {
	t.Errorf("TODO")
}

//-------------------------------------------------------------------------------------------------

func Test_1_ConfigWrite(t *testing.T) {
	config1 := ConfigType1{ AppName : "HelloWorld" }
	config1.Web.Hostname = "localhost"
	config1.Web.Port = 9000
	dbconfig := DatabaseConfig{Hostname: "127.0.0.1", Port: 9000, Name: "testdb", 
		Login: "guest", Password: "guest"}
	config1.Database = dbconfig
	err := WriteConfig("tests_output/config1.txt", config1)
	if err != nil {	t.Errorf(err.Error()) }
}

//-------------------------------------------------------------------------------------------------
