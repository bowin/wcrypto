package wcrypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestEncodeOk (t *testing.T) {
	wc := New("exx11133", "vt16ul0py8ekh5w2rhy8n0zfr2tkh9ba4933ntroe21", "wx83f34254af1b48b4")
	raw := "message"
	assert.Equal(t, raw, wc.Decrypt(wc.Encrypt(raw)))
}
