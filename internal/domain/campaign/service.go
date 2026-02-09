package campaign

import (
	"campainmail/internal/contract"
	"campainmail/internal/exceptions"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := Create(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", exceptions.ErrInternal
	}

	return campaign.ID, nil
}
