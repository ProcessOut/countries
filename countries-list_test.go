package traveler

import "testing"

func TestList(t *testing.T) {
	t.Parallel()

	list := Countries()

	if list == nil || len(list) == 0 {
		t.Error("List is empty?")
	}

	for k, v := range list {
		if v.CountryAlpha3 == "" || v.FullName == "" || v.Numeric == 0 {
			t.Errorf("Missing information with country `%s'", k)
		}
		if len(k) != 2 || len(v.CountryAlpha3) != 3 {
			t.Errorf("Problem with the length of Alpha(s): %s & %s", k, v.CountryAlpha3)
		}
	}
}
