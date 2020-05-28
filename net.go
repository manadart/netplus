package netplus

import "net"

// Addr represents a network end point address.
// It extends the `Addr` interface from the standard library and includes an
// additional method for retrieving address flags.
type Addr interface {
	net.Addr
}
