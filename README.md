GoSimpleConfigLib
=================

development in progress, no stable release yet


Purpose
-------

Allow to export a struct into a file 
Allow to import a file into a struct

notes/restrictions: 
- it works only for the following types : string, int, uint, bool and struct (fields with other types are ignored)
- for import, if configuration structure does not contain a field defined in the config file, the setting is ignored
- utf8 only

go get github.com/thomassilvi/GoSimpleConfigLib


Configuration file syntax
-------------------------

```
# comment
FieldName = value

```

Export overview
---------------



Example:

```

type MyConfigType struct {
	Name	string
	Age	uint8
	Address struct {
		City	string
	}
}

myconfig := MyConfigType { Name: "Roger Rabbit", Age: 20, Address { City: "ToonTown" } }
err := simple_config.WriteConfig("conf/myappname.conf", myconfig)

```

will produce

```

Name = Roger Rabbit
Age = 20
Address.City = ToonTown

```

Import overview
---------------

You must pass a reference to the config in order fields can be set.

Example:

```

type MyConfigType struct {
	Name	string
	Age	uint8
	Address struct {
		City	string
	}
}

myconfig := MyConfigType{}
err := simple_config.ReadConfig("conf/myappname.conf", &myconfig)

```




