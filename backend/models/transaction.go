package models

import (
	"backend/modules/crypto"
	"backend/modules/mail"
	"backend/modules/session"
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)


type MailRequest struct {
	Email			string 	`json:"email"`
	OrganizationID	string	`json:"organization_id"`
	AuthorityID		int		`json:"authority_id"`
}
func InviteUser(r *http.Request) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		cookie, err := r.Cookie("_cookie");if err != nil {
			return err
		}
		s, err := session.CheckSession(cookie.Value, project)
		if err != nil {
			return err
		}
		requestMail, _ := GetMailJson(r)
		var u User
		// 送信する
		err = tx.FirstOrCreate(&u, User{Email: requestMail.Email}).Error; if err != nil {
			return err
		}
		verification, _ := crypto.MakeRandomStr(30)
		var oa = OrganizationAuthority{
			UserID: u.ID,
			OrganizationID: requestMail.OrganizationID,
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
		} else {
			if newOA.Active {
				err = errors.New("既に組織に登録されているようです。")
				return err
			}
			newOA.Verification = oa.Verification
			result = tx.Save(&newOA); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
		}
		m := email
		m.Sub = s.Name + "さんから招待が送られました"
		m.To = requestMail.Email
		url := mainMessage + origin + "/register/verification?code=" + verification
		m.Message = url
		err = mail.SendEmail(m); if err !=nil {
			errorlog.Print(err)
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

func (pur *ProjectUpdateRequest)BulkUpdateProject() error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if pur.FieldDelete {
			var f Field
			if err := tx.Delete(&f, "project_id = ?", pur.Project.ID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if pur.MilestoneDelete {
			var m Milestone
			if err := tx.Delete(&m, "project_id = ?", pur.Project.ID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if pur.VersionDelete {
			var v Version
			if err := tx.Delete(&v, "project_id = ?", pur.Project.ID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if pur.Project.ImageData != "" {
			deleteImageName := pur.Project.Image
			fileName, err := StoreImage("projects", pur.Project.ImageData); if err != nil {
				return errors.New("couldn't save the image")
			}
			pur.Project.Image = fileName
			result := DB.Debug().Omit("Organization", "Tasks", "AuthorityUsers.User", "AuthorityUsers.Type", "AuthorityUsers", "AuthorityUsers").Session(&gorm.Session{FullSaveAssociations: true}).Save(&pur.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			DeleteImage(deleteImageName, "projects")
			return nil
		}
		result := tx.Omit("Organization", "Tasks", "AuthorityUsers.User", "AuthorityUsers.Type", "AuthorityUsers.Project").Session(&gorm.Session{FullSaveAssociations: true}).Save(&pur.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return nil
	})
	return err
}


const (
	mainMessage = "下記のURLより参加できます。\n"
)