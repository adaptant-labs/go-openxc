// Copyright 2018 Adaptant Labs
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package openxc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type traceDriver struct {
	traceFile *os.File
	dec       *json.Decoder
}

func (d *traceDriver) Open(dsn string) error {
	traceFile, err := os.Open(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d.traceFile = traceFile
	d.dec = json.NewDecoder(d.traceFile)

	return nil
}

func (d *traceDriver) Read(msg *VehicleMessage) (error) {
	err := d.dec.Decode(msg)
	if err == io.EOF {
		return err
	} else if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return nil
}

func (d *traceDriver) Reset() error {
	_, err := d.traceFile.Seek(io.SeekStart, 0)
	d.dec = json.NewDecoder(d.traceFile)
	return err
}

func (d *traceDriver) Close() error {
	return d.traceFile.Close()
}

func init() {
	RegisterDataSource("trace", &traceDriver{})
}
