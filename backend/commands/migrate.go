package commands

import (
	"backend/config"
	md "backend/models"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
	"gorm.io/gorm"
)

var (
	DB = config.DB
)

var structSlice []interface{}

func SeriesOfCreation(MQDB *gorm.DB) {
	structSlice = append(structSlice, &md.Organization{},&md.Project{},&md.User{},&md.OrganizationAuthority{},&md.ProjectAuthority{},&md.Task{},&md.Field{},&md.Type{},&md.Status{},&md.Priority{},&md.Milestone{},&md.Comment{},&md.Authority{},&md.Activity{},&md.ActivityContent{})
	for _, v := range structSlice {
		Migrate(MQDB, v)
	}
	CreateActivityContent(MQDB)
	CreateOrganization(MQDB)
	CreateProject(MQDB)
	CreateUser(MQDB)
	CreateOrganizationAuthority(MQDB)
	CreateProjectAuthority(MQDB)
	CreateField(MQDB)
	CreateMilestone(MQDB)
	CreateStatus(MQDB)
	CreateType(MQDB)
	CreatePriority(MQDB)
	CreateTask(MQDB)
	// CreateComment(MQDB)
	CreateAuthority(MQDB)
}

func AutoMigrate(MQDB *gorm.DB) {
	MQDB.AutoMigrate(&md.Organization{},&md.Project{},&md.User{},&md.OrganizationAuthority{},&md.ProjectAuthority{},&md.Task{},&md.Field{},&md.Type{},&md.Status{},&md.Priority{},&md.Milestone{},&md.Comment{},&md.Authority{},&md.Activity{},&md.ActivityContent{})
}

func Migrate(MQDB *gorm.DB, tableStruct interface{}) {
	MQDB.AutoMigrate(&tableStruct)
	MQDB.Migrator().DropTable(&tableStruct)
	MQDB.Migrator().CreateTable(&tableStruct)
}

func CreateActivityContent(MQDB *gorm.DB) {
	result := make([]md.ActivityContent, 0)
	for _, value := range []string{"タスクを作成しました", "タスクを変更しました", "コメントを作成しました", "コメントを変更しました", "プロジェクトを作成しました", "プロジェクトを変更しました"} {
		var newActivityContent = md.ActivityContent{
			Content: value,
		}
		result = append(result, newActivityContent)
	}
	MQDB.Create(&result)
}

func CreateOrganization(MQDB *gorm.DB) md.Organization {
	var newOrganization = md.Organization{
		ID:					"prygen4fDISDVgSYDjxZ5uICD",
		Name:				"+base",
		Address:			"東京都渋谷区円山町5番5号 Navi渋谷Ⅴ3階",
		Number:				"080-1234-5678",
		Founded:			"2021年4月1日",
		Description:		"看護師を経験し、すべての医療従事者に「心のコップを満たす習慣」を広めようと思った",
		Image:				"base.png",
		Plan:				"standard",
	}
	MQDB.Create(&newOrganization)
	return newOrganization
}

func CreateUser(MQDB *gorm.DB) md.User {
	var newUser = md.User{
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
	MQDB.Create(&newUser)
	return newUser
}

func CreateProject(MQDB *gorm.DB) {
	var newProject = md.Project{
		Name:      		"性格診断",
		Image:     		"diagnosis.png",
		OrganizationID: "prygen4fDISDVgSYDjxZ5uICD",
	}
	MQDB.Create(&newProject)
}

func CreateProjectAuthority(MQDB *gorm.DB) {
	var newProjectAuthority = md.ProjectAuthority{
		ProjectID:		1,
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	MQDB.Create(&newProjectAuthority)
}

func CreateOrganizationAuthority(MQDB *gorm.DB) {
	var newOrganizationAuthority = md.OrganizationAuthority{
		OrganizationID:	"prygen4fDISDVgSYDjxZ5uICD",
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	MQDB.Create(&newOrganizationAuthority)
}

func CreateStatus(MQDB *gorm.DB) {
	result := make([]md.Status, 0)
	for _, value := range []string{"相談", "依頼", "再議", "未対応", "対応中", "中断", "確認", "調整", "完了"} {
		var newStatus = md.Status{
			Name: value,
		}
		result = append(result, newStatus)
	}
	MQDB.Create(&result)
}

func CreateField(MQDB *gorm.DB) {
	var fields []md.Field
	fieldTypes := []string{"エンジニア", "医療", "データサイエンス", "営業"}
	for _, fieldType := range fieldTypes {
		newField := md.Field{
			Name: fieldType,
			ProjectID: 1,
		}
		fields = append(fields, newField)
	}
	MQDB.Create(&fields)
}

func CreateMilestone(MQDB *gorm.DB) {
	var milestones []md.Milestone
	var projects []md.Project
	milestoneTypes := []string{"フェーズ1", "フェーズ2", "フェーズ3", "フェーズ4", "テスト", "ベータ版公開"}
	DB.Find(&projects)
	for _, milestoneType := range milestoneTypes {
		milestone := md.Milestone{
			Name: milestoneType,
			ProjectID: 1,
		}
		milestones = append(milestones, milestone)
	}
	MQDB.Create(&milestones)
}

func CreateType(MQDB *gorm.DB) {
	result := make([]md.Type, 0)
	for _, value := range []string{"追加", "変更", "バグ", "その他"} {
		var newType = md.Type{
			Name: value,
		}
		result = append(result, newType)
	}
	MQDB.Create(&result)
}

func CreatePriority(MQDB *gorm.DB) {
	result := make([]md.Priority, 0)
	for _, value := range []string{"低", "中", "高"} {
		var newPriority = md.Priority{
			Name: value,
		}
		result = append(result, newPriority)
	}
	MQDB.Create(&result)
}

func CreateTask(MQDB *gorm.DB) {
	fieldID := 1
	milestoneID := 1
	var newTask = md.Task{
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
	MQDB.Create(&newTask)
}

// func CreateComment(MQDB *gorm.DB) {
// 	var newComment1 = Comment{
// 		Content: "これで大丈夫ですか？",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: true,
// 	}
// 	SQDB.Create(&newComment1)
// 	MQDB.Create(&newComment1)
// 	var newComment2 = Comment{
// 		Content: "あとこれはどうしましょうか。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: true,
// 	}
// 	SQDB.Create(&newComment2)
// 	MQDB.Create(&newComment2)
// 	var newComment3 = Comment{
// 		Content: "これで大丈夫です。ありがとうございます。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: false,
// 	}
// 	SQDB.Create(&newComment3)
// 	MQDB.Create(&newComment3)
// 	SQDB.Create(&CommentReply{1,2,3})
// 	MQDB.Create(&CommentReply{2,2,3})
// 	var newComment4 = Comment{
// 		Content: "かしこまりました。",
// 		TaskID: 1,
// 		UserID: 1,
// 		ParentID: false,
// 	}
// 	SQDB.Create(&newComment4)
// 	MQDB.Create(&newComment4)
// 	SQDB.Create(&CommentReply{3,3,4})
// 	MQDB.Create(&CommentReply{4,3,4})
// }

func CreateAuthority(MQDB *gorm.DB) {
	result := make([]md.Authority, 0)
	for _, value := range []string{"管理者", "一般"} {
		var newAuthority = md.Authority{
			Name: value,
		}
		result = append(result, newAuthority)
	}
	MQDB.Create(&result)
}

func DeleteUsersImages() {
	files, _ := ioutil.ReadDir("./upload/users")
	if len(files) != 0 {
		for _, v := range files {
			if v.Name() != "rio1.png" || v.Name() != "rio2.jpg" || v.Name() != "rio3.png" || v.Name() != "teiko.png" {
				os.Remove("./upload/user/" + v.Name())
			}
		}
	}
}

func DeleteOrganizationsImages() {
	files, _ := ioutil.ReadDir("./upload/organizations")
	if len(files) != 0 {
		for _, v := range files {
			if v.Name() != "base.png"{
				os.Remove("./upload/organizations/" + v.Name())
			}
		}
	}
}

func DeleteProjectsImages() {
	files, _ := ioutil.ReadDir("./upload/projects")
	if len(files) != 0 {
		for _, v := range files {
			if v.Name() != "panda.png" {
				os.Remove("./upload/projects/" + v.Name())
			}
		}
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}