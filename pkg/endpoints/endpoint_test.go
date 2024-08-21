package endpoints

import (
	"testing"
)

// TestCase represents a single test case for NewLink functionality
type TestCase struct {
	name         string
	endpoint     *Endpoint
	replacements map[string]string
	want         string
}

// Run the tests
func TestNewLink(t *testing.T) {
	endpoints := createEndpoints()
	tests := createTestCases(endpoints)

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runTest(t, test)
		})
	}
}

// createEndpoints Creates dummy endpoints.
func createEndpoints() []*Endpoint {
	return []*Endpoint{
		{SourceID: 1000, Category: "artist", Type: "web", UnformattedURL: "https://test.com/artist/{artistID}"},
		{SourceID: 1000, Category: "track", Type: "web", UnformattedURL: "https://test.com/track/{trackID}"},
		{SourceID: 2000, Category: "artist", Type: "web", UnformattedURL: "https://test.com/{artistID}/tabs"},
		{SourceID: 2000, Category: "artist", Type: "api", UnformattedURL: "https://test.com/api/{artistID}/tabs?from={from}&size={size}"},
	}
}

// createTestCases Creates the test cases.
func createTestCases(endpoints []*Endpoint) []TestCase {
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
func runTest(t *testing.T, testCase TestCase) {
	link := testCase.endpoint.CreateLink(testCase.replacements)
	if link != testCase.want {
		t.Errorf("CreateLink() for endpoint %v got = %v, want = %v", testCase.endpoint, link, testCase.want)
	}
}
