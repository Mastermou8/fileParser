package menu

import "fmt"

func SelectOption() (error, int) {

	println("Select an option: Extract first page(1) Crop file(2)")
	var x int
	x, err := fmt.Scanln()
	if err != nil {
		return fmt.Errorf("Failed reading input: %w", err), 0
	}
	return nil, x
}
