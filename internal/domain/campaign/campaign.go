package campaign

import (
	"campainmail/internal/exceptions"
	"time"

	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=25"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func Create(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for i, value := range emails {
		contacts[i] = Contact{Email: value}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := exceptions.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}

	return nil, err
}
