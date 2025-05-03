package filesystem

import (
	"io"
	"os"
)

type Tape struct {
	file *os.File
}

func NewTape(file *os.File) *Tape {
	return &Tape{file: file}
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
