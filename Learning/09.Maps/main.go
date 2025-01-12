package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")
	
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all the langauges: ", languages)
	fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all the langauges: ", languages)

	// LOOPS
	for key, value := range(languages) {
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}
}
