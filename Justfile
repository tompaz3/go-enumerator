# MIT License
#
# Copyright (c) 2024 Tomasz Paździurek
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.
#


linters-run:
  @gofumpt -l -w .
  @golangci-lint run --fix -j 3 ./...
  @nilaway -include-pkgs="github.com/tompaz3/go-enumerator" ./...

go-install:
  @go install go.uber.org/nilaway/cmd/nilaway@latest
  @go install mvdan.cc/gofumpt@latest

go-install-enumerator:
  @go install github.com/tompaz3/go-enumerator@v0.0.1

go-test:
  @ginkgo -r -p ./...

go-build:
  @go build -o ./bin/enumerator ./enumerator/main/main.go
