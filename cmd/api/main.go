package main

import (
	"campainmail/internal/domain/campaign"
	"campainmail/internal/endpoints"
	"campainmail/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(endpoints.Auth)

	db := database.NewDb()

	campaignService := campaign.ServiceImpl{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Patch("/campaigns/{id}/cancel", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaigns/{id}", endpoints.HandlerError(handler.CampaignDeleteById))

	http.ListenAndServe(":3000", r)
}
