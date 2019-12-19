package dbmodel

type (
	Administrator struct {
		ID    int    `json:"id" gorm:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	AdministratorStore interface {
		Create(administrator *Administrator) error
		UpdateEmail(id int, email string) error
		Find(id int) (*Administrator, error)
		Delete(id int) error
	}
)
