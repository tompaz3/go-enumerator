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

package generator

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
)

const generatorPackageName = "github.com/tompaz3/go-enumerator"

type generator struct {
	enum   generationEnum
	buf    *bytes.Buffer
	writer *Writer
}

func newGenerator(enum Enum) *generator {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)
	return &generator{
		enum:   newGenerationEnum(enum),
		buf:    &buf,
		writer: NewWriter(writer),
	}
}

type generationEnum struct {
	Enum
	baseStruct                  string
	marshallableStruct          string
	invalidNameError            string
	invalidNameErrorConstructor string
}

func newGenerationEnum(enum Enum) generationEnum {
	return generationEnum{
		Enum:                        enum,
		baseStruct:                  "base" + enum.Type,
		marshallableStruct:          "Marshallable" + enum.Type,
		invalidNameError:            "Invalid" + enum.Type + "NameError",
		invalidNameErrorConstructor: "newInvalid" + enum.Type + "NameError",
	}
}

func Generate(enum Enum) error {
	src, srcErr := generateSource(enum)
	if srcErr != nil {
		return srcErr
	}

	return save(src, enum.Destination)
}

func save(sourceCode []byte, destination *string) error {
	file, resolveDestErr := resolveDestination(destination)
	if resolveDestErr != nil {
		return resolveDestErr
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if file == nil {
		return nil
	}

	if _, writeErr := file.Write(sourceCode); writeErr != nil {
		return newSaveFileError(writeErr)
	}

	return nil
}

func generateSource(enum Enum) ([]byte, error) {
	if err := enum.validate(); err != nil {
		return nil, err
	}

	gen := newGenerator(enum)
	gen.generateCopyright()
	gen.generateHeader()
	gen.generateImports()
	gen.generateInterface()
	gen.generateBaseImpl()
	gen.generateValues()
	gen.generateOfString()
	gen.generateJSONMarshalling()
	gen.generateInvalidNameError()

	if err := gen.writer.Flush(); err != nil {
		return nil, err
	}

	return gen.buf.Bytes(), nil
}

func (g *generator) generateCopyright() {
	newCopyrightGenerator(g.enum, g.writer).
		generateCopyrightClause()
}

func (g *generator) generateHeader() {
	w := g.writer
	e := g.enum
	w.Line("package " + e.Package)
	w.LineBreak()
	w.Line("// Code generated by " + generatorPackageName + " DO NOT EDIT.")
	inputArgs := ""
	if e.InputArgs != "" {
		inputArgs = " " + e.InputArgs
	}
	w.Line("//go:generate enumerator" + inputArgs)
	w.LineBreak()
}

func (g *generator) generateImports() {
	w := g.writer

	if !g.enum.Marshalling.JSONOptions.Generate {
		return
	}
	w.Line("import (")
	w.Line("\t\"bytes\"")
	newJSONMarshallerGenerator(g.enum, g.writer).
		generateImports()
	w.Line(")")
	w.LineBreak()
}

func (g *generator) generateInterface() {
	w := g.writer
	e := g.enum
	w.Line("type " + e.Type + " interface {")
	w.Line("\tsealed" + e.Type + "()")
	w.Line("\tString() string")
	newJSONMarshallerGenerator(g.enum, g.writer).
		generateToMarshallerDeclaration()
	w.Line("}")
	w.LineBreak()
}

func (g *generator) generateBaseImpl() {
	w := g.writer
	e := g.enum
	w.Line("type " + e.baseStruct + " struct {")
	w.Line("\tname string")
	w.Line("}")
	w.LineBreak()
	w.Line("func (b " + e.baseStruct + ") sealed" + e.Type + "() {}")
	w.LineBreak()
	w.Line("func (b " + e.baseStruct + ") String() string {")
	w.Line("\treturn b.name")
	w.Line("}")
	w.LineBreak()
}

func (g *generator) generateValues() {
	w := g.writer
	e := g.enum
	w.Line("var (")

	for _, value := range e.Values {
		w.Line("\t" + value + " = " + e.baseStruct + "{name: \"" + value + "\"}")
	}

	w.LineBreak()
	w.Line("\tallValuesByString = map[string]" + e.Type + "{")

	for _, value := range e.Values {
		w.Line("\t\t" + value + ".String(): " + value + ",")
	}
	w.Line("\t}")
	w.Line(")")
	w.LineBreak()
}

func (g *generator) generateOfString() {
	newOfStringGenerator(g.enum, g.writer).
		generateOfStringMethods()
}

func (g *generator) generateJSONMarshalling() {
	newJSONMarshallerGenerator(g.enum, g.writer).
		generateJSONMarshalling()
}

func (g *generator) generateInvalidNameError() {
	newInvalidNameErrorGenerator(g.enum, g.writer).
		generateInvalidNameError()
}

func resolveDestination(destination *string) (*os.File, error) {
	if destination == nil || len(*destination) == 0 {
		return os.Stdout, nil
	}

	if err := os.MkdirAll(filepath.Dir(*destination), os.ModePerm); err != nil {
		return nil, newSaveFileError(err)
	}
	file, err := os.Create(*destination)
	if err != nil {
		return nil, newSaveFileError(err)
	}
	return file, nil
}
