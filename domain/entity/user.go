package entity

type User struct {
	ID        uint64   `gorm:"primaryKey" json:"id,omitempty"`
	Name      string   `gorm:"not null" json:"name,omitempty" validate:"nonzero"`
	Email     string   `gorm:"not null;unique" json:"email,omitempty"`
	Password  string   `gorm:"not null" json:"password,omitempty"`
	BirthDate string   `gorm:"not null" json:"birthDate,omitempty"`
	Age       int      `json:"age,omitempty"`
	Profile   string   `gorm:"not null" json:"Profile,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
