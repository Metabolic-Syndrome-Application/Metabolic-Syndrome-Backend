package seeds

import (
	"encoding/json"

	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

func SeedQuiz(db *gorm.DB) {
	quizData := []models.QuizChallenge{
		//1
		{
			Question: "โรคในข้อใดไม่เกี่ยวข้องกับภาวะเมทาบอลิกซินโดรม",
			Choices: json.RawMessage(`[
                {
					"option": "หัวใจ", 
					"isCorrect": true
				},
                {
					"option": "อ้วน",
					"isCorrect": false
				},
                {
					"option": "ความดันโลหิตสูง", 
					"isCorrect": false
				}
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//2
		{
			Question: "ค่าน้ำตาลปกติต้องไม่เกินเท่าไหร่",
			Choices: json.RawMessage(`[
                {
                    "option": "100",
                    "isCorrect": true
                },
                {
                    "option": "70",
                    "isCorrect": false
                },
                {
                    "option": "150",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//3
		{
			Question: "โรคเบาหวานแบ่งออกเป็นกี่ชนิด",
			Choices: json.RawMessage(`[
                {
                    "option": "2",
                    "isCorrect": false
                },
                {
                    "option": "3",
                    "isCorrect": false
                },
                {
                    "option": "4",
                    "isCorrect": true
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//4
		{
			Question: "โรคที่ใช้ BMI เป็นตัววัดว่าเป็นโรคนี้หรือไม่ คือโรคอะไร",
			Choices: json.RawMessage(`[
                {
                    "option": "โรคเบาหวาน",
                    "isCorrect": false
                },
                {
                    "option": "โรคอ้วน",
                    "isCorrect": true
                },
                {
                    "option": "ภาวะไขมันในเลือดสูง",
                    "isCorrect": false
                },
				{
                    "option": "โรคความดันโลหิตสูง",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//5
		{
			Question: "ภาวะเมตาบอลิกซินโดรมเกิดจากอะไร",
			Choices: json.RawMessage(`[
                {
                    "option": "พันธุกรรม",
                    "isCorrect": false
                },
                {
                    "option": "พฤติกรรม",
                    "isCorrect": true
                },
                {
                    "option": "เกิดขึ้นเองตามธรรมชาติ",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//6
		{
			Question: "อินซูลินคืออะไร",
			Choices: json.RawMessage(`[
                {
                    "option": "ฮอร์โมนที่สร้างจากตับอ่อนเพื่อใช้ควบคุมระดับน้ำตาลในเลือด",
                    "isCorrect": true
                },
                {
                    "option": "เอนไซม์ในร่างกาย",
                    "isCorrect": false
                },
				{
                    "option": "สารกระตุ้นความหิว",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//7
		{
			Question: "ไขมันในเลือดสูงสามารถเสี่ยงต่อโรคใดได้บ้าง",
			Choices: json.RawMessage(`[
                {
                    "option": "โรคหัวใจ",
                    "isCorrect": false
                },
                {
                    "option": "โรคอัมพฤกษ์",
                    "isCorrect": false
                },
				{
                    "option": "โรคหลอดเลือดอุดตัน",
                    "isCorrect": true
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//8
		{
			Question: "พฤติกรรมใดเสี่ยงต่อภาวะเมตาบอลิกซินโดรม",
			Choices: json.RawMessage(`[
                {
                    "option": "สูบบุหรี่",
                    "isCorrect": false
                },
                {
                    "option": "ออกกำลังกายไม่เพียงพอ",
                    "isCorrect": false
                },
				{
                    "option": "รับประทานอาหารรสจัด",
                    "isCorrect": false
                },
				{
                    "option": "ถูกทุกข้อ",
                    "isCorrect": true
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//9
		{
			Question: "โรคในข้อใดเกิดจากปริมาณคลอเรสเตอรอลและไตรกลีเซอไรด์ผิดปกติ",
			Choices: json.RawMessage(`[
                {
                    "option": "โรคเบาหวาน",
                    "isCorrect": false
                },
                {
                    "option": "โรคอ้วน",
                    "isCorrect": false
                },
                {
                    "option": "ภาวะไขมันในเลือดสูง",
                    "isCorrect": true
                },
				{
                    "option": "โรคความดันโลหิตสูง",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
		//10
		{
			Question: "อาหารใดมีคลอเรสเตอรอลสูงที่สุด",
			Choices: json.RawMessage(`[
                {
                    "option": "ถั่วเหลือง",
                    "isCorrect": false
                },
                {
                    "option": "เนื้อปลา",
                    "isCorrect": false
                },
                {
                    "option": "อาหารทะเล",
                    "isCorrect": true
                },
				{
                    "option": "กระเทียม",
                    "isCorrect": false
                }
            ]`),
			Points:    150,
			LimitTime: 1,
		},
	}
	db.Create(&quizData)

}
