package csvparser

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/betas-in/logger"
)

// CSVParser ...
type CSVParser interface {
	Parse(string, chan []string) error
}

type csvParser struct {
	log *logger.Logger
}

// NewParser ...
func NewParser(log *logger.Logger) CSVParser {
	return &csvParser{log: log}
}

// #TODO add elaborate test cases for csv parser
func (c *csvParser) Parse(body string, lines chan []string) error {
	csvReader := csv.NewReader(strings.NewReader(body))
	csvReader.LazyQuotes = true
	defer close(lines)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.log.Error("csvparser.parse").Msgf("%+v", err)
			return err
		}
		lines <- line
	}
	return nil
}
