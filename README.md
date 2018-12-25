# go-openxc

[![godoc](https://godoc.org/github.com/adaptant-labs/go-openxc?status.svg)](http://godoc.org/github.com/adaptant-labs/go-openxc)
[![Build Status](https://travis-ci.com/adaptant-labs/go-openxc.svg?branch=master)](https://travis-ci.com/adaptant-labs/go-openxc)
[![Go Report Card](https://goreportcard.com/report/github.com/adaptant-labs/go-openxc)](https://goreportcard.com/report/github.com/adaptant-labs/go-openxc)

go-openxc is a simple Go library for working with the [OpenXC](https://github.com/openxc) Platform.

<img src="https://raw.githubusercontent.com/adaptant-labs/go-openxc/master/go-openxc-logo.png" width="200">

This implementation is far less mature than the official libraries and bindings
provided by the OpenXC project, but is a first step in enabling the use of
OpenXC in Go applications.

The interfaces have further been designed with extensibility in mind, allowing
for additional data sources to be added, as well as to allow OpenXC messaging
and vehicle information embedding in more complex applications while remaining
compliant with the OpenXC messaging specification.

## Logo

The go-openxc gopher logo is a modification of the gopher vector image by Takuya Ueda ([https://twitter.com/tenntenn](https://github.com/golang-samples/gopher-vector/blob/master/README.md#gopher)), superimposed with the OpenXC logo from the OpenXC Platform Media Kit ([http://openxcplatform.com/overview/media.html](http://openxcplatform.com/overview/media.html)). The resulting work is licensed under Creative Commons Attribution 4.0 International (CC BY 4.0).

<a rel="license" href="http://creativecommons.org/licenses/by/4.0/deed.en">
	<img alt="Creative Commons licensing" style="border-width:0" src="http://i.creativecommons.org/l/by/4.0/88x31.png" />
</a>
