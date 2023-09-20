package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
	"github.com/spudtrooper/strunpack"
)

var (
	fileSpecRegex  = regexp.MustCompile(`^([^:]+):(\d+)`)
	fileSpecRegex2 = regexp.MustCompile(`^(?P<File>[^:]+):(?P<Line>\d+)`)
)

type displaySpec struct {
	File string
	Line int
}

func getDisplaySpec(spec string, res *displaySpec) error {
	if err := strunpack.Unpack(spec, fileSpecRegex2, res); err != nil {
		return err
	}
	if res.Line <= 0 {
		return errors.Errorf("invalid line number: %d", res.Line)
	}
	return nil
}

//go:generate genopts --prefix=ProcessFile linesAbove:int:10 linesBelow:int:10 numberLines fromStart toEnd keepGoing
func ProcessFile(spec string, optss ...ProcessFileOption) error {
	opts := MakeProcessFileOptions(optss...)

	var ds displaySpec
	if err := getDisplaySpec(spec, &ds); err != nil {
		return err
	}
	f, displayedLine := ds.File, ds.Line

	file, err := os.Open(f)
	if err != nil {
		return errors.Errorf("error opening file: %s: %v", f, err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for curLine := 1; s.Scan(); curLine++ {
		line := s.Text()

		if !((opts.FromStart() && curLine < displayedLine) ||
			(opts.ToEnd() && curLine > displayedLine) ||
			(curLine >= displayedLine-opts.LinesAbove() && curLine <= displayedLine+opts.LinesBelow())) {
			continue
		}

		// Highlight current line?
		if curLine == displayedLine {
			line = fmt.Sprintf("\033[1;3;33m%s\033[0m", line)
		}

		// Prefix the line with the line number?
		if opts.NumberLines() {
			if curLine == displayedLine {
				line = fmt.Sprintf("\033[1;3;33m%4d:\033[0m %s", curLine, line)
			} else {
				line = fmt.Sprintf("%4d: %s", curLine, line)
			}
		}

		fmt.Println(line)
	}

	if err := s.Err(); err != nil {
		return errors.Errorf("error scanning file: %s: %v", f, err)
	}

	return nil
}
