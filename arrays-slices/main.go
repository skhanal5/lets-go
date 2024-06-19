package main

import "fmt"

func main() {

	// exercise 1-4
	vals := getHobbies()
	outputHobbies(vals[:])
	
	// exercise 5-6
	dynamic := getCourseGoals()
	outputGoals(dynamic)
}

func getHobbies() [3]string {
	hobbyOne := getInput("What is your hobby: ")
	hobbyTwo := getInput("What is your hobby: ")
	hobbyThree := getInput("What is your hobby: ")

	return [3]string{hobbyOne, hobbyTwo, hobbyThree}
}

func outputHobbies(vals []string) {
	fmt.Println("Array of hobbies ", vals)
	fmt.Println("First element ", vals[0])
	fmt.Println("Combined hobby new list ", vals[1:3])

	slce := vals[:2]
	fmt.Println("First two elements slice ", slce)
	anotherSlce := vals[:1]
	anotherSlce = vals[:2]
	fmt.Println("First two elements slice ", anotherSlce)

	alt := vals[1:3]
	fmt.Println("Reslicing the slice ", alt)
}

func getCourseGoals() []string {
	dynamic := []string{}
	goalOne := getInput("What is your goal: ")
	goalTwo := getInput("What is your goal: ")

	return append(dynamic, goalOne, goalTwo)
}

func outputGoals(dynamic []string) {
	fmt.Println(dynamic)
	dynamic[1] = "changed goal 2"
	fmt.Println(dynamic)
	dynamic = append(dynamic, "a third one")
	fmt.Println(dynamic)
}

func mapExercise() {
	courseRating := map[string]float64{}
	// each time we add onto the map, go will have to reallocate memory
	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	fmt.Println(courseRating)

	/*
		We can define capacity of the map by using
		map and passing in a param.
		
		Note: different from array version of map.
		We cannot define a "predefined" slot in maps
		like arrays, only capacity can be defined
	*/
	courseRating2 := make(map[string]float64, 10)
	fmt.Println(courseRating2)
}

func makeExercise() {
	/*
		Makes a slice with predefined slots that are empty (second param)
		With the third parameter, you can define the capacity of the underlying array
		- This is more efficient, each append call won't dynamically make new arrays until we reach capacity
	*/
	userNames := make([]string, 2, 5)
	userNames[0] = "Jen"
	fmt.Println(userNames)
}

func getInput(input string) string {
	fmt.Print(input)
	hobby := ""
	fmt.Scanln(&hobby)
	return hobby
}

