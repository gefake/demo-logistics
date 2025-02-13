package database

type Position struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (p *Position) CreatePosition(position Position) (uint, error) {
	db := DataSource.Context
	result := db.Create(&position)
	if result.Error != nil {
		return 0, result.Error
	}
	return position.ID, nil
}

func (p *Position) GetPositionByID(positionID uint) (Position, error) {
	db := DataSource.Context
	var position Position
	result := db.First(&position, positionID)
	if result.Error != nil {
		return Position{}, result.Error
	}
	return position, nil
}

func (p *Position) UpdatePosition(positionID uint, position Position) error {
	db := DataSource.Context
	result := db.Model(&position).Where("id = ?", positionID).Updates(position)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Position) GetPositionByName(name string) (Position, error) {
	db := DataSource.Context
	var position Position
	result := db.Where("name = ?", name).First(&position)
	if result.Error != nil {
		return Position{}, result.Error
	}
	return position, nil
}

func (p *Position) DeletePosition(positionID uint) error {
	db := DataSource.Context
	result := db.Delete(&Position{}, positionID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
