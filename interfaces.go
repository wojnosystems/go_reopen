package go_reopen

import "io"

// ReOpener is an interface type that says this
// object supports re-opening some resource
type ReOpener interface {
	// Signal that the file need to be closed and re-opened
	ReOpen() error
}

// Below are convenience interfaces for mixing and matching ReOpeners
type Closer interface {
	ReOpener
	io.Closer
}

type Reader interface {
	ReOpener
	io.Reader
}

type Writer interface {
	ReOpener
	io.Writer
}

type ReadWriter interface {
	ReOpener
	io.ReadWriter
}

type ReadCloser interface {
	ReOpener
	io.ReadCloser
}

type WriteCloser interface {
	ReOpener
	io.WriteCloser
}

type ReadWriteCloser interface {
	ReOpener
	io.ReadWriteCloser
}
