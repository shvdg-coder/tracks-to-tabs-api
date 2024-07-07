package common

import (
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"testing"
)

// Test the creation of a new Link instance, with formatting the endpoint URL and the corresponding values.
func TestNewLink(t *testing.T) {
	sources := []*src.Source{
		{ID: 1000, Name: "MusicProvider1", Category: "music"},
		{ID: 2000, Name: "TabProvider1", Category: "tabs"},
	}

	endpoints := []*end.Endpoint{
		{SourceID: 1000, Category: "artist", Type: "web", URL: "https://test.com/artist/{artistID}"},
		{SourceID: 1000, Category: "track", Type: "web", URL: "https://test.com/track/{trackID}"},
		{SourceID: 2000, Category: "artist", Type: "web", URL: "https://test.com/a/wsa/{artistName}-tabs-a{artistID}"},
		{SourceID: 2000, Category: "tab", Type: "web", URL: "https://test.com/a/wsa/{artistName}-{trackTitle}-tab-s{trackID}"},
		{SourceID: 2000, Category: "artist", Type: "api", URL: "https://test.com/api/artist/{artistID}/songs?from={from}&size={size}"},
	}

	tests := []struct {
		name         string
		endpoint     *end.Endpoint
		replacements map[string]string
		want         string
	}{
		{
			name:         "MusicProvider1_Artist",
			endpoint:     endpoints[0],
			replacements: map[string]string{"{artistID}": "123456"},
			want:         "https://test.com/artist/123456",
		},
		{
			name:         "MusicProvider1_Track",
			endpoint:     endpoints[1],
			replacements: map[string]string{"{trackID}": "78910"},
			want:         "https://test.com/track/78910",
		},
		{
			name:         "TabProvider1_Artist",
			endpoint:     endpoints[2],
			replacements: map[string]string{"{artistID}": "111213", "{artistName}": "testArtist"},
			want:         "https://test.com/a/wsa/testArtist-tabs-a111213",
		},
		{
			name:         "TabProvider1_Tab",
			endpoint:     endpoints[3],
			replacements: map[string]string{"{trackID}": "141516", "{artistName}": "testArtist", "{trackTitle}": "testTrack"},
			want:         "https://test.com/a/wsa/testArtist-testTrack-tab-s141516",
		},
		{
			name:         "TabProvider1_API",
			endpoint:     endpoints[4],
			replacements: map[string]string{"{artistID}": "171819", "{from}": "0", "{size}": "20"},
			want:         "https://test.com/api/artist/171819/songs?from=0&size=20",
		},
	}

	for _, tt := range tests {
		var source *src.Source
		for _, s := range sources {
			if int(s.ID) == tt.endpoint.SourceID {
				source = s
				break
			}
		}
		if source == nil {
			t.Errorf("No source found for endpoint %v", tt.endpoint)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			link := NewLink(source, tt.endpoint, tt.replacements)
			if link.FormattedURL != tt.want {
				t.Errorf("NewLink() for endpoint %v got = %v, want = %v", tt.endpoint, link.FormattedURL, tt.want)
			}
		})
	}
}
