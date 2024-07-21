// MIT License
//
// Copyright (c) 2024 Tomasz Paździurek
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

	color "github.com/tompaz3/go-enumerator/internal/generator/colorwithundefinedandjsonmarshalling"
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
			name:  `GIVEN "Undefined" WHEN Of THEN Undefined`,
			value: "Undefined",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Undefined, r.color)
			},
		},
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

func Test_Color_OfOrUndefined(t *testing.T) {
	t.Parallel()

	type result struct {
		color color.Color
	}

	tests := []struct {
		name  string
		value string
		then  func(t *testing.T, r result)
	}{
		{
			name:  `GIVEN "Undefined" WHEN Of THEN Undefined`,
			value: "Undefined",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, color.Undefined, r.color)
			},
		},
		{
			name:  `GIVEN "Red" WHEN Of THEN Red`,
			value: "Red",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, color.Red, r.color)
			},
		},
		{
			name:  `GIVEN "Green" WHEN Of THEN Green`,
			value: "Green",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, color.Green, r.color)
			},
		},
		{
			name:  `GIVEN "Blue" WHEN Of THEN Blue`,
			value: "Blue",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, color.Blue, r.color)
			},
		},
		{
			name:  `GIVEN "" WHEN Of THEN Undefined`,
			value: "",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, r.color, color.Undefined)
			},
		},
		{
			name:  `GIVEN "InvalidColor" WHEN Of THEN Undefined`,
			value: "InvalidColor",
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Equal(t, r.color, color.Undefined)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			clr := color.OfOrUndefined(tt.value)
			tt.then(t, result{
				color: clr,
			})
		})
	}
}

func Test_MarshallableColor_MarshalJSON(t *testing.T) {
	t.Parallel()

	type result struct {
		marhsalled []byte
		err        error
	}

	tests := []struct {
		name  string
		color color.Color
		then  func(t *testing.T, r result)
	}{
		{
			name:  `GIVEN Undefined WHEN MarshalJSON THEN "Undefined"`,
			color: color.Undefined,
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, []byte(`"Undefined"`), r.marhsalled)
			},
		},
		{
			name:  `GIVEN Red WHEN MarshalJSON THEN "Red"`,
			color: color.Red,
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, []byte(`"Red"`), r.marhsalled)
			},
		},
		{
			name:  `GIVEN Green WHEN MarshalJSON THEN "Green"`,
			color: color.Green,
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, []byte(`"Green"`), r.marhsalled)
			},
		},
		{
			name:  `GIVEN Blue WHEN MarshalJSON THEN "Blue"`,
			color: color.Blue,
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, []byte(`"Blue"`), r.marhsalled)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			marshalled, err := tt.color.ToJSONMarshallable().MarshalJSON()
			tt.then(t, result{
				marhsalled: marshalled,
				err:        err,
			})
		})
	}
}

func Test_MarshallableColor_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type result struct {
		color color.Color
		err   error
	}

	tests := []struct {
		name string
		json []byte
		then func(t *testing.T, r result)
	}{
		{
			name: `GIVEN "Undefined" WHEN UnmarshalJSON THEN Undefined`,
			json: []byte(`"Undefined"`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Undefined, r.color)
			},
		},
		{
			name: `GIVEN "Red" WHEN UnmarshalJSON THEN Red`,
			json: []byte(`"Red"`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Red, r.color)
			},
		},
		{
			name: `GIVEN "Green" WHEN UnmarshalJSON THEN Green`,
			json: []byte(`"Green"`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Green, r.color)
			},
		},
		{
			name: `GIVEN "Blue" WHEN UnmarshalJSON THEN Blue`,
			json: []byte(`"Blue"`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Equal(t, color.Blue, r.color)
			},
		},
		{
			name: `GIVEN empty JSON WHEN UnmarshalJSON THEN zero`,
			json: make([]byte, 0),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Zero(t, r.color)
			},
		},
		{
			name: `GIVEN null JSON WHEN UnmarshalJSON THEN zero`,
			json: []byte("null"),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.NoError(t, r.err)
				assert.Zero(t, r.color)
			},
		},
		{
			name: `GIVEN "" JSON WHEN UnmarshalJSON THEN error`,
			json: []byte(`""`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Error(t, r.err)
				assert.Equal(t, "could not unmarshal Color from JSON\ninvalid Color name: \"\"", r.err.Error())
				assert.Zero(t, r.color)
			},
		},
		{
			name: `GIVEN "InvalidColor" JSON WHEN UnmarshalJSON THEN error`,
			json: []byte(`"InvalidColor"`),
			then: func(t *testing.T, r result) {
				t.Helper()
				assert.Error(t, r.err)
				assert.Equal(t, "could not unmarshal Color from JSON\ninvalid Color name: \"InvalidColor\"", r.err.Error())
				assert.Zero(t, r.color)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var marshallable color.MarshallableColor
			err := marshallable.UnmarshalJSON(tt.json)
			tt.then(t, result{
				color: marshallable.ToEnum(),
				err:   err,
			})
		})
	}
}
