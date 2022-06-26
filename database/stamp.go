package database

import(
	"log"
	"go_practiceapp/model"
)

func Stamp_create(Stamp *model.Stamps) error {
	db := Connection()
	user := model.Stamps{UsersID: Stamp.UsersID, Type: Stamp.Type, Stamp_datetime: Stamp.Stamp_datetime}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Stamp_read_by_id(userid int) ([]model.Stamps, error) {
	var stamps []model.Stamps
	db := Connection()
	result := db.Where("users_id = ?", userid).Find(&stamps)
	log.Println(result.RowsAffected)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return stamps, result.Error
	}
}