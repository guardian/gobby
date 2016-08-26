# gobby

gobby is a simple wrapper around Golang gob for tracking the status of jobs

---

## Usage

```go
package main

import (
	"fmt"
	"github.com/guardian/gobby"
)

func main() {
	g := gobby.New("/var/tmp/gobbyfile")
	g.Set("meep", gobby.JobStatus{"a", "done", nil})
	g.Save()

	g.Load()
	j, _ := g.Get("meep")
	fmt.Println(j)
}
```
