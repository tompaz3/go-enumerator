package color

// Code generated by github.com/tompaz3/go-enumerator DO NOT EDIT.
// generate

import (
	"bytes"
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
	Undefined = baseColor{name: "Undefined"}
	Red = baseColor{name: "Red"}
	Green = baseColor{name: "Green"}
	Blue = baseColor{name: "Blue"}

	allValuesByString = map[string]Color{
		Undefined.String(): Undefined,
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

func OfOrUndefined(name string) Color {
	if value, ok := allValuesByString[name]; ok {
		return value
	}
	return Undefined
}

type MarshallableColor struct {
	en Color
}

func (b MarshallableColor) MarshalJSON() ([]byte, error) {
	return []byte("\"" + b.en.String() + "\""), nil
}

func (b *MarshallableColor) UnmarshalJSON(jsonBytes []byte) error {
	if len(jsonBytes) == 0 {
		b.en = Undefined
	}

	jsonString := bytes.NewBuffer(jsonBytes).String();
	if jsonString == "null" {
		b.en = Undefined
	}

	trimmedString := strings.Trim(jsonString, "\"")
	orUndefined := OfOrUndefined(trimmedString)
	b.en = orUndefined

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

