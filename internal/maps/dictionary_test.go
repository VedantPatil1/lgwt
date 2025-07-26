package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "This is just a test"}

	t.Run("known key", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "This is just a test"
		assertError(t, err, nil)
		assertString(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {
		got, err := dictionary.Search("unknown")

		assertError(t, err, ErrorNotFound)
		assertString(t, got, "")
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		err := dictionary.Add("test", "test is just a test")
		want := "test is just a test"

		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", want)
	})
	t.Run("word exists", func(t *testing.T) {

		dictionary := Dictionary{"test": "existing"}
		err := dictionary.Add("test", "test is just a test")
		assertError(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, "test", "existing")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test": "existing"}
		err := dictionary.Update("test", "updated")
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", "updated")
	})
	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "existing"}
		err := dictionary.Update("unknown", "updated")
		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "existing"}
		err := dictionary.Delete("test")
		assertError(t, err, nil)
		_, got := dictionary.Search("test")
		assertError(t, got, ErrorNotFound)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("unknown")
		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func assertError(t testing.TB, err, want error) {
	if err != want {
		t.Errorf("Got %q want %q", err, want)
	}
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q given %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, definition string) {
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}
	assertString(t, got, definition)
}
