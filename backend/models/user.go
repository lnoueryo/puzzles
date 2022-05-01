package models

import (
	"backend/config"
	"backend/modules/crypto"
	"backend/modules/session"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID						int						`gorm:"AUTO_INCREMENT"json:"id"`
	Name					string					`json:"name"`
	Age						int						`json:"age"`
	Sex						string					`json:"sex"`
	Email					string					`json:"email"`
	Address					string					`json:"address"`
	Password				string					`json:"password"`
	Image					string					`json:"image"`
	ImageData				string					`gorm:"<-:false;-:migration;"json:"image_data"`
	Description				string					`json:"description"`
	Organization			string					`gorm:"-:all;"json:"organization"`
	Projects				[]ProjectAuthority		`json:"projects"`
	Organizations			[]OrganizationAuthority	`gorm:"foreignkey:UserID;"json:"organizations"`
	Tasks					[]Task					`gorm:"foreignKey:AssigneeID;references:ID"json:"tasks"`
	CreatedAt				time.Time				`gorm:"->:false;<-:create;autoCreateTime;"json:"-"`
	UpdatedAt				time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}


func NewUser(r *http.Request) (User, error) {
	user, _ := GetUserJson(r)
	randStr, _ := crypto.MakeRandomStr(20)
	filename := randStr + ".png"
	password := crypto.Encrypt(user.Password)
	user = User{Name: user.Name, Email: user.Email, Image: filename, Password: password}
	return user, nil
}

func UserAll() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users, result.Error
	}
	return users, nil
}

func UserLatest(limit int) ([]User, error) {
	var users []User
	result := DB.Order("created_at desc").Limit(limit).Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users, result.Error
	}
	return users, nil
}

func (u *User) Create() error {
	result := DB.Create(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (u *User) Update() error {
	if u.Password != "" {
		u.Password = crypto.Encrypt(u.Password)
		result := DB.Omit("Organizations").Save(u)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return nil
	}
	result := DB.Omit("Organizations", "Password").Save(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (u *User)GetMainUser(userID int, orgID string) error {
	result := DB.Preload("Organizations.Organization.Projects.AuthorityUsers." + clause.Associations).
	Preload("Organizations.Organization.Projects.AuthorityUsers", "active = ?", true).
	Preload("Organizations.Organization.Projects." + clause.Associations).
	Preload("Organizations.Organization.Projects").
	Preload("Organizations.Organization.Users.User").
	Preload("Organizations.Organization.Users").
	Preload("Organizations." + clause.Associations).
	Preload("Organizations", "organization_id = ?", orgID).
	Preload("Organizations").
	First(&u, "id = ?", userID); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (u *User)CheckUser() error {

	result := DB.FirstOrCreate(&u, User{Email: u.Email}); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func GetProjectAuthorities(uid int) error {
	var pas []ProjectAuthority
	result := DB.Preload("Project.Milestones").Preload("Project.Fields").Preload("ProjectUsers.Type").Preload("ProjectUsers.User").Preload("ProjectUsers", "active = ?", true).Preload(clause.Associations).Find(&pas, "user_id = ? and active = true", uid); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}


func SearchUserLike(r *http.Request, column string) ([]User, int64, error) {
	var users []User
	query := r.URL.Query()
	page, _ := strconv.Atoi(query["page"][0])
	if query[column][0] == "" {
		users, count, _ := ChunkUser(page)
		return users, count, nil
	}
	split := 10
	offset := (page - 1) * split
	textSlice := strings.Split(query[column][0], " ")
	var tx *gorm.DB
	var count int64
	for _, text := range textSlice {
		likeText := "%" + text + "%"
		tx = DB.Model(&User{}).Limit(split).Offset(offset).Where(fmt.Sprintf("%s LIKE ? ", column), likeText)
	}
	result := tx.Find(&users).Limit(-1).Offset(-1).Count(&count)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users, 0, result.Error
	}
	return users, count, nil
}

func ChunkUser(page int) ([]User, int64, error) {
	var users []User
	var count int64
	split := 10
	offset := (page - 1) * split
	result := DB.Limit(split).Offset(offset).Find(&users).Limit(-1).Offset(-1).Count(&count)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return users, 0, result.Error
	}
	return users, count, nil
}

func (u *User) Validate(r *http.Request) error {
	err := u.CheckBlank(r)
	if err != nil {
		return err
	}

	err = u.ComparePassword(r)
	if err != nil {
		return err
	}

	err = CheckEmailFormat(u.Email)
	if err != nil {
		return err
	}

	err = u.CheckLength(r)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateValidate(r *http.Request) error {
	err := u.CheckBlankForUpdate(r)
	if err != nil {
		return err
	}

	err = u.MatchPassword(r)
	if err != nil {
		return err
	}

	err = u.SearchSameEmail(r)
	if err != nil {
		return err
	}

	err = u.CheckImage(r)
	if err != nil {
		return err
	}

	err = u.CheckLengthForUpdate(r)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) CheckBlank(r *http.Request) error {
	name := r.FormValue("name")
	if name == "" {
		message := "name is blank"
		err := errors.New(message)
		return err
	}

	email := r.FormValue("email")
	if email == "" {
		message := "email address is blank"
		err := errors.New(message)
		return err
	}

	password := r.FormValue("password")
	if password == "" {
		message := "password is blank"
		err := errors.New(message)
		return err
	}

	confirmation := r.FormValue("confirmation")
	if confirmation == "" {
		message := "confirmation is blank"
		err := errors.New(message)
		return err
	}
	return nil
}
	// login, _ := GetJsonForm(r)
	// var login Login
	// err := json.NewDecoder(r.Body).Decode(&login)
	// fmt.Print(err)
	// fmt.Print("err")
	// fmt.Print(login)
func (u *User) CheckBlankForUpdate(r *http.Request) error {
	name := r.FormValue("name")
	if name == "" {
		message := "name is blank"
		err := errors.New(message)
		return err
	}

	email := r.FormValue("email")
	if email == "" {
		message := "email address is blank"
		err := errors.New(message)
		return err
	}

	password := r.FormValue("current-password")
	if password == "" {
		message := "password is blank"
		err := errors.New(message)
		return err
	}

	return nil
}

func (u *User) ComparePassword(r *http.Request) error {
	password := r.FormValue("password")
	confirmation := r.FormValue("confirmation")
	if password != confirmation {
		message := "password and confirmation must be the same"
		err := errors.New(message)
		return err
	}
	return nil
}

func CheckEmailFormat(email string) error {
	regex := `^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$`

	isEnabled := regexp.MustCompile(regex).Match([]byte(email))
	if !isEnabled {
		message := "invalid email address pattern"
		err := errors.New(message)
		return err
	}
	return nil
}

func (u *User) CheckLength(r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(name) > 50 {
		message := "name must be less than 50 characters"
		err := errors.New(message)
		return err
	}
	index := strings.Index(email, "@")
	localPart := email[:index]
	if len(localPart) > 64 {
		message := "invalid email address pattern"
		err := errors.New(message)
		return err
	}

	if 8 > len(password) {
		message := "password must be more than 8 characters"
		err := errors.New(message)
		return err
	}
	return nil
}

func (u *User) CheckLengthForUpdate(r *http.Request) error {
	name := r.FormValue("name")
	email := r.FormValue("email")

	if len(name) > 50 {
		message := "name must be less than 50 characters"
		err := errors.New(message)
		return err
	}
	index := strings.Index(email, "@")
	localPart := email[:index]
	if len(localPart) > 64 {
		message := "invalid email address pattern"
		err := errors.New(message)
		return err
	}

	return nil
}

func (u *User) MatchPassword(r *http.Request) error {
	currentPassword := crypto.Encrypt(r.FormValue("current-password"))
	if u.Password != currentPassword {
		message := "current password is wrong"
		err := errors.New(message)
		return err
	}
	return nil
}

func (u *User) CheckImage(r *http.Request) error {
	currentPassword := crypto.Encrypt(r.FormValue("current-password"))
	if u.Password != currentPassword {
		message := "current password is wrong"
		err := errors.New(message)
		return err
	}
	return nil
}

func (u *User) SearchSameEmail(r *http.Request) error {
	var user User
	result := DB.Where("email = ?", r.Form.Get("email")).First(&user)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if u.Email == user.Email {
			return nil
		}
		err := errors.New("email address is already registered")
		return err
	}
	return nil
}

func GetJsonForm(r *http.Request) (Login, error) {
	var login Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		fmt.Println(err)
		return login, err
	}
	return login, err
}

func (u *User)CreateSession(w http.ResponseWriter) (session.Session, error) {
	s := session.Session{
		UserID:			u.ID,
		Name:			u.Name,
		Age:			u.Age,
		Sex:			u.Sex,
		Email:			u.Email,
		Address:		u.Address,
		Image:			u.Image,
		Description:	u.Description,
		Organization:	u.Organization,
		CreatedAt:		time.Now(),
	}
	s.CreateSession(config.App.Project)
	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    s.ID,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: 4,
	}
	http.SetCookie(w, &cookie)
	return s, nil
}
// func (u *User)CreateSession(w http.ResponseWriter) (Session, error) {
// 	// sessionID 作成
// 	var s Session
// 	var oa OrganizationAuthority
// 	sessionId := string(u.ID) + timeToString(u.CreatedAt) + timeToString(time.Now())
// 	hashedByteSessionId := sha256.Sum256([]byte(sessionId))
// 	hashedSessionId := fmt.Sprintf("%x", (hashedByteSessionId))
// 	result := DB.Preload("Type").Preload(clause.Associations).Find(&oa, "user_id = ? and organization_id = ?", u.ID, u.Organization)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		return s, result.Error
// 	}
// 	fmt.Println(oa)
// 	u.Authority = oa.Type.Name
// 	// session用ファイル作成
// 	s = Session{
// 		ID:				hashedSessionId,
// 		UserID:			u.ID,
// 		Name:			u.Name,
// 		Age:			u.Age,
// 		Sex:			u.Sex,
// 		Email:			u.Email,
// 		Address:		u.Address,
// 		Image:			u.Image,
// 		Description:	u.Description,
// 		Organization:	u.Organization,
// 		Authority:		u.Authority,
// 		CreatedAt:		time.Now(),
// 	}
// 	// sessionフォルダの有無判定
// 	_, err := os.Stat("session")
// 	if err != nil {
// 		os.Mkdir("session", 0777)
// 	}
// 	filepath := fmt.Sprintf("./session/%v.txt", hashedSessionId)
// 	f, err := os.Create(filepath)
// 	if err != nil {
// 		return s, err
// 	}
// 	defer f.Close()
// 	enc := gob.NewEncoder(f)

// 	if err := enc.Encode(s); err != nil {
// 		return s, err
// 	}
// 	// Path is needed for all path
// 	cookie := http.Cookie{
// 		Name:     "_cookie",
// 		Value:    hashedSessionId,
// 		HttpOnly: true,
// 		Secure:   true,
// 		Path:     "/",
// 		SameSite: 4,
// 	}
// 	http.SetCookie(w, &cookie)
// 	return s, nil
// }

func GetUserJson(r *http.Request) (User, error) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
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