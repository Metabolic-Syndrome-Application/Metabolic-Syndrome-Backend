package seeds

import (
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

func SeedPlan(db *gorm.DB) {
	planData := []models.Plan{
		{
			Name:  "โปรแกรมอาหารความเสี่ยงต่ำ",
			Type:  "food",
			Photo: "โปรแกรมอาหารความเสี่ยงต่ำ.png",
		},
		{
			Name:  "โปรแกรมออกกำลังกายความเสี่ยงต่ำ",
			Type:  "exercise",
			Photo: "โปรแกรมออกกำลังกายความเสี่ยงต่ำ.png",
		},
		{
			Name:  "โปรแกรมการพักผ่อนความเสี่ยงต่ำ",
			Type:  "rest",
			Photo: "โปรแกรมการพักผ่อนความเสี่ยงต่ำ.png",
		},
	}
	db.Create(&planData)

}
