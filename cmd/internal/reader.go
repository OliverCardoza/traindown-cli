package internal

import (
	"fmt"
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

func validate(session *traindown.Session) error {
	if len(session.Errors) != 0 {
		errMsg := "error writing guidance to storage:"
		for _, err := range session.Errors {
			errMsg = errMsg + fmt.Sprintf("\n\t%v", err)
		}
		return fmt.Errorf(errMsg)
	}
	if session.Date.IsZero() {
		return fmt.Errorf("session missing date")
	}
	if len(session.Movements) == 0 {
		return fmt.Errorf("session missing movements")
	}
	for _, movement := range session.Movements {
		if movement.Name == "" {
			return fmt.Errorf("movement missing name")
		}
		if len(movement.Performances) == 0 {
			return fmt.Errorf("movement %s missing performances", movement.Name)
		}
		for _, performance := range movement.Performances {
			if performance.Reps == 0 {
				return fmt.Errorf("performaance missing reps")
			}
			if performance.Sets == 0 {
				return fmt.Errorf("performance missing sets")
			}
		}
	}
	return nil
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
			err = validate(session)
			if err != nil {
				return errors.Wrapf(err, "validation error in file: %s", path)
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
