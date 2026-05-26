package embeddedpostgres

import (
	"io"
	"os"
)

type syncedLogger struct {
	offset int64
	logger io.Writer
	file   *os.File
}

func newSyncedLogger(dir string, logger io.Writer) (*syncedLogger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *syncedLogger) flush() error { _ = "STUB: not implemented"; return nil }

func readLogsOrTimeout(logger *os.File) (logContent []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
