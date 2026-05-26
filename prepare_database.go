package embeddedpostgres

import (
	"io"
	"os"

	"github.com/lib/pq"
)

const (
	fmtCloseDBConn = "unable to close database connection: %w"
	fmtAfterError  = "%v happened after error: %w"
)

type initDatabase func(binaryExtractLocation, runtimePath, pgDataDir, username, password, locale string, encoding string, logger *os.File) error
type createDatabase func(port uint32, username, password, database string) error

func defaultInitDatabase(binaryExtractLocation, runtimePath, pgDataDir, username, password, locale string, encoding string, logger *os.File) error {
	_ = "STUB: not implemented"
	return nil
}

// we want to preserve the original error

func createPasswordFile(runtimePath, password string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func defaultCreateDatabase(port uint32, username, password, database string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// connectionClose closes the database connection and handles the error of the function that used the database connection
func connectionClose(db io.Closer, err error) error { _ = "STUB: not implemented"; return nil }

func healthCheckDatabaseOrTimeout(config Config) error { _ = "STUB: not implemented"; return nil }

func healthCheckDatabase(port uint32, database, username, password string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func openDatabaseConnection(port uint32, username string, password string, database string) (*pq.Connector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func errorCustomDatabase(database string, err error) error { _ = "STUB: not implemented"; return nil }
