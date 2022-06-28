package database

import (
	"errors"
	"go_practiceapp/model"
	"time"
)

func CreateStamp(Stamp model.Stamps) (model.Stamps, error) {
	db := Connection()
	result := db.Create(&Stamp)
	if result.Error != nil {
		return Stamp, result.Error
	}
	return Stamp, nil
}

func UpdateStampUpTime(userid int, up_datetime time.Time) error {
	db := Connection()
	var stamps []model.Stamps

	if find := db.Debug().Where("up_datetime IS NULL AND users_id = ? AND in_datetime BETWEEN ? AND ?", userid, up_datetime.AddDate(0, 0, -1), up_datetime).Find(&stamps); find.RowsAffected < 1 {
		return errors.New("couldn't find record where up_datetime is null")
	} else if find.RowsAffected > 1 {
		return errors.New("couldn't select record to edit because there are too many stamps today")
	}

	update := db.Debug().Table("stamps").Where("id = ?", stamps[0].ID).Select("up_datetime").Updates(map[string]interface{}{"up_datetime": up_datetime})
	if update.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}

func ReadStampById(userid int) ([]model.Stamps, error) {
	var stamps []model.Stamps
	db := Connection()
	result := db.Where("users_id = ?", userid).Find(&stamps)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return stamps, result.Error
	}
}

func DeleteStamp(id int) error {
	db := Connection()
	var stamp model.Stamps

	result := db.Debug().Model(&stamp).Where("id = ?", id).Delete(stamp)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("no records updated")
	}
	return nil
}
