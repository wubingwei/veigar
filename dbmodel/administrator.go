package dbmodel

type (
	Administrator struct {
		ID    int    `gorm:"id"`
		Name  string `gorm:"name"`
		Email string `gorm:"email"`
	}

	AdministratorStore interface {
		Create(administrator *Administrator) error
		UpdateEmail(id int, email string) error
		Find(id int) (*Administrator, error)
		Delete(id int) error
	}
)
