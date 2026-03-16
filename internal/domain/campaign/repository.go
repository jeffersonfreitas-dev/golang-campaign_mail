package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Update(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetById(id string) (*Campaign, error)
	Delete(campaing *Campaign) error
}
