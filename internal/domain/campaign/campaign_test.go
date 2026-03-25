package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Campaingn X"
	createdBy = "teste@email.com"
	content   = "Body's"
	contacts  = []string{"teste@email.com"}
	fake      = faker.New()
)

func Test_Create_Successfuly(t *testing.T) {
	assert := assert.New(t)

	campaign, err := Create(name, content, contacts, createdBy)

	assert.Nil(err)
	assert.NotNil(campaign)
	assert.Equal(campaign.Name, name, "name should be equal")
	assert.Equal(campaign.Content, content, "content should be equal")
	assert.Equal(campaign.CreatedBy, createdBy)
	assert.NotEmpty(campaign.Contacts)
	assert.Equal(len(campaign.Contacts), 1, "campaign contacts size should be 1")
}

func Test_Create_IDNotNil(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, content, contacts, createdBy)

	assert.Nil(err)
	assert.NotNil(result)
	assert.NotNil(result.ID)
}

func Test_Create_NewCampaignStatusPending(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, content, contacts, createdBy)

	assert.Nil(err)
	assert.NotNil(result)
	assert.Equal(result.Status, Pending)
}

func Test_Create_NameValidateMin(t *testing.T) {
	assert := assert.New(t)

	campaign, err := Create("", content, contacts, createdBy)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Name deve conter no mínimo 5 caracteres")
}

func Test_Create_NameValidateMax(t *testing.T) {
	assert := assert.New(t)

	campaign, err := Create(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Nil(campaign)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Name deve conter no máximo 25 caracteres")
}

func Test_Create_ContentValidateMin(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, "", contacts, createdBy)

	assert.Nil(result)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Content deve conter no mínimo 5 caracteres")
}

func Test_Create_ContentValidateMax(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, fake.Lorem().Text(1050), contacts, createdBy)

	assert.Nil(result)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Content deve conter no máximo 1024 caracteres")
}

func Test_Create_ContactsValidateEmail(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, content, []string{"email_invalid"}, createdBy)

	assert.Nil(result)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Email deve ser válido")
}

func Test_Create_ContactsValidateMin(t *testing.T) {
	assert := assert.New(t)

	result, err := Create(name, content, nil, createdBy)

	assert.Nil(result)
	assert.NotNil(err)
	assert.Equal(err.Error(), "Contacts deve conter no mínimo 1 caracteres")
}

func Test_Create_CreateOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, err := Create(name, content, contacts, createdBy)

	assert.Nil(err)
	assert.NotNil(campaign)
	assert.Greater(campaign.CreatedOn, now)
}
