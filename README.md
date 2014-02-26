GoSimpleConfigLib
=================

development in progress, no stable release yet


Purpose
-------

Allow to export a struct into a file and to import a file into a struct

notes/restrictions: 
- field name in the struct must begin by an uppercase (otherwise the field is not settable)
- it works only for the following types : string, int, uint, bool and struct (fields with other types are ignored)
- for import, if configuration structure does not contain a field defined in the config file, the setting is ignored
- utf8 only

go get github.com/thomassilvi/GoSimpleConfigLib


Configuration file syntax
-------------------------

```
# comment
FieldName = value
FieldName1.FieldName11 = value

```

Export overview
---------------



Example:

```

package main

import (
        "github.com/thomassilvi/GoSimpleConfigLib"
        "log"
)

type MyConfigType struct {
        DebugEnabled    bool
        Log struct {
                File            string
                Verbosity       uint8
        }
}


func main() {
        myconfig := MyConfigType {}
        myconfig.DebugEnabled = true
        myconfig.Log.File = "/var/log/myapp.log"
        myconfig.Log.Verbosity = 3
        err := simple_config.WriteConfig("myappname.conf", myconfig)
        if err != nil {
                log.Println(err)
        }

}

```

will generate the file myappname.conf with the following content

```
DebugEnabled = true
Log.File = /var/log/myapp.log
Log.Verbosity = 3
```

Import overview
---------------

You must pass the reference of the configuration in order fields can be set.

Example:

```
package main

import (
        "github.com/thomassilvi/GoSimpleConfigLib"
        "log"
        "fmt"
)

type MyConfigType struct {
        DebugEnabled    bool
        Log struct {
                File            string
                Verbosity       uint8
        }
}


func main() {
        myconfig := MyConfigType {}
        err := simple_config.ReadConfig("myappname.conf", &myconfig)
        if err != nil {
                log.Println(err)
        }
        fmt.Println(myconfig)
}

```




