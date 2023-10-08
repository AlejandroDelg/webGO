package main

/*

package main

import (
	"Application/helpers"
	"fmt"
	"log"
	"sort"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	user := User{
		FirstName: "Alejandro",
		LastName:  "Delgado",
	}

	changeName(&user)
	log.Println(user.FirstName)

	myMap := make(map[string]User)

	myMap["primero"] = User{
		FirstName: "Alejandro", LastName: "Delgado",
	}

	myMap["segundo"] = User{FirstName: "hola", LastName: "adios"}
	log.Println(myMap)

	var myString []string

	myString = append(myString, "")

	myString = append(myString, "itemToDelete")

	log.Println(myString)

	var myInt []int

	myInt = append(myInt, 1)

	myInt = append(myInt, 3)

	myInt = append(myInt, 2)

	sort.Ints(myInt)

	users := []User{}

	users = append(users, User{
		FirstName: "Alex",
		LastName:  "Delgado",
	})
	users = append(users, User{
		FirstName: "Alejandro",
		LastName:  "Delgado",
	})

	for i, l := range users {
		log.Println(i, l)
	}

	dog := Dog{
		Name: "Cuqui", Breed: "Yorksire",
	}
	printInfo(&dog)

	gorilla := Gorilla{
		Name: "Gorilla", Color: "Red", NumberOfTeeth: 20,
	}
	printInfo(&gorilla)

	var myVar helpers.SomeTipe

	println(myVar.TypeName)

}

func printInfo(a Animal) {
	fmt.Println("This animal says: ", a.Says(), " and has ", a.NumberOfLegs(), " legs. ")
}

func changeName(user *User) {
	user.FirstName = "este es mi verdadero nombre."
}

func (d *Dog) Says() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Gorilla) Says() string {
	return "AUG!"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}




















package main

import (
	"Application/helpers"
	"encoding/json"
	"fmt"
	"log"
)

const numPools = 10

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	num := <-intChan
	log.Println(num)

	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error ")
	}

	log.Println("Unmarshalled: ", unmarshalled)

	// write JSON

	var mySlice []Person

	var m1 Person

	m1.FirstName = "name"

	m1.HairColor = "red"

	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m1.FirstName = "name2"

	m1.HairColor = "red2"

	m1.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")

	if err != nil {
		log.Println("Error")
	}

	fmt.Println(string(newJson))

}

func calculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPools)
	intChan <- randomNumber

}

*/
