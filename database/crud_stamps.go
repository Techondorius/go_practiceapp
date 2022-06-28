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

// stampsテーブルからup_datetimeの前24時間以内にINがStampされてて、up_datetimeがNULLのレコードを返します。
func ReadStampsByUserID(userid int, up_datetime time.Time) ([]model.Stamps) {
	db := Connection()
	var stamps []model.Stamps
	db.Where("up_datetime IS NULL AND users_id = ? AND in_datetime BETWEEN ? AND ?", userid, up_datetime.AddDate(0, 0, -1), up_datetime).Find(&stamps)
	return stamps
}

// 
func UpdateStampUpTime(s model.Stamps, up_datetime time.Time) (model.Stamps, error){
	db := Connection()
	s.Up_datetime = &up_datetime
	if update := db.Debug().Select("up_datetime").Updates(s); update.RowsAffected == 0 {
		return model.Stamps{}, errors.New("no records updated")
	}
	return s, nil
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
