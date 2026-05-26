package embeddedpostgres

// VersionStrategy provides a strategy that can be used to determine which version of Postgres should be used based on
// the operating system, architecture and desired Postgres version.
type VersionStrategy func() (operatingSystem string, architecture string, postgresVersion PostgresVersion)

func defaultVersionStrategy(config Config, goos, arch string, linuxMachineName func() string, shouldUseAlpineLinuxBuild func() bool) VersionStrategy {
	_ = "STUB: not implemented"
	return *new(VersionStrategy)
}

// the zonkyio/embedded-postgres-binaries project produces
// arm binaries with the following name schema:
// 32bit: arm32v6 / arm32v7
// 64bit (aarch64): arm64v8

// postgres below version 14.2 is not available for macos on arm

func linuxMachineName() string { _ = "STUB: not implemented"; return "" }

func shouldUseAlpineLinuxBuild() bool { _ = "STUB: not implemented"; return false }
