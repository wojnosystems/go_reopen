package go_reopen

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFile_ReOpen(t *testing.T) {
	// generate temp file name
	tf := tmpFileName(t)
	// filename to rotate tf into
	tf1 := tf + ".1"
	defer func() {
		_ = os.Remove(tf)
		_ = os.Remove(tf1)
	}()

	// create the re-openable file
	f, err := Create(tf)
	if err != nil {
		t.Fatal(err)
	}
	// write a message
	_, err = f.WriteString("first")
	if err != nil {
		t.Fatal(err)
	}

	// rotate the file
	_ = os.Rename(tf, tf1)

	// re-open
	err = f.ReOpen()
	if err != nil {
		t.Error(err)
	}

	// write after re-open, should be in the tf file, not tf1
	_, err = f.WriteString("second")
	if err != nil {
		t.Fatal(err)
	}
	_ = f.Close()

	// Read both files and compare to ensure that the writes
	// were performed correctly after the reopen
	first, err := ioutil.ReadFile(tf1)
	second, err := ioutil.ReadFile(tf)
	if !strings.Contains(string(first), "first") {
		t.Error(`expected first file to contain "first"`)
	}
	if !strings.Contains(string(second), "second") {
		t.Error(`expected second file to contain "second"`)
	}

	if strings.Contains(string(first), "second") {
		t.Error(`expected first file not to contain "second"`)
	}
	if strings.Contains(string(second), "first") {
		t.Error(`expected second file not to contain "first"`)
	}
}

// tmpFileName generates a temporary file
func tmpFileName(t *testing.T) string {
	tf, err := ioutil.TempFile("", "go-test-reopen-file-*")
	if err != nil {
		t.Fatal(err)
	}
	_ = tf.Close()
	return t.Name()
}
