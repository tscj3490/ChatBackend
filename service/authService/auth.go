package authService

import (
	"errors"
	"fmt"

	"../../config"
	"../../db"
	"../../model"
	"../../util/email"
	"../../util/random"
	twilio "github.com/carlosdp/twiliogo"
	"golang.org/x/crypto/bcrypt"
)

// ForgotPassword handle client email to recovery password
func ForgotPassword(email_ string) bool {

	// generate verify code to reset password
	code := random.GenerateRandomDigitString(6)

	admin := &model.Admin{}
	res := db.ORM.Table("admins").Where("email = ?", email_).Find(&admin).RecordNotFound()
	if res {
		return false
	}

	password, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		return false
	}

	db.ORM.Table("admins").Where("email = ?", email_).Update("password", string(password))

	go email.SendForgotEmail(email_, code)

	return true
}

// ChangePassword
func ChangePassword(chgpass *model.ChangePass) (*model.ChangePass, error) {
	admin := &model.Admin{}
	var err error
	var password []byte

	if res := db.ORM.Table("admins").First(&admin, "email = ?", chgpass.Email).RecordNotFound(); res {
		err = errors.New(chgpass.Email + "doesn't exist, Please input correctly.")
		fmt.Println("This email doesn't exist, Please input correctly.")
		return nil, err
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(chgpass.OldPass))
		if err != nil {
			fmt.Println("Here-1")
			err = errors.New("Old password doesn't correct, Please input correctly.")
			return nil, err
		} else {
			fmt.Println("Here-2")
			password, err = bcrypt.GenerateFromPassword([]byte(chgpass.NewPass), 10)
			chgpass.NewPass = string(password)
			db.ORM.Table("admins").Where("email = ?", chgpass.Email).Update("password", chgpass.NewPass)
		}
	}

	return chgpass, err
}

// SendCode
func SendCode(phone string) (string, error) {
	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)
	fmt.Println(verifyCode)

	client := twilio.NewClient(config.Sid, config.Token)
	fmt.Println("-----", config.Sid, config.Token, config.ServerPhone, phone)
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Here’s your Sellaboo Verification Code : "+verifyCode))

	if err != nil {
		fmt.Println(err)
		err := errors.New("The phone number is invalid.")
		return "", err
	} else {
		fmt.Println(message.Status)
	}

	user := &model.User{}
	if res := db.ORM.Where("phone = ?", phone).First(&user).RecordNotFound(); !res {
		db.ORM.Table("users").Where("phone = ?", phone).UpdateColumn("code", verifyCode)
		return verifyCode, nil
	}
	user.Phone = phone
	user.Code = verifyCode
	fmt.Println("+++", verifyCode)
	if err := db.ORM.Create(&user).Error; err != nil {
		return verifyCode, err
	}

	return verifyCode, err
}

// CheckPhone
func CheckPhone(phone string) (string, error) {
	user := &model.User{}

	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)

	client := twilio.NewClient(config.Sid, config.Token)
	fmt.Println("-----", config.Sid, config.Token, config.ServerPhone, phone)
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Here’s your Sellaboo Verification Code : "+verifyCode))

	if err != nil {
		fmt.Println(err)
		err := errors.New("The phone number is invalid.")
		return "", err
	} else {
		fmt.Println(message.Status)
	}

	if res := db.ORM.Table("users").Where("phone = ?", phone).First(&user).RecordNotFound(); res {
		fmt.Println("---------+++-", user, res)
		err := errors.New("You are an unregistered user.")
		return "", err
	}
	fmt.Println("------------")
	db.ORM.Table("users").Where("phone = ?", phone).UpdateColumn("code", verifyCode)

	return verifyCode, err
}

// AddOnlyPhone
func AddOnlyPhone(phone string, teamID uint) (string, error) {
	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)
	fmt.Println(verifyCode)

	client := twilio.NewClient(config.Sid, config.Token)
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Here’s your Sellaboo Verification Code : "+verifyCode))

	if err != nil {
		fmt.Println(err)
		err := errors.New("The phone number is invalid.")
		return "", err
	} else {
		fmt.Println(message.Status)
	}

	user := &model.User{}

	if res := db.ORM.Where("phone = ?", phone).First(&user).RecordNotFound(); !res {
		db.ORM.Table("users").Where("phone = ?", phone).UpdateColumn("code", verifyCode)
		return verifyCode, nil
	} else {
		user.Phone = phone
		user.TeamID = teamID
		user.Role = "seller"
		user.IsVerified = false
		user.Code = verifyCode
		if err := db.ORM.Create(&user).Error; err != nil {
			return verifyCode, err
		}
	}

	return verifyCode, err
}

// VerifyCode
func VerifyCode(code string) (uint, bool, *model.User, error) {

	// generate verify code to reset password
	var objid uint
	var result bool
	var err error
	user := &model.User{}

	err = db.ORM.Table("users").Where("code = ?", code).Find(&user).Error
	if err == nil {
		objid = user.ID
		result = true
		db.ORM.Table("users").UpdateColumn("is_verified", true)
		// db.ORM.Table("users").UpdateColumn("role", "seller")
	} else {
		result = false
	}

	return objid, result, user, err
}

// AddManager
func AddManager(managerInfo *model.ManagerInfo) (*model.User, error) {

	var err error
	user := &model.User{}

	user.TeamID = managerInfo.Team.ID
	user.Avatar = managerInfo.Team.Logo
	user.Name = managerInfo.Name
	user.Phone = managerInfo.Phone
	user.Email = managerInfo.Email
	user.Role = "manager"
	user.IsVerified = false

	if res := db.ORM.Where("phone = ?", user.Phone).First(&user).RecordNotFound(); !res {
		err := errors.New(user.Phone + " is already registered")
		return nil, err
	}

	// Insert Data
	if err := db.ORM.Create(&user).Error; err != nil {
		return nil, err
	}
	err = db.ORM.Last(&user).Error

	return user, err
}

// VerifyRole
func VerifyRole(ID uint) (*model.User, error) {
	user := &model.User{}
	if err := db.ORM.Where("role = ?", "manager").First(&user, ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}
