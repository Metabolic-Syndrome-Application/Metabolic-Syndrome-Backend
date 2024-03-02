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
			Photo:       "ภารกิจออกกำลังกาย.png",
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
			Photo:       "ภารกิจนอนหลับพักผ่อน.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},
		//3
		{
			Name:        "ภารกิจลดน้ำหวานและคาเฟอีน",
			Description: "ภารกิจลดน้ำหวาน ภารกิจที่จะทำให้คุณสามารถลดน้ำหวาน น้ำอัดลม ชานมไข่มุกได้ เพียงคุณดื่มน้ำหวานและคาเฟอีนเพียงสัปดาห์ละ 2 ครั้ง คือวันอังคารและวันพฤหัส",
			Photo:       "ภารกิจลดน้ำหวานและคาเฟอีน.png",
			Detail: models.Detail{
				Name: []string{"ไม่ดื่มน้ำหวานและคาเฟอีน เช่น น้ำอัดลม ชานมไข่มุก กาแฟ"},
				Day:  []string{"monday", "wednesday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},
		//4
		{
			Name:        "ภารกิจดื่มน้ำวันละ 8 แก้ว",
			Description: "ภารกิจดื่มน้ำวันละ 8 แก้ว ภารกิจที่จะทำให้คุณดื่มน้ำให้เพียงพอต่อวันจนเกิดเป็นนิสัย เพียงคุณดื่มวันละ 2 ลิตรหรือ 8 แก้วทุกๆวัน",
			Photo:       "ภารกิจดื่มน้ำวันละ 8 แก้ว.png",
			Detail: models.Detail{
				Name: []string{"ดื่มน้ำ 8 แก้วหรือ 2 ลิตร"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},
		//5
		{
			Name:        "ภารกิจผลไม้ดีต่อสุขภาพ",
			Description: "ภารกิจผลไม้ดีต่อสุขภาพ ภารกิจที่จะทำให้คุณมีสุขภาพร่างกายที่ดีขึ้น เพียงคุณกินผลไม้อะไรก็ได้ในทุกๆวัน ยกเว้นผลไม้ที่มีแคลลอรี่สูง เช่น ทุเรียน",
			Photo:       "ภารกิจผลไม้ดีต่อสุขภาพ.png",
			Detail: models.Detail{
				Name: []string{"กินผลไม้ เช่น ส้ม ฝรั่ง กล้วย"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			Points:  200,
			NumDays: 7,
		},
	}
	db.Create(&dailyData)

}
