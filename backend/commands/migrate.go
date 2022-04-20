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

func SeriesOfCreation(SQDB *gorm.DB, MQDB *gorm.DB) {
	structSlice = append(structSlice, &md.Organization{},&md.Project{},&md.User{},&md.OrganizationAuthority{},&md.ProjectAuthority{},&md.Task{},&md.Field{},&md.Type{},&md.Status{},&md.Priority{},&md.Milestone{},&md.Comment{},&md.Authority{},&md.Activity{},&md.ActivityContent{})
	for _, v := range structSlice {
		Migrate(SQDB, MQDB, v)
	}
	CreateActivityContent(SQDB, MQDB)
	CreateOrganization(SQDB, MQDB)
	CreateProject(SQDB, MQDB)
	CreateUser(SQDB, MQDB)
	CreateOrganizationAuthority(SQDB, MQDB)
	CreateProjectAuthority(SQDB, MQDB)
	CreateField(SQDB, MQDB)
	CreateMilestone(SQDB, MQDB)
	CreateStatus(SQDB, MQDB)
	CreateType(SQDB, MQDB)
	CreatePriority(SQDB, MQDB)
	CreateTask(SQDB, MQDB)
	// CreateComment(SQDB, MQDB)
	CreateAuthority(SQDB, MQDB)
}

func AutoMigrate(SQDB *gorm.DB, MQDB *gorm.DB) {
	SQDB.AutoMigrate(&md.Organization{},&md.Project{},&md.User{},&md.OrganizationAuthority{},&md.ProjectAuthority{},&md.Task{},&md.Field{},&md.Type{},&md.Status{},&md.Priority{},&md.Milestone{},&md.Comment{},&md.Authority{},&md.Activity{},&md.ActivityContent{})
	MQDB.AutoMigrate(&md.Organization{},&md.Project{},&md.User{},&md.OrganizationAuthority{},&md.ProjectAuthority{},&md.Task{},&md.Field{},&md.Type{},&md.Status{},&md.Priority{},&md.Milestone{},&md.Comment{},&md.Authority{},&md.Activity{},&md.ActivityContent{})
}

func Migrate(SQDB *gorm.DB, MQDB *gorm.DB, tableStruct interface{}) {
	SQDB.AutoMigrate(&tableStruct)
	MQDB.AutoMigrate(&tableStruct)
	SQDB.Migrator().DropTable(&tableStruct)
	SQDB.Migrator().CreateTable(&tableStruct)
	MQDB.Migrator().DropTable(&tableStruct)
	MQDB.Migrator().CreateTable(&tableStruct)
}

func CreateActivityContent(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]md.ActivityContent, 0)
	for _, value := range []string{"タスクを作成しました", "タスクを変更しました", "コメントを作成しました", "コメントを変更しました", "プロジェクトを作成しました", "プロジェクトを変更しました"} {
		var newActivityContent = md.ActivityContent{
			Content: value,
		}
		result = append(result, newActivityContent)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreateOrganization(SQDB *gorm.DB, MQDB *gorm.DB) md.Organization {
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
	SQDB.Create(&newOrganization)
	MQDB.Create(&newOrganization)
	return newOrganization
}

func CreateUser(SQDB *gorm.DB, MQDB *gorm.DB) md.User {
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
	SQDB.Create(&newUser)
	MQDB.Create(&newUser)
	return newUser
}

func CreateProject(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newProject = md.Project{
		Name:      		"性格診断",
		Image:     		"diagnosis.png",
		OrganizationID: "prygen4fDISDVgSYDjxZ5uICD",
	}
	SQDB.Create(&newProject)
	MQDB.Create(&newProject)
}

func CreateProjectAuthority(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newProjectAuthority = md.ProjectAuthority{
		ProjectID:		1,
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	SQDB.Create(&newProjectAuthority)
	MQDB.Create(&newProjectAuthority)
}

func CreateOrganizationAuthority(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newOrganizationAuthority = md.OrganizationAuthority{
		OrganizationID:	"prygen4fDISDVgSYDjxZ5uICD",
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	SQDB.Create(&newOrganizationAuthority)
	MQDB.Create(&newOrganizationAuthority)
}

func CreateStatus(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]md.Status, 0)
	for _, value := range []string{"相談", "依頼", "再議", "未対応", "対応中", "中断", "確認", "調整", "完了"} {
		var newStatus = md.Status{
			Name: value,
		}
		result = append(result, newStatus)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreateField(SQDB *gorm.DB, MQDB *gorm.DB) {
	var fields []md.Field
	fieldTypes := []string{"エンジニア", "医療", "データサイエンス", "営業"}
	for _, fieldType := range fieldTypes {
		newField := md.Field{
			Name: fieldType,
			ProjectID: 1,
		}
		fields = append(fields, newField)
	}
	SQDB.Create(&fields)
	MQDB.Create(&fields)
}

func CreateMilestone(SQDB *gorm.DB, MQDB *gorm.DB) {
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
	SQDB.Create(&milestones)
	MQDB.Create(&milestones)
}

func CreateType(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]md.Type, 0)
	for _, value := range []string{"追加", "変更", "バグ", "その他"} {
		var newType = md.Type{
			Name: value,
		}
		result = append(result, newType)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreatePriority(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]md.Priority, 0)
	for _, value := range []string{"低", "中", "高"} {
		var newPriority = md.Priority{
			Name: value,
		}
		result = append(result, newPriority)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreateTask(SQDB *gorm.DB, MQDB *gorm.DB) {
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
	SQDB.Create(&newTask)
	MQDB.Create(&newTask)
}

// func CreateComment(SQDB *gorm.DB, MQDB *gorm.DB) {
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

func CreateAuthority(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]md.Authority, 0)
	for _, value := range []string{"管理者", "一般"} {
		var newAuthority = md.Authority{
			Name: value,
		}
		result = append(result, newAuthority)
	}
	SQDB.Create(&result)
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