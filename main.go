package main

import (
	"encoding/json"
	"fmt"

	bnkgopackage "github.com/stonetanatip/golang-playground/package/bnk-gopackage"
	"github.com/stonetanatip/golang-playground/package/school"
)

type customer struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func main() {
	fmt.Println("Hello, playground")
	name := []string{}
	name = append(name, "stone")
	fmt.Println(name)
	fmt.Println(len(name))

	model := make(map[string]int)
	model["stone"] = 22
	fmt.Println(model)

	printFirst()
	defer printFinish()
	printSecond()

	cus := customer{
		Firstname: "stone",
		Lastname:  "tanatip",
		Code:      12123,
		Phone:     "0999999999",
	}
	fmt.Println(cus)
	fmt.Println(cus.Firstname)

	customerLists := []customer{}

	customer1 := customer{
		Firstname: "stone",
		Lastname:  "happy",
		Code:      12123,
		Phone:     "0999999999",
	}

	customer2 := customer{
		Firstname: "stone",
		Lastname:  "tanatip",
		Code:      12123,
		Phone:     "0999999999",
	}

	customerLists = append(customerLists, customer1)
	customerLists = append(customerLists, customer2)
	fmt.Println(customerLists)

	customerListsJson, err := json.Marshal(customerLists)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(customerListsJson))

	// call package
	println("School Name", school.SchoolName)
	println("School Address:", school.GetSchoolAddress())
	println("Member Name : ", bnkgopackage.GetFullNameCherprang())

}

func printFirst() {
	fmt.Println("First")
}

func printSecond() {
	fmt.Println("Second")
}

func printFinish() {
	fmt.Println("Close")
}
