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

package color

// Code generated by github.com/tompaz3/go-enumerator DO NOT EDIT.
// generate

import (
	"bytes"
	"errors"
	"strings"
)

type Color interface {
	sealedColor()
	String() string
	ToJSONMarshallable() MarshallableColor
}

type baseColor struct {
	name string
}

func (b baseColor) sealedColor() {}

func (b baseColor) String() string {
	return b.name
}

var (
	Red = baseColor{name: "Red"}
	Green = baseColor{name: "Green"}
	Blue = baseColor{name: "Blue"}

	allValuesByString = map[string]Color{
		Red.String(): Red,
		Green.String(): Green,
		Blue.String(): Blue,
	}
)

func Of(name string) (Color, error) {
	if value, ok := allValuesByString[name]; ok {
		return value, nil
	}
	return nil, newInvalidColorNameError(name)
}

type MarshallableColor struct {
	en Color
}

func (b MarshallableColor) MarshalJSON() ([]byte, error) {
	if b.en == nil {
		return []byte("null"), nil
	}
	return []byte("\"" + b.en.String() + "\""), nil
}

func (b *MarshallableColor) UnmarshalJSON(jsonBytes []byte) error {
	if len(jsonBytes) == 0 {
		return nil
	}

	jsonString := bytes.NewBuffer(jsonBytes).String()
	if jsonString == "null" {
		return nil
	}

	trimmedString := strings.Trim(jsonString, "\"")
	value, err := Of(trimmedString)
	if err != nil {
		return errors.Join(errors.New("could not unmarshal Color from JSON"), err)
	}
	b.en = value

	return nil
}

func (b baseColor) ToJSONMarshallable() MarshallableColor {
	return MarshallableColor{en: b}
}

func (m MarshallableColor) ToEnum() Color {
	return m.en
}

type InvalidColorNameError struct {
	name string
}

func (e InvalidColorNameError) Error() string {
	return "invalid Color name: \"" + e.name + "\""
}

func newInvalidColorNameError(name string) InvalidColorNameError {
	return InvalidColorNameError{name: name}
}

