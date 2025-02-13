package database

type Role struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
}

func (r *Role) CreateRole(role Role) (uint, error) {
	db := DataSource.Context
	result := db.Create(role)
	if result.Error != nil {
		return 0, result.Error
	}
	return r.ID, nil
}

func (r *Role) GetRoleByID(roleID uint) (Role, error) {
	var role Role
	result := DataSource.Context.First(&role, roleID)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, nil
}

func (r *Role) UpdateRole(roleID uint, newRole Role) error {
	result := DataSource.Context.Model(r).Where("id = ?", roleID).Updates(newRole)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Role) GetRoleByName(roleName string) (Role, error) {
	var role Role
	result := DataSource.Context.Where("name = ?", roleName).First(&role)
	if result.Error != nil {
		return Role{}, result.Error
	}
	return role, nil
}

func (r *Role) DeleteRole(roleID uint) error {
	result := DataSource.Context.Delete(r, roleID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
