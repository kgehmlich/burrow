// Copyright 2019 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path"

	"github.com/hyperledger/burrow/deploy/compile"
	"github.com/hyperledger/burrow/logging"
)

func main() {
	for _, solfile := range os.Args[1:] {
		resp, err := compile.Compile(solfile, false, "", nil, logging.NewNoopLogger())
		if err != nil {
			fmt.Printf("failed compile solidity: %v\n", err)
			os.Exit(1)
		}

		if resp.Error != "" {
			fmt.Print(resp.Error)
			os.Exit(1)
		}

		if resp.Warning != "" {
			fmt.Print(resp.Warning)
			os.Exit(1)
		}

		f, err := os.Create(solfile + ".go")
		if err != nil {
			fmt.Printf("failed to create go file: %v\n", err)
			os.Exit(1)
		}

		f.WriteString(fmt.Sprintf("package %s\n\n", path.Base(path.Dir(solfile))))
		f.WriteString("import hex \"github.com/tmthrgd/go-hex\"\n\n")

		for _, c := range resp.Objects {
			f.WriteString(fmt.Sprintf("var Bytecode_%s = hex.MustDecodeString(\"%s\")\n",
				c.Objectname, c.Contract.Evm.Bytecode.Object))
			f.WriteString(fmt.Sprintf("var Abi_%s = []byte(`%s`)\n",
				c.Objectname, c.Contract.Abi))
		}
	}
}
