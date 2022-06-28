package database

import (
	"errors"
	"go_practiceapp/model"
	"time"
)

func CreateStamp(stamp model.Stamps) (model.Stamps, error) {
	db := Connection()
	result := db.Create(&stamp)
	if result.Error != nil {
		return stamp, result.Error
	}
	return stamp, nil
}

func UpdateStampUpTime(userid int, up_datetime time.Time) (model.Stamps, error) {
	db := Connection()
	var stamps []model.Stamps

	if find := db.Debug().Where("up_datetime IS NULL AND users_id = ? AND in_datetime BETWEEN ? AND ?", userid, up_datetime.AddDate(0, 0, -1), up_datetime).Find(&stamps); find.RowsAffected < 1 {
		return model.Stamps{}, errors.New("couldn't find record where up_datetime is null")
	} else if find.RowsAffected > 1 {
		return model.Stamps{}, errors.New("couldn't select record to edit because there are too many stamps today")
	}

	update := db.Debug().Model(&stamps).Where("id = ?", stamps[0].ID).Select("up_datetime").Updates(map[string]any{"up_datetime": up_datetime})
	if update.RowsAffected == 0 {
		return model.Stamps{}, errors.New("no records updated")
	}
	return stamps[0], nil
}

func UpdateStampTimestamp(stampid int, in_datetime time.Time, up_datetime time.Time) (model.Stamps, error) {
	db := Connection()
	stamp := model.Stamps{
		ID:          stampid,
		In_datetime: in_datetime,
		Up_datetime: &up_datetime,
	}
	update := db.Debug().Table("stamps").Select("in_datetime", "up_datetime").Where("id = ?", stamp.ID).Updates(map[string]any{"in_datetime": stamp.In_datetime, "up_datetime": stamp.Up_datetime})
	if update.RowsAffected == 0 {
		return model.Stamps{}, errors.New("no records updated")
	} else {
		return stamp, nil
	}
}

func ReadStampByUserId(userid int) ([]model.Stamps, error) {
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
