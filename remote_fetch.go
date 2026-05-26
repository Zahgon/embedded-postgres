package embeddedpostgres

import (
	"archive/zip"
	"net/http"
)

// RemoteFetchStrategy provides a strategy to fetch a Postgres binary so that it is available for use.
type RemoteFetchStrategy func() error

//nolint:funlen
func defaultRemoteFetchStrategy(remoteFetchHost string, versionStrategy VersionStrategy, cacheLocator CacheLocator) RemoteFetchStrategy {
	_ = "STUB: not implemented"
	return *new(RemoteFetchStrategy)
}

func closeBody(resp *http.Response) func() { _ = "STUB: not implemented"; return nil }

func decompressResponse(bodyBytes []byte, contentLength int64, cacheLocator CacheLocator, downloadURL string) error {
	_ = "STUB: not implemented"
	return nil

	// if the content length is not set (i.e. chunked encoding),
	// we need to use the length of the bodyBytes otherwise
	// the unzip operation will fail
}

// we have successfully found the file, return early

func decompressSingleFile(file *zip.File, cacheLocation string) error {
	_ = "STUB: not implemented"
	return nil
}

// if multiple processes attempt to extract
// to prevent file corruption when multiple processes attempt to extract at the same time
// first to a cache location, and then move the file into place.

// if anything failed before the rename then the temporary file should be cleaned up.
// if the rename was successful then there is no temporary file to remove.

// Windows cannot rename a file if is it still open.
// The file needs to be manually closed to allow the rename to happen

func errorExtractingPostgres(err error) error { _ = "STUB: not implemented"; return nil }

func errorFetchingPostgres(err error) error { _ = "STUB: not implemented"; return nil }
