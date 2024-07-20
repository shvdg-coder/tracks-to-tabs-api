package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/playlists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Database *logic.DatabaseManager

	artistsAPI      *artists.API
	artistTrackAPI  *artisttrack.API
	difficultiesAPI *difficulties.API
	endpointsAPI    *endpoints.API
	instrumentsAPI  *instruments.API
	playlistsAPI    *playlists.API
	referencesAPI   *references.API
	sessionsAPI     *sessions.API
	sourcesAPI      *sources.API
	tabsAPI         *tabs.API
	tracksAPI       *tracks.API
	trackTabAPI     *tracktab.API
	usersAPI        *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// ArtistsAPI initiates upon first use and returns the artists.API.
func (api *API) ArtistsAPI() *artists.API {
	if api.artistsAPI == nil {
		api.artistsAPI = artists.NewAPI(api.Database, api.ArtistTrackAPI(), api.TracksAPI())
	}
	return api.artistsAPI
}

// ArtistTrackAPI initiates upon first use and returns the artisttrack.API.
func (api *API) ArtistTrackAPI() *artisttrack.API {
	if api.artistTrackAPI == nil {
		api.artistTrackAPI = artisttrack.NewAPI(api.Database)
	}
	return api.artistTrackAPI
}

// DifficultiesAPI initiates upon first use and returns the difficulties.API.
func (api *API) DifficultiesAPI() *difficulties.API {
	if api.difficultiesAPI == nil {
		api.difficultiesAPI = difficulties.NewAPI(api.Database)
	}
	return api.difficultiesAPI
}

// EndpointsAPI initiates upon first use and returns the endpoints.API.
func (api *API) EndpointsAPI() *endpoints.API {
	if api.endpointsAPI == nil {
		api.endpointsAPI = endpoints.NewAPI(api.Database)
	}
	return api.endpointsAPI
}

// InstrumentsAPI initiates upon first use and returns the instruments.API.
func (api *API) InstrumentsAPI() *instruments.API {
	if api.instrumentsAPI == nil {
		api.instrumentsAPI = instruments.NewAPI(api.Database)
	}
	return api.instrumentsAPI
}

// PlaylistsAPI initiates upon first use and returns the playlists.API.
func (api *API) PlaylistsAPI() *playlists.API {
	if api.playlistsAPI == nil {
		api.playlistsAPI = playlists.NewAPI(api.Database)
	}
	return api.playlistsAPI
}

// ReferencesAPI initiates upon first use and returns the references.API.
func (api *API) ReferencesAPI() *references.API {
	if api.referencesAPI == nil {
		api.referencesAPI = references.NewAPI(api.Database)
	}
	return api.referencesAPI
}

// SessionsAPI initiates upon first use and returns the sessions.API.
func (api *API) SessionsAPI() *sessions.API {
	if api.sessionsAPI == nil {
		api.sessionsAPI = sessions.NewAPI(api.Database)
	}
	return api.sessionsAPI
}

// SourcesAPI initiates upon first use and returns the sources.API.
func (api *API) SourcesAPI() *sources.API {
	if api.sourcesAPI == nil {
		api.sourcesAPI = sources.NewAPI(api.Database)
	}
	return api.sourcesAPI
}

// TabsAPI initiates upon first use and returns the tabs.API.
func (api *API) TabsAPI() *tabs.API {
	if api.tabsAPI == nil {
		api.tabsAPI = tabs.NewAPI(api.Database)
	}
	return api.tabsAPI
}

// TracksAPI initiates upon first use and returns the tracks.API.
func (api *API) TracksAPI() *tracks.API {
	if api.tracksAPI == nil {
		api.tracksAPI = tracks.NewAPI(api.Database, api.TrackTabAPI(), api.TabsAPI())
	}
	return api.tracksAPI
}

// TrackTabAPI initiates upon first use and returns the tracktab.API.
func (api *API) TrackTabAPI() *tracktab.API {
	if api.trackTabAPI == nil {
		api.trackTabAPI = tracktab.NewAPI(api.Database)
	}
	return api.trackTabAPI
}

// UsersAPI initiates upon first use and returns the users.API.
func (api *API) UsersAPI() *users.API {
	if api.usersAPI == nil {
		api.usersAPI = users.NewAPI(api.Database)
	}
	return api.usersAPI
}
