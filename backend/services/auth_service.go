package services

import (
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/mail"
	"backend/modules/session"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type LoginUser struct {
	Organization	string    `json:"organization"`
	Email			string    `json:"email"`
	Password		string    `json:"password"`
}

type MailRequest struct {
	// リクエストで送られてきたメールアドレス
	Email			string 	`json:"email"`
	// 加入する組織のID
	OrganizationID	string	`json:"organization_id"`
	// 加入するユーザーの権限
	AuthorityID		int		`json:"authority_id"`
}

const (
	// メールのメッセージ
	mainMessage = "下記のURLより参加できます。\n"
)

func Login(w http.ResponseWriter, r *http.Request) (*http.Request, error) {
	// バリデーション
	u, err := tryToLogin(w, r);if err != nil {
		return r, err
	}
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
	ses := session.Session{}
	r = r.WithContext(context.WithValue(
		r.Context(),
		ses,
		s,
	))
	s.CreateSession(projectID)
	SetCookie(w, s.ID)

	return r, nil
}

func Logout(w *http.ResponseWriter, r *http.Request) error {
	fmt.Print("Hello")
	var ses session.Session
	s := r.Context().Value(ses).(session.Session)
	fmt.Print("2")

	session.DeleteSession(s.ID, projectID)

	cookie, err := r.Cookie(cookieKey);if err != nil {
		return err
	}

	// クッキーを無効にする
	cookie.MaxAge = -1
	http.SetCookie(*w, cookie)
	return nil
}

func InviteUser(r *http.Request) error {

	err := DB.Transaction(func(tx *gorm.DB) error {

		cookie, err := r.Cookie(cookieKey);if err != nil {
			return err
		}

		s, err := session.CheckSession(cookie.Value, projectID)
		if err != nil {
			return err
		}

		// メール送信用の情報をjsonから取得
		requestMail, _ := getMailJson(r)

		// ユーザーが既に登録されているか確認。いない場合はユーザー作成
		var u models.User
		err = tx.FirstOrCreate(&u, models.User{Email: requestMail.Email}).Error; if err != nil {
			return err
		}

		// verification用のパスワード作成
		verification, _ := crypto.MakeRandomStr(30)
		var oa = models.OrganizationAuthority{
			UserID: u.ID,
			OrganizationID: requestMail.OrganizationID,
			AuthorityID: requestMail.AuthorityID,
			Active: false,
			Verification: verification,
		}

		newoa, err := models.FindByIDs(tx, oa.OrganizationID, oa.UserID);if err != nil {
			err = oa.Create(DB);if err != nil {
				return err
			}
		} else {
			// 招待されており、既に組織に登録されている場合
			if newoa.Active {
				err = errors.New("既に組織に登録されているようです。")
				return err
			}

			// 招待されているが、ユーザーがまだ組織に登録されている場合
			newoa.Verification = oa.Verification
			err = newoa.Update(tx);if err != nil {
				err = errors.New("エラーが発生しました。時間をおいて再度お試しください。")
				return err
			}
		}

		// 招待メールの送信
		m := email
		m.Sub = s.Name + "さんから招待が送られました"
		m.To = requestMail.Email
		url := mainMessage + origin + "/register/verification?code=" + verification
		m.Message = url
		err = mail.SendEmail(m); if err !=nil {
			err = errors.New("存在しないメールアドレスです。")
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ValidateVerification(r *http.Request, oa *models.OrganizationAuthority) (string, error) {
    q := r.URL.Query()
	code, ok := q["code"];if !ok {
		url := allowOrigin + "/expiry"
		err := errors.New("not found")
		return url, err
	}

	// URLのパラメーターがない場合
	verification := code[0]
	if verification == "" {
		url := allowOrigin + "/expiry"
		err := errors.New("not from browser")
		return url, err
	}

	err := oa.Find(DB, verification);if err != nil {
		url := allowOrigin + "/expiry"
		return url, err
	}

	// 有効期限の有無確認
	before24hours := -time.Hour * 24
	oneDay := time.Now().Add(before24hours)
	if !oneDay.Before(oa.UpdatedAt) {
		url := allowOrigin + "/expiry"
		err := errors.New("verification expired")
		return url, err
	}

	// activeを有効に変更
	err = oa.ChangeActive(DB); if err != nil {
		url := allowOrigin + "/bad-connection"
		return url, err
	}

	url := allowOrigin + "/success"
	m := email
	m.Sub = "承認されました"
	m.To = oa.User.Email
	m.Message = "組織ID: " + oa.OrganizationID + "\nメールアドレス: " + oa.User.Email

	// ユーザー登録が初めての場合
	if oa.User.Name == "" {
		var passwordLength uint32 = 20
		password, _ := crypto.MakeRandomStr(passwordLength)
		oa.User.ChangePassword = password
		m.Message += "\n初回パスワード: " + password
	}

	// ユーザー登録成功のメール送信
	err = mail.SendEmail(m); if err !=nil {
		url := allowOrigin + "/bad-connection"
		return url, err
	}
	return url, nil
}

// ログインのバリデーション
func tryToLogin(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	var u = models.NewUser()
	l, err := GetLoginJson(r);if err != nil {
		return u, err
	}

	// ログインフォーム空白確認
	err = l.CheckLoginFormBlank();if err != nil {
		return u, err
	}
	// メールアドレスのフォームチェック
	err = CheckEmailFormat(l.Email);if err != nil {
		return u, err
	}

	// 全ての項目を踏まえログイン情報が正しい確認
	err = u.FindLoginUser(DB, l.Email, l.Password, l.Organization);if err != nil {
		return u, err
	}

	if len(u.Organizations) == 0 {
		message := "organization is wrong"
		err := errors.New(message)
		return u, err
	}

	return u, nil
}

func (l *LoginUser) CheckLoginFormBlank() error {

	if l.Organization == "" {
		message := "organization is blank"
		err := errors.New(message)
		return err
	}

	if l.Email == "" {
		message := "email address is blank"
		err := errors.New(message)
		return err
	}

	if l.Password == "" {
		message := "password is blank"
		err := errors.New(message)
		return err
	}
	return nil
}

func GetLoginJson(r *http.Request) (LoginUser, error) {
	var login LoginUser
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return login, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &login)
		if err != nil {
			return login, err
		}
		return login, nil
	}

	message := "request body is empty"
	err = errors.New(message)
	return login, err
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

// JSONを構造体MailRequestに変換
func getMailJson(r *http.Request) (MailRequest, error) {
	var mailRequest MailRequest
	json.NewDecoder(r.Body).Decode(&mailRequest)
	return mailRequest, nil
}

func SetCookie(w http.ResponseWriter, id string) {
	cookie := http.Cookie{
		Name:     cookieKey,
		Value:    id,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: 4,
	}
	http.SetCookie(w, &cookie)
}