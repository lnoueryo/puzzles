package commands

import (
	"backend/modules/crypto"
	"backend/modules/image"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	// "gorm.io/gorm/clause"
)

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
	SQDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	dbconf := `root:@/puzzle?parseTime=true&charset=utf8&loc=Local`
	MQDB, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(SQDB)
		panic("failed to connect database")
	}
	// SeriesOfCreation(SQDB, MQDB)
	start := time.Now()
	AutoMigrate(SQDB, MQDB)
	var u User
	MQDB.Preload(clause.Associations).Preload("Organizations", "id = ?", "prygen4fDISDVgSYDjxZ5uICD").First(&u, "email = ?", "popo62520908@gmail.com")
	// createOrganizations(SQDB, MQDB)
	// createUsers(SQDB, MQDB)
	// createProjects(SQDB, MQDB)
	// createProjectAuthority(SQDB, MQDB)
	// CreateFields(SQDB, MQDB)
	// CreateMilestones(SQDB, MQDB)
	// createFieldUsers(SQDB, MQDB)
	createTasks(SQDB, MQDB)
	// createProjectUsers(SQDB, MQDB)
	// ReadOrganization()
	fmt.Println((time.Now()).Sub(start))
	// readTask(MQDB)
}

func (pa *ProjectAuthority) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println(pa)
	return
  }
func createOrganizations(SQDB *gorm.DB, MQDB *gorm.DB) {
	DeleteOrganizationsImages()
	f, err := os.Open("data/organizations.csv")
    if err != nil {
        fmt.Println(err)
    }
	r := csv.NewReader(f)
	var organizations []Organization
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

        var organization Organization
        for i, v := range record {
            switch i {
            case 0:
                organization.Founded = v
            case 2:
                organization.Number = v
            case 3:
                organization.Address = v
            case 4:
                organization.Name = v
            case 5:
                organization.CreditCard = v
            case 6:
                organization.Expiry = v
            }
			// fmt.Println(user)
        }
		organization.GetImage()
		organization.Plan = "standard"
		organization.ID, _ = crypto.MakeRandomStr(25)
		organizations = append(organizations, organization)
		fmt.Println(organization)
    }
	SQDB.Create(&organizations)
	MQDB.Create(&organizations)

}

func createProjects(SQDB *gorm.DB, MQDB *gorm.DB) {
	// Migrate(SQDB, MQDB, Project{})
	// CreateProject(SQDB, MQDB)
	DeleteProjectsImages()
	var organizations []Organization
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
		var projects []Project
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
	
			var project Project
			for _, v := range record {
				project.Name = v
			}
			project.GetImage()
			project.OrganizationID = organizationKeys[orgRandomInt]
			projects = append(projects, project)
			fmt.Println(project)
		}
		SQDB.Create(&projects)
		MQDB.Create(&projects)
	}
}

func createUsers(SQDB *gorm.DB, MQDB *gorm.DB) {
	DeleteUsersImages()
	f, err := os.Open("data/users.csv")
    if err != nil {
        fmt.Println(err)
    }
	r := csv.NewReader(f)
	// var users [][]User
	users := make([][]User, 2, 2)
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

        var user User
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
		SQDB.Create(&user)
		MQDB.Create(&user)
	}
}

func createOrganizationUsers(SQDB *gorm.DB, MQDB *gorm.DB) {
	Migrate(SQDB, MQDB, OrganizationAuthority{})
	var organizations []Organization
	DB.Preload(clause.Associations).Find(&organizations)
	for j, organization := range organizations {
		var organizationAuthorities []OrganizationAuthority
		if j == 500 {
			for i := 1; i < 21; i++ {
				rand.Seed(time.Now().UnixNano())
				num := rand.Intn(10001)
				userID := num
				organizationAuthority := OrganizationAuthority{
					OrganizationID: organization.ID,
					UserID: userID,
					AuthID: 2,
				}
				organizationAuthorities = append(organizationAuthorities, organizationAuthority)
			}
		} else {
			for i := 1; i < 21; i++ {
				userID := (j * 20) + i
				if i == 1 {
					if userID == 1 {
						CreateOrganizationAuthority(SQDB, MQDB)
						continue
					}
					organizationAuthority := OrganizationAuthority{
						OrganizationID: organization.ID,
						UserID: userID,
						AuthID: 1,
					}
					organizationAuthorities = append(organizationAuthorities, organizationAuthority)
					continue
				}
				organizationAuthority := OrganizationAuthority{
					OrganizationID: organization.ID,
					UserID: userID,
					AuthID: 2,
				}
				organizationAuthorities = append(organizationAuthorities, organizationAuthority)
			}
		}
		DB.Create(&organizationAuthorities)
	}
}

func createProjectUsers(SQDB *gorm.DB, MQDB *gorm.DB) {
	Migrate(SQDB, MQDB, ProjectAuthority{})
	CreateProjectAuthority(SQDB, MQDB)
	var organizations []Organization
	DB.Preload(clause.Associations).Find(&organizations)
	for _, organization := range organizations {
		var projectAuthorities []ProjectAuthority
		if len(organization.Projects) == 0 {
			continue
		}
		for _, project := range organization.Projects {
			// limit := 0
			// userIDSlice := []int{}
			// for _, user := range organization.Users {
			// 	userIDSlice = append(userIDSlice, user.ID)
			// }
			// rand.Seed(time.Now().UnixNano())
			// rand.Shuffle(len(userIDSlice), func(i, j int) { userIDSlice[i], userIDSlice[j] = userIDSlice[j], userIDSlice[i] })
			// for i, id := range userIDSlice {
			// 	rand.Seed(time.Now().UnixNano())
			// 	authType := 2
			// 	if i == 0 {
			// 		authType = 1
			// 	}
			// 	projectAuthority := ProjectAuthority{
			// 		ProjectID: project.ID,
			// 		UserID: id,
			// 		AuthID: authType,
			// 	}
			// 	projectAuthorities = append(projectAuthorities, projectAuthority)
			// }
			for i, user := range organization.Users {
				boolean := []bool{true,false}
				rand.Seed(time.Now().UnixNano())
				active := rand.Intn(2)
				authType := 2
				if i == 0 {
					authType = 1
				}
				projectAuthority := ProjectAuthority{
					ProjectID: project.ID,
					UserID: user.ID,
					AuthID: authType,
					Active: boolean[active],
				}
				projectAuthorities = append(projectAuthorities, projectAuthority)
			}
		}
		fmt.Println(projectAuthorities)
		DB.Create(&projectAuthorities)
	}
}

func createTasks(SQDB *gorm.DB, MQDB *gorm.DB) {
	Migrate(SQDB, MQDB, Task{})
	CreateTask(SQDB, MQDB)
	var projects []Project
	var pas []ProjectAuthority
	DB.Preload(clause.Associations).Find(&projects)
	DB.Preload(clause.Associations).Find(&pas, "active = true")
	for _ ,project := range projects {
		if project.ID == 1 {
			continue
		}
		var joinPaSlice []ProjectAuthority
		for _ ,pa := range pas {
			if pa.ProjectID ==  project.ID{
				joinPaSlice = append(joinPaSlice, pa)
			}
		}
		project.AuthorityUsers = joinPaSlice
		random := 0
		var tasks []Task
		count := 1
		for {
			rand.Seed(time.Now().UnixNano())
			assigneeNum := rand.Intn(len(project.AuthorityUsers))
			assignerNum := rand.Intn(len(project.AuthorityUsers))
			random = rand.Intn(600)
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
			task := Task{
				AssigneeID: project.AuthorityUsers[assigneeNum].User.ID,
				AssignerID: project.AuthorityUsers[assignerNum].User.ID,
				FieldID: &fieldID,
				MilestoneID: &milestoneID,
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
		}
		fmt.Println(tasks)
		MQDB.Create(&tasks)
		SQDB.Create(&tasks)
	}
	// for _, project := range projects {
	// 	if project.ID == 1 {
	// 		continue
	// 	}
	// 	random := 0
	// 	var tasks []Task
	// 	count := 1
	// 	for {
	// 		rand.Seed(time.Now().UnixNano())
	// 		assigneeNum := rand.Intn(len(project.AuthorityUsers))
	// 		assignerNum := rand.Intn(len(project.AuthorityUsers))
	// 		random = rand.Intn(500)
	// 		milestoneNum := rand.Intn(len(project.Milestones))
	// 		statusNum := rand.Intn(9)
	// 		statusID := statusNum + 1
	// 		fieldNum := rand.Intn(len(project.Fields))
	// 		priorityNum := rand.Intn(3)
	// 		priorityID := priorityNum + 1
	// 		typeNum := rand.Intn(5)
	// 		typeID := typeNum + 1
	// 		fieldID := project.Fields[fieldNum].ID
	// 		milestoneID := project.Milestones[milestoneNum].ID
	// 		task := Task{
	// 			AssigneeID: project.AuthorityUsers[assigneeNum].User.ID,
	// 			AssignerID: project.AuthorityUsers[assignerNum].User.ID,
	// 			FieldID: &fieldID,
	// 			MilestoneID: &milestoneID,
	// 			StatusID: statusID,
	// 			PriorityID: priorityID,
	// 			TypeID: typeID,
	// 			ProjectID: project.ID,
	// 			Key: project.Name + "_" + strconv.Itoa(count),
	// 			Title: "テスト",
	// 			Detail: "テストです",
	// 			ParentID: 1,
	// 			EstimatedTime: 1,
	// 			ActualTime: 0,
	// 			StartTime: time.Now(),
	// 			Deadline: time.Now(),
	// 		}
	// 		tasks = append(tasks, task)
	// 		count += 1
	// 		if random == 0 {
	// 			break
	// 		}
	// 	}
	// 	MQDB.Create(&tasks)
	// 	SQDB.Create(&tasks)
	// }
}

// func createTasks(SQDB *gorm.DB, MQDB *gorm.DB) {
// 	Migrate(SQDB, MQDB, Task{})
// 	CreateTask(SQDB, MQDB)
// 	var organizations []Organization
// 	var pa []ProjectAuthority
// 	DB.Preload("Projects.AuthorityUsers.User").Preload("Projects", func(db *gorm.DB) *gorm.DB {
// 		return DB.Preload(clause.Associations)
// 	  }).Preload(clause.Associations).Find(&organizations)
// 	for _, organization := range organizations {
// 		if len(organization.Projects) == 0 {
// 			continue
// 		}
// 		for _, project := range organization.Projects {
// 			if project.ID == 1 {
// 				continue
// 			}
// 			random := 0
// 			var tasks []Task
// 			count := 1
// 			for {
// 				rand.Seed(time.Now().UnixNano())
// 				assigneeNum := rand.Intn(len(project.AuthorityUsers))
// 				assignerNum := rand.Intn(len(project.AuthorityUsers))
// 				random = rand.Intn(500)
// 				milestoneNum := rand.Intn(len(project.Milestones))
// 				statusNum := rand.Intn(9)
// 				statusID := statusNum + 1
// 				fieldNum := rand.Intn(len(project.Fields))
// 				priorityNum := rand.Intn(3)
// 				priorityID := priorityNum + 1
// 				typeNum := rand.Intn(5)
// 				typeID := typeNum + 1
// 				fieldID := project.Fields[fieldNum].ID
// 				milestoneID := project.Milestones[milestoneNum].ID
// 				task := Task{
// 					AssigneeID: project.AuthorityUsers[assigneeNum].User.ID,
// 					AssignerID: project.AuthorityUsers[assignerNum].User.ID,
// 					FieldID: &fieldID,
// 					MilestoneID: &milestoneID,
// 					StatusID: statusID,
// 					PriorityID: priorityID,
// 					TypeID: typeID,
// 					ProjectID: project.ID,
// 					Key: project.Name + "_" + strconv.Itoa(count),
// 					Title: "テスト",
// 					Detail: "テストです",
// 					ParentID: 1,
// 					EstimatedTime: 1,
// 					ActualTime: 0,
// 					StartTime: time.Now(),
// 					Deadline: time.Now(),
// 				}
// 				tasks = append(tasks, task)
// 				count += 1
// 				if random == 0 {
// 					break
// 				}
// 			}
// 			MQDB.Create(&tasks)
// 			SQDB.Create(&tasks)
// 		}
// 	}
// }

func CreateFields(SQDB *gorm.DB, MQDB *gorm.DB) {
	Migrate(SQDB, MQDB, Field{})
	CreateField(SQDB, MQDB)
	var organizations []Organization
	fieldTypes := []string{"フロントエンド", "バックエンド", "デザイナー"}
	DB.Preload(clause.Associations).Find(&organizations)
	for _, organization := range organizations {
		var fields []Field
		if organization.ID == "prygen4fDISDVgSYDjxZ5uICD" {
			continue
		}
		for _, project := range organization.Projects {
			for _, fieldType := range fieldTypes {
				field := Field{
					Name: fieldType,
					ProjectID: project.ID,
				}
				fields = append(fields, field)
			}
		}
		SQDB.Create(&fields)
		MQDB.Create(&fields)
	}
}

func CreateMilestones(SQDB *gorm.DB, MQDB *gorm.DB) {
	Migrate(SQDB, MQDB, Milestone{})
	CreateMilestone(SQDB, MQDB)
	var milestones []Milestone
	var projects []Project
	milestoneTypes := []string{"フェーズ1", "フェーズ2", "フェーズ3", "フェーズ4", "フェーズ5"}
	DB.Find(&projects)
	for _, project := range projects {
		if project.ID == 1 {
			continue
		}
		for _, milestoneType := range milestoneTypes {
			milestone := Milestone{
				Name: milestoneType,
				ProjectID: project.ID,
			}
			milestones = append(milestones, milestone)
		}
	}
	SQDB.Create(&milestones)
	MQDB.Create(&milestones)
}
