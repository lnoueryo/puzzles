package commands

import (
	"time"
	"gorm.io/gorm"
)


func SQlite(SQDB *gorm.DB) {
	SQCreateOrganization(SQDB)
	SQCreateProject(SQDB)
	SQCreateUser(SQDB)
	SQCreateOrganizationAuthority(SQDB)
	SQCreateProjectAuthority(SQDB)
	SQCreateField(SQDB)
	SQCreateMilestone(SQDB)
	SQCreateStatus(SQDB)
	SQCreateType(SQDB)
	SQCreatePriority(SQDB)
	SQCreateTask(SQDB)
	// SQCreateComment(SQDB)
	SQCreateAuthority(SQDB)
}


func SQCreateOrganization(SQDB *gorm.DB) Organization {
	var newOrganization = Organization{
		ID:					"prygen4fDISDVgSYDjxZ5uICD",
		Name:				"+base",
		Address:			"東京都渋谷区円山町5番5号 Navi渋谷Ⅴ3階",
		Number:				"080-1234-5678",
		Founded:			"2021年4月1日",
		Description:		"看護師を経験し、すべての医療従事者に「心のコップを満たす習慣」を広めようと思った",
		Image:				"base.png",
		Plan:				"standard",
	}
	SQDB.Create(&newOrganization)
	return newOrganization
}

func SQCreateUser(SQDB *gorm.DB) User {
	var newUser = User{
		ID: 1,
		Name:      "井上領",
		Age:      31,
		Sex:      "男",
		Email:     "popo62520908@gmail.com",
		Address:     "東京都世田谷区松原1-43-14",
		Password:  "52f96e51831c8229413e28b0e58fa3b992f7571e4ff5bf5ccfc1a21f391e4f05",
		Image:     "CWCM67iUYAAZ1Kp.png",
		Description: "hello",
	}
	SQDB.Create(&newUser)

	return newUser
}

func SQCreateProject(SQDB *gorm.DB) {
	var newProject = Project{
		Name:      		"性格診断",
		Image:     		"base.png",
		OrganizationID: "prygen4fDISDVgSYDjxZ5uICD",
	}
	SQDB.Create(&newProject)
}

func SQCreateProjectAuthority(SQDB *gorm.DB) {
	var newProjectAuthority = ProjectAuthority{
		ProjectID:		1,
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	SQDB.Create(&newProjectAuthority)
}

func SQCreateOrganizationAuthority(SQDB *gorm.DB) {
	var newOrganizationAuthority = OrganizationAuthority{
		OrganizationID:	"prygen4fDISDVgSYDjxZ5uICD",
		UserID:			1,
		AuthorityID:	1,
	}
	SQDB.Create(&newOrganizationAuthority)
}

func SQCreateStatus(SQDB *gorm.DB) {
	result := make([]Status, 0)
	for _, value := range []string{"相談", "依頼", "再議", "未対応", "対応中", "中断", "確認", "調整", "完了"} {
		var newStatus = Status{
			Name: value,
		}
		result = append(result, newStatus)
	}
	SQDB.Create(&result)
}

func SQCreateField(SQDB *gorm.DB) {
	var fields []Field
	fieldTypes := []string{"エンジニア", "医療", "データサイエンス", "営業"}
	for _, fieldType := range fieldTypes {
		newField := Field{
			Name: fieldType,
			ProjectID: 1,
		}
		fields = append(fields, newField)
	}
	SQDB.Create(&fields)
}

func SQCreateMilestone(SQDB *gorm.DB) {
	var milestones []Milestone
	var projects []Project
	milestoneTypes := []string{"フェーズ1", "フェーズ2", "フェーズ3", "フェーズ4", "テスト", "ベータ版公開"}
	DB.Find(&projects)
	for _, milestoneType := range milestoneTypes {
		milestone := Milestone{
			Name: milestoneType,
			ProjectID: 1,
		}
		milestones = append(milestones, milestone)
	}
	SQDB.Create(&milestones)
}

func SQCreateType(SQDB *gorm.DB) {
	result := make([]Type, 0)
	for _, value := range []string{"追加", "変更", "バグ", "その他"} {
		var newType = Type{
			Name: value,
		}
		result = append(result, newType)
	}
	SQDB.Create(&result)
}

func SQCreatePriority(SQDB *gorm.DB) {
	result := make([]Priority, 0)
	for _, value := range []string{"低", "中", "高"} {
		var newPriority = Priority{
			Name: value,
		}
		result = append(result, newPriority)
	}
	SQDB.Create(&result)
}

func SQCreateTask(SQDB *gorm.DB) {
	fieldID := 1
	milestoneID := 1
	var newTask = Task{
		AssigneeID:  	1,
		AssignerID:  	1,
		StatusID:     1,
		FieldID:     &fieldID,
		MilestoneID:  &milestoneID,
		TypeID:     	1,
		PriorityID:     1,
		ProjectID:     1,
		Key: 	   "性格診断_1",
		Title:     "課題を作成",
		Detail:     "<strong>課題を作成してください</strong>",
		Deadline: time.Now(),
	}
	SQDB.Create(&newTask)
}

// func SQCreateComment(SQDB *gorm.DB) {
// 	var newComment1 = Comment{
// 		Content: "これで大丈夫ですか？",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: true,
// 	}
// 	SQDB.Create(&newComment1)
// 	.Create(&newComment1)
// 	var newComment2 = Comment{
// 		Content: "あとこれはどうしましょうか。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: true,
// 	}
// 	SQDB.Create(&newComment2)
// 	.Create(&newComment2)
// 	var newComment3 = Comment{
// 		Content: "これで大丈夫です。ありがとうございます。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: false,
// 	}
// 	SQDB.Create(&newComment3)
// 	.Create(&newComment3)
// 	SQDB.Create(&CommentReply{1,2,3})
// 	.Create(&CommentReply{2,2,3})
// 	var newComment4 = Comment{
// 		Content: "かしこまりました。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: false,
// 	}
// 	SQDB.Create(&newComment4)
// 	.Create(&newComment4)
// 	SQDB.Create(&CommentReply{3,3,4})
// 	.Create(&CommentReply{4,3,4})
// }

func SQCreateAuthority(SQDB *gorm.DB) {
	result := make([]Authority, 0)
	for _, value := range []string{"管理者", "一般"} {
		var newAuthority = Authority{
			Name: value,
		}
		result = append(result, newAuthority)
	}
	SQDB.Create(&result)
}
