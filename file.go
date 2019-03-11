// go_reopen Copyright © 2019 Chris Wojno. All rights reserved.

package go_reopen

import "os"

// File implements a file the can be ReOpened
// This structure embeds the *os.File, so all methods are common to both
// When reopen.OpenFile is called, it creates a File object that stores
// the parameters used when creating the file. These are then
// replayed later when ReOpen is called.
type File struct {
	// File is the file handle
	*os.File

	// mode is just like with os.OpenFile
	mode os.FileMode

	// flags are just like os.OpenFile
	flags int
}

// OpenFile creates a new reopen.File which functions just
// like an os.File, but the original values used to create
// the reopen.File are stored for later use with ReOpen().
// @see os.OpenFile for the full description of the parameters
// Note: os.NewFile function don't exist because we need
//   the file path, flags, and modes to re-open the file.
//
// @param name the path to the file to open
// @param flags used to create the file, if missing.
//   os.O_CREATE is automatically added to ReOpen calls
// @param mode are the permissions and other flags. Log files
//   usually use 0600 to allow the owner to read & write,
//   while nobody else can read, but you can also do 0644 to
//   allow any one to read the file as well
func OpenFile(name string, flags int, mode os.FileMode) (f *File, err error) {
	f = &File{
		mode:  mode,
		flags: flags, // Always create the file when re-opening
	}
	f.File, err = os.OpenFile(name, flags, mode)
	return
}

// Create works just like os.Create, but returns a reopen.File
func Create(name string) (f *File, err error) {
	return OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// ReOpen performs the file close and re-open.
// If you didn't specify os.O_CREATE when creating the
// reopen.File, it will be added for you to prevent
// getting errors when opening files which may no longer
// exist, as with log rotation.
// @return err if any errors were encountered while closing or opening the file
func (f *File) ReOpen() (err error) {
	err = f.File.Close()
	if err != nil {
		return
	}
	f.File, err = os.OpenFile(f.File.Name(), f.flags|os.O_CREATE, f.mode)
	return
}
