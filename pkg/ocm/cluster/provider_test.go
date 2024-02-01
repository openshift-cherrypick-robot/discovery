// Copyright Contributors to the Open Cluster Management project

package cluster

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	getRequestFunc func(*http.Request) (*http.Response, error)
)

// Mocking the ClusterGetInterface
type getClientMock struct{}

func (cm *getClientMock) Get(request *http.Request) (*http.Response, error) {
	return getRequestFunc(request)
}

// When the everything is good
func TestProviderGetClustersNoError(t *testing.T) {
	getRequestFunc = func(*http.Request) (*http.Response, error) {
		file, err := os.Open("testdata/ocm_mock.json")
		if err != nil {
			t.Error(err)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(file),
		}, nil
	}
	httpClient = &getClientMock{} //without this line, the real api is fired

	response, err := ClusterProvider.GetClusters(ClusterRequest{})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, len(response.Items))
}
