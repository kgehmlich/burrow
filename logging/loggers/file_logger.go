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
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/hyperledger/burrow/logging/structure"
)

type FileLogger struct {
	path         string
	file         *os.File
	formatName   string
	streamLogger log.Logger
}

type FileTemplateParams struct {
	Date time.Time
}

func NewFileTemplateParams() *FileTemplateParams {
	return &FileTemplateParams{
		Date: time.Now(),
	}
}

const timeFormat = "2006-01-02_15h04m05s"

func (ftp *FileTemplateParams) Timestamp() string {
	return ftp.Date.Format(timeFormat)
}

func NewFileLogger(path string, formatName string) (*FileLogger, error) {
	tmpl, err := template.New("file-logger").Parse(path)
	if err != nil {
		return nil, fmt.Errorf("could not parse path string '%s' as Go text/template: %v", path, err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, NewFileTemplateParams())
	if err != nil {
		return nil, err
	}
	fl := &FileLogger{
		path:       buf.String(),
		formatName: formatName,
	}
	err = fl.Reload()
	if err != nil {
		return nil, err
	}
	return fl, nil
}

func (fl *FileLogger) Log(keyvals ...interface{}) error {
	switch structure.Signal(keyvals) {
	case structure.SyncSignal:
		return fl.file.Sync()
	case structure.ReloadSignal:
		return fl.Reload()
	default:
		return fl.streamLogger.Log(keyvals...)
	}
}

func (fl *FileLogger) Reload() error {
	if fl.file != nil {
		err := fl.file.Close()
		if err != nil {
			return fmt.Errorf("could not close file %v: %v", fl.file, err)
		}
	}
	var err error
	fl.file, err = os.OpenFile(fl.path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	fl.streamLogger, err = NewStreamLogger(fl.file, fl.formatName)
	if err != nil {
		return err
	}
	return nil
}
