package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T)  {
		got := Hello("Sean", "English")
		want := "Hello, Sean"
		assetCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func (t *testing.T)  {
		got := Hello("", "English")
		want := "Hello, World"
		assetCorrectMessage(t, got, want)
	})

	t.Run(("in Spanish"), func (t *testing.T)  {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}