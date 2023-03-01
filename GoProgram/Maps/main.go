package main

import "fmt"

func main() {

	//create empty map
	// var colors map[string]string

	//another way create empty map using ":="
	colorsAnother := make(map[string]string) //the same line 8

	//declaration maps with value
	colors := map[string]string{
		"red":    "#FF0000",
		"yellow": "#FFFF00",
	}

	//add to map (used by line's way)
	colorsAnother["white"] = "#FFFFFF"

	//print map
	fmt.Println("Map before:", colorsAnother)

	//access map by key
	valueOfWhiteKey := colorsAnother["white"]
	fmt.Println("Value of White Key:", valueOfWhiteKey)

	//retrieve key and value of map by key (first value stored key if exist and zero value if not exists, second variable stored true if key exsist and back)
	key, isExist := colorsAnother["white"]
	fmt.Println("Key:", key, "isExists:", isExist)
	newKey, newIsExist := colorsAnother["another"]
	fmt.Println("Key:", newKey, "IsExists:", newIsExist)

	//retrive isExist key
	_, keyIsExist := colorsAnother["black"]
	fmt.Println("Black is Exists:", keyIsExist)

	//check exists map by key
	valueOfAnotherKey := colorsAnother["another"]
	fmt.Println("Has Another key:", valueOfAnotherKey) //return zero value if key not exsist

	//len of map
	lengthOfMap := len(colorsAnother)
	fmt.Println("Len of Map:", lengthOfMap)

	//delete map by key
	delete(colorsAnother, "white")

	fmt.Println(colors)
	fmt.Println(colorsAnother)

	//interating over map
	interatingOverMaps(colors)

}

func interatingOverMaps(maps map[string]string) {
	fmt.Println("Interating over Maps")
	for key, value := range maps {
		fmt.Println("\tHex code for", key, "is", value)
	}
}
