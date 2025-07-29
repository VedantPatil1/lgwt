package racer

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want:= fastURL
	got:= Racer(slowURL,fastURL)

	if got := want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
