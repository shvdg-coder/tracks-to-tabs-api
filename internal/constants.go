package internal

const (
	KeyDatabaseURL                         = "DATABASE_URL"
	KeyAdminInitialEmail                   = "ADMIN_INITIAL_EMAIL"
	KeyAdminInitialPassword                = "ADMIN_INITIAL_PASSWORD"
	KeyDatabaseEnableCreatingCommand       = "DATABASE_ENABLE_CREATING"
	KeyDatabaseEnablePurgingCommand        = "DATABASE_ENABLE_PURGING"
	KeyDatabaseEnableMinimumSeedingCommand = "DATABASE_ENABLE_MINIMUM_SEEDING"
	KeyDatabaseEnableDummySeedingCommand   = "DATABASE_ENABLE_DUMMY_SEEDING"
)

const (
	PathConfig = "config.yaml"
)

const (
	ValueDatabaseDriver = "postgres"
	CommandCreate       = "create"
	CommandPurge        = "purge"
	CommandSeed         = "seed"
	CategoryMusic       = "music"
	CategoryTabs        = "tabs"
	CategoryTab         = "tab"
	CategoryArtist      = "artist"
	CategoryTrack       = "track"
)
