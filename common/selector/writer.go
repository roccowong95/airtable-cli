package selector

import (
	"bytes"
	"strings"
)

type multilineWriter struct {
	b     *strings.Builder
	lines []string
}

func newMultilineWriter() *multilineWriter {
	var ret multilineWriter
	ret.b = &strings.Builder{}
	return &ret
}

func (w *multilineWriter) Write(p []byte) (n int, e error) {
	for {
		idx := bytes.IndexByte(p, '\n')
		if -1 == idx {
			return w.b.Write(p)
		}
		w.b.Write(p[:idx])
		w.lines = append(w.lines, w.b.String())
		w.b.Reset()
		if idx == len(p)-1 {
			break
		}
		if idx < len(p)-1 {
			p = p[idx+1:]
		}
	}
	return len(p), nil
}

func (w *multilineWriter) Dump() []string {
	if w.b.Len() != 0 {
		w.lines = append(w.lines, w.b.String())
		w.b.Reset()
	}
	return w.lines
}
