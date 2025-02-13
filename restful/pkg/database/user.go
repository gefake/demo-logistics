package database

type User struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Username   string   `json:"username" binding:"required"`
	Firstname  string   `json:"firstname" binding:"required"`
	Lastname   string   `json:"lastname" binding:"required"`
	Phone      string   `json:"phone" binding:"required"`
	Email      string   `json:"email" binding:"required"`
	PositionID uint     `json:"position_id"`
	Position   Position `gorm:"references:ID"`
	RoleID     uint     `json:"role_id"`
	Role       Role     `gorm:"references:ID"`
	Password   string   `json:"password" binding:"required"`
}

func (u *User) GetAllCargos() ([]Cargo, error) {
	var cargos []Cargo

	result := DataSource.Context.Find(&cargos, "client_id = ?", u.ID).Preload("CargoProducts").Preload("Client")

	if result.Error != nil {
		return []Cargo{}, result.Error
	}

	return cargos, nil
}

// CreateUser creates a new user in the database.
func (u *User) CreateUser(user User) (uint, error) {
	result := DataSource.Context.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

// GetUserByID retrieves a user by its ID from the database.
func (u *User) GetUserByID(userID uint) (User, error) {
	var user User
	result := DataSource.Context.First(&user, userID)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (u *User) GetByLogin(login string) (User, error) {
	var user User
	result := DataSource.Context.Where("username = ?", login).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

// UpdateUser updates a user by its ID in the database.
func (u *User) UpdateUser(userID uint, user User) error {
	result := DataSource.Context.Model(&user).Where("id = ?", userID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user by its ID from the database.
func (u *User) DeleteUser(userID uint) error {
	result := DataSource.Context.Delete(&User{}, userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
