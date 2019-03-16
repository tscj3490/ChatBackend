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
)

// ForgotPassword handle client email to recovery password
func ForgotPassword(email_ string, role string) bool {
	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)
	fmt.Println(verifyCode)

	customer := &model.Customer{}
	vendor := &model.Vendor{}
	db.ORM.Table("customers").Where("username = ?", email_).Find(&customer)
	db.ORM.Table("vendors").Where("username = ?", email_).Find(&vendor)
	var fullname string
	if role == "customer" {
		fullname = customer.UserName
		go email.SendForgotEmail(customer.Email, fullname, verifyCode)
	}
	if role == "vendor" {
		fullname = vendor.UserName
		go email.SendForgotEmail(vendor.Email, fullname, verifyCode)
	}
	// send forgot email to user email

	return true
}

// ChangePassword
func ChangePassword(chgpass *model.ChangePass) (*model.ChangePass, error) {
	customer := &model.Customer{}
	vendor := &model.Vendor{}
	var err error
	//	var password []byte
	fmt.Println(chgpass.Email, chgpass.OldPass)

	//	password, err = bcrypt.GenerateFromPassword([]byte(chgpass.OldPass), 10)
	//	chgpass.OldPass = string(password)
	//	password, err = bcrypt.GenerateFromPassword([]byte(chgpass.NewPass), 10)
	//	chgpass.NewPass = string(password)

	if chgpass.Role == "customer" {
		if res := db.ORM.Table("customers").First(&customer, "email = ?", chgpass.Email).RecordNotFound(); res {
			err = errors.New(chgpass.Email + "doesn't exist, Please input correctly.")
			fmt.Println("This email doesn't exist, Please input correctly.")
			return nil, nil
		} else {
			//			if customer.Password == chgpass.OldPass {
			customer.Password = chgpass.NewPass
			db.ORM.Table("customers").Where("email = ?", chgpass.Email).Update("password", chgpass.NewPass)
			err = errors.New("Password is changed correctly.")
			//			} else {
			//				err = errors.New(chgpass.OldPass + "doesn't correct, Please input correctly.")
			//			}
		}
	}
	if chgpass.Role == "vendor" {
		if res := db.ORM.Table("vendors").First(&vendor, "email = ?", chgpass.Email).RecordNotFound(); res {
			err = errors.New(chgpass.Email + "doesn't exist, Please input correctly.")
			fmt.Println("This email doesn't exist, Please input correctly.")
			return nil, nil
		} else {
			//			if vendor.Password == chgpass.OldPass {
			vendor.Password = chgpass.NewPass
			db.ORM.Table("vendors").Where("email = ?", chgpass.Email).Update("password", chgpass.NewPass)
			err = errors.New("Password is changed correctly.")
			//			} else {
			//				err = errors.New(chgpass.OldPass + "doesn't correct, Please input correctly.")
			//			}
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
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Verification code :"+verifyCode))

	if err != nil {
		fmt.Println(err)
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
	if res := db.ORM.Where("phone = ?", phone).First(&user).RecordNotFound(); res {
		err := errors.New("You are an unregistered user.")
		return "", err
	}

	db.ORM.Table("users").Where("phone = ?", phone).UpdateColumn("code", verifyCode)

	client := twilio.NewClient(config.Sid, config.Token)
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Verification code :"+verifyCode))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message.Status)
	}

	return verifyCode, err
}

// AddOnlyPhone
func AddOnlyPhone(phone string, teamID uint) (string, error) {
	// generate verify code to reset password
	verifyCode := random.GenerateRandomDigitString(6)
	fmt.Println(verifyCode)

	client := twilio.NewClient(config.Sid, config.Token)
	message, err := twilio.NewMessage(client, config.ServerPhone, "+"+phone, twilio.Body("Verification code :"+verifyCode))

	fmt.Println(message.Status)

	user := &model.User{}
	if res := db.ORM.Where("phone = ?", phone).First(&user).RecordNotFound(); !res {
		db.ORM.Table("users").Where("phone = ?", phone).UpdateColumn("code", verifyCode)
		return verifyCode, nil
	}
	user.Phone = phone
	user.TeamID = teamID
	user.Code = verifyCode
	user.Role = "seller"
	user.IsVerified = false
	fmt.Println("+++", verifyCode)
	if err := db.ORM.Create(&user).Error; err != nil {
		return verifyCode, err
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
