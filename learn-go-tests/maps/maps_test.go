package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}

	}

	assetError := func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
	}

	assetNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("expected to get no error")
		}
	}

	t.Run("Search", func(t *testing.T) {
		got, err := dict.Search("test")
		want := "this is a test"
		if err != nil {
			t.Fatal("Expected no error got an error")
		}
		assertStrings(t, got, want)
	})

	t.Run("missing word", func(t *testing.T) {
		_, err := dict.Search("missing")

		if err == nil {
			t.Fatal("Expected error got nil")
		}

		assetError(t, err, ErrNotFound)

	})

	t.Run("add to map", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "add test")

		want := "add test"

		got, err := dict.Search("test")

		assetNoError(t, err)
		assertStrings(t, got, want)

	})

}
