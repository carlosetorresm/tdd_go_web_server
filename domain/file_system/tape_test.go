package filesystem_test

import (
	"io"
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := test.CreateTempFile(t, "12345")
	defer clean()

	tape := filesystem.NewTape(file)

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
