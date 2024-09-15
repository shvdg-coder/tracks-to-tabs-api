package internal

const (
	KeyDatabaseURL    = "DATABASE_URL"
	KeySshUser        = "SSH_USER"
	KeySshPassword    = "SSH_PASSWORD"
	KeySshServer      = "SSH_SERVER"
	KeySshDestination = "SSH_DESTINATION"
	KeySshLocalPort   = "SSH_LOCAL_PORT"
)

const (
	PathSeedConfig = "seed-config.yaml"
)

const (
	ValueDatabaseDriver = "postgres"
	CommandCreate       = "create"
	CommandPurge        = "purge"
	CommandSeed         = "seed"
)
