package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/traindown/traindown-go"
)

// TraindownReader reads a file or directory and returns Traindown data.
type TraindownReader struct {
	suffix string
}

// Read reads the traindown files referenced in the input file path.
// If the input is a directory then the reader will traverse it and
// read all files.
func (t *TraindownReader) Read(input string) ([]*traindown.Session, error) {
	var sessions []*traindown.Session
	err := filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, t.suffix) {
			data, err := os.ReadFile(path)
			if err != nil {
				return errors.Wrapf(err, "error reading file: %s", path)
			}
			session, err := traindown.ParseByte(data)
			if err != nil {
				return errors.Wrapf(err, "error parsing file: %s", path)
			}
			sessions = append(sessions, session)
			return nil
		}
		return nil
	})
	return sessions, err
}

// NewTraindownReader returns a new TraindownReader. All files that do not have
// the provided suffix are ignored.
func NewTraindownReader(suffix string) *TraindownReader {
	return &TraindownReader{
		suffix: suffix,
	}
}
