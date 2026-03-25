package campaign

import (
	"campainmail/internal/contract"
	"campainmail/internal/exceptions"
	"errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := Create(newCampaign.Name, newCampaign.Content, newCampaign.Emails, newCampaign.CreatedBy)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", exceptions.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImpl) GetById(id string) (*contract.CampaignResponse, error) {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, exceptions.ProcessErrorToReturn(err)
	}

	return &contract.CampaignResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
		CreatedBy:            campaign.CreatedBy,
	}, nil
}

func (s *ServiceImpl) Cancel(id string) error {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return exceptions.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Cancel()
	err = s.Repository.Update(campaign)

	return nil
}

func (s *ServiceImpl) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return exceptions.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("Campaign may not be deleted because its status is pendidng")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	return nil
}
