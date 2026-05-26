package embeddedpostgres

import (
	"archive/tar"
	"io"

	"github.com/xi2/xz"
)

func defaultTarReader(xzReader *xz.Reader) (func() (*tar.Header, error), func() io.Reader) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decompressTarXz(tarReader func(*xz.Reader) (func() (*tar.Header, error), func() io.Reader), path, extractPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func errorUnableToExtract(cacheLocation, binariesPath string, err error) error {
	_ = "STUB: not implemented"
	return nil
}
