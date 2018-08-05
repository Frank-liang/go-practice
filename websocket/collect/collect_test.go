package collect

import "testing"

func TestGetPS(t *testing.T) {
	res, err := GetPS()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}
