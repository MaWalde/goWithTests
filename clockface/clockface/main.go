package main

import (
	"os"
	"time"

	"github.com/MaWalde/goWithTests.git/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
