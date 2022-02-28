package csvparser

import (
	"testing"

	"github.com/betas-in/logger"
	"github.com/betas-in/utils"
)

func TestCSVParser(t *testing.T) {
	log := logger.NewLogger(3, true)
	c := NewParser(log)

	data := `id, Indonesia
	in, India`

	lines := make(chan []string, 100)
	err := c.Parse(data, lines)
	utils.Test().Nil(t, err)

	for loop := true; loop; {
		var line []string
		line, loop = <-lines
		if loop {
			utils.Test().Equals(t, 2, len(line))
		}
	}
}
