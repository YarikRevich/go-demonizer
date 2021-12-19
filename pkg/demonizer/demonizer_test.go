package demonizer

import (
	"testing"

	"github.com/franela/goblin"
)

func TestDemonize(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Testing demonizer", func() {
		g.It("Test demonize function", func() {
			DemonizeProcess("sleep", "500")
		})
	})
}
