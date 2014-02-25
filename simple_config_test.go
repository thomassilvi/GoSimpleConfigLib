package simple_config

import (
	"testing"
)

// test types definition

type DatabaseConfig struct {
	Hostname string
	Port     uint16
	Name     string
	Login    string
	Password string
}

type ConfigType1 struct {
	AppName string
	Web     struct {
		Hostname string
		Port     uint16
	}
	Database DatabaseConfig
}

type ConfigType2 struct {
	b	bool
	i	int
	i8	int8
	i16	int16
	i32	int32
	i64	int64
	u	uint
	u8	uint8
	u16	uint16
	u32	uint32
	u64	uint64
	s	string
}

//-------------------------------------------------------------------------------------------------

func Test_1_ConfigWrite(t *testing.T) {
	config1 := ConfigType1{AppName: "HelloWorld"}
	config1.Web.Hostname = "localhost"
	config1.Web.Port = 9000
	dbconfig := DatabaseConfig{Hostname: "127.0.0.1", Port: 9000, Name: "testdb",
		Login: "guest", Password: "guest"}
	config1.Database = dbconfig
	err := WriteConfig("tests_output/config1.txt", config1)
	if err != nil {
		t.Errorf(err.Error())
	}
}

//-------------------------------------------------------------------------------------------------

func Test_2_ConfigWrite(t *testing.T) {
	config2 := ConfigType2{}
	config2.b = true
	config2.i = 0
	config2.i8 = 127
	config2.i16 = -32768
	config2.i32 = 2147483647
	config2.i64 = -9223372036854775808
	config2.u = 0
	config2.u8 = 255
	config2.u16 = 65535
	config2.u32 = 4294967295
	config2.u64 = 18446744073709551615
	config2.s = "Hello World !"

	err := WriteConfig("tests_output/config2.txt", config2)
	if err != nil {
		t.Errorf(err.Error())
	}
}

//-------------------------------------------------------------------------------------------------

func Test_1_ConfigRead(t *testing.T) {
	t.Errorf("TODO")
}

//-------------------------------------------------------------------------------------------------


