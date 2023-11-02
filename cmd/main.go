package main

import (
	"encoding/json"
	"fmt"
	//REMOVE Dot Import after thini=king of an appropriate name for the models package (Discouraged in GOlang)
	. "github.com/shubhexists/go-json-db/models" 
)
//EVERY THING HERE SHOULD BE IN MODELS FOR THE USER TO CREATE THESE DYNAMICALLY, THESE ARE JUST THE EXAMPLES.
//ALSO THINK OF A BASIC WAY TO ALLOW USERS TO CREATE THESE STRUCTS DYNAMICALLY
//THIS MAY BE ADDED INTO THE EXAMPLES FOLDER (CREATE LATER MAYBE)
type User struct{
	//@todo Add implementation for custom tags
	Name string       `json:"name" db:"main"` //Change this custom tag name maybe?
	Age json.Number   `json:"age"`
	Contact string    `json:"contact"`
	Company string    `json:"company"`
	Address Address   `json:"address"`
}

type Address struct{
	City string		     `json:"city"`
	State string         `json:"state"`
	Country string       `json:"country"`
	Pincode json.Number  `json:"pincode"`
}

func main(){
	//All the collections would be in the /database directory
	dir := "./database"
	db, err := New(dir, nil)
	if err != nil{
		fmt.Println(err)
		return 
	}
	//Sample Lists 
	//We would have a test directory for testing and CI/CD Later..
	employees := []User{
		{"John4","23","9585394030","Humanize",Address{
			"Delhi",
			"Delhi",
			"India",
			"110092",
		}},
		{"John10","23","9585394030","Humanize",Address{
			"Delhi",
			"Delhi",
			"India",
			"110092",
		}},
		{"John20","23","9585394030","Humanize",Address{
			"Delhi",
			"Delhi",
			"India",
			"110092",
		}},
		{"John30","23","9585394030","Humanize",Address{
			"Delhi",
			"Delhi",
			"India",
			"110092",
		}},
	}

	//Writing into the database Example
	for _,items := range employees {
		db.Write("users", items.Name, User{
			Name: items.Name,
			Age: items.Age,
			Contact: items.Contact,
			Company: items.Company,
			Address: items.Address,
		})
	}

	//Read All Data in a Collection
	records, err := db.ReadAll("users");
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	//Update Complete Record Example
	db.UpdateRecord("users", "John4", User{
		Name: "Shubham",
		Age: "18",
		Contact: "9585394030",
		Company: "Humanize",
		Address: Address{
			City: "Delhi",
			State: "Delhi",
			Country: "India",
			Pincode: "110092",
		},
	})

	//Read a specific record from file name
	record2, err := db.Read("users", "John10")
	if err != nil{
		fmt.Println("Error", err)
	}
	fmt.Println(record2)

	//Delete any particular record from collection
	err4 := db.Delete("users","John10")
	if err4 != nil{
		fmt.Println("Error", err4)
	}

	// Update a particular field in a record
	updatedFields := map[string]interface{}{
		"name": "Ragnar",
		"age":  25,
		"address": map[string]interface{}{
			"city": "Rajasthan",
		},
	}
	s := User{}
	err5 := db.Update("users", "John20", s, updatedFields)
	if err5 != nil{
		fmt.Println("Error", err5)
	}
}
