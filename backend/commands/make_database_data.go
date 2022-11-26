package commands

import (
	"backend/config"
	"backend/models"
	md "backend/models"
	"backend/modules/crypto"
	"backend/modules/image"
	"backend/modules/storage"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const BUCKET_NAME = "puzzles-media"

func MakeDBData(name string, arg1 string) error {
	if arg1 != "" {
		err := fmt.Errorf(fmt.Sprintf(`no command "go run main.go %v"`, CreateArgsText()))
		return err
	}
	allDatabase()
	// err := fmt.Errorf(fmt.Sprintf(`no command "go run main.go %v"`, CreateArgsText()))
	return nil
}

func allDatabase() {
	// dbconf := `puzzles:password@tcp(localhost:3306)/puzzle?parseTime=true&charset=utf8&loc=Local`
	// MQDB, _ := gorm.Open(mysql.Open(dbconf), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// })
	// SeriesOfCreation(MQDB)
	start := time.Now()
	// AutoMigrate(MQDB)
	// SQlite(SQDB)
	// fmt.Println(c)
	// var t Task
	// DB = DB.Preload("Comments", "parent_id = ?", 1)
	// DB = RecursivePreload(DB)
	// DB.First(&t, 1)
	// DB.Preload(RecursivePreload()).Preload("Comments", "parent_id = ?", 1).Preload(clause.Associations).First(&t, 1)
	// CreateOrganizations(MQDB)
	// CreateOrganizationUsers(MQDB)
	// CreateProjectUsers(MQDB)
	// CreateUser(MQDB)
	// CreateUsers(MQDB)
	// CreateProjects(MQDB)
	// createProjectAuthority(MQDB)
	// CreateFields(MQDB)
	// CreateMilestones(MQDB)
	// CreateVersions(MQDB)
	// CreateTasks(MQDB)
	// createFieldUsers(MQDB)
	// CreateComments(MQDB)
	// createProjectUsers(MQDB)
	// ReadOrganization()
	// MQDB.Migrator().DropTable(&OrganizationAuth{})
	// MQDB.Migrator().CreateTable(&OrganizationAuth{})
	// Task(MQDB)
	// createUserImage()
	// createOrganizationImage()
	createProjectImage()
	fmt.Println((time.Now()).Sub(start))
	// readTask(MQDB)
}

func Task(MQDB *gorm.DB) {
    f, err := os.Open("task1.csv")
    if err != nil {
        fmt.Print(err)
    }

    r := csv.NewReader(f)
    var tasks []md.Task
	var timestamp = "2006-01-02 15:04:05"
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
		id, _ := strconv.Atoi(record[0])
		assignee_id, _ := strconv.Atoi(record[1])
		assigner_id, _ := strconv.Atoi(record[2])
		status_id, _ := strconv.Atoi(record[3])
		field_id, _ := strconv.Atoi(record[4])
		milestone_id, _ := strconv.Atoi(record[5])
		version_id, _ := strconv.Atoi(record[6])
		priority_id, _ := strconv.Atoi(record[7])
		type_id, _ := strconv.Atoi(record[8])
		project_id, _ := strconv.Atoi(record[9])
		parent_id, _ := strconv.Atoi(record[10])
		EstimatedTime, _ := strconv.ParseFloat(record[14], 32)
		ActualTime, _ := strconv.ParseFloat(record[15], 32)
		StartTime, _ := time.Parse(timestamp, record[16])
		Deadline, _ := time.Parse(timestamp, record[17])
		CreatedAt, _ := time.Parse(timestamp, record[18])
		UpdatedAt, _ := time.Parse(timestamp, record[19])
        task := md.Task{
			ID: id,
			AssigneeID: assignee_id,
			AssignerID: assigner_id,
			StatusID: status_id,
			FieldID: &field_id,
			MilestoneID: &milestone_id,
			VersionID: &version_id,
			PriorityID: priority_id,
			TypeID: type_id,
			ProjectID: project_id,
			ParentID: parent_id,
			Key: record[11],
			Title: record[12],
			Detail: record[13],
			EstimatedTime: float32(EstimatedTime),
			ActualTime: float32(ActualTime),
			StartTime: StartTime,
			Deadline: Deadline,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		tasks = append(tasks, task)
    }
	DB.Save(&tasks)
}

func RecursivePreload(DB *gorm.DB) *gorm.DB {
	column := "Comments"
	DB.Preload(column + ".User")
	for i := 0; i < 100; i++{
		DB.Preload(column + ".User")
		column += ".Replies"
		DB.Preload(column)
	}
	return DB
}

func CreateOrganizations(MQDB *gorm.DB) {
	DeleteOrganizationsImages()
	f, err := os.Open("data/users.csv")
    if err != nil {
        fmt.Println(err)
    }
	r := csv.NewReader(f)
	var organizations []md.Organization
    for i := 0; ; i++ {
		record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(err)
        }
        if i == 0 {
            continue
        }
		if record[12] == "" {
			continue
		}

        var organization md.Organization
        for i, v := range record {
            switch i {
            case 4:
                organization.Founded = v
            case 8:
                organization.Number = v
            case 11:
                organization.Address = v
            case 12:
                organization.Name = v
            case 13:
                organization.CreditCard = v
            case 14:
                organization.Expiry = v
            }
        }
		organization.GetImage()
		organization.Plan = "standard"
		organization.ID, _ = crypto.MakeRandomStr(25)
		organizations = append(organizations, organization)
		fmt.Println(organization)
		if len(organizations) == 100 {
			break
		}
    }
	MQDB.Create(&organizations)

}

func CreateProjects(MQDB *gorm.DB) {
	// Migrate(MQDB, Project{})
	// CreateProject(MQDB)
	DeleteProjectsImages()
	var organizations []md.Organization
	var organizationKeys []string
	MQDB.Find(&organizations)
	for i, v := range organizations {
		if i == 0 {
			continue
		}
		organizationKeys = append(organizationKeys, v.ID)
	}
	for i := 0; i < 2; i++ {
		f, err := os.Open("data/text.csv")
		if err != nil {
			fmt.Println(err)
		}
		r := csv.NewReader(f)
		var projects []md.Project
		for i := 0; ; i++ {
			rand.Seed(time.Now().UnixNano())
			orgRandomInt := rand.Intn(len(organizationKeys))
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			if i == 0 {
				continue
			}
	
			var project md.Project
			for _, v := range record {
				project.Name = v
			}
			project.GetImage()
			project.OrganizationID = organizationKeys[orgRandomInt]
			projects = append(projects, project)
			fmt.Println(project)
		}
		MQDB.Create(&projects)
	}
}

func CreateUsers(MQDB *gorm.DB) {
	DeleteUsersImages()
	f, err := os.Open("data/users.csv")
    if err != nil {
        fmt.Println(err)
    }
	r := csv.NewReader(f)
	// var users [][]User
	users := make([][]md.User, 2, 2)
    for i := 0; ; i++ {
		record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(err)
        }
        if i == 0 {
            continue
        }

        var user md.User
        for i, v := range record {
            switch i {
            case 0:
                user.Name = v
            case 3:
                user.Age, _ = strconv.Atoi(v)
            case 5:
                user.Sex = v
            case 7:
                user.Email = v
            case 11:
                user.Address = v
            }
			// fmt.Println(user)
        }
		randStr, _ := crypto.MakeRandomStr(20)
		user.Image = randStr + ".png"
		image.CreateImage(user.Name, user.Image)
		user.Password = "52f96e51831c8229413e28b0e58fa3b992f7571e4ff5bf5ccfc1a21f391e4f05"
		if i < 6000 {
			users[0] = append(users[0], user)
		} else {
			users[1] = append(users[1], user)
		}
    }
	for _, user := range users {
		MQDB.Create(&user)
	}
}

func CreateOrganizationUsers(MQDB *gorm.DB) {
	Migrate(MQDB, md.OrganizationAuthority{})
	var organizations []md.Organization
	DB.Preload(clause.Associations).Find(&organizations)
	rand.Seed(time.Now().UnixNano())
	for j, organization := range organizations {
		var organizationAuthorities []md.OrganizationAuthority
		if j == 500 {
			for i := 1; i < 21; i++ {
				num := rand.Intn(10001)
				userID := num
				organizationAuthority := md.OrganizationAuthority{
					OrganizationID: organization.ID,
					UserID: userID,
					AuthorityID: 2,
					Active: true,
				}
				organizationAuthorities = append(organizationAuthorities, organizationAuthority)
			}
		} else {
			for i := 1; i < 21; i++ {
				userID := (j * 20) + i
				if i == 1 {
					if userID == 1 {
						CreateOrganizationAuthority(MQDB)
						continue
					}
					organizationAuthority := md.OrganizationAuthority{
						OrganizationID: organization.ID,
						UserID: userID,
						AuthorityID: 1,
						Active: true,
					}
					organizationAuthorities = append(organizationAuthorities, organizationAuthority)
					continue
				}
				organizationAuthority := md.OrganizationAuthority{
					OrganizationID: organization.ID,
					UserID: userID,
					AuthorityID: 2,
					Active: true,
				}
				organizationAuthorities = append(organizationAuthorities, organizationAuthority)
			}
		}
		MQDB.Create(&organizationAuthorities)
	}
}

func CreateProjectUsers(MQDB *gorm.DB) {
	// Migrate(MQDB, md.ProjectAuthority{})
	// var organizations []md.Organization
	// DB.Preload("Users." + clause.Associations).Preload(clause.Associations).Find(&organizations)
	// fmt.Println(organizations[0])
	// return
	Migrate(MQDB, md.ProjectAuthority{})
	CreateProjectAuthority(MQDB)
	var organizations []md.Organization
	DB.Preload(clause.Associations).Find(&organizations)
	rand.Seed(time.Now().UnixNano())
	for _, organization := range organizations {
		var projectAuthorities []md.ProjectAuthority
		if len(organization.Projects) == 0 {
			continue
		}
		for _, project := range organization.Projects {
			for i, user := range organization.Users {
				boolean := []bool{true,false}
				active := rand.Intn(2)
				authType := 2
				if i == 0 {
					authType = 1
				}
				projectAuthority := md.ProjectAuthority{
					ProjectID: project.ID,
					UserID: user.UserID,
					AuthorityID: authType,
					Active: boolean[active],
				}
				projectAuthorities = append(projectAuthorities, projectAuthority)
			}
		}
		DB.Create(&projectAuthorities)
	}
}

func CreateTasks(MQDB *gorm.DB) {
	// Migrate(MQDB, md.Task{})
	CreateTask(MQDB)
	var projects []md.Project
	DB.Preload("AuthorityUsers", "active = ?", true).Preload(clause.Associations).Find(&projects)
	rand.Seed(time.Now().UnixNano())
	for i, project := range projects {
		if project.ID == 1 {
			continue
		}
		random := 0
		var tasks []md.Task
		count := 1
		for {
			assigneeNum := rand.Intn(len(project.AuthorityUsers))
			assignerNum := rand.Intn(len(project.AuthorityUsers))
			random = rand.Intn(100000)
			statusNum := rand.Intn(9)
			statusID := statusNum + 1
			priorityNum := rand.Intn(3)
			priorityID := priorityNum + 1
			typeNum := rand.Intn(4)
			typeID := typeNum + 1
			fieldID := 0
			if len(project.Fields) != 0 {
				fieldNum := rand.Intn(len(project.Fields))
				fieldID = project.Fields[fieldNum].ID
			}
			milestoneID := 0
			if len(project.Milestones) != 0 {
				milestoneNum := rand.Intn(len(project.Milestones))
				milestoneID = project.Milestones[milestoneNum].ID
			}
			versionID := 0
			if len(project.Milestones) != 0 {
				versionNum := rand.Intn(len(project.Versions))
				versionID = project.Versions[versionNum].ID
			}
			task := md.Task{
				AssigneeID: project.AuthorityUsers[assigneeNum].UserID,
				AssignerID: project.AuthorityUsers[assignerNum].UserID,
				FieldID: &fieldID,
				MilestoneID: &milestoneID,
				VersionID: &versionID,
				StatusID: statusID,
				PriorityID: priorityID,
				TypeID: typeID,
				ProjectID: project.ID,
				Key: project.Name + "_" + strconv.Itoa(count),
				Title: "テスト",
				Detail: "テストです",
				ParentID: 1,
				EstimatedTime: 1,
				ActualTime: 0,
				StartTime: time.Now(),
				Deadline: time.Now(),
			}
			tasks = append(tasks, task)
			count += 1
			if random == 0 {
				break
			}
			if len(tasks) == 500 {
				MQDB.Create(&tasks)
				tasks = []md.Task{}
			}
		}
		MQDB.Create(&tasks)
		fmt.Println(i)
		// SQDB.Create(&tasks)
	}
}

func CreateTasksCSV(MQDB *gorm.DB) {
	// Migrate(MQDB, md.Task{})
	// CreateTask(MQDB)
	var projects []md.Project
	DB.Preload("AuthorityUsers", "active = ?", true).Preload(clause.Associations).Find(&projects)
	rand.Seed(time.Now().UnixNano())
	for _ , project := range projects {
		if project.ID == 1 {
			continue
		}
		random := 0
		var tasks [][]string
		count := 1
		for {
			assigneeNum := rand.Intn(len(project.AuthorityUsers))
			assignerNum := rand.Intn(len(project.AuthorityUsers))
			random = rand.Intn(10000)
			statusNum := rand.Intn(9)
			statusID := statusNum + 1
			priorityNum := rand.Intn(3)
			priorityID := priorityNum + 1
			typeNum := rand.Intn(4)
			typeID := typeNum + 1
			fieldID := 0
			if len(project.Fields) != 0 {
				fieldNum := rand.Intn(len(project.Fields))
				fieldID = project.Fields[fieldNum].ID
			}
			milestoneID := 0
			if len(project.Milestones) != 0 {
				milestoneNum := rand.Intn(len(project.Milestones))
				milestoneID = project.Milestones[milestoneNum].ID
			}
			versionID := 0
			if len(project.Milestones) != 0 {
				versionNum := rand.Intn(len(project.Versions))
				versionID = project.Versions[versionNum].ID
			}
			AssigneeID := strconv.Itoa(project.AuthorityUsers[assigneeNum].UserID)
			AssignerID := strconv.Itoa(project.AuthorityUsers[assignerNum].UserID)
			FieldID := strconv.Itoa(fieldID)
			MilestoneID :=  strconv.Itoa(milestoneID)
			VersionID := strconv.Itoa(versionID)
			StatusID := strconv.Itoa(statusID)
			PriorityID := strconv.Itoa(priorityID)
			TypeID := strconv.Itoa(typeID)
			ProjectID := strconv.Itoa(project.ID)
			task := []string{
				AssigneeID,
				AssignerID,
				FieldID,
				MilestoneID,
				VersionID,
				StatusID,
				PriorityID,
				TypeID,
				ProjectID,
				project.Name + "_" + strconv.Itoa(count),
				"テスト",
				"テストです",
				"1",
				"1",
				"0",
				timeToString(time.Now()),
				timeToString(time.Now()),
			}
			tasks = append(tasks, task)
			count += 1
			if random == 0 {
				break
			}
			if len(tasks) == 20000 {
				break
			}
		}
		os.Create("tasks.csv")
		MQDB.Create(&tasks)
		// SQDB.Create(&tasks)
	}
}

func CreateFields(MQDB *gorm.DB) {
	Migrate(MQDB, md.Field{})
	CreateField(MQDB)
	var fields []md.Field
	var projects []md.Project
	fieldTypes := []string{"フロントエンド", "バックエンド", "デザイナー"}
	DB.Find(&projects)
	for _, project := range projects {
		if project.ID == 1 {
			continue
		}
		for _, fieldType := range fieldTypes {
			field := md.Field{
				Name: fieldType,
				ProjectID: project.ID,
			}
			fields = append(fields, field)
		}
	}
	MQDB.Create(&fields)
}

func CreateMilestones(MQDB *gorm.DB) {
	Migrate(MQDB, md.Milestone{})
	CreateMilestone(MQDB)
	var milestones []md.Milestone
	var projects []md.Project
	milestoneTypes := []string{"フェーズ1", "フェーズ2", "フェーズ3", "フェーズ4", "フェーズ5"}
	DB.Find(&projects)
	for _, project := range projects {
		if project.ID == 1 {
			continue
		}
		for _, milestoneType := range milestoneTypes {
			milestone := md.Milestone{
				Name: milestoneType,
				ProjectID: project.ID,
			}
			milestones = append(milestones, milestone)
		}
	}
	MQDB.Create(&milestones)
}

func CreateComments(MQDB *gorm.DB) {
	Migrate(MQDB, md.Comment{})
	var comments []md.Comment
	var projects []md.Project
	DB.Preload(clause.Associations).Find(&projects)
	rand.Seed(time.Now().UnixNano())
	for _, project := range projects {
		for i, task := range project.Tasks {
			end := 0
			for {
				userNum := rand.Intn(len(project.AuthorityUsers))
				end = rand.Intn(2)
				parentID := 0
				comment := md.Comment{
					Content: strconv.Itoa(i),
					TaskID: task.ID,
					UserID: project.AuthorityUsers[userNum].UserID,
					ParentID: &parentID,
				}
				comments = append(comments, comment)
				if end == 0 {
					break
				}
			}
			MQDB.Create(&comments)
			TreeComments(MQDB, project, comments, 2)
			comments = nil
		}
	}
}

func CreateVersions(MQDB *gorm.DB) {
	Migrate(MQDB, md.Version{})
	var versions []md.Version
	var projects []md.Project
	versionTypes := []string{"release1.0", "release1.1", "release1.2", "release1.3", "release1.4"}
	DB.Find(&projects)
	for _, project := range projects {
		if project.ID == 1 {
			continue
		}
		for _, milestoneType := range versionTypes {
			version := md.Version{
				Name: milestoneType,
				ProjectID: project.ID,
			}
			versions = append(versions, version)
		}
	}
	MQDB.Create(&versions)
}

func TreeComments(MQDB *gorm.DB, project md.Project, comments []md.Comment, endNum int) {
	rand.Seed(time.Now().UnixNano())
	treeEnd := rand.Intn(endNum)
	if treeEnd != 0 {
		return
	}
	var newComments []md.Comment
	for i, comment := range comments {
		end := 0
		for {
			end = rand.Intn(1)
			userNum := rand.Intn(len(project.AuthorityUsers))
			comment := md.Comment{
				Content: strconv.Itoa(i),
				TaskID: comment.TaskID,
				UserID: project.AuthorityUsers[userNum].UserID,
				ParentID: &comment.ID,
			}
			newComments = append(newComments, comment)
			if end == 0 {
				break
			}
		}
		MQDB.Create(&newComments)
		TreeComments(MQDB, project, comments, endNum*2)
		newComments = nil
	}
}
var layout = "2006-01-02 15:04:05"
func timeToString(t time.Time) string {
    str := t.Format(layout)
    return str
}

func createUserImage() {
	var users []models.User
	config.DB.Find(&users)
	for  _, u := range users {
		buf, err := image.CreateImage(u.Name, u.Image); if err != nil {
			fmt.Println(err)
		}

		// 環境による保存場所の変更
		path := "users/" + u.Image
		err = storage.StoreImageToGCS(buf.Bytes(), path, BUCKET_NAME); if err != nil {
			fmt.Println(err)
		}
		fmt.Println(err)
	}
}

func createOrganizationImage() {

	var org []models.Organization
	config.DB.Find(&org)
	var wg sync.WaitGroup
	for  _, o := range org {
		o := o
		wg.Add(1)
		go func() {
			url := "https://loremflickr.com/320/240?random=1"
			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			// // 環境による保存場所の変更
			path := "organizations/" + o.Image
			buf := new(bytes.Buffer)
			io.Copy(buf, response.Body)
			err = storage.StoreImageToGCS(buf.Bytes(), path, BUCKET_NAME); if err != nil {
				fmt.Println(err)
			}
		}()

	}
	wg.Wait()
}

func createProjectImage() {

	var pr []models.Project
	config.DB.Find(&pr)
	var wg sync.WaitGroup
	defer wg.Done()
	for  _, p := range pr {
		p := p
		wg.Add(1)
		go func() {
			url := "https://loremflickr.com/320/240?random=1"
			response, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			// // 環境による保存場所の変更
			path := "projects/" + p.Image
			buf := new(bytes.Buffer)
			io.Copy(buf, response.Body)
			err = storage.StoreImageToGCS(buf.Bytes(), path, BUCKET_NAME); if err != nil {
				fmt.Println(err)
			}
		}()

	}
	wg.Wait()
}