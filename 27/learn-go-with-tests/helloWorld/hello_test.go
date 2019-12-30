// write a test for Hello() function

package main

import "testing"

// func TestHello(t *testing.T) {
// 	// got := Hello()
// 	// want := "Hello, world"

// 	got := Hello("Go")
// 	want := "Hello, Go"

// 	// We are calling the Errorf method on our t which will print out a message and fail the test.
// 	// %q is very useful as it wraps your values in double quotes.
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }
// // // -
// func TestHello(t *testing.T) {
// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Go")
// 		want := "Hello, Go"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("say 'Hello, world when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, world"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})

// }

// // // --

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Go")
		want := "Hello, Go"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
