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

package types

import (
	"net/http"
	"strconv"
)

// From JSONRPC 2.0 spec
type RPCErrorCode int

const (
	RPCErrorCodeParseError     RPCErrorCode = -32700
	RPCErrorCodeInvalidRequest RPCErrorCode = -32600
	RPCErrorCodeMethodNotFound RPCErrorCode = -32601
	RPCErrorCodeInvalidParams  RPCErrorCode = -32602
	RPCErrorCodeInternalError  RPCErrorCode = -32603
	RPCErrorCodeServerError    RPCErrorCode = -32000
)

func (code RPCErrorCode) String() string {
	switch code {
	case RPCErrorCodeParseError:
		return "Parse Error"
	case RPCErrorCodeInvalidRequest:
		return "Parse Error"
	case RPCErrorCodeMethodNotFound:
		return "Method Not Found"
	case RPCErrorCodeInvalidParams:
		return "Invalid Params"
	case RPCErrorCodeInternalError:
		return "Internal Error"
	case RPCErrorCodeServerError:
		return "Server Error"
	default:
		return strconv.FormatInt(int64(code), 10)
	}
}

func (code RPCErrorCode) HTTPStatusCode() int {
	switch code {
	case RPCErrorCodeInvalidRequest:
		return http.StatusBadRequest
	case RPCErrorCodeMethodNotFound:
		return http.StatusMethodNotAllowed
	default:
		return http.StatusInternalServerError
	}
}

func (code RPCErrorCode) Error() string {
	return code.String()
}
