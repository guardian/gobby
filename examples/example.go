package main

import (
	"fmt"
	"github.com/guardian/gobby"
)

func main() {
	g := gobby.NewGobby("/var/tmp/gobbyfile")
	g.Set("meep", gobby.JobStatus{"a", "done", nil})
	g.Save()

	g.Load()
	j, _ := g.Get("meep")
	fmt.Println(j)
}
