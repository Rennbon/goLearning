package dbBase

import "testing"

func Test(t *testing.T) {
	if loadGlobalSession() != nil {
		t.Error()
	}
}
