package main

import (
	"flag"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/showline/lib"
)

var (
	linesAbove  = flag.Int("a", 10, "Number of lines above the specified line")
	linesBelow  = flag.Int("b", 10, "Number of lines below the specified line")
	numberLines = flag.Bool("n", false, "Number the output lines")
	fromStart   = flag.Bool("s", false, "Start from the beginning of the file")
	toEnd       = flag.Bool("e", false, "Go to the end of the file")
	keepGoing   = flag.Bool("k", false, "Keep going even if there are errors")
)

func realMain() error {
	for _, spec := range flag.Args() {
		if err := lib.ProcessFile(spec,
			lib.ProcessFileLinesAboveFlag(linesAbove),
			lib.ProcessFileLinesBelowFlag(linesBelow),
			lib.ProcessFileNumberLinesFlag(numberLines),
			lib.ProcessFileFromStartFlag(fromStart),
			lib.ProcessFileToEndFlag(toEnd),
			lib.ProcessFileKeepGoingFlag(keepGoing),
		); err != nil {
			if *keepGoing {
				fmt.Printf("error processing file: %s: %v\n", spec, err)
				continue
			}
			return errors.Errorf("error processing file: %s: %v", spec, err)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	check.Err(realMain())
}
