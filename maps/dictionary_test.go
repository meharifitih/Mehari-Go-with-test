package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	defination := "this is just a test"
	dictionary.Add(word, defination)
	assertdefination(t, dictionary, word, defination)
}

func assertdefination(t testing.TB, dictionart Dictionary, word, defination string) {
	t.Helper()
	got, err := dictionart.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if defination != got {
		t.Errorf("got %q want %q", got, defination)
	}
}
