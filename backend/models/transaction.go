package models

import (
	"backend/modules/crypto"
	"backend/modules/mail"
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)


type MailRequest struct {
	Email string `json:"email"`
	AuthorityID int `json:"authority_id"`
}
func InviteUser(r *http.Request) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		s, _ := CheckSession(r)
		requestMail, _ := GetMailJson(r)
		var u User
		// 送信する
		err := tx.FirstOrCreate(&u, User{Email: requestMail.Email}).Error; if err != nil {
			return err
		}
		verification, _ := crypto.MakeRandomStr(30)
		var oa = OrganizationAuthority{
			UserID: u.ID,
			OrganizationID: s.Organization,
			AuthorityID: requestMail.AuthorityID,
			Active: false,
			Verification: verification,
		}
		condition := OrganizationAuthority{
			OrganizationID: oa.OrganizationID,
			UserID: oa.UserID,
		}
		var newOA OrganizationAuthority
		result := tx.Where(condition).First(&newOA); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result = tx.Create(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			return nil
		}
		if newOA.Active {
			err = errors.New("既に組織に登録されているようです。")
			return err
		}
		newOA.Verification = oa.Verification
		result = tx.Save(&newOA); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		m := email
		m.Sub = s.Name + "さんから招待が送られました"
		m.To = requestMail.Email
		url := mainMessage + origin + "/register/verification?code=" + verification
		m.Message = url
		err = mail.SendEmail(m); if err !=nil {
			err = errors.New("存在しないメールアドレスです。")
			return err
		}
		// nilが返却されるとトランザクション内の全処理がコミットされる
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetMailJson(r *http.Request) (MailRequest, error) {
	var mailRequest MailRequest
	json.NewDecoder(r.Body).Decode(&mailRequest)
	return mailRequest, nil
}


const (
	mainMessage = "下記のURLより参加できます。\n"
)