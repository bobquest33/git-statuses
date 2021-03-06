package util

import (
	"io"
	"strings"
)

type SimpleLabeledWriter struct {
	label  string
	base   io.Writer
	output bool
}

func NewSimpleLabeledWriter(label string, base io.Writer) io.WriteCloser {
	writer := &SimpleLabeledWriter{
		base: base,
	}
	writer.SetLabel(label)
	return writer
}

func (w *SimpleLabeledWriter) SetLabel(label string) {
	w.label = label
}

func (w *SimpleLabeledWriter) Write(p []byte) (int, error) {
	w.output = true
	str := string(p)

	w.writeCore(strings.Replace(str, newline, "", -1))

	return len(p), nil
}

func (w *SimpleLabeledWriter) Close() error {
	if w.output {
		w.writeCore(w.label + "\n")
	}
	//FIXME: if w.base is io.Closer
	return nil
}

func (w *SimpleLabeledWriter) writeCore(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}
	return w.base.Write([]byte(str))
}
