package database

import (
	"campainmail/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}

func (c *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	for _, item := range c.campaigns {
		if item.ID == id {
			return &item, nil
		}
	}

	return nil, errors.New("Campanha com o id informado não existe")
}
