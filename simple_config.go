package simple_config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//---------------------------------------------------------------------------------------------------------------------

func ReadConfig(filename string) (error) {
	fmt.Println("TODO")

	return nil
}

//---------------------------------------------------------------------------------------------------------------------

func WriteConfig(filename string, config interface{}) (error) {
	c := reflect.ValueOf(config)
	if c.Kind() == reflect.Ptr { c = c.Elem() }
	if c.Type().Kind() != reflect.Struct {
		return errors.New("config parameter is not a struct")
	}

	fileTmp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil { return err }
	defer fileTmp.Close()

	configLines := generateConfigLines ("", c)

	for _, v := range configLines {
		_, err = fileTmp.WriteString (v + "\n")
		if err!= nil {
			fileTmp.Close()
			return err
		}
	}

	return nil
}

//---------------------------------------------------------------------------------------------------------------------

func generateConfigLines (keyprefix string, c reflect.Value) (result []string) {
	var prefixTmp, valueTmp string
	typeOfConfig := c.Type() 

	for i := 0; i < c.NumField(); i++ {
		fieldTmp := c.Field(i)
		prefixTmp = keyprefix + typeOfConfig.Field(i).Name 
		switch fieldTmp.Kind() {
			case reflect.Uint16 : { 
				valueTmp = strconv.FormatUint(fieldTmp.Uint() , 10)
				result = append(result, prefixTmp + " = " + valueTmp ) 
			}
			case reflect.String : { 
				result = append(result, prefixTmp + " = " + fieldTmp.String() ) 
			}
			case reflect.Struct : { 
				result = append(result, generateConfigLines(prefixTmp + ".", fieldTmp)...)
			}
			// default not handled 
		}
	}

	return result
}

//---------------------------------------------------------------------------------------------------------------------





