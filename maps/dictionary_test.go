package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		if got == nil {
			t.Fatal("expected an error")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dict.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "test definition"}

		err := dict.Delete(word)

		assertError(t, err, nil)

		_, err = dict.Search(word)

		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dict Dictionary, word string, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertString(t, got, definition)
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
		t.Errorf("got error %q want %q", got, want)
	}
}
