package api_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/favonia/cloudflare-ddns-go/internal/api"
)

var ttlTests = []int{1, 2, 30, 293, 842, 8, 37284789}

func TestTTLDescribe(t *testing.T) {
	t.Parallel()

	for _, i := range ttlTests {
		if i == 1 {
			continue
		}
		ttl := api.TTL(i)
		assert.Equalf(t, ttl.Describe(), strconv.Itoa(i), "TTL = %d should be described directly.", i)
	}
}

func TestTTLOneDescribe(t *testing.T) {
	t.Parallel()

	ttl := api.TTL(1)
	expected := "1 (automatic)"

	assert.Equal(t, ttl.Describe(), expected, "TTL = 1 should be described as %q.", expected)
}

func TestTTLString(t *testing.T) {
	t.Parallel()

	for _, i := range ttlTests {
		ttl := api.TTL(i)
		assert.Equalf(t, ttl.String(), strconv.Itoa(i), "TTL.String() should display TTL directly.", i)
	}
}

func TestTTLInt(t *testing.T) {
	t.Parallel()

	for _, i := range ttlTests {
		ttl := api.TTL(i)
		assert.Equalf(t, ttl.Int(), i, "TTL.Int() should recover the underlying number.")
	}
}