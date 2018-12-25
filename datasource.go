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

// The datasources package provides a generic interface for the management of
// OpenXC data sources.
//
// The datasources package must be used in combination with a datasource
// driver, some of which are included within this distribution.
package openxc

import (
	"sync"
	"fmt"
)

var (
	drivers_mutex	sync.RWMutex
	drivers		= make(map[string]DataSource_Driver)
)

type dataSourceConnector struct {
	name		string
	driver		DataSource_Driver
}

func (ds dataSourceConnector) Connect() (error) {
	return ds.driver.Open(ds.name)
}

func (ds dataSourceConnector) Disconnect() (error) {
	return ds.driver.Close()
}

type DataSource struct {
	driver		DataSource_Driver
}

type DataSource_Driver interface {
	Open(name string)(error)
	Close()(error)
	Read()(VehicleMessage, error)
}

func openDataSourceConnection(conn dataSourceConnector) *DataSource {
	ds := &DataSource{driver: conn.driver};

	err := conn.driver.Open(conn.name)
	if err != nil {
		fmt.Errorf("failed to open connection to %s", conn.name)
		return nil
	}

	return ds
}

func OpenDataSource(driverName, dataSourceName string) (*DataSource, error) {
	drivers_mutex.RLock()
	drv, ok := drivers[driverName]
	drivers_mutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("unknown driver %s", driverName);
	}

	return openDataSourceConnection(dataSourceConnector{name: dataSourceName, driver: drv}), nil
}

func (ds *DataSource) ReadDataSource() (VehicleMessage, error) {
	return ds.driver.Read()
}

func (ds *DataSource) CloseDataSource() error {
	return ds.driver.Close()
}

// RegisterDataSource makes a data source driver available under the specified
// name. Only a single registration per unique name is supported, multiple
// registrations under the same driver name will result in a panic.
func RegisterDataSource(name string, drv DataSource_Driver) {
	drivers_mutex.Lock()
	defer drivers_mutex.Unlock()

	if drv == nil {
		panic("Register datasource not defined")
	}

	if _, dup := drivers[name]; dup {
		panic("Datasource already registered")
	}

	drivers[name] = drv
}
