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
			Photo:       "assets/images/diabetes_article.png",
			Details:     "โรคเบาหวาน คือ โรคที่เซลล์ร่างกายมีความผิดปกติในขบวนการเปลี่ยนน้ำตาลในเลือดให้เป็นพลังงาน โดยขบวนการนี้เกี่ยวข้องกับอินซูลินซึ่งเป็นฮอร์โมนที่สร้างจากตับอ่อนเพื่อใช้ควบคุมระดับน้ำตาลในเลือด เมื่อน้ำตาลไม่ได้ถูกใช้จึงทำให้ระดับน้ำตาลในเลือดสูงขึ้นกว่าระดับปกติ โดยปกติค่าน้ำตาลในเลือดไม่ควรเกิน 100 มิลลิกรัมต่อเดซิลิตร โรคเบาหวานแบ่งเป็น 4 ชนิด ตามสาเหตุของการเกิดโรค\nเบาหวานชนิดที่ 1 - เกิดจากภูมิคุ้มกันร่างกายทำลายเซลล์สร้างอินสุลินในตับอ่อน ทำให้เกิดภาวะขาดอินสุลิน\nเบาหวานชนิดที่ 2 - เกิดจากภาวะการลดลงของการสร้างอินสุลิน ร่วมกับภาวะดื้ออินสุลิน และมักเป็นกรรมพันธุ์\nเบาหวานชนิดพิเศษ - สาเหตุของโรคเบาหวานชนิดนี้อาจเกิดจากความความผิดปกติของตับอ่อน หรือเป็นโรคที่เกี่ยวกับการทำงานที่ผิดปกติของอินสุลินโดยกำเนิด\nเบาหวานขณะตั้งครรภ์ - เบาหวานชนิดนี้เกิดขึ้นขณะตั้งครรภ์และหายไปได้หลังคลอดบุตร แต่มีโอกาสเสี่ยงที่จะเกิดโรคเบาหวานชนิดที่ 2 ในอนาคต ",
			Symptoms:    "1. ปัสสาวะบ่อย ถึงบ่อยมาก เนื่องจากกระบวนการกรองน้ำตาลในเลือดที่สูงมากออกมากทางปัสสาวะ ไตจำเป็นต้องดึงน้ำออกมาด้วย ดังนั้น ผู้ป่วยยิ่งมีระดับน้ำตาลสูงมากเท่าใดก็ยิ่งปัสสาวะบ่อยและมากขึ้นเท่านั้น ทำให้ต้องตื่นมาเข้าห้องน้ำตอนกลางคืนหลายครั้ง\n2. คอแห้ง กระหายน้ำ เป็นผลมาจากการที่ร่างกายเสียน้ำไปจากการปัสสาวะบ่อยและมาก ทำให้เกิดภาวะขาดน้ำจึงต้องชดเชยด้วยการดื่มน้ำบ่อยๆ\n3. หิวบ่อย กินจุ เนื่องจากร่างกายขาดพลังงาน จึงทำให้รู้สึกหิวบ่อย และรับประทานจุ\n4. น้ำหนักลด เนื่องจากในภาวะที่ขาดอินซูลิน ร่างกายไม่สามารถนำน้ำตาลในเลือดไปใช้เป็นพลังงานได้ ร่วมกับการขาดน้ำจากการปัสสาวะบ่อย ร่างกายจึงจำเป็นต้องนำเอาโปรตีนและไขมันที่เก็บสะสมไว้ในเนื้อเยื่อมาใช้แทน จึงทำให้ร่างกายรู้สึกอ่อนเพลีย และน้ำหนักตัวลดลงโดยไม่ทราบสาเหตุ\n5. คันตามผิวหนัง\n6. ตามัว ชาปลายมือ ปลายเท้า\n7. เป็นแผลง่าย หายยาก",
			Medications: "1. ยาฉีดอินซูลิน สำหรับรักษาโรคเบาหวาน\n1.1 อินซูลินชนิดออกฤทธิ์สั้น (Short acting insulin หรือ Regular insulin, RI) เช่น ยา แอ็คทราพิด (Actrapid®), ฮิวมูลินอาร์ (Humulin R®), เจ็นซูลินอาร์ (Gensulin R®)\n1.2 อินซูลินชนิดออกฤทธิ์นานปานกลาง (Intermediate acting insulin หรือ NPH Insulin) เช่น ยา ฮิวมูลินเอ็น (Humulin N®) อินซูลาทาร์ด (Insulatard HM® ) เจ็นซูลินเอ็น (Gensulin N®)\n1.3 อินซูลินชนิดออกฤทธิ์นาน (Long acting insulin) เช่น ยา ดีทีเมียร์ (Detemir), กลาร์จีน (Glargine)\n1.4 อินซูลินชนิดออกฤทธิ์ไวมาก (Rapid acting insulin หรือ Ultrashort acting insulin) เช่น ยา ลิสโปร (Lispro), แอสพาร์ท (Aspart), กลูไลซีน (Glulisine)\n2. ยากินสำหรับรักษาโรคเบาหวาน\n2.1 ยารักษาโรคเบาหวาน กลุ่มที่ออกฤทธิ์ทำให้น้ำตาลในเลือดต่ำ (Hypoglycemic drugs)\n2.1.1ยากลุ่มซัลโฟนิลยูเรีย (Sulfonylureas): เช่นยา อะซีโตเฮกซาไมด์ (Acetohexamide), คลอร์โพรพาไมด์ (Chlorpropamide), โทลาซาไมด์ (Tolazamide), ไกลเมพิไรด์ (Glimepiride), ไกลพิไซด์ (Glipizide), ไกลเบนคลาไมด์ (Glibenclamide) หรือ อีกชื่อคือ ไกลบูไรด์ (Glyburide) ยาในกลุ่มนี้รับประทานวันละ1-2 ครั้ง โดยยาจะออกฤทธิ์กระตุ้นตับอ่อนให้ผลิตอินซูลินเพิ่มขึ้น และมีอาการข้างเคียงที่พบได้บ่อย ได้แก่ ภาวะน้ำตาลในเลือดต่ำ ผื่นที่ผิวหนัง แน่นท้อง ดังนั้นข้อควรจะปฏิบัติเมื่อกินยากลุ่มนี้ จึงควรจะพกน้ำผลไม้ หรืออาหารพวกแป้งติดตัว และกินอาหาร การออกกำลังกายให้ตรงเวลา\n2.1.2 ยากลุ่มที่ไม่ใช่ซัลโฟนิลยูเรีย (Non - sulfonylureas หรือ Glinides หรือ Meglitinides): เช่นยา รีพาไกลไนด์ (Repaglinide), นาทิไกลไนด์ (Nateglinide), มิทิไกลไนด์ (Mitiglinide)\n2.2 ยารักษาโรคเบาหวาน #ที่ออกฤทธ์ต้านการเพิ่มระดับน้ำตาลในเลือด (Antihyper glycemic drugs)\n2.2.1 ยากลุ่มไบกัวไนด์ (Biguanides): เช่นยา เมทฟอร์มิน (Metformin) ยาเมทฟอร์มิน จะรับประทานวันละ 2 ครั้ง เช้าและเย็น ยาชนิดนี้จะออกฤทธิ์ลดการสร้างน้ำตาลจากตับ ผลข้างเคียงของยาอาจจะทำให้เกิดอาการ แน่นท้อง เบื่ออาหาร การรับประทานพร้อมอาหารจะลดอาการข้างเคียงของยา สำหรับผู้ที่เป็นโรคตับ หรือโรคไต อาจจะเกิดภาวะกรดในเลือด\n2.2.2 ยากลุ่มไธอะโซลิดีนไดโอน (Thiazolidinediones): เช่นยา ไพโอกลิทาโซน (Pioglitazone) ยากลุ่มนี้จะช่วยเพิ่มการตอบสนองของร่างกายต่ออินซูลิน ทำให้ร่างกายสามารถนำอินซูลินไปใช้ได้ดีขึ้น โดยไม่ทำให้ระดับน้ำตาลในเลือดต่ำผิดปกติ แต่เนื่องจากยากลุ่มนี้อาจสามารถส่งผลเสียต่อตับ จึงควรหมั่นตรวจเลือดเพื่อดูค่าเอนไซม์ตับเป็นประจำ เพื่อเฝ้าระวังการทำงานของตับที่ผิดปกติ\n2.3 ยารักษาโรคเบาหวาน กลุ่มที่ยับยั้งการทำงานของเอนไซม์ที่ย่อยสลายคาร์โบไฮเดรตในลำไส้ ยากลุ่มที่ยับยั้งการทำงานของเอนไซม์อัลฟ่ากลูโคซิเดส/เอนไซม์ยับยั้งการย่อยสลายคาร์โบไฮเดรตในลำไส้ (Alpha - glucosidase inhibitor) เช่นยา อะคาร์โบส (Acarbose), โวกลิโบส (Voglibose), ไมกลิทอล (Miglitol)ยากลุ่มนี้เป็นยาที่นำมาใช้รักษาผู้ป่วยเบาหวานโดยลดการดูดซึมสารอาหารที่ลำไส้เล็กส่วนต้น สามารถลด FPG ได้16-20 มก.% ลด HbA1c ได้ 0.59 % และลดน้ำตาลหลังอาหาร [post prandrial glucose] ได้ 51 มก.% มีอาการข้างเคียงที่พบได้บ่อย ได้แก่ ท้องอืด ท้องเดิน ปวดท้อง ควรเริ่มยาแต่น้อยเคี้ยวพร้อมอาหาร\n2.4 ยารักษาโรคเบาหวาน กลุ่มอินครีตินฮอร์โมน (Incretin hormones) Incretin mimetics (อินครีตินฮอร์โมน สร้างจากลำไส้เล็กเพื่อกระตุ้นการหลั่งอินซูลินเมื่อมีการบริโภคอาหาร)\n2.5 ยากลุ่มที่ออกฤทธิ์เลียนแบบการทำงานของฮอร์โมน Glucagon (Glucagon-like peptide-1 receptor agonists ย่อว่า GLP-1 receptor agonist) เช่นยา เอ็กซีนาไทด์ (Exenatide)\n2.6 ยารักษาโรคเบาหวาน กลุ่มที่ออกฤทธิ์ยับยั้งเอนไซม์ Dipeptidyl peptidase เอนไซม์ที่เกี่ยวข้องกับการย่อยสลายของน้ำตาล Glucose (Dipeptidyl peptidase-4 Inhibitor, DPP-4 Inhibitor)\n2.7 ยากลุ่มใหม่ซึ่งเป็น ยาออกฤทธิ์ควบคุมการดูดซึมกลับของน้ำตาลกลูโคสบริเวณท่อไต ยากลุ่ม Sodium - glucose Cotransporter inhibitors (SGLT2 inhibitors): เช่นยา ดาพากลิโฟลซิน (Dapagliflozin), คานากลิโฟลซิน (Canaglifloziin), เอ็มพากลิโฟลซิน (Empagliflozin)",
			Behaviors:   "1. จำกัดปริมาณอาหารคาร์โบไฮเดรต ได้แก่ ข้าว แป้ง ขนมปัง สปาเก็ตตี้ วุ้นเส้น ข้าวโพด เผือก มัน ผลไม้ และเครื่องดื่มต่างๆ ควรได้รับปริมาณที่เหมาะสมในแต่ละมื้อ\n2. เลือกคาร์โบไฮเดรตเชิงซ้อน เป็นคาร์โบไฮเดรตที่อุดมด้วยใยอาหาร เช่น ข้าวแป้งไม่ขัดสี ธัญพืช และผักต่าง ๆ แต่ผักบางชนิดมีคาร์โบไฮเดรตอยู่มาก เช่น ฟักทอง ไม่ควรรับประทานมากเกินไป\n3. ผลไม้ต่อวันต้องไม่มากเกินไป เพราะผลไม้เป็นอาหารที่มีน้ำตาล เรียกว่า Fructose มีอยู่ในผลไม้ทุกชนิด ไม่ว่าจะรสเปรี้ยวหรือรสหวาน แนะนำให้จำกัดการรับประทาน 3 - 4 ส่วนต่อวัน หรือ 1 ส่วนต่อมื้อ ตัวอย่าง 1 ส่วนบริโภค ได้แก่ แอปเปิ้ล 1 ลูก, กล้วยหนึ่งลูก, เมลอน 6 ชิ้นพอดีคำ\n4. เลี่ยงดื่มน้ำผลไม้ทุกชนิด เนื่องจากไม่มีใยอาหารและมีน้ำตาลค่อนข้างสูง ใยอาหารในผลไม้สามารถขัดขวางการดูดซึมน้ำตาลบางส่วนได้ ผู้ป่วยบางคนอาจเคยได้รับข้อมูลว่า ควรเลือกกินผลไม้ที่ไม่หวาน ดัชนีน้ำตาลต่ำ (Glycemic Index) แต่ที่จริงแล้วสิ่งสำคัญกว่าคือ การควบคุมปริมาณให้เหมาะสม เพราะหากกินผลไม้รสจืดในปริมาณมากก็สามารถทำให้น้ำตาลเพิ่มขึ้นสูงได้อยู่ดี\n5. เลี่ยงการรับประทานขนมหวาน จะสามารถช่วยลดระดับน้ำตาลในเลือดได้มาก เนื่องจากน้ำตาลทรายที่ใส่ในขนมหวานสามารถถูกดูดซึมได้เร็ว ดังนั้นระดับน้ำตาลในเลือดจะเพิ่มขึ้นอย่างรวดเร็ว\n6. เลี่ยงเครื่องดื่มน้ำตาลสูง เช่น น้ำอัดลม ชา กาแฟ นมเปรี้ยวและนมรสต่าง ๆ น้ำสมุนไพร เป็นต้น เพื่อเป็นการจำกัดปริมาณแคลอรีที่รับประทาน หรือ เพื่อเป็นการควบคุมระดับน้ำตาลในเลือด ควรเปลี่ยนมาเป็นการใช้น้ำตาลเทียม (non-nutritive sweeteners) เช่น aspartame, saccharin, sucralose หรือหญ้าหวานแทนน้ำตาลทรายขาว หรือ น้ำตาลทรายแดง โดยสามารถอ่านรายละเอียดฉลากโภชนาการทุกครั้งก่อนเลือกซื้อ\n7. จำกัดปริมาณแอลกอฮอล์ ไม่เกิน 1 Drink ในผู้หญิง และ 2 Drink ในผู้ชาย เนื่องจากแอลกอฮอล์สามารถทำให้ระดับน้ำตาลในเลือดสูงได้เช่นกัน และไม่ควรดื่มตอนท้องว่าง เพราะจะทำให้น้ำตาลต่ำได้\n8. งดสูบบุหรี่โดยเด็ดขาด เพื่อลดความเสี่ยงของโรคหัวใจและหลอดเลือด\n9. ควรออกกำลังกายอย่างสม่ำเสมอ เช่น การเดินเร็ว ปั่นจักรยาน หรือเต้นแอโรบิก ครั้งละประมาณ 30 นาที ให้ได้ 3-5 ครั้ง/สัปดาห์",
		},
		{
			Disease:     "hypertension",
			Name:        "โรคความดันโลหิตสูง",
			Photo:       "assets/images/hypertension_article.png",
			Details:     "โรคความดันโลหิตสูง คือ สภาวะผิดปกติที่บุคคลมีระดับความดันโลหิตสูงขึ้นกว่าระดับปกติของคนส่วนใหญ่ และถือว่าเป็นสภาวะที่ต้องควบคุม เนื่องจากความดันโลหิตสูงทำให้เกิดความเสียหาย และการเสื่อมสภาพของหลอดเลือดแดง นำไปสู่สภาวะการแข็งตัวของหลอดเลือด การอุดตันของหลอดเลือด หรือหลอดเลือดแตกได้ นอกจากนี้ความดันโลหิตสูงยังเป็นปัจจัยเสี่ยงทำให้เกิดโรคแทรกซ้อนอื่นๆ ตามมาได้เช่น โรคหัวใจวาย โรคอัมพาต โรคสมองเสื่อมหรือโรคไตวายเรื้อรัง",
			Symptoms:    "โรคความดันโลหิตสูงระยะแรกส่วนใหญ่ไม่มีอาการมีเพียงส่วนน้อยที่มีอาการและที่พบได้บ่อยคือ ปวดมึนท้ายทอยตึงที่ต้นคอ ปวดศีรษะ สำหรับผู้ที่มีความดันสูงรุนแรงอาจมีอาการ เช่น อ่อนเพลีย เหนื่อยง่าย ใจสั่น มือเท้าชา ตามัว อัมพาต หรือเสียชีวิตเฉียบพลัน เป็นต้น",
			Medications: "1. ยาขับปัสสาวะกลุ่มไธอะไซด์ไดยูเรติก (Diuretics) เช่น Hydrochlorothiazide (HCTZ), Furosemide ออกฤทธิ์โดยการขับน้ำออกจากหลอดเลือด ทำให้ปริมาตรเลือดในหลอดเลือดลดลง ความดันโลหิตจึงลดลง\n2. ยาลดความดันโลหิตกลุ่ม ACE inhibitors เช่น Captopril , Enalapril , Perindopril, Ramipril ยานี้มีฤทธิ์ยับยั้งเอนไซม์แองจิโอเทนซินคอนเวอร์ติ้ง (angiotensin converting enzymes) ออกฤทธิ์ด้วยกลไกหลายอย่าง เช่น ขยายหลอดเลือด, ขับโซเดียมและน้ำออกจากร่างกาย เป็นต้น จึงส่งผลโดยรวมทำให้สามารถลดความดันโลหิตได้\n3. ยากลุ่ม angiotensin-II receptor antagonists เช่น Candesartan , Losartan , Valsatan ออกฤทธิ์ลดความดัน และมีอาการข้างเคียงจากยาคล้ายยากลุ่ม ACE inhibitors แต่ทำให้เกิดอาการไอน้อยกว่า\n4. ยากลุ่มแคลเซียมแชลแนลบล็อกเกอร์ (Calcium channel blockers)เช่น Amlodipine , Diltiazem , Felodipine, Nifedipine, Verapamil, ออกฤทธิ์โดยการปิดกั้นการไหลเข้าของเกลือแคลเซียมทำให้กล้ามเนื้อหัวใจบีบตัวลดลง ช่วยขยายหลอดเลือด ส่งผลให้ความดันโลหิตลดลง อาการข้างเคียงจากการใช้ยา ปวดศีรษะ, หน้าแดง (flushing) ข้อเท้าบวม (ankle-edema) ท้องผูก อาจพบภาวะหัวใจเต้นเร็ว ยกเว้นในยา Verapamil, Diltiazem",
			Behaviors:   "1. รับประทานอาหารที่มีคุณภาพ โดยการลดอาหารเค็มจัด ลดอาหารมัน เพิ่มผักผลไม้ เน้นอาหารพวกธัญพืช ปลา นมไขมันต่ำ ถั่ว รับประทานอาหารที่มีไขมันอิ่มตัวต่ำ หลีกเลี่ยงเนื้อแดง น้ำตาล เครื่องดื่มที่มีรสหวานจะทำให้ระดับความดันโลหิตลดลงได้\n2. ควบคุมน้ำหนักให้อยู่ในเกณฑ์ปกติ\n3. ออกกำลังกายแบบแอโรบิก หมายถึงการออกกำลังกาย ที่มีการเคลื่อนไหวร่างกายตลอดเวลา เช่น วิ่ง เดินเร็ว ว่ายนํ้า อย่างสม่ำเสมอ อย่างน้อย วันละ 15-30 นาที 3-6 วันต่อสัปดาห์ และการควบคุมน้ำหนักให้อยู่ในเกณฑ์ปกติ\n4. จำกัดปริมาณการบริโภคเครื่องดื่มแอลกอฮอล์\n5. ห้ามสูบบุหรี่\n6. พยายามจัดการหรือลดความเครียด\n7. ติดตามความดันโลหิตที่บ้าน\n8. รักษาระดับความดันโลหิตในช่วงตั้งครรภ์",
		},
		{
			Disease:     "obesity",
			Name:        "โรคอ้วน",
			Photo:       "assets/images/obesity_article.png",
			Details:     "โรคอ้วน คือ ภาวะที่มีน้ำหนักตัว หรือสัดส่วนไขมันในร่างกายมากผิดปกติ โดยใช้เกณฑ์ดัชนีมวลกาย (Body mass index: BMI) เป็นตัวกำหนด",
			Symptoms:    "1. หายใจลำบาก เหนื่อยหอบง่าย\n2. หยุดหายใจขณะหลับ\n3. แผลหายช้า มีอาการอักเสบง่าย\n4. เจ็บปวดหน้าอก\n5. เวียนหัว หน้ามืดบ่อย ๆ",
			Medications: "1. อริสแตท (orlistat) : ขัดขวางการดูดซึมไขมันจากทางเดินอาหาร\n2. ลิรากลูไทด์ (liraglutide) : ทำให้เบื่ออาหารและน้ำหนักตัวลด\n3. เฟนเทอร์มีน (phentermine) และยาอื่นที่ออกฤทธิ์คล้ายกัน\n4. ยาสูตรผสมบูโพรพิออน/นาลเทรกโซน (bupropion/naltrexone) มีฤทธิ์กระตุ้นระบบประสาทส่วนกลางและทำให้ไม่อยากรับประทานอาหาร",
			Behaviors:   "1. การลดปริมาณแคลอรี่\n2. ทำให้อิ่มโดยการเลือกรับประทาน เช่น การบริโภคผักและผลไม้จำนวนมากที่มีแคลอรี่น้อย\n3. เลือกอาหารที่ดีต่อสุขภาพ\n4. การจำกัดปริมาณอาหารบางกลุ่ม เช่น คาร์โบไฮเดรตและไขมัน\n5. การทดแทนมื้ออาหาร",
		},
		{
			Disease:     "hyperlipidemia",
			Name:        "ภาวะไขมันในเลือดสูง",
			Photo:       "assets/images/hyperlipidemia_article.png",
			Details:     "ภาวะไขมันในเลือดสูง คือ ภาวะที่ร่างกายมีระดับไขมันในเลือดสูงกว่าเกณฑ์ปกติ โดยอาจมีความผิดปกติทั้งไขมัน “คลอเรสเตอรอล” และ “ไตรกลีเซอไรด์” ซึ่งเป็นปัจจัยเสี่ยงต่อการเกิดโรคต่างๆ เช่น โรคหัวใจและหลอดเลือด เส้นเลือดตีบ อุดตัน โรคหลอดเลือดสมอง หรือเลือดไหลเวียนไปเลี้ยงอวัยวะต่างๆ ได้ไม่มีดีพอ รวมถึงความดันโลหิตสูงได้",
			Symptoms:    "1. ความดันโลหิตสูงขึ้น\n2. เวียนศีรษะ หน้ามืด วิงเวียนเมื่อลุกนั่งเร็วๆ หรือก้มหน้านานๆ เวียนหัว\n3. ปวดหัวบ่อยๆ\n4. ใจสั่น หัวใจเต้นเร็วผิดปกติ",
			Medications: "1. ลดปริมาณอาหารที่มีโคเลสเตอรอลสูง ได้แก่ สมองสัตว์ ไข่แดง หอยนางรม ปลาหมึก กุ้ง เครื่องในสัตว์ โดยจำกัดให้ได้รับได้ไม่เกินละ 300 มิลลิกรัม ต่อวัน\n2. ลดอาหารที่มีปริมาณไขมันสูง เช่น ข้าวขาหมู ข้าวมันไก่ หมูสามชั้น หนังเป็ด - ไก่\n3. หลีกเลี่ยงการใช้น้ำมันหมู น้ำมันมะพร้าว\n4. การปรุงอาหารที่ใช้วิธี นึ่ง ต้ม  อบ ย่าง แทนการทอดหรือการใช้น้ำมันผัด\n5. งดเครื่องดื่มที่มีแอลกอฮอล์ เช่น เหล้า เบียร์\n6. ควบคุมน้ำหนักตัวให้อยู่ในเกณฑ์ปกติ\n7. งดการสูบบุหรี่\n8. ออกกำลังกายอย่างสม่ำเสมอ",
			Behaviors:   "1. กลุ่มยาสแตติน (Statins) เช่น  atorvastatin (อะทอร์วาสะแตติน), fluvastatin (ฟลูวาสะแตติน), pitavastatin (พิทาวาสะแตติน), pravastatin (พราวาสะแตติน), rosuvastatin (โรซูวาสะแตติน) และ simvastatin (ซิมวาสะแตติน)  กลุ่มยาสแตตินเป็นกลุ่มยาชนิดแรก ๆ ที่แพทย์มักใช้ในการรักษาไขมันในเลือด ซึ่งยากลุ่มนี้จะช่วยลดระดับคอเลสเตอรอลชนิดไม่ดี โดยการยับยั้งการทำงานของเอนไซม์ชนิดหนึ่งในร่างกายที่มีหน้าที่ช่วยผลิตคอเลสเตอรอลชนิดไม่ดี นอกจากนี้ ยาในกลุ่มสแตตินยังช่วยลดระดับไตรกลีเซอร์ไรด์ เพิ่มระดับคอเลสเตอรอลชนิดดี และช่วยลดความเสี่ยงในการเกิดหลอดเลือดอุดตัน\n2. ยากลุ่มไฟเบรต (Fibrates) เช่น ยาฟีโนไฟเบรต (Fenofibrate) และ ยาเจมไฟโบรซิล (Gemfibrozil) เป็นต้น ยากลุ่มไฟเบรตเป็นกลุ่มยาที่ใช้สำหรับลดระดับไขมันในเลือด โดยเฉพาะไตรกลีเซอร์ไรด์ โดยยาจะออกฤทธิ์ลดกระบวนการผลิตไลโปโปรตีนชนิดหนึ่งที่มีหน้าที่ลำเลียงไตรกลีเซอร์ไรด์ในกระแสเลือด และเร่งกระบวนการกำจัดไตรกลีเซอร์ไรด์จากเลือด รวมถึงช่วยเพิ่มระดับคอเลสเตอรอลชนิดดีในเลือด",
		},
	}
	db.Create(&knowledgeData)

}
