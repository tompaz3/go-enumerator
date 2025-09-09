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

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/tompaz3/go-enumerator/internal/generator"
)

// version - can be overwritten by -ldflags "-X main.version=v<major.minor.patch>"
var version = "v0.0.7"

func main() {
	inputArgs := strings.Join(os.Args, " ")
	copyrightFile := flag.String("copyright", "", "license file")
	destination := flag.String("destination", "", "destination file")
	packageName := flag.String("package", "", "package name")
	typeName := flag.String("type", "", "type name")
	valueNames := flag.String("values", "", "comma-separated values")
	undefinedValue := flag.String("undefined", "", "undefined value name - must be one of the values")
	marshalJSON := flag.Bool("marshal-json", false, "generate JSON marshalling")
	unmarshalUnknownToUndefined := flag.Bool(
		"unmarshal-json-to-undefined",
		false,
		"unmarshal unknown or null values to undefined",
	)
	checkSumType := flag.Bool(
		"go-check-sumtype",
		false,
		"add go-check-sumtype comment for exhaustiveness check using https://github.com/alecthomas/go-check-sumtype",
	)
	versionPrintRequested := flag.Bool("version", false, "print version")
	flag.Parse()

	if *versionPrintRequested {
		printVersion()
		return
	}

	values := stripValueNames(*valueNames)

	enum := generator.Enum{
		InputArgs:      inputArgs,
		CopyrightFile:  *copyrightFile,
		Destination:    destination,
		Package:        *packageName,
		Type:           *typeName,
		Values:         values,
		UndefinedValue: *undefinedValue,
		Marshalling: generator.MarshalOptions{
			JSONOptions: generator.JSONMarshalOptions{
				Generate:       *marshalJSON,
				NilToUndefined: *unmarshalUnknownToUndefined,
			},
		},
		CheckSumType: *checkSumType,
	}
	err := generator.Generate(enum)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Generated %s to %s!\n", enum.Type, *enum.Destination)
}

func stripValueNames(valueNames string) []string {
	if valueNames == "" {
		return []string{}
	}
	return strings.Split(valueNames, ",")
}

func printVersion() {
	fmt.Println(resolveVersion())
}

func resolveVersion() string {
	if version != "" { // overwritten by -ldflags
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return version
	}

	ver := info.Main.Version
	if ver != "" {
		return ver
	}

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			return s.Value
		}
	}

	return version
}
