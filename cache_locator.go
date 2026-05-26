package embeddedpostgres

// CacheLocator retrieves the location of the Postgres binary cache returning it to location.
// The result of whether this cache is present will be returned to exists.
type CacheLocator func() (location string, exists bool)

func defaultCacheLocator(cacheDirectory string, versionStrategy VersionStrategy) CacheLocator {
	_ = "STUB: not implemented"
	return *new(CacheLocator)
}
