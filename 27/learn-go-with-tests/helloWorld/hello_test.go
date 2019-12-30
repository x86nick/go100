// write a test for Hello() function

package main

import "testing"

func TestHello(t *testing.T) {
	// got := Hello()
	// want := "Hello, world"

	got := Hello("Go")
	want := "Hello, Go"

	// We are calling the Errorf method on our t which will print out a message and fail the test.
	// %q is very useful as it wraps your values in double quotes.
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
