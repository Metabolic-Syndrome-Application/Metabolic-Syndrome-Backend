package seeds

import (
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

func SeedPlan(db *gorm.DB) {
	planData := []models.Plan{

		// low
		//1
		{
			Name:        "โปรแกรมออกกำลังกายความเสี่ยงต่ำ",
			Type:        "exercise",
			Description: "โปรแกรมออกกำลังกายความเสี่ยงต่ำ สำหรับผู้ที่มีความเสี่ยงโรคต่ำ",
			Photo:       "โปรแกรมออกกำลังกายความเสี่ยงต่ำ.png",
			Detail: models.Detail{
				Name: []string{"ออกกำลังกาย 30 นาที"},
				Day:  []string{"monday", "wednesday", "friday"},
			},
		},
		//2
		{
			Name:        "โปรแกรมอาหารความเสี่ยงต่ำ",
			Type:        "food",
			Description: "โปรแกรมอาหารความเสี่ยงต่ำ สำหรับผู้ที่มีความเสี่ยงโรคต่ำ",
			Photo:       "โปรแกรมอาหารความเสี่ยงต่ำ.png",
			Detail: models.Detail{
				Name: []string{"กินผักและผลไม้ เช่น ส้ม ฝรั่ง กล้วย", "ดื่มน้ำวันละ 8 แก้ว"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
		//3
		{
			Name:        "โปรแกรมการพักผ่อนความเสี่ยงต่ำ",
			Type:        "rest",
			Description: "โปรแกรมการพักผ่อนความเสี่ยงต่ำ สำหรับผู้ที่มีความเสี่ยงโรคต่ำ",
			Photo:       "โปรแกรมการพักผ่อนความเสี่ยงต่ำ.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม หรือนอนครบ 8 ชั่วโมง"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},

		// medium
		//1
		{
			Name:        "โปรแกรมออกกำลังกายความเสี่ยงกลาง",
			Type:        "exercise",
			Description: "โปรแกรมออกกำลังกายความเสี่ยงกลาง สำหรับผู้ที่มีความเสี่ยงโรคกลาง",
			Photo:       "โปรแกรมออกกำลังกายความเสี่ยงกลาง.png",
			Detail: models.Detail{
				Name: []string{"ออกกำลังกาย 30 นาที"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday"},
			},
		},
		//2
		{
			Name:        "โปรแกรมอาหารความเสี่ยงกลาง",
			Type:        "food",
			Description: "โปรแกรมอาหารความเสี่ยงกลาง สำหรับผู้ที่มีความเสี่ยงโรคกลาง",
			Photo:       "โปรแกรมอาหารความเสี่ยงกลาง.png",
			Detail: models.Detail{
				Name: []string{"กินผักและผลไม้ เช่น ส้ม ฝรั่ง กล้วย", "ดื่มน้ำวันละ 8 แก้ว", "ไม่ทานอาหารรสจัด เช่น ส้มตำ"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
		//3
		{
			Name:        "โปรแกรมการพักผ่อนความเสี่ยงกลาง",
			Type:        "rest",
			Description: "โปรแกรมการพักผ่อนความเสี่ยงกลาง สำหรับผู้ที่มีความเสี่ยงโรคกลาง",
			Photo:       "โปรแกรมการพักผ่อนความเสี่ยงกลาง.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม หรือนอนครบ 8 ชั่วโมง"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},

		// high
		//1
		{
			Name:        "โปรแกรมออกกำลังกายความเสี่ยงสูง",
			Type:        "exercise",
			Description: "โปรแกรมออกกำลังกายความเสี่ยงสูง สำหรับผู้ที่มีความเสี่ยงโรคสูง",
			Photo:       "โปรแกรมออกกำลังกายความเสี่ยงสูง.png",
			Detail: models.Detail{
				Name: []string{"ออกกำลังกาย 30 นาที"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
		//2
		{
			Name:        "โปรแกรมอาหารความเสี่ยงสูง",
			Type:        "food",
			Description: "โปรแกรมอาหารความเสี่ยงสูง สำหรับผู้ที่มีความเสี่ยงโรคสูง",
			Photo:       "โปรแกรมอาหารความเสี่ยงสูง.png",
			Detail: models.Detail{
				Name: []string{"กินผักและผลไม้ เช่น ส้ม ฝรั่ง กล้วย", "ดื่มน้ำวันละ 8 แก้ว", "ไม่ทานอาหารรสจัด เช่น ส้มตำ", "ไม่ทานอาหารที่มีคลอเรสเตอรอลสูง เช่น อาการทะเล เครื่องในสัตว์"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
		//3
		{
			Name:        "โปรแกรมการพักผ่อนความเสี่ยงสูง",
			Type:        "rest",
			Description: "โปรแกรมการพักผ่อนความเสี่ยงสูง สำหรับผู้ที่มีความเสี่ยงโรคสูง",
			Photo:       "โปรแกรมการพักผ่อนความเสี่ยงสูง.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม หรือนอนครบ 8 ชั่วโมง"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},

		// metabolic
		//1
		{
			Name:        "โปรแกรมออกกำลังกายเมตาบอลิกซินโดรม",
			Type:        "exercise",
			Description: "โปรแกรมออกกำลังกายเมตาบอลิกซินโดรม สำหรับผู้ที่มีความเสี่ยงเมตาบอลิกซินโดรม",
			Photo:       "โปรแกรมออกกำลังกายเมตาบอลิกซินโดรม.png",
			Detail: models.Detail{
				Name: []string{"ออกกำลังกาย 30 นาที"},
				Day:  []string{"monday", "wednesday", "friday"},
			},
		},
		//2
		{
			Name:        "โปรแกรมอาหารเมตาบอลิกซินโดรม",
			Type:        "food",
			Description: "โปรแกรมอาหารเมตาบอลิกซินโดรม สำหรับผู้ที่มีความเสี่ยงเมตาบอลิกซินโดรม",
			Photo:       "โปรแกรมอาหารเมตาบอลิกซินโดรม.png",
			Detail: models.Detail{
				Name: []string{"ดื่มน้ำวันละ 8 แก้ว"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
		//3
		{
			Name:        "โปรแกรมการพักผ่อนเมตาบอลิกซินโดรม",
			Type:        "rest",
			Description: "โปรแกรมการพักผ่อนเมตาบอลิกซินโดรม สำหรับผู้ที่มีความเสี่ยงเมตาบอลิกซินโดรม",
			Photo:       "โปรแกรมการพักผ่อนเมตาบอลิกซินโดรม.png",
			Detail: models.Detail{
				Name: []string{"เข้านอนก่อน 4 ทุ่ม หรือนอนครบ 8 ชั่วโมง"},
				Day:  []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
		},
	}
	db.Create(&planData)

}
