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
	"bufio"
	"errors"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

//#################################################################################################
// definition of error with comment
//#################################################################################################

type Error interface {
	error
	Comment() string
}

type SimpleConfigError struct {
	err     string
	comment string
}

//-------------------------------------------------------------------------------------------------

func (e *SimpleConfigError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.err
}

//-------------------------------------------------------------------------------------------------

func (e *SimpleConfigError) Comment() string {
	if e == nil {
		return "<nil>"
	}
	return e.comment
}

//-------------------------------------------------------------------------------------------------

func toError(e error) Error {
	return &SimpleConfigError{e.Error(), ""}
}

//-------------------------------------------------------------------------------------------------

var (
	ErrArgNotStructOrPtrStruct = errors.New("config parameter is not a struct or a pointer on a struct")
	ErrArgNotPtrOnStruct       = errors.New("config parameter is not a pointer on a struct")
	ErrParse                   = &SimpleConfigError{"parse error", ""}
	ErrFieldNotFound           = &SimpleConfigError{"field not found", ""}
	ErrNotStructField          = &SimpleConfigError{"field is not a struct", ""}
	ErrFieldNotSettable        = &SimpleConfigError{"field is not settable", ""}
)

//#################################################################################################
// Config writer & reader
//#################################################################################################

func WriteConfig(filename string, config interface{}) error {
	c := reflect.ValueOf(config)
	if c.Kind() == reflect.Ptr {
		c = c.Elem()
	}
	if c.Type().Kind() != reflect.Struct {
		return ErrArgNotStructOrPtrStruct
	}

	fileTmp, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fileTmp.Close()

	configLines := generateConfigLines("", c)

	for _, v := range configLines {
		_, err = fileTmp.WriteString(v + "\n")
		if err != nil {
			fileTmp.Close()
			return err
		}
	}

	return nil
}

//-------------------------------------------------------------------------------------------------

func generateConfigLines(keyprefix string, c reflect.Value) (result []string) {
	var prefixTmp, valueTmp string
	typeOfConfig := c.Type()

	for i := 0; i < c.NumField(); i++ {
		fieldTmp := c.Field(i)
		prefixTmp = keyprefix + typeOfConfig.Field(i).Name
		k := fieldTmp.Kind()

		switch {
		case k == reflect.Bool:
			{
				valueTmp = strconv.FormatBool(fieldTmp.Bool())
				result = append(result, prefixTmp+" = "+valueTmp)
			}
		case k >= reflect.Int && k <= reflect.Int64:
			{
				valueTmp = strconv.FormatInt(fieldTmp.Int(), 10)
				result = append(result, prefixTmp+" = "+valueTmp)
			}
		case k >= reflect.Uint && k <= reflect.Uint64:
			{
				valueTmp = strconv.FormatUint(fieldTmp.Uint(), 10)
				result = append(result, prefixTmp+" = "+valueTmp)
			}
		case k == reflect.String:
			{
				result = append(result, prefixTmp+" = "+fieldTmp.String())
			}
		case k == reflect.Struct:
			{
				result = append(result, generateConfigLines(prefixTmp+".", fieldTmp)...)
			}
			// default not handled
		}

	}

	return result
}

//-------------------------------------------------------------------------------------------------

func getKeyValue(line string) (found bool, key, value string) {
	if line[0] == '#' {
		return false, "", ""
	}
	indexEqual := strings.Index(line, "=")
	if indexEqual == -1 {
		return false, "", ""
	}
	keyTmp := line[:indexEqual-1]
	if keyTmp[0] == '#' {
		return false, "", ""
	}
	valueTmp := line[indexEqual+1 : len(line)-1]
	keyTmp = strings.Trim(keyTmp, " ")
	valueTmp = strings.Trim(valueTmp, " ")
	return true, keyTmp, valueTmp
}

//-------------------------------------------------------------------------------------------------

func getParseError(k reflect.Kind, line string) Error {
	ErrParse.comment = "Parse error for " + k.String() + " with line:" + line
	return ErrParse
}

//-------------------------------------------------------------------------------------------------

func getBitSizeFromKind(k reflect.Kind) int {
	result := 0

	switch k {
	case reflect.Int:
		{
			result = 0
		}
	case reflect.Int8:
		{
			result = 8
		}
	case reflect.Int16:
		{
			result = 16
		}
	case reflect.Int32:
		{
			result = 32
		}
	case reflect.Int64:
		{
			result = 64
		}
	case reflect.Uint:
		{
			result = 0
		}
	case reflect.Uint8:
		{
			result = 8
		}
	case reflect.Uint16:
		{
			result = 16
		}
	case reflect.Uint32:
		{
			result = 32
		}
	case reflect.Uint64:
		{
			result = 64
		}
	}

	return result
}

//-------------------------------------------------------------------------------------------------

func ReadConfig(filename string, configuration interface{}) error {
	readConfigError := new(SimpleConfigError)

	if reflect.ValueOf(configuration).Kind() != reflect.Ptr {
		return ErrArgNotPtrOnStruct
	}
	s := reflect.ValueOf(configuration).Elem()
	if s.Kind() != reflect.Struct {
		return ErrArgNotPtrOnStruct
	}

	fileTmp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileTmp.Close()

	var lineTmp, keyTmp, keyPartTmp, valueTmp string
	var isConfig, fieldNotFound bool
	var keysTmp []string
	var fieldTmp reflect.Value

	readerTmp := bufio.NewReader(fileTmp)
	lineTmp, err = readerTmp.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}
	for err == nil {

		isConfig, keyTmp, valueTmp = getKeyValue(lineTmp)
		if isConfig {
			keysTmp = strings.Split(keyTmp, ".")
			fieldNotFound = true
			fieldTmp = s
			for i := 0; i < len(keysTmp) && fieldNotFound; i++ {
				keyPartTmp = keysTmp[i]
				if keyPartTmp[0] != strings.ToUpper(keyPartTmp)[0] {
					readConfigError = ErrFieldNotSettable
					readConfigError.comment = "Field name >" + keyPartTmp + "< is not settable as first letter is not uppercase"
					return readConfigError
				}

				if fieldTmp.Kind() == reflect.Struct {
					fieldTmp = fieldTmp.FieldByName(keysTmp[i])
					if i == len(keysTmp)-1 {
						fieldNotFound = !fieldTmp.IsValid()
					}
				} else {
					readConfigError = ErrNotStructField
					readConfigError.comment = keyPartTmp + " does not correspond to a struct field"
					return readConfigError
				}
			}

			if fieldNotFound {
				readConfigError = ErrFieldNotFound
				readConfigError.comment = "Field >" + keyTmp + "< can not be found"
				return readConfigError
			}

			if !fieldTmp.CanSet() {
				readConfigError = ErrFieldNotSettable
				readConfigError.comment = "Field >" + keyTmp + "< is not settable"
				return readConfigError
			}

			fieldKindTmp := fieldTmp.Kind()
			switch {

			case fieldKindTmp == reflect.String:
				{
					fieldTmp.SetString(valueTmp)
				}
			case fieldKindTmp == reflect.Bool:
				{
					boolTmp, err := strconv.ParseBool(valueTmp)
					if err != nil {
						return getParseError(fieldKindTmp, lineTmp)
					}
					fieldTmp.SetBool(boolTmp)
				}
			case fieldKindTmp >= reflect.Int && fieldKindTmp <= reflect.Int64:
				{
					bitSize := getBitSizeFromKind(fieldKindTmp)
					intTmp, err := strconv.ParseInt(valueTmp, 10, bitSize)
					if err != nil {
						return getParseError(fieldKindTmp, lineTmp)
					}
					fieldTmp.SetInt(intTmp)
				}
			case fieldKindTmp >= reflect.Uint && fieldKindTmp <= reflect.Uint64:
				{
					bitSize := getBitSizeFromKind(fieldKindTmp)
					uintTmp, err := strconv.ParseUint(valueTmp, 10, bitSize)
					if err != nil {
						return getParseError(fieldKindTmp, lineTmp)
					}
					fieldTmp.SetUint(uintTmp)
				}
			}
		}

		// next line
		lineTmp, err = readerTmp.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
	}

	return nil
}

//-------------------------------------------------------------------------------------------------
