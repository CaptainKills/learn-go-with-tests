package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Test": "This is just a test."}

	t.Run("known error", func(t *testing.T) {
		got, _ := dictionary.Search("Test")
		want := "This is just a test."

		assertString(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dictionary.Search("Unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("Test", "This is just a test.")

		want := "This is just a test."

		assertDefinition(t, dictionary, "Test", want)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "Test"
		value := "This is just a test."

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "New value.")

		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "Test"
		value := "This is just a test."

		dictionary := Dictionary{key: value}
		newValue := "This is a new value."

		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "Test"
		value := "This is a new value."

		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		dictionary := Dictionary{"Test": "This is just a test."}

		dictionary.Delete("Test")

		_, err := dictionary.Search("Test")

		assertError(t, err, ErrorNotFound)
	})

	t.Run("non-existing key", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("Test")

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("Expected key to be found: ", err)
	}

	assertString(t, got, value)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
