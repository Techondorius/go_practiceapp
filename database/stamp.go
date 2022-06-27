package database

import(
	"time"
	"errors"
	"go_practiceapp/model"
)

func Stamp_create(Stamp *model.Stamps) (model.Stamps, error) {
	db := Connection()
	stamp := model.Stamps{UsersID: Stamp.UsersID, In_datetime: Stamp.In_datetime, Up_datetime: nil}
	result := db.Create(&stamp)
	if result.Error != nil {
		return stamp, result.Error
	}
	return stamp, nil
}

func Put_up_datetime(userid int, up_datetime time.Time) error {
	db := Connection()
	var stamps []model.Stamps

	if result := db.Debug().Model(&stamps).Where("users_id = ? AND up_datetime IS NULL AND in_datetime BETWEEN ? AND ?", userid, up_datetime.AddDate(0, 0, -1), up_datetime).Find(&stamps); result.RowsAffected < 1{
		return errors.New("couldn't find record where up_datetime is null")
	} else if result.RowsAffected > 1{
		return errors.New("couldn't select record to edit because there are too many stamps today")
	}

	result := db.Debug().Model(&stamps).Where("id = ?", stamps[0].ID).Select("up_datetime").Updates(map[string]interface{}{"up_datetime": up_datetime})

	if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}

func Stamp_read_by_id(userid int) ([]model.Stamps, error) {
	var stamps []model.Stamps
	db := Connection()
	result := db.Where("users_id = ?", userid).Find(&stamps)
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

	result := db.Debug().Model(&stamp).Where("id = ?", id).Delete(stamp)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}