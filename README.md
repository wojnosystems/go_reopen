# Overview

Hey friendo, I'm Chris. Have you ever needed to re-open a file? I have. We've all needed log files to get rotated, but never before did we have such a re-usable way to do it.

```go
package main

import (
	reopen "github.com/wojnosystems/go_reopen"
)

func main() { 
	// Open the file:
    f, err := reopen.Create("/var/log/my-service.log")
    if err != nil {
    	// do a thing with the error
    }
    _, err = f.Write( []byte("write this!") )
    
    // Close and Re-open
    err = f.ReOpen()
    if err != nil {
    	// do a thing with the error
    }
    
    _, err = f.Write( []byte("write something else!") )
    
    // CLose the file and release resources, 
    // but this cannot be re-opened any more.
    _ = f.Close()
}
```

# Copyright

Copyright © 2019 Chris Wojno. All rights reserved.

No Warranties. Use this software at your own risk.

# License

Attribution 4.0 International https://creativecommons.org/licenses/by/4.0/
