package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known world", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, "this is just a test")
	})

	t.Run("unknown world", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {
	t.Run("new world", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		definition := "this is just a test"
		err := dictionary.Add(key, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("existing world", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{key: definition}
		err := dictionary.Add(key, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating existing key", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		newDefinition := "hello"
		dictionary := Dictionary{key: definition}
		err := dictionary.Update(key, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("updating inexistent key", func(t *testing.T) {
		key := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(key, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	key := "test"
	definition := "this is just a test"
	dictionary := Dictionary{key: definition}
	dictionary.Delete(key)
	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("should delete key %q", key)
	}
}

func assertStrings(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)

	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, definition string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if got != definition {
		t.Errorf("got: %q, want: %q", got, definition)
	}
}
