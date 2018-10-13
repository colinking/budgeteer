package mock

import (
	"testing"
)

func TestToken(t *testing.T) {
	d := New()
	token := "hello world"
	d.SaveToken(token)

	if d.GetToken() != token {
		t.Errorf("Invalid token returned")
	}
}
