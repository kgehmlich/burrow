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

package evm

import (
	"github.com/hyperledger/burrow/execution/errors"
	"github.com/hyperledger/burrow/execution/exec"
)

type EventSink interface {
	Call(call *exec.CallEvent, exception *errors.Exception) error
	Log(log *exec.LogEvent) error
}

type noopEventSink struct {
}

func NewNoopEventSink() *noopEventSink {
	return &noopEventSink{}
}

func (es *noopEventSink) Call(call *exec.CallEvent, exception *errors.Exception) error {
	return nil
}

func (es *noopEventSink) Log(log *exec.LogEvent) error {
	return nil
}

type logFreeEventSink struct {
	EventSink
	error error
}

func NewLogFreeEventSink(eventSink EventSink) *logFreeEventSink {
	return &logFreeEventSink{
		EventSink: eventSink,
	}
}

func (esc *logFreeEventSink) Log(log *exec.LogEvent) error {
	return errors.ErrorCodef(errors.ErrorCodeIllegalWrite,
		"Log emitted from contract %v, but current call should be log-free", log.Address)
}
