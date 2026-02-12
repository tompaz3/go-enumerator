// MIT License
//
// Copyright (c) 2024-2026 Tomasz Paździurek
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package color_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	color "github.com/tompaz3/go-enumerator/internal/generator/colorplain"
)

func Test_Color_Of(t *testing.T) {
	t.Parallel()

	type result struct {
		color color.Color
		err   error
	}

	tests := []struct {
		name  string
		value string
		then  func(t *testing.T, r result)
	}{
		{
			name:  `GIVEN "Red" WHEN Of THEN Red`,
			value: "Red",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Red, r.color)
			},
		},
		{
			name:  `GIVEN "Green" WHEN Of THEN Green`,
			value: "Green",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Green, r.color)
			},
		},
		{
			name:  `GIVEN "Blue" WHEN Of THEN Blue`,
			value: "Blue",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Blue, r.color)
			},
		},
		{
			name:  `GIVEN "" WHEN Of THEN error`,
			value: "",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Error(t, r.err)
				var invalidColorNameError color.InvalidColorNameError
				assert.ErrorAs(t, r.err, &invalidColorNameError)
				assert.Equal(t, `invalid Color name: ""`, r.err.Error())
				assert.Nil(t, r.color)
			},
		},
		{
			name:  `GIVEN "InvalidColor" WHEN Of THEN error`,
			value: "InvalidColor",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Error(t, r.err)
				var invalidColorNameError color.InvalidColorNameError
				assert.ErrorAs(t, r.err, &invalidColorNameError)
				assert.Equal(t, `invalid Color name: "InvalidColor"`, r.err.Error())
				assert.Nil(t, r.color)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			clr, err := color.Of(tt.value)
			tt.then(t, result{
				color: clr,
				err:   err,
			})
		})
	}
}

func Test_Color_Values(t *testing.T) {
	t.Parallel()

	assert.Equal(t, []color.Color{color.Red, color.Green, color.Blue}, color.Values())
}
