package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/ployns/Metabolic-Syndrome-Backend/models"
	"gorm.io/gorm"
)

type ScreeningController struct {
	DB *gorm.DB
}

func NewScreeningController(DB *gorm.DB) ScreeningController {
	return ScreeningController{DB}
}

func (sc *ScreeningController) MetabolicRisk(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := sc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var payload = struct {
		Occupation             string  `json:"occupation,omitempty"`
		Height                 float32 `json:"height,omitempty"`
		Weight                 float32 `json:"weight,omitempty"`
		BMI                    float32 `json:"bmi,omitempty"`
		Waistline              float32 `json:"waistline,omitempty"`
		SystolicBloodPressure  int     `json:"systolicBloodPressure,omitempty"`
		DiastolicBloodPressure int     `json:"diastolicBloodPressure,omitempty"`
		PulseRate              int     `json:"pulseRate,omitempty"`
		BloodGlucose           float32 `json:"bloodGlucose,omitempty"`
		Triglyceride           float32 `json:"triglyceride,omitempty"`
		HDL                    float32 `json:"hdl,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var risk = 0

	// screening Waistline and HDL
	if ProfilePatient.Gender == "female" {
		if payload.Waistline >= 80 {
			risk += 1
		}
		if payload.HDL < 50 {
			risk += 1
		}
	} else if ProfilePatient.Gender == "male" {
		if payload.Waistline >= 90 {
			risk += 1
		}
		if payload.HDL < 40 {
			risk += 1
		}
	}

	// screening BloodGlucose
	if payload.BloodGlucose > 100 {
		risk += 1
	}

	// screening BloodPressure
	if payload.SystolicBloodPressure > 130 || payload.DiastolicBloodPressure > 85 {
		risk += 1
	}

	// screening Triglyceride
	if payload.Triglyceride > 150 {
		risk += 1
	}

	// result
	metabolicRisk := ""

	if risk <= 1 {
		metabolicRisk = "low"
		var planExercise models.Plan
		var planFood models.Plan
		var planRest models.Plan
		sc.DB.First(&planExercise, "name = 'โปรแกรมออกกำลังกายเมตาบอลิกซินโดรม'")
		sc.DB.First(&planFood, "name = 'โปรแกรมอาหารเมตาบอลิกซินโดรม'")
		sc.DB.First(&planRest, "name = 'โปรแกรมการพักผ่อนเมตาบอลิกซินโดรม'")

		// connect hospital and have planID => not add plan default
		var planID pq.StringArray
		if ProfilePatient.HN != nil && ProfilePatient.PlanID != nil {
			planID = ProfilePatient.PlanID
		} else {
			planID = pq.StringArray{planExercise.ID.String(), planFood.ID.String(), planRest.ID.String()}
		}

		updateRisk := &models.Patient{
			DiseaseRisk: models.DiseaseRisk{
				Diabetes:       "metabolicLow",
				Hyperlipidemia: "metabolicLow",
				Hypertension:   "metabolicLow",
				Obesity:        "metabolicLow"},
			PlanID: planID,
		}

		result := sc.DB.Model(&ProfilePatient).Updates(updateRisk)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
			return
		}

		// table patient_plan fotr patient.Plan
		var plans []models.Plan
		var plansToRemove []models.Plan

		sc.DB.Find(&plansToRemove)
		sc.DB.Model(&ProfilePatient).Association("Plan").Delete(&plansToRemove)

		for _, planID := range ProfilePatient.PlanID {

			sc.DB.Where("id = ?", planID).Find(&plans)
			sc.DB.Model(&ProfilePatient).Association("Plan").Append(&plans)

		}

	} else if risk == 2 {
		metabolicRisk = "medium"
		updateRisk := &models.Patient{
			DiseaseRisk: models.DiseaseRisk{
				Diabetes:       "metabolicMedium",
				Hyperlipidemia: "metabolicMedium",
				Hypertension:   "metabolicMedium",
				Obesity:        "metabolicMedium"},
		}

		result := sc.DB.Model(&ProfilePatient).Updates(updateRisk)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
			return
		}
	} else if risk >= 3 {
		metabolicRisk = "high"
		updateRisk := &models.Patient{
			DiseaseRisk: models.DiseaseRisk{
				Diabetes:       "metabolicHigh",
				Hyperlipidemia: "metabolicHigh",
				Hypertension:   "metabolicHigh",
				Obesity:        "metabolicHigh"},
		}

		result := sc.DB.Model(&ProfilePatient).Updates(updateRisk)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
			return
		}
	}

	// add data Occupation
	updatePatient := &models.Patient{
		Occupation: payload.Occupation,
	}
	result2 := sc.DB.Model(&ProfilePatient).Updates(updatePatient)
	if result2.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update profile Staff"})
		return
	}

	now := time.Now()
	newRecordHealth := &models.RecordHealth{
		PatientID:              ProfilePatient.ID,
		Height:                 payload.Height,
		Weight:                 payload.Weight,
		Waistline:              payload.Waistline,
		SystolicBloodPressure:  payload.SystolicBloodPressure,
		DiastolicBloodPressure: payload.DiastolicBloodPressure,
		PulseRate:              payload.PulseRate,
		BloodGlucose:           payload.BloodGlucose,
		Triglyceride:           payload.Triglyceride,
		HDL:                    payload.HDL,
		RecordBy:               currentUser.Role,
		Timestamp:              now,
	}
	fmt.Println("time:", now)
	result3 := sc.DB.Create(&newRecordHealth)

	if result3.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not create RecordHealth"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"metabolicRisk": metabolicRisk}})

}

func (sc *ScreeningController) DiseaseRisk(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := sc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	var RecordHealthPatient models.RecordHealth
	result2 := sc.DB.Order("timestamp DESC").First(&RecordHealthPatient, "patient_id = ?", currentUser.ID)
	if result2.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	payload := struct {
		Disease       []string `json:"disease,omitempty"`
		FamilyDisease []string `json:"familyDisease,omitempty"`
	}{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var planExercise models.Plan
	var planFood models.Plan
	var planRest models.Plan

	//screening diabetes
	calDiabetes := func() string {
		countDiabetes := 0
		//age
		currentYear := time.Now().Year()
		age := (currentYear + 543) - ProfilePatient.YearOfBirth
		if age >= 35 && age < 45 {
			countDiabetes += 1
		} else if age >= 45 && age < 50 {
			countDiabetes += 2
		} else if age >= 50 && age < 60 {
			countDiabetes += 3
		} else if age >= 60 {
			countDiabetes += 4
		}
		//bmi
		bmi := RecordHealthPatient.Weight / ((float32(RecordHealthPatient.Height) / 100) * (float32(RecordHealthPatient.Height) / 100))
		if bmi >= 23 && bmi < 27.5 {
			countDiabetes += 1
		} else if bmi >= 27.5 {
			countDiabetes += 3
		}

		//waist to height ratio
		ratio := RecordHealthPatient.Waistline / float32(RecordHealthPatient.Height)
		if ratio > 0.5 && ratio <= 0.6 {
			countDiabetes += 3
		} else if ratio > 0.6 {
			countDiabetes += 5
		}

		//BloodPressure
		if (RecordHealthPatient.SystolicBloodPressure >= 120 && RecordHealthPatient.SystolicBloodPressure < 140) && (RecordHealthPatient.DiastolicBloodPressure < 90) {
			countDiabetes += 2
		} else if RecordHealthPatient.SystolicBloodPressure >= 140 || RecordHealthPatient.DiastolicBloodPressure >= 90 {
			countDiabetes += 4
		}

		//family diabetes
		for _, value := range payload.FamilyDisease {
			if value == "diabetes" {
				countDiabetes += 2
				break
			}
		}

		//BloodGlucose
		if RecordHealthPatient.BloodGlucose >= 100 {
			countDiabetes += 5
		}

		//diabetes
		for _, value := range payload.Disease {
			if value == "diabetes" {
				countDiabetes += 15
				break
			}
		}

		//result
		if countDiabetes <= 7 {
			return "low"
		} else if countDiabetes > 7 && countDiabetes <= 14 {
			return "medium"
		} else {
			return "high"
		}

	}

	//screening Hyperlipidemia
	calHyperlipidemia := func() string {
		countHyperlipidemia := 0
		if RecordHealthPatient.HDL >= 60 {
			countHyperlipidemia += 1
		} else if RecordHealthPatient.HDL >= 45 {
			countHyperlipidemia += 2
		} else if RecordHealthPatient.HDL < 45 {
			countHyperlipidemia += 4
		}
		if RecordHealthPatient.Triglyceride < 150 {
			countHyperlipidemia += 1
		} else if RecordHealthPatient.Triglyceride < 200 {
			countHyperlipidemia += 2
		} else if RecordHealthPatient.Triglyceride >= 200 {
			countHyperlipidemia += 4
		}

		if countHyperlipidemia <= 2 {
			return "low"
		} else if countHyperlipidemia <= 4 {
			return "medium"
		} else {
			return "high"
		}

	}

	//screening Hypertension
	calHypertension := func() string {
		if RecordHealthPatient.SystolicBloodPressure < 130 && RecordHealthPatient.DiastolicBloodPressure < 85 {
			return "low"
		} else if (RecordHealthPatient.SystolicBloodPressure >= 130 && RecordHealthPatient.SystolicBloodPressure < 140) || (RecordHealthPatient.DiastolicBloodPressure >= 85 && RecordHealthPatient.DiastolicBloodPressure < 90) {
			return "medium"
		} else {
			return "high"
		}
	}

	//screening Obesity
	calObesity := func() string {
		bmi := RecordHealthPatient.Weight / ((float32(RecordHealthPatient.Height) / 100) * (float32(RecordHealthPatient.Height) / 100))
		if bmi < 23 {
			return "low"
		} else if bmi < 25 {
			return "medium"
		} else {
			return "high"
		}
	}

	if calDiabetes() == "high" || calHyperlipidemia() == "high" || calHypertension() == "high" || calObesity() == "high" {
		sc.DB.First(&planExercise, "name = 'โปรแกรมออกกำลังกายความเสี่ยงสูง'")
		sc.DB.First(&planFood, "name = 'โปรแกรมอาหารความเสี่ยงสูง'")
		sc.DB.First(&planRest, "name = 'โปรแกรมการพักผ่อนความเสี่ยงสูง'")
	} else if calDiabetes() == "medium" || calHyperlipidemia() == "medium" || calHypertension() == "medium" || calObesity() == "medium" {
		sc.DB.First(&planExercise, "name = 'โปรแกรมออกกำลังกายความเสี่ยงกลาง'")
		sc.DB.First(&planFood, "name = 'โปรแกรมอาหารความเสี่ยงกลาง'")
		sc.DB.First(&planRest, "name = 'โปรแกรมการพักผ่อนความเสี่ยงกลาง'")
	} else {
		sc.DB.First(&planExercise, "name = 'โปรแกรมออกกำลังกายความเสี่ยงต่ำ'")
		sc.DB.First(&planFood, "name = 'โปรแกรมอาหารความเสี่ยงต่ำ'")
		sc.DB.First(&planRest, "name = 'โปรแกรมการพักผ่อนความเสี่ยงต่ำ'")
	}

	// connect hospital and have planID => not add plan default
	var planID pq.StringArray
	if ProfilePatient.HN != nil && ProfilePatient.PlanID != nil {
		planID = ProfilePatient.PlanID
	} else {
		planID = pq.StringArray{planExercise.ID.String(), planFood.ID.String(), planRest.ID.String()}
	}

	updateDiseaseRisk := &models.Patient{
		DiseaseRisk: models.DiseaseRisk{
			Diabetes:       calDiabetes(),
			Hyperlipidemia: calHyperlipidemia(),
			Hypertension:   calHypertension(),
			Obesity:        calObesity()},
		PlanID: planID,
	}

	a := sc.DB.Model(&ProfilePatient).Updates(updateDiseaseRisk)
	if a.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Can not update DiseaseRisk"})
		return
	}
	// table patient_plan fotr patient.Plan
	var plans []models.Plan
	var plansToRemove []models.Plan

	sc.DB.Find(&plansToRemove)
	sc.DB.Model(&ProfilePatient).Association("Plan").Delete(&plansToRemove)

	for _, planID := range ProfilePatient.PlanID {

		sc.DB.Where("id = ?", planID).Find(&plans)
		sc.DB.Model(&ProfilePatient).Association("Plan").Append(&plans)

	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"diseaseRisk": ProfilePatient.DiseaseRisk}})
}

func (sc *ScreeningController) GetDiseaseRisk(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var ProfilePatient models.Patient
	result := sc.DB.First(&ProfilePatient, "id = ?", currentUser.ID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"diseaseRisk": ProfilePatient.DiseaseRisk}})
}
