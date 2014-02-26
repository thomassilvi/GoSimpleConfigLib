/*
GoSimpleConfigLib allow you to export/import settings to/from a config file.
Copyright (C) 2014 Thomas Silvi

This file is part of the lib GoSimpleConfigLib.

GoSimpleConfigLib is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

GoSimpleConfigLib is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar. If not, see <http://www.gnu.org/licenses/>.
*/

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

type ConfigAllType struct {
        S    string
        I    int
        B    bool
        I8   int8
        I16  int16
        I32  int32
        I64  int64
        Ui   uint
        Ui8  uint8
        Ui16 uint16
        Ui32 uint32
        Ui64 uint64
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

func Test_1_ConfigRead(t *testing.T) {

	expectedConfig := ConfigType1{AppName: "MyWebApp"}
	expectedConfig.Web.Hostname = "127.0.0.1"
	expectedConfig.Web.Port = 8000
	expectedConfig.Database.Hostname = "127.0.0.1"
	expectedConfig.Database.Name = "mywebapp_dev_db"
	expectedConfig.Database.Login = "guest"
	expectedConfig.Database.Password = "test"
	expectedConfig.Database.Port = 6000

	config := ConfigType1{}
	err := ReadConfig("tests_input/config1.txt", &config)
	if err != nil {
		t.Errorf(err.Error())
	}
	if config != expectedConfig {
		t.Errorf("config is :\n", config, "\nand not\n", expectedConfig)
	}
}

//-------------------------------------------------------------------------------------------------

func Test_1_ExportImportConfigType1 (t *testing.T) {

	filename := "tests_output/configtype1.conf"

	expectedConfig := ConfigType1{AppName: "TestApp"}
	expectedConfig.Web.Hostname = "www.test.org"
	expectedConfig.Web.Port = 8765
	dbconfig := DatabaseConfig{Hostname: "localhost", Port: 9000, Name: "testdb",
		Login: "guest", Password: "guest"}
	expectedConfig.Database = dbconfig
	err := WriteConfig(filename, expectedConfig)
	if err != nil {
		t.Errorf(err.Error())
	}

	config := ConfigType1{}
	err = ReadConfig(filename, &config)
	if err != nil {
		t.Errorf(err.Error())
	}

	if config != expectedConfig {
		t.Errorf("config is :\n", config, "\nand not\n", expectedConfig)
	}
}

//-------------------------------------------------------------------------------------------------

func Test_2_ExportImport (t *testing.T) {

        filename := "tests_output/config2.txt"

        expectedConfig := ConfigAllType{}

        expectedConfig.S = "Hello World !"
        expectedConfig.B = true
        expectedConfig.I = 0
        expectedConfig.I8 = 127
        expectedConfig.I16 = -32768
        expectedConfig.I32 = 2147483647
        expectedConfig.I64 = -9223372036854775808
        expectedConfig.Ui = 0
        expectedConfig.Ui8 = 255
        expectedConfig.Ui16 = 65535
        expectedConfig.Ui32 = 4294967295
        expectedConfig.Ui64 = 18446744073709551615

        err := WriteConfig(filename, expectedConfig)
        if err != nil {
                t.Errorf(err.Error())
        }

        config := ConfigAllType{}
        err = ReadConfig(filename, &config)
        if err != nil {
                t.Errorf(err.Error())
        }

        if config != expectedConfig {
                t.Errorf("config is :\n", config, "\nand not\n", expectedConfig)
        }

}

//-------------------------------------------------------------------------------------------------

