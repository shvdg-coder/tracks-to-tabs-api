package internal

const (
	KeyDatabaseURL                        = "DATABASE_URL"
	KeyAdminInitialEmail                  = "ADMIN_INITIAL_EMAIL"
	KeyAdminInitialPassword               = "ADMIN_INITIAL_PASSWORD"
	KeyDatabaseAllowCreatingCommand       = "DATABASE_ALLOW_CREATING"
	KeyDatabaseAllowPurgingCommand        = "DATABASE_ALLOW_PURGING"
	KeyDatabaseAllowMinimumSeedingCommand = "DATABASE_ALLOW_MINIMUM_SEEDING"
	KeyDatabaseAllowDummySeedingCommand   = "DATABASE_ALLOW_DUMMY_SEEDING"
)

const (
	ValueDatabaseDriver      = "postgres"
	CommandCreate            = "create"
	CommandPurge             = "purge"
	CommandSeed              = "seed"
	InstrumentElectricGuitar = "Electric Guitar"
	InstrumentAcousticGuitar = "Acoustic Guitar"
	InstrumentBassGuitar     = "Bass Guitar"
	InstrumentDrums          = "Drums"
	DifficultyEasy           = "Easy"
	DifficultyIntermediate   = "Intermediate"
	DifficultyHard           = "Hard"
	DifficultyExpert         = "Expert"
)
