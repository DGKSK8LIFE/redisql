package utils

import "fmt"

// PrintRow prettyprints a SQL row
func PrintRow(id string, m map[string]string) {
	fmt.Println(fmt.Sprintf("%s: \n", id))
	for key, value := range m {
		fmt.Printf("\t%s: %s\n", key, value)
	}
}
