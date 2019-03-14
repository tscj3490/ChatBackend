package teamService

import (
	"errors"

	"../../db"
	"../../model"
)

// InitService inits service
func InitService() {

}

// CreateTeam creates a team
func CreateTeam(team *model.Team) (*model.Team, error) {
	t := &model.Team{}
	if res := db.ORM.Where("name = ?", team.Name).First(&t).RecordNotFound(); !res {
		err := errors.New(team.Name + " is already registered")
		return nil, err
	}
	// Insert Data
	if err := db.ORM.Create(&team).Error; err != nil {
		return nil, err
	}
	err := db.ORM.Last(&team).Error
	return team, err
}

// ReadTeam reads a team
func ReadTeam(id uint) (*model.Team, error) {
	team := &model.Team{}
	// Read Data
	// err := db.ORM.Table("teams").Select("teams.*, companies.name as company_name").
	// 	Joins("left join companies on companies.id = teams.company_id").
	// 	First(&team, "teams.id = ?", id).Error
	err := db.ORM.First(&team, "id = ?", id).Error
	return team, err
}

// UpdateTeam reads a team
func UpdateTeam(team *model.Team) (*model.Team, error) {
	// Create change info
	err := db.ORM.Model(team).Updates(team).Error
	return team, err
}

// DeleteTeam deletes team with object id
func DeleteTeam(id uint) error {
	err := db.ORM.Delete(&model.Team{ID: id}).Error
	return err
}

// ReadTeams return teams after retreive with params
func ReadTeams(query string, offset int, count int, field string, sort int, userID uint) ([]*model.Team, int, error) {
	teams := []*model.Team{}
	totalCount := 0

	res := db.ORM
	// res = res.Table("teams").Select("teams.*, companies.name as company_name").
	// 	Joins("left join companies on companies.id = teams.company_id")

	if query != "" {
		query = "%" + query + "%"
		res = res.Where("name LIKE ?", query)
	}
	// get total count of collection with initial query
	res.Find(&teams).Count(&totalCount)

	// add page feature
	if offset != 0 || count != 0 {
		res = res.Offset(offset)
		res = res.Limit(count)
	}
	// add sort feature
	if field != "" && sort != 0 {
		if sort > 0 {
			res = res.Order(field)
		} else {
			res = res.Order(field + " desc")
		}
	}
	err := res.Find(&teams).Error

	return teams, totalCount, err
}
