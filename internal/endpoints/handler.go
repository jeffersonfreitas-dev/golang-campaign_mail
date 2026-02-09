package endpoints

import "campainmail/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
