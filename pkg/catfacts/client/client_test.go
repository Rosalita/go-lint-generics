package catfacts_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	catfacts "github.com/Rosalita/go-lint-generics/pkg/catfacts/client"

	"github.com/stretchr/testify/assert"
)

// mockCatFactHandler mocks the behaviour of the real catfacts API and is used for testing.
func mockCatFactHandler(rw http.ResponseWriter, req *http.Request) {
	switch {
	case req.URL.String() == "/fact" && req.Method == "GET":
		responseJSON := `{ 
			"fact": "Cats eat grass to aid their digestion and to help them get rid of any fur in their stomachs.",
			"length": 92
			}`
		rw.WriteHeader(200)
		rw.Write([]byte(responseJSON))

	case req.URL.String() == "/foo" && req.Method == "GET":
		responseJSON := `{"message":"Not Found","code":404}`
		rw.WriteHeader(404)
		rw.Write([]byte(responseJSON))

	default:
		rw.WriteHeader(500)
	}
}

func TestGetCatFact(t *testing.T) {
	tests := []struct {
		path           string
		expectedResult *catfacts.CatFact
		expectedErr    error
	}{
		{
			path: "/fact",
			expectedResult: &catfacts.CatFact{
				Fact:   "Cats eat grass to aid their digestion and to help them get rid of any fur in their stomachs.",
				Length: 92,
			},
		},
		{
			path:           "/foo",
			expectedErr:    errors.New("status code not ok"),
		},
	}

	// Create a new client and configure it to use test server instead of the real API endpoint.
	testServer := httptest.NewServer(http.HandlerFunc(mockCatFactHandler))

	client, err := catfacts.NewClient(testServer.URL)
	assert.NoError(t, err)

	for _, test := range tests {
		ctx := context.Background()
		result, err := client.Get(ctx, test.path)
		assert.Equal(t, test.expectedResult, result)
		assert.Equal(t, test.expectedErr, err)
	}
}
