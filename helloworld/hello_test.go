package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone.", func(t *testing.T) {
		got := Hello("Danick", "")
		want := "Hello Danick!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello World!' when an empty string is specified.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello World!' in Spanish.", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hello World!' in French.", func(t *testing.T) {
		got := Hello("Sophie", "French")
		want := "Bonjour Sophie!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
