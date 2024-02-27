package seeds

import (
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

func SeedKnowledge(db *gorm.DB) {
	knowledgeData := []models.Knowledge{
		{
			Disease:     "diabetes",
			Name:        "โรคเบาหวาน",
			Details:     "",
			Symptoms:    "Increased thirst, frequent urination, extreme hunger, unexplained weight loss, fatigue, blurred vision, etc.",
			Medications: "Insulin, Metformin, Sulfonylureas, etc.",
			Behaviors:   "Regular exercise, healthy diet, monitoring blood sugar levels, etc.",
		},
		{
			Disease:     "hypertension",
			Name:        "โรคความดันโลหิตสูง",
			Details:     "",
			Symptoms:    "Headaches, shortness of breath, nosebleeds, chest pain, vision changes, etc.",
			Medications: "Diuretics, ACE inhibitors, Calcium channel blockers, etc.",
			Behaviors:   "Low-sodium diet, regular physical activity, stress management, etc.",
		},
		{
			Disease:     "obesity",
			Name:        "โรคอ้วน",
			Details:     "",
			Symptoms:    "Increased body weight, fatigability, joint pain, breathlessness, snoring, etc.",
			Medications: "Orlistat, Phentermine, Liraglutide, etc.",
			Behaviors:   "Balanced diet, regular exercise, portion control, behavioral therapy, etc.",
		},
		{
			Disease:     "hyperlipidemia",
			Name:        "ภาวะไขมันในเลือดสูง",
			Details:     "",
			Symptoms:    "Chest pain, yellowish skin, fatty deposits under the skin, etc.",
			Medications: "Statins, Fibrates, Niacin, etc.",
			Behaviors:   "Healthy diet, regular exercise, weight management, smoking cessation, etc.",
		},
	}

	// เพิ่มข้อมูลลงในฐานข้อมูล
	// for _, data := range knowledgeData {
	//     if err := db.Create(&data).Error; err != nil {
	//         panic(err)
	//     }
	// }
	db.Create(&knowledgeData)

}
