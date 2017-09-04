// Invokes the perigord driver application

package main

// Relative import here forces inclusion of our migrations
import (
	_ "{{.project}}/migrations"
	"github.com/swarmdotmarket/perigord/stub"
)

func main() {
	stub.StubMain()
}