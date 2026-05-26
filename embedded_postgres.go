package embeddedpostgres

import (
	"errors"
	"sync"
)

var mu sync.Mutex

var (
	ErrServerNotStarted     = errors.New("server has not been started")
	ErrServerAlreadyStarted = errors.New("server is already started")
)

// EmbeddedPostgres maintains all configuration and runtime functions for maintaining the lifecycle of one Postgres process.
type EmbeddedPostgres struct {
	config              Config
	cacheLocator        CacheLocator
	remoteFetchStrategy RemoteFetchStrategy
	initDatabase        initDatabase
	createDatabase      createDatabase
	started             bool
	syncedLogger        *syncedLogger
}

// NewDatabase creates a new EmbeddedPostgres struct that can be used to start and stop a Postgres process.
// When called with no parameters it will assume a default configuration state provided by the DefaultConfig method.
// When called with parameters the first Config parameter will be used for configuration.
func NewDatabase(config ...Config) *EmbeddedPostgres { _ = "STUB: not implemented"; return nil }

func newDatabaseWithConfig(config Config) *EmbeddedPostgres { _ = "STUB: not implemented"; return nil }

// Start will try to start the configured Postgres process returning an error when there were any problems with invocation.
// If any error occurs Start will try to also Stop the Postgres process in order to not leave any sub-process running.
//
//nolint:funlen
func (ep *EmbeddedPostgres) Start() error { _ = "STUB: not implemented"; return nil }

func (ep *EmbeddedPostgres) downloadAndExtractBinary(cacheExists bool, cacheLocation string) error {
	_ = "STUB: not implemented"
	// lock to prevent collisions with duplicate downloads
	return nil
}

func (ep *EmbeddedPostgres) cleanDataDirectoryAndInit() error {
	_ = "STUB: not implemented"
	return nil
}

// Stop will try to stop the Postgres process gracefully returning an error when there were any problems.
func (ep *EmbeddedPostgres) Stop() error { _ = "STUB: not implemented"; return nil }

func encodeOptions(port uint32, parameters map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// Double-quote parameter values - they may have spaces.
// Careful: CMD on Windows uses only double quotes to delimit strings.
// It treats single quotes as regular characters.

func startPostgres(ep *EmbeddedPostgres) error { _ = "STUB: not implemented"; return nil }

func stopPostgres(ep *EmbeddedPostgres) error { _ = "STUB: not implemented"; return nil }

func ensurePortAvailable(port uint32) error { _ = "STUB: not implemented"; return nil }

func dataDirIsValid(dataDir string, version PostgresVersion) bool {
	_ = "STUB: not implemented"
	return false
}
