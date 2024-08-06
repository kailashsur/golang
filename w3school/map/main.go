package main

import "fmt"

func main() {
	/*
			Go Maps
		Maps are used to store data values in key:value pairs.

		Each element in a map is a key:value pair.

		A map is an unordered and changeable collection that does not allow duplicates.

		The length of a map is the number of its elements. You can find it using the len() function.

		The default value of a map is nil.

		Maps hold references to an underlying hash table.

		Go has multiple ways for creating maps.
	*/

	map_a := map[string]int{"kaialsh": 20, "sili": 21}
	var map_b = map[string]string{"kailash": "Gopalpur", "sili": "Jajpur"}

	fmt.Println("map a = ", map_a)
	fmt.Println("Map b = ", map_b)

	// __________ map using make()

	c := make(map[string]int) // now map is empty
	c["kailash sur"] = 21
	c["Sili chand"] = 20

	fmt.Println("Map c = ", c)
	delete(c, "kailash sur")
	fmt.Println("Map c = ", c)
}
