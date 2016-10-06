package countries

import "testing"

func TestList(t *testing.T) {
	t.Parallel()

	list := List()

	if list == nil || len(list) == 0 {
		t.Error("List is empty?")
	}

	for k, v := range list {
		if v.Alpha3 == "" || v.FullName == "" || v.Numeric == 0 {
			t.Errorf("Missing information with country `%s'", k)
		}
		if len(k) != 2 || len(v.Alpha3) != 3 {
			t.Errorf("Problem with the length of Alpha(s): %s & %s", k, v.Alpha3)
		}
	}
}
