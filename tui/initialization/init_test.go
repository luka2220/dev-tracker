package initialization

import (
	"fmt"
	"testing"
)

func TestValidateProjectNameInput(t *testing.T) {
	// Defining invalid test values
	invalid1 := "#test name"
	invalid2 := "!@--namw  $"
	invalid3 := "~!@#$%^&*()+invalid"
	invalid4 := ""

	// Defining valid test values
	valid1 := "json-parser"
	valid2 := "interesting interpreter"
	valid3 := "rest_API7777"
	valid4 := "web_socket-client server"

	t1 := ValidateProjectNameInput(invalid1)
	t2 := ValidateProjectNameInput(invalid2)
	t3 := ValidateProjectNameInput(invalid3)
	t4 := ValidateProjectNameInput(invalid4)
	t5 := ValidateProjectNameInput(valid1)
	t6 := ValidateProjectNameInput(valid2)
	t7 := ValidateProjectNameInput(valid3)
	t8 := ValidateProjectNameInput(valid4)

	if t1 {
		validChard := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Project name can only contain %s => %s, should return false", validChard, invalid1)
		t.Error(msg)
	}

	if t2 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Project name can only contain %s => %s, should return false", validChars, invalid2)
		t.Error(msg)
	}

	if t3 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Project name can only contain %s => %s, should return false", validChars, invalid3)
		t.Error(msg)
	}

	if t4 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Project name can only contain %s => %s, should return false", validChars, invalid4)
		t.Error(msg)
	}

	if !t5 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid1, validChars)
		t.Error(msg)
	}

	if !t6 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid2, validChars)
		t.Error(msg)
	}

	if !t7 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid3, validChars)
		t.Error(msg)
	}

	if !t8 {
		validChars := "\033[1;32malph-numeric characters or '-', '_', ''\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid4, validChars)
		t.Error(msg)
	}
}

func TestValidateActiveBoardInput(t *testing.T) {
	// Defning invalid test values
	invalid1 := "tr"
	invalid2 := ""
	invalid3 := "-"
	invalid4 := "a"

	// Defining valid test values
	valid1 := "y"
	valid2 := "n"
	valid3 := "Y"
	valid4 := "N"

	t1 := ValidateActiveBoardInput(invalid1)
	t2 := ValidateActiveBoardInput(invalid2)
	t3 := ValidateActiveBoardInput(invalid3)
	t4 := ValidateActiveBoardInput(invalid4)
	t5 := ValidateActiveBoardInput(valid1)
	t6 := ValidateActiveBoardInput(valid2)
	t7 := ValidateActiveBoardInput(valid3)
	t8 := ValidateActiveBoardInput(valid4)

	if t1 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return false (valid chars: %s)", invalid1, validChars)
		t.Error(msg)
	}

	if t2 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return false (valid chars: %s)", invalid2, validChars)
		t.Error(msg)
	}

	if t3 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return false (valid chars: %s)", invalid3, validChars)
		t.Error(msg)
	}

	if t4 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return false (valid chars: %s)", invalid4, validChars)
		t.Error(msg)
	}

	if !t5 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid1, validChars)
		t.Error(msg)
	}

	if !t6 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid2, validChars)
		t.Error(msg)
	}

	if !t7 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid3, validChars)
		t.Error(msg)
	}

	if !t8 {
		validChars := "\033[1;32my, Y, n, N\033[0m" // adds bolded green ANSI colouring
		msg := fmt.Sprintf("❌ Something went wrong => %s, should return true (valid chars: %s)", valid4, validChars)
		t.Error(msg)
	}
}
