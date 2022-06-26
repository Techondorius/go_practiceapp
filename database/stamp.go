package database

import(
	"log"
	"time"
	"errors"
	"go_practiceapp/model"
)

func Stamp_create(Stamp *model.Stamps) error {
	db := Connection()
	user := model.Stamps{UsersID: Stamp.UsersID, In_datetime: Stamp.In_datetime, Up_datetime: nil}
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	log.Println(user)
	return nil
}

func Put_up_datetime(userid int, up_datetime time.Time) error {
	db := Connection()
	var stamp model.Stamps

	if result := db.Debug().Model(&stamp).Find("users_id = ? AND up_datetime IS NULL AND in_datetime BETWEEN ? AND ?", up_datetime.AddDate(0, 0, -1), up_datetime, userid).Select("up_datetime").Updates(map[string]interface{}{"up_datetime":up_datetime}); result.RowsAffected != 1{
		return errors.New("there are too many ins today")
	}

	result := db.Debug().Model(&stamp).Where("users_id = ? AND up_datetime IS NULL AND in_datetime BETWEEN ? AND ?", up_datetime.AddDate(0, 0, -1), up_datetime, userid).Select("up_datetime").Updates(map[string]interface{}{"up_datetime":up_datetime})

	if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	log.Println(stamp)
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

func Delete_stamp(id int) error{
	db := Connection()
	var stamp model.Stamps
	// var stamp model.Stamps

	// result2 := db.Model(&stamp).Where("users_id = ?", User.ID).Unscoped().Delete(&stamp)
	// log.Println("result2", result2)
	result := db.Debug().Model(&stamp).Where("id = ?", id).Delete(stamp)
	log.Println(result)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}