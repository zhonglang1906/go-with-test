package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	asserCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() //告知测试套件，这个方法是辅助函数。且测试失败时，报告的行号将在函数调用中，而不是辅助函数内部。
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("chris", "")
		want := "Hello, chris"

		//if got != want {
		//	t.Errorf("got %q want %q", got, want)
		//}
		asserCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		//if got != want {
		//	t.Errorf("got %q want %q", got, want)
		//}
		asserCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		asserCorrectMessage(t, got, want)
	})
}
