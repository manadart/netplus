// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2020 Joseph Phillips. All rights reserved.

package netplus

import (
	"errors"
	"net"
)

var (
	errInvalidInterface         = errors.New("invalid network interface")
	//errInvalidInterfaceIndex    = errors.New("invalid network interface index")
	//errInvalidInterfaceName     = errors.New("invalid network interface name")
	errNoSuchInterface          = errors.New("no such network interface")
	//errNoSuchMulticastInterface = errors.New("no such multicast network interface")
)

// Interface represents a mapping between network interface name
// and index. It also represents network interface facility
// information.
type Interface struct {
	*net.Interface
}

// Addrs returns a list of unicast interface addresses for a specific
// interface.
func (ifi *Interface) Addrs() ([]Addr, error) {
	if ifi == nil {
		return nil, &net.OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: errInvalidInterface}
	}
	ifat, err := interfaceAddrTable(ifi)
	if err != nil {
		err = &net.OpError{Op: "route", Net: "ip+net", Source: nil, Addr: nil, Err: err}
	}
	return ifat, err
}

func interfaceByIndex(ift []Interface, index int) (*Interface, error) {
	for _, ifi := range ift {
		if index == ifi.Index {
			return &ifi, nil
		}
	}
	return nil, errNoSuchInterface
}
