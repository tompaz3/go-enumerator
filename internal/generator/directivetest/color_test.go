package directivetest_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_ "embed" // embed package is imported for go:embed directive
)

//go:generate go-enumerator -destination ./color/color.go -package color -type Color -values Undefined,Red,Green,Blue -undefined Undefined -marshal-json -unmarshal-json-to-undefined -copyright ../../../LICENSE -go-checksum-type

//go:embed color/expected_color.txt
var expectedColor []byte

func Test_Generated(t *testing.T) {
	t.Parallel()

	content, err := os.ReadFile("./color/color.go")
	assert.NoError(t, err)
	assert.Equal(t, expectedColor, content)
}
