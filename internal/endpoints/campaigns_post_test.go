package endpoints

import (
	"bytes"
	"campainmail/internal/contract"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func Test_CampaignsPost_shouldSaveNewCampaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)

	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hello everybody",
		Emails:  []string{"teste@email.com"},
	}

	handler := Handler{CampaignService: service}
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("3232", nil)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(201, status)
	assert.Nil(err)
}

func Test_CampaignsPost_should_throw_error(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)

	body := contract.NewCampaign{
		Name:    "teste",
		Content: "Hello everybody",
		Emails:  []string{"teste@email.com"},
	}

	handler := Handler{CampaignService: service}
	service.On("Create", mock.Anything).Return("", fmt.Errorf("Error"))

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
	assert.Equal(err.Error(), "Error")
}
