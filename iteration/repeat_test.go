package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("got %q want %q", repeated, expected)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
