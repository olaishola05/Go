package main

import "fmt"

type Part string

func printAssemlyLine(title string, slice []Part) {
	fmt.Println()
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		assemblies := slice[i]
		fmt.Println(assemblies)
	}
}

func main() {
	assembliesPart := make([]Part, 3)
	assembliesPart[0] = "Head Lamp"
	assembliesPart[1] = "Tyre"
	assembliesPart[2] = "Brake pads"
	printAssemlyLine("Part in store", assembliesPart)

	assembliesPart = append(assembliesPart, "Nitro", "Cylinder")
	printAssemlyLine("Updated assembliesPart assemblies", assembliesPart)

	assembliesPart = assembliesPart[3:]
	printAssemlyLine("New parts", assembliesPart)
}
