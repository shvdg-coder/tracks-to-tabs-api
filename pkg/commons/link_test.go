package commons

import (
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"testing"
)

// TestCase represents a single test case for NewLink functionality
type TestCase struct {
	name         string
	endpoint     *end.Endpoint
	replacements map[string]string
	want         string
}

// Run the tests
func TestNewLink(t *testing.T) {
	sources := createSources()
	endpoints := createEndpoints()
	tests := createTestCases(endpoints)

	// Run tests
	for _, tt := range tests {
		var source *src.Source
		for _, s := range sources {
			if s.ID == tt.endpoint.SourceID {
				source = s
				break
			}
		}
		if source == nil {
			t.Errorf("No source found for endpoint %v", tt.endpoint)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			runTest(t, tt, source)
		})
	}
}

// createSources Creates dummy sources.
func createSources() []*src.Source {
	return []*src.Source{
		{ID: 1000, Name: "MusicProvider1", Category: "music"},
		{ID: 2000, Name: "TabProvider2", Category: "tabs"},
	}
}

// createEndpoints Creates dummy endpoints.
func createEndpoints() []*end.Endpoint {
	return []*end.Endpoint{
		{SourceID: 1000, Category: "artist", Type: "web", URL: "https://test.com/artist/{artistID}"},
		{SourceID: 1000, Category: "track", Type: "web", URL: "https://test.com/track/{trackID}"},
		{SourceID: 2000, Category: "artist", Type: "web", URL: "https://test.com/{artistID}/tabs"},
		{SourceID: 2000, Category: "artist", Type: "api", URL: "https://test.com/api/{artistID}/tabs?from={from}&size={size}"},
	}
}

// createTestCases Creates the test cases.
func createTestCases(endpoints []*end.Endpoint) []TestCase {
	return []TestCase{
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
			name:         "TabProvider2_ArtistTabs",
			endpoint:     endpoints[2],
			replacements: map[string]string{"{artistID}": "111213"},
			want:         "https://test.com/111213/tabs",
		},
		{
			name:         "TabProvider2_TabAPI",
			endpoint:     endpoints[3],
			replacements: map[string]string{"{artistID}": "141516", "{from}": "0", "{size}": "20"},
			want:         "https://test.com/api/141516/tabs?from=0&size=20",
		},
	}
}

// Runs the test for creating a new Link.
func runTest(t *testing.T, tt TestCase, source *src.Source) {
	link := NewLink(source, tt.endpoint, tt.replacements)
	if link.FormattedURL != tt.want {
		t.Errorf("NewLink() for endpoint %v got = %v, want = %v", tt.endpoint, link.FormattedURL, tt.want)
	}
}
