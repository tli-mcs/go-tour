// Exercise: JSON!!

// Create a struct called "Dimensions" with 2 string attributes: Height and Weight
// Create a struct "Human", with 2 string attributes: Name and Description and one called "Dimensions" of type Dimensions
// Now in the json object we are going to have a nested object
// Print out the values ONLY of that nested object

package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name        string
	Description string
	Dimensions  Dimensions
}

type Dimensions struct {
	Height string
	Weight string
}

func main() {
	humanJson := `{"name": "Rick",
                  "description": "has a grandson called Morty",
                  "dimensions": { 
                    "height": "1.80m",
                    "weight": "50kg"
                  }
                }`
	human := Human{}
	json.Unmarshal([]byte(humanJson), &human)
	fmt.Println(human.Dimensions)
}
