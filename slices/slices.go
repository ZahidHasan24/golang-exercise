//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printContent(itemList []Part) {
	for i := 0; i < len(itemList); i++ {
		item := itemList[i]
		fmt.Println(item)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	assemblyLine := make([]Part, 3)

	//  - Create an assembly line having any three parts
	assemblyLine[0] = "Book"
	assemblyLine[1] = "Note"
	assemblyLine[2] = "Pen"

	fmt.Println("3 Parts: ")
	printContent(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Eraser", "Sharpner")
	fmt.Println("\nAdded two parts:")
	printContent(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\nSliced:")
	printContent(assemblyLine)
}
