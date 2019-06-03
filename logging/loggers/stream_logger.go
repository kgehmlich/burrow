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

package loggers

import (
	"fmt"
	"io"
	"text/template"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/term"
	"github.com/hyperledger/burrow/logging/structure"
)

const (
	JSONFormat        = "json"
	LogfmtFormat      = "logfmt"
	TerminalFormat    = "terminal"
	defaultFormatName = TerminalFormat
)

const (
	newline = '\n'
)

type Syncable interface {
	Sync() error
}

func NewStreamLogger(writer io.Writer, format string) (log.Logger, error) {
	var logger log.Logger
	var err error
	switch format {
	case "":
		return NewStreamLogger(writer, defaultFormatName)
	case JSONFormat:
		logger = log.NewJSONLogger(writer)
	case LogfmtFormat:
		logger = log.NewLogfmtLogger(writer)
	case TerminalFormat:
		logger = term.NewLogger(writer, log.NewLogfmtLogger, func(keyvals ...interface{}) term.FgBgColor {
			switch structure.Value(keyvals, structure.ChannelKey) {
			case structure.TraceChannelName:
				return term.FgBgColor{Fg: term.DarkGreen}
			default:
				return term.FgBgColor{Fg: term.Yellow}
			}
		})
	default:
		logger, err = NewTemplateLogger(writer, format, []byte{})
		if err != nil {
			return nil, fmt.Errorf("did not recognise format '%s' as named format and could not parse as "+
				"template: %v", format, err)
		}
	}
	return log.LoggerFunc(func(keyvals ...interface{}) error {
		switch structure.Signal(keyvals) {
		case structure.SyncSignal:
			if s, ok := writer.(Syncable); ok {
				return s.Sync()
			}
			// Don't log signals
			return nil
		default:
			return logger.Log(keyvals...)
		}
	}), nil
}

func NewTemplateLogger(writer io.Writer, textTemplate string, recordSeparator []byte) (log.Logger, error) {
	tmpl, err := template.New("template-logger").Parse(textTemplate)
	if err != nil {
		return nil, err
	}
	return log.LoggerFunc(func(keyvals ...interface{}) error {
		err := tmpl.Execute(writer, structure.KeyValuesMap(keyvals))
		if err == nil {
			_, err = writer.Write(recordSeparator)
		}
		return err
	}), nil

}
