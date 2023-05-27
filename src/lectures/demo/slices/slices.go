package main

import "fmt"

func printRoutes(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		routes := slice[i]
		fmt.Println(routes)

	}
}

func main() {
	routes := []string{"Grocery", "Department Store", "Salon"}
	printRoutes("Route 1", routes)

	routes = append(routes, "Jamaica", "Lucia", "Bermuda")
	printRoutes("Route 2", routes)

	fmt.Println()
	fmt.Println(routes[0], "visited")
	fmt.Println(routes[1], "visited")

	routes = routes[2:]
	printRoutes("Remaining locations", routes)

}
