package commands

import (
	"backend/config"
	"backend/modules/crypto"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
	"gorm.io/gorm"
)

type Organization struct {
	ID					string		`gorm:"primaryKey;type:varchar(100);unique_index"json:"id"`
	Name				string		`json:"name"`
	Address				string		`json:"address"`
	Number				string		`json:"number"`
	Founded				string		`json:"founded"`
	Image				string		`json:"image"`
	Description			string		`json:"description"`
	Plan				string		`json:"plan"`
	CreditCard			string		`json:"creditcard"`
	Expiry				string		`json:"expiry"`
	Projects			[]Project	`json:"projects"`
	Users				[]OrganizationAuthority	`gorm:"foreignkey:OrganizationID;migrate;"json:"users"`
	CreatedAt			time.Time	`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt			time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}
type OrganizationAuthority struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
	UserID			int				`json:"user_id"`
	OrganizationID	string			`json:"organization_id"`
    AuthorityID		int				`json:"auth_id"`
	Active			bool			`json:"active"`
	Verification	string			`json:"verification"`
	User 			User			`gorm:"foreignkey:UserID;migrate"json:"user"`
	Organization 	Organization	`gorm:"foreignkey:OrganizationID;migrate"json:"organization"`
	Type			Authority		`gorm:"foreignkey:AuthorityID;migrate"json:"type"`
	CreatedAt		time.Time		`gorm:"<-:false;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time		`gorm:"<-;autoUpdateTime;"json:"-"`
}

type Project struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
    OrganizationID	string			`gorm:"<-:create;"json:"organization_id"`
	Name			string			`json:"name"`
	Description		string			`json:"description"`
	Image			string			`json:"image"`
	Organization	Organization	`gorm:"->;references:ID;"json:"organization"`
	ImageData		string			`gorm:"migration;"json:"image_data"`
	Authority		string			`gorm:"migration"json:"authority"`
	Tasks			[]Task			`json:"tasks"`
	Milestones		[]Milestone		`json:"milestones"`
	Fields			[]Field			`json:"fields"`
	AuthorityUsers  []ProjectAuthority	`json:"authority_users"`
	Users			[]User			`gorm:"many2many:project_authorities;"json:"users"`
	CreatedAt		time.Time		`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time		`gorm:"autoUpdateTime;"json:"updated_at"`
}

type ProjectAuthority struct {
	ID				int					`gorm:"AUTO_INCREMENT"json:"id"`
    ProjectID		int					`json:"project_id"`
    UserID			int					`json:"user_id"`
    AuthorityID			int					`json:"auth_id"`
	Active			bool				`json:"active"`
	Type			Authority			`gorm:"foreignkey:AuthorityID;migrate"json:"type"`
	User			User				`gorm:"foreignkey:UserID;migrate"json:"user"`
	Project			Project				`gorm:"foreignkey:ProjectID;migrate"json:"project"`
	ProjectUsers	[]ProjectAuthority	`gorm:"foreignkey:ProjectID;references:ProjectID;migrate;"json:"project_users"`
	CreatedAt		time.Time			`gorm:"<-:create;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time			`json:"-"`
}


type User struct {
	ID						int						`gorm:"AUTO_INCREMENT"json:"id"`
	Name					string					`json:"name"`
	Age						int						`json:"age"`
	Sex						string					`json:"sex"`
	Email					string					`json:"email"`
	Address					string					`json:"address"`
	Password				string					`gorm:"->:false;<-:create"json:"password"`
	Image					string					`json:"image"`
	Description				string					`json:"description"`
	Organization			string					`gorm:"migration"json:"organization"`
	Authority				string					`gorm:"migration"json:"authority"`
	Projects				[]ProjectAuthority		`gorm:"migration"json:"projects"`
	Organizations			[]OrganizationAuthority	`gorm:"foreignkey:UserID;migrate;"json:"organizations"`
	Tasks					[]Task					`gorm:"foreignKey:AssigneeID;references:ID"json:"tasks"`
	CreatedAt				time.Time				`gorm:"->:false;<-:create;autoCreateTime;"json:"-"`
	UpdatedAt				time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}

type Field struct {
	ID			int			`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID	int			`json:"project_id"`
	Name		string		`json:"name"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}

type Task struct {
	ID        	  int       `gorm:"AUTO_INCREMENT"json:"id"`
	AssigneeID	  int       `gorm:"<-;not null"json:"assignee_id"`
	AssignerID	  int       `gorm:"<-;not null"json:"assigner_id"`
	StatusID      int       `gorm:"<-;not null"json:"status_id"`
	FieldID       *int		`gorm:"<-;not null;"json:"field_id"`
	MilestoneID   *int		`gorm:"<-;not null;"json:"milestone_id"`
	PriorityID    int       `gorm:"<-;not null"json:"priority_id"`
	TypeID        int       `gorm:"<-;not null"json:"type_id"`
	ProjectID     int       `gorm:"<-;not null"json:"project_id"`
	ParentID	  int    	`gorm:"<-;not null"json:"parent_id"`
	Key        	  string    `json:"key"`
	Title         string    `gorm:"<-;not null"json:"title" sql:"CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	Detail        string    `gorm:"<-;"json:"detail"`
	EstimatedTime float32   `json:"estimated_time"`
	ActualTime    float32   `json:"actual_time"`
	StartTime     time.Time `gorm:"default:null;"json:"start_time"`
	Deadline      time.Time `gorm:"default:null;"json:"deadline"`
	// Status        Status    `gorm:"embedded;embeddedPrefix:status_"`
	Status        Status    `gorm:"references:ID;"json:"status"`
	Field         Field    	`gorm:"references:ID"json:"field"`
	Milestone     Milestone `gorm:"references:ID"json:"milestone"`
	Type          Type    	`gorm:"references:ID"json:"type"`
	Priority      Priority  `gorm:"references:ID"json:"priority"`
	Assignee  	  User      `gorm:"references:ID;foreignKey:AssigneeID"json:"assignee"`
	Assigner  	  User      `gorm:"references:ID;foreignKey:AssignerID"json:"assigner"`
	Comments	  []Comment `json:"comments"`
	CreatedAt 	  time.Time `gorm:"<-:create;autoCreateTime;"json:"created_at"`
	UpdatedAt 	  time.Time `gorm:"autoUpdateTime;"json:"updated_at"`
}

type Status struct {
	ID   int    `gorm:"AUTO_INCREMENT"json:"id"`
	Name string `json:"name"`
}

type Type struct {
	ID   int    `gorm:"AUTO_INCREMENT"json:"id"`
	Name string `json:"name"`
}

type Milestone struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	Name 	  string	`json:"name"`
	ProjectID int		`json:"project_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"-"`
}

type Priority struct {
	ID   int    `gorm:"AUTO_INCREMENT"json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	Content	  string	`json:"content"`
	TaskID 	  int		`json:"task_id"`
	UserID 	  int		`json:"user_id"`
	User 	  User		`gorm:"foreignkey:UserID;"json:"user"`
	ParentID  int		`json:"parent_id"`
	Replies	[]Comment 	`gorm:"foreignKey:ParentID"json:"replies"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"updated_at"`
}

type Authority struct {
	ID int
	Name string
}

type Activity struct {
	ID			int				`gorm:"AUTO_INCREMENT"json:"id"`
	UserID		int				`json:"user_id"`
	ProjectID	int				`json:"project_id"`
	ContentID	int				`json:"content_id"`
	User		User			`gorm:"references:ID;"json:"user"`
	Content		ActivityContent	`gorm:"references:ID;"json:"content"`
	CreatedAt	time.Time 		`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt	time.Time 		`gorm:"autoUpdateTime;"json:"updated_at"`
}

type ActivityContent struct {
	ID			int		`gorm:"AUTO_INCREMENT"json:"id"`
	Content		string	`json:"content"`
}

var (
	DB = config.DB
)

var structSlice []interface{}

func SeriesOfCreation(SQDB *gorm.DB, MQDB *gorm.DB) {
	structSlice = append(structSlice, &Organization{},&Project{},&User{},&OrganizationAuthority{},&ProjectAuthority{},&Task{},&Field{},&Type{},&Status{},&Priority{},&Milestone{},&Comment{},&Authority{},&Activity{},&ActivityContent{})
	for _, v := range structSlice {
		Migrate(SQDB, MQDB, v)
	}
	// CreateOrganization(SQDB, MQDB)
	// CreateProject(SQDB, MQDB)
	// CreateUser(SQDB, MQDB)
	// CreateOrganizationAuthority(SQDB, MQDB)
	// CreateProjectAuthority(SQDB, MQDB)
	// CreateField(SQDB, MQDB)
	// CreateMilestone(SQDB, MQDB)
	// CreateStatus(SQDB, MQDB)
	// CreateType(SQDB, MQDB)
	// CreatePriority(SQDB, MQDB)
	// CreateTask(SQDB, MQDB)
	// // CreateComment(SQDB, MQDB)
	// CreateAuthority(SQDB, MQDB)
}

func AutoMigrate(SQDB *gorm.DB, MQDB *gorm.DB) {
	SQDB.AutoMigrate(&Organization{},&Project{},&User{},&OrganizationAuthority{},&ProjectAuthority{},&Task{},&Field{},&Type{},&Status{},&Priority{},&Milestone{},&Comment{},&Authority{},&Activity{},&ActivityContent{})
	MQDB.AutoMigrate(&Organization{},&Project{},&User{},&OrganizationAuthority{},&ProjectAuthority{},&Task{},&Field{},&Type{},&Status{},&Priority{},&Milestone{},&Comment{},&Authority{},&Activity{},&ActivityContent{})
}

func Migrate(SQDB *gorm.DB, MQDB *gorm.DB, tableStruct interface{}) {
	SQDB.AutoMigrate(&tableStruct)
	MQDB.AutoMigrate(&tableStruct)
	SQDB.Migrator().DropTable(&tableStruct)
	SQDB.Migrator().CreateTable(&tableStruct)
	MQDB.Migrator().DropTable(&tableStruct)
	MQDB.Migrator().CreateTable(&tableStruct)
}


func CreateOrganization(SQDB *gorm.DB, MQDB *gorm.DB) Organization {
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
	MQDB.Create(&newOrganization)
	return newOrganization
}

func CreateUser(SQDB *gorm.DB, MQDB *gorm.DB) User {
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
	MQDB.Create(&newUser)
	return newUser
}

func CreateProject(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newProject = Project{
		Name:      		"性格診断",
		Image:     		"base.png",
		OrganizationID: "prygen4fDISDVgSYDjxZ5uICD",
	}
	SQDB.Create(&newProject)
	MQDB.Create(&newProject)
}

func CreateProjectAuthority(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newProjectAuthority = ProjectAuthority{
		ProjectID:		1,
		UserID:			1,
		AuthorityID:	1,
		Active: true,
	}
	SQDB.Create(&newProjectAuthority)
	MQDB.Create(&newProjectAuthority)
}

func CreateOrganizationAuthority(SQDB *gorm.DB, MQDB *gorm.DB) {
	var newOrganizationAuthority = OrganizationAuthority{
		OrganizationID:	"prygen4fDISDVgSYDjxZ5uICD",
		UserID:			1,
		AuthorityID:	1,
	}
	SQDB.Create(&newOrganizationAuthority)
	MQDB.Create(&newOrganizationAuthority)
}

func CreateStatus(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]Status, 0)
	for _, value := range []string{"相談", "依頼", "再議", "未対応", "対応中", "中断", "確認", "調整", "完了"} {
		var newStatus = Status{
			Name: value,
		}
		result = append(result, newStatus)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreateField(SQDB *gorm.DB, MQDB *gorm.DB) {
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
	MQDB.Create(&fields)
}

func CreateMilestone(SQDB *gorm.DB, MQDB *gorm.DB) {
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
	MQDB.Create(&milestones)
}

func CreateType(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]Type, 0)
	for _, value := range []string{"追加", "変更", "バグ", "その他"} {
		var newType = Type{
			Name: value,
		}
		result = append(result, newType)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func CreatePriority(SQDB *gorm.DB, MQDB *gorm.DB) {
	result := make([]Priority, 0)
	for _, value := range []string{"低", "中", "高"} {
		var newPriority = Priority{
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
	result := make([]Authority, 0)
	for _, value := range []string{"管理者", "一般"} {
		var newAuthority = Authority{
			Name: value,
		}
		result = append(result, newAuthority)
	}
	SQDB.Create(&result)
	MQDB.Create(&result)
}

func (u *User) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/user/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	u.Image = filename
}

func (o *Organization) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/organizations/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	o.Image = filename
}

func (p *Project) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/projects/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	p.Image = filename
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