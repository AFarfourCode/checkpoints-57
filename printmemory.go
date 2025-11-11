package main

import (
	"fmt"     // For printing to console
	"unicode" // For checking if a character is printable/graphic
)

// PrintMemory prints the contents of a 10-byte array in both hexadecimal and ASCII format.
// Non-printable characters are replaced by '.' in the ASCII output.
func PrintMemory(arr [10]byte) {
	// -------------------------------
	// 1️⃣ Print hexadecimal representation
	// -------------------------------
	for i := 0; i < 10; i++ {
		// %02x formats the byte in lowercase hexadecimal with 2 digits
		// e.g., 0 -> "00", 10 -> "0a", 255 -> "ff"
		fmt.Printf("%02x", arr[i])

		// Add a space after every byte except the last one
		if i != 9 {
			fmt.Print(" ")
		}

		// Add a line break to match the expected grouping:
		// After the 4th byte (index 3) and 8th byte (index 7)
		// This formats the output exactly like the example:
		// 68 65 6c 6c
		// 6f 10 15 2a
		// 00 00
		if i == 3 || i == 7 {
			fmt.Println()
		}
	}

	// Extra newline for spacing before ASCII output
	fmt.Println()

	// -------------------------------
	// 2️⃣ Print ASCII representation
	// -------------------------------
	for i := 0; i < 10; i++ {
		// Convert byte to rune to properly handle Unicode checks
		r := rune(arr[i])

		// Check if the character is printable (graphic)
		// If printable, print the character
		// Otherwise, print '.' for non-printable bytes
		if unicode.IsGraphic(r) {
			fmt.Printf("%c", r)
		} else {
			fmt.Print(".")
		}
	}

	// Newline after ASCII output
	fmt.Println()
}

func main() {
	// Call PrintMemory with a 10-byte array
	// The array is partially initialized: {'h','e','l','l','o',16,21,'*'}
	// The remaining bytes are automatically 0 (Go zero-initializes arrays)
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
