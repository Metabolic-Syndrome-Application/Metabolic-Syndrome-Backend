package seeds

import (
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

func SeedDaily(db *gorm.DB) {
	dailyData := []models.DailyChallenge{
		//1
		{
			Name:        "ภารกิจออกกำลังกาย",
			Description: "ภารกิจออกกำลังกาย ภารกิจที่จะทำให้คุณมีสุขภาพที่ดีขึ้น เพียงออกกำลังกายทุกวันวันละ 30 นาที",
			Photo:       "exercise.png",
			Detail: models.Detail{
				Name: []string{"ออกกำลังกาย 30 นาที"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},

		//2
		{
			Name:        "ภารกิจนอนหลับพักผ่อน",
			Description: "ภารกิจนอนหลับพักผ่อน ภารกิจที่จะทำให้คุณนอนหลับเต็มอิ่มในทุกๆวัน เพียงคุณเข้านอนก่อน 4 ทุ่มทุกวัน",
			Photo:       "rest.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},
	}
	db.Create(&dailyData)

}
