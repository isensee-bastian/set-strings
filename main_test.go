package main

import "testing"

func TestSetStrings(t *testing.T) {
	set := NewSet("red", "green", "yellow", "red")
	checkInt(t, 3, set.Size())
	checkBool(t, true, set.Contains("red"))
	checkBool(t, true, set.Contains("green"))
	checkBool(t, true, set.Contains("yellow"))
	checkBool(t, false, set.Contains("blue"))
	checkContains(t, set.Slice(), "red", "green", "yellow")
	checkString(t, "green, red, yellow", set.String())

	set.Add("blue")
	checkInt(t, 4, set.Size())
	checkBool(t, true, set.Contains("blue"))
	checkContains(t, set.Slice(), "red", "green", "yellow", "blue")
	checkString(t, "blue, green, red, yellow", set.String())

	set.Add("green")
	checkInt(t, 4, set.Size())

	set.Remove("red")
	checkInt(t, 3, set.Size())
	checkContains(t, set.Slice(), "green", "yellow", "blue")
	checkString(t, "blue, green, yellow", set.String())

	set.Remove("red")
	checkInt(t, 3, set.Size())
}

func checkContains(t *testing.T, slice []string, expectedElements ...string) {
	t.Helper()
	// Since our use case covers rather small element counts, searching with a nested
	// loop is fast enough.
	for _, expected := range expectedElements {
		found := false

		for _, element := range slice {
			if expected == element {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("Expected to find %v but slice is %v", expected, slice)
		}
	}
}

func checkBool(t *testing.T, expected, actual bool) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func checkString(t *testing.T, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func checkInt(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}
