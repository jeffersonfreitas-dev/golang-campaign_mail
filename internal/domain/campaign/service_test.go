package campaign

import (
	"campainmail/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Update(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	// args := r.Called(campaign)
	return nil, nil
}

func (r *repositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), args.Error(1)
}

var (
	newCampaign = contract.NewCampaign{
		Name:      "Campanha X",
		Content:   "Body mail",
		Emails:    []string{"teste@email.com", "email@teste.com"},
		CreatedBy: "teste@email.com",
	}
	service = ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(contract.NewCampaign{})

	assert.Empty(id)
	assert.NotNil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || len(campaign.Contacts) != len(newCampaign.Emails) ||
			campaign.Content != newCampaign.Content {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repositoryMock
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	camp, _ := Create(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == camp.ID
	})).Return(camp, nil)
	service.Repository = repositoryMock

	campaign, _ := service.GetById(camp.ID)
	assert.Equal(camp.ID, campaign.ID)
	assert.Equal(camp.Name, campaign.Name)
	assert.Equal(camp.Content, campaign.Content)
	assert.Equal(camp.CreatedBy, campaign.CreatedBy)
}

func Test_GetById_ReturnErrorWhenWrong(t *testing.T) {
	assert := assert.New(t)
	camp, _ := Create(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("error"))
	service.Repository = repositoryMock

	_, err := service.GetById(camp.ID)
	assert.NotNil(err)
}
