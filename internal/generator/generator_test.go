// MIT License
//
// Copyright (c) 2024-2025 Tomasz Paździurek
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

package generator_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tompaz3/go-enumerator/internal/generator"

	_ "embed" // embed package is imported for go:embed directive
)

const licenseFilePath = "../../LICENSE"

//go:embed colorplain/expected_color.txt
var expectedColor []byte

//go:embed colorwithundefined/expected_color.txt
var expectedColorWithUndefined []byte

//go:embed colorwithjsonmarshalling/expected_color.txt
var expectedColorWithJSONMarshalling []byte

//go:embed colorwithundefinedandjsonmarshalling/expected_color.txt
var expectedColorWithUndefinedAndJSONMarshalling []byte

//go:embed colorwithundefinedandjsonmarshallingniltoundefined/expected_color.txt
var expectedColorWithUndefinedAndJSONMarshallingNilToUndefined []byte

//go:embed colorwithoutcopyright/expected_color.txt
var expectedColorWithoutCopyrightClause []byte

//nolint:funlen // this test prepares a few test cases, which tend to be lengthy
func Test_Generator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		enum     func() generator.Enum
		expected []byte
	}{
		{
			name: `generate plain`,
			enum: func() generator.Enum {
				destination := "./colorplain/color.go"
				return generator.Enum{
					Destination:   &destination,
					CopyrightFile: licenseFilePath,
					Package:       "color",
					Type:          "Color",
					Values:        []string{"Red", "Green", "Blue"},
					CheckSumType:  true,
				}
			},
			expected: expectedColor,
		},
		{
			name: `generate with undefined`,
			enum: func() generator.Enum {
				destination := "./colorwithundefined/color.go"
				return generator.Enum{
					Destination:    &destination,
					CopyrightFile:  licenseFilePath,
					Package:        "color",
					Type:           "Color",
					Values:         []string{"Undefined", "Red", "Green", "Blue"},
					UndefinedValue: "Undefined",
				}
			},
			expected: expectedColorWithUndefined,
		},
		{
			name: `generate with JSON marshalling`,
			enum: func() generator.Enum {
				destination := "./colorwithjsonmarshalling/color.go"
				return generator.Enum{
					Destination:   &destination,
					CopyrightFile: licenseFilePath,
					Package:       "color",
					Type:          "Color",
					Values:        []string{"Red", "Green", "Blue"},
					Marshalling: generator.MarshalOptions{
						JSONOptions: generator.JSONMarshalOptions{
							Generate: true,
						},
					},
				}
			},
			expected: expectedColorWithJSONMarshalling,
		},
		{
			name: `generate with undefined and JSON marshalling`,
			enum: func() generator.Enum {
				destination := "./colorwithundefinedandjsonmarshalling/color.go"
				return generator.Enum{
					Destination:    &destination,
					CopyrightFile:  licenseFilePath,
					Package:        "color",
					Type:           "Color",
					Values:         []string{"Undefined", "Red", "Green", "Blue"},
					UndefinedValue: "Undefined",
					Marshalling: generator.MarshalOptions{
						JSONOptions: generator.JSONMarshalOptions{
							Generate: true,
						},
					},
				}
			},
			expected: expectedColorWithUndefinedAndJSONMarshalling,
		},
		{
			name: `generate with undefined and JSON marshalling and nil to undefined`,
			enum: func() generator.Enum {
				destination := "./colorwithundefinedandjsonmarshallingniltoundefined/color.go"
				return generator.Enum{
					Destination:    &destination,
					CopyrightFile:  licenseFilePath,
					Package:        "color",
					Type:           "Color",
					Values:         []string{"Undefined", "Red", "Green", "Blue"},
					UndefinedValue: "Undefined",
					Marshalling: generator.MarshalOptions{
						JSONOptions: generator.JSONMarshalOptions{
							Generate:       true,
							NilToUndefined: true,
						},
					},
				}
			},
			expected: expectedColorWithUndefinedAndJSONMarshallingNilToUndefined,
		},
		{
			name: `generate with undefined and JSON marshalling and nil to undefined`,
			enum: func() generator.Enum {
				destination := "./colorwithoutcopyright/color.go"
				return generator.Enum{
					Destination:    &destination,
					Package:        "color",
					Type:           "Color",
					Values:         []string{"Undefined", "Red", "Green", "Blue"},
					UndefinedValue: "Undefined",
					Marshalling: generator.MarshalOptions{
						JSONOptions: generator.JSONMarshalOptions{
							Generate:       true,
							NilToUndefined: true,
						},
					},
				}
			},
			expected: expectedColorWithoutCopyrightClause,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// given
			enum := tt.enum()

			// when
			err := generator.Generate(enum)

			// then
			assert.NoError(t, err)
			// and
			content, err := os.ReadFile(*enum.Destination)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, content)
		})
	}
}
