// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package lib

//go:generate genopts --prefix=ProcessFile --outfile=/Users/jeff/Projects/spudtrooper/showline/lib/processfileoptions.go "linesAbove:int:10" "linesBelow:int:10" "numberLines" "fromStart" "toEnd" "keepGoing"

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type ProcessFileOption struct {
	f func(*processFileOptionImpl)
	s string
}

func (o ProcessFileOption) String() string { return o.s }

type ProcessFileOptions interface {
	FromStart() bool
	HasFromStart() bool
	KeepGoing() bool
	HasKeepGoing() bool
	LinesAbove() int
	HasLinesAbove() bool
	LinesBelow() int
	HasLinesBelow() bool
	NumberLines() bool
	HasNumberLines() bool
	ToEnd() bool
	HasToEnd() bool
}

func ProcessFileFromStart(fromStart bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_fromStart = true
		opts.fromStart = fromStart
	}, fmt.Sprintf("lib.ProcessFileFromStart(bool %+v)", fromStart)}
}
func ProcessFileFromStartFlag(fromStart *bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if fromStart == nil {
			return
		}
		opts.has_fromStart = true
		opts.fromStart = *fromStart
	}, fmt.Sprintf("lib.ProcessFileFromStart(bool %+v)", fromStart)}
}

func ProcessFileKeepGoing(keepGoing bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_keepGoing = true
		opts.keepGoing = keepGoing
	}, fmt.Sprintf("lib.ProcessFileKeepGoing(bool %+v)", keepGoing)}
}
func ProcessFileKeepGoingFlag(keepGoing *bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if keepGoing == nil {
			return
		}
		opts.has_keepGoing = true
		opts.keepGoing = *keepGoing
	}, fmt.Sprintf("lib.ProcessFileKeepGoing(bool %+v)", keepGoing)}
}

func ProcessFileLinesAbove(linesAbove int) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_linesAbove = true
		opts.linesAbove = linesAbove
	}, fmt.Sprintf("lib.ProcessFileLinesAbove(int %+v)", linesAbove)}
}
func ProcessFileLinesAboveFlag(linesAbove *int) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if linesAbove == nil {
			return
		}
		opts.has_linesAbove = true
		opts.linesAbove = *linesAbove
	}, fmt.Sprintf("lib.ProcessFileLinesAbove(int %+v)", linesAbove)}
}

func ProcessFileLinesBelow(linesBelow int) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_linesBelow = true
		opts.linesBelow = linesBelow
	}, fmt.Sprintf("lib.ProcessFileLinesBelow(int %+v)", linesBelow)}
}
func ProcessFileLinesBelowFlag(linesBelow *int) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if linesBelow == nil {
			return
		}
		opts.has_linesBelow = true
		opts.linesBelow = *linesBelow
	}, fmt.Sprintf("lib.ProcessFileLinesBelow(int %+v)", linesBelow)}
}

func ProcessFileNumberLines(numberLines bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_numberLines = true
		opts.numberLines = numberLines
	}, fmt.Sprintf("lib.ProcessFileNumberLines(bool %+v)", numberLines)}
}
func ProcessFileNumberLinesFlag(numberLines *bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if numberLines == nil {
			return
		}
		opts.has_numberLines = true
		opts.numberLines = *numberLines
	}, fmt.Sprintf("lib.ProcessFileNumberLines(bool %+v)", numberLines)}
}

func ProcessFileToEnd(toEnd bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		opts.has_toEnd = true
		opts.toEnd = toEnd
	}, fmt.Sprintf("lib.ProcessFileToEnd(bool %+v)", toEnd)}
}
func ProcessFileToEndFlag(toEnd *bool) ProcessFileOption {
	return ProcessFileOption{func(opts *processFileOptionImpl) {
		if toEnd == nil {
			return
		}
		opts.has_toEnd = true
		opts.toEnd = *toEnd
	}, fmt.Sprintf("lib.ProcessFileToEnd(bool %+v)", toEnd)}
}

type processFileOptionImpl struct {
	fromStart       bool
	has_fromStart   bool
	keepGoing       bool
	has_keepGoing   bool
	linesAbove      int
	has_linesAbove  bool
	linesBelow      int
	has_linesBelow  bool
	numberLines     bool
	has_numberLines bool
	toEnd           bool
	has_toEnd       bool
}

func (p *processFileOptionImpl) FromStart() bool      { return p.fromStart }
func (p *processFileOptionImpl) HasFromStart() bool   { return p.has_fromStart }
func (p *processFileOptionImpl) KeepGoing() bool      { return p.keepGoing }
func (p *processFileOptionImpl) HasKeepGoing() bool   { return p.has_keepGoing }
func (p *processFileOptionImpl) LinesAbove() int      { return or.Int(p.linesAbove, 10) }
func (p *processFileOptionImpl) HasLinesAbove() bool  { return p.has_linesAbove }
func (p *processFileOptionImpl) LinesBelow() int      { return or.Int(p.linesBelow, 10) }
func (p *processFileOptionImpl) HasLinesBelow() bool  { return p.has_linesBelow }
func (p *processFileOptionImpl) NumberLines() bool    { return p.numberLines }
func (p *processFileOptionImpl) HasNumberLines() bool { return p.has_numberLines }
func (p *processFileOptionImpl) ToEnd() bool          { return p.toEnd }
func (p *processFileOptionImpl) HasToEnd() bool       { return p.has_toEnd }

func makeProcessFileOptionImpl(opts ...ProcessFileOption) *processFileOptionImpl {
	res := &processFileOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeProcessFileOptions(opts ...ProcessFileOption) ProcessFileOptions {
	return makeProcessFileOptionImpl(opts...)
}
