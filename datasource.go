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
	"fmt"
	"sync"
)

var (
	driversMutex sync.RWMutex
	drivers       = make(map[string]DataSourceDriver)
)

type dataSourceConnector struct {
	name   string
	driver DataSourceDriver
}

func (ds dataSourceConnector) Connect() error {
	return ds.driver.Open(ds.name)
}

func (ds dataSourceConnector) Disconnect() error {
	return ds.driver.Close()
}

// DataSource provides a generic interface for the management of OpenXC data
// sources.
type DataSource struct {
	driver DataSourceDriver
}

// DataSourceDriver provides the backing driver interface for each DataSource.
type DataSourceDriver interface {
	Open(name string) error
	Close() error
	Read() (VehicleMessage, error)
}

func openDataSourceConnection(conn dataSourceConnector) *DataSource {
	ds := &DataSource{driver: conn.driver}

	err := conn.driver.Open(conn.name)
	if err != nil {
		fmt.Errorf("failed to open connection to %s", conn.name)
		return nil
	}

	return ds
}

// OpenDataSource opens the named DataSource
func OpenDataSource(driverName, dataSourceName string) (*DataSource, error) {
	driversMutex.RLock()
	drv, ok := drivers[driverName]
	driversMutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("unknown driver %s", driverName)
	}

	return openDataSourceConnection(dataSourceConnector{name: dataSourceName, driver: drv}), nil
}

// ReadDataSource reads a single VehicleMessage from the data source stream.
// This may be called multiple times to advance across the stream, and will
// return with an EOF when the stream runs out.
func (ds *DataSource) ReadDataSource() (VehicleMessage, error) {
	return ds.driver.Read()
}

// CloseDataSource closes the data source. Each caller that has opened the data
// source is responsible for closing it directly.
func (ds *DataSource) CloseDataSource() error {
	return ds.driver.Close()
}

// RegisterDataSource makes a data source driver available under the specified
// name. Only a single registration per unique name is supported, multiple
// registrations under the same driver name will result in a panic.
func RegisterDataSource(name string, drv DataSourceDriver) {
	driversMutex.Lock()
	defer driversMutex.Unlock()

	if drv == nil {
		panic("Register datasource not defined")
	}

	if _, dup := drivers[name]; dup {
		panic("Datasource already registered")
	}

	drivers[name] = drv
}
