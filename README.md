GoSimpleConfigLib
=================

development in progress, no stable release yet


Purpose
-------

Allow to export a struct into a file 
Allow to import a file into a struct

notes: 
- it works only for the following types : string, int, uint, bool and struct
- fields with other types are ignored

go get github.com/thomassilvi/GoSimpleConfigLib

Export overview
---------------

type MyConfigType struct {
	Name	string
	Age	uint8
	Address struct {
		City	string
	}
}

myconfig := MyConfigType { Name: "Roger Rabbit", Age: 20, Address { City: "ToonTown" } }
err := WriteConfig("conf/myappname.conf", myconfig)

will produce

Name = Roger Rabbit
Age = 20
Address.City = ToonTown


Import overview
---------------

 TODO





