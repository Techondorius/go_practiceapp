package database

import(
	"go_practiceapp/model"
)

func Stamp_in(Stamp *model.Stamps, Type string) error {
	db := Connection()
	user := model.Stamps{UsersID: Stamp.UsersID, Type: Type, Stamp_datetime: Stamp.Stamp_datetime}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}