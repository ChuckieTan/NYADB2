package booter_test

import (
	"bytes"
	"nyadb2/backend/utils/booter"
	"testing"
)

func TestBooter(t *testing.T) {
	bt := booter.Create("/tmp/booter_test")
	raw := []byte("123456jksadhfjksadflkwejflk;n")

	bt.Update(raw)

	if !bytes.Equal(raw, bt.Load()) {
		t.Fatal(raw, " ", bt.Load())
	}
}
