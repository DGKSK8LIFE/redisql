package utils

import "fmt"

// PrintRow prettyprints a SQL row
func PrintRow(id string, m map[string]string) {
	fmt.Printf(fmt.Sprintf("%s: \n", id))
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}
}
