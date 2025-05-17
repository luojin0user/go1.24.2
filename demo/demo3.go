package demo

type Conn1 interface {
	// Read reads data from the connection.
	// Read can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetReadDeadline.
}

type Conn struct {
	// constant
	conn        Conn1
	isClient    bool
	handshakeFn func(any) error // (*Conn).clientHandshake or serverHandshake

	// isHandshakeComplete is true if the connection is currently transferring
	// application data (i.e. is not currently processing a handshake).
	// isHandshakeComplete is true implies handshakeErr == nil.

	handshakeErr error   // error resulting from handshake
	vers         uint16  // TLS version
	haveVers     bool    // version has been negotiated
	config       *Config // configuration passed to constructor
	// handshakes counts the number of handshakes performed on the
	// connection so far. If renegotiation is disabled then this is either
	// zero or one.

}

type Config struct {
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.

}

func (c *Conn) clientHandshake(ctx any) (err error) {
	return err
}

func Client(conn Conn, config *Config) *Conn {
	c := &Conn{
		conn:     conn,
		config:   config,
		isClient: true,
	}
	c.handshakeFn = c.clientHandshake
	return c
}
