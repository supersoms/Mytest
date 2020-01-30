package models

import (
	orm "CXCProject/database"
	"CXCProject/message"
	"errors"
	"log"
)

type Oauth struct {
	Result  User   `json:"result"`
	Code    string `json:"code"`
	Success bool   `json:"success"`
}

type User struct {
	ID      int    `form:"id"`
	Cid     string `json:"cid" form:"address" binding:"required`
	Address string `json:"address" form:"cid" binding:"required`
}

var Users []User

//用户注册
func (u User) OauthLogin() (err error) {
	user := User{}
	orm.DB.Where("cid = ?", u.Cid).Find(&user) //根据cid检查用户是否已注册，记住这个Find()方法必须指定跟表名一致的名字是user，不能是u
	if user.Cid == u.Cid {
		err = errors.New(message.USER_REGISTER_FAIL_EXISTED_MSG)
		return
	}
	result := orm.DB.Create(&u)
	if result.Error != nil {
		log.Println(result.Error)
		err = errors.New(message.USER_REGISTER_FAIL_MSG)
		return
	}
	//defer orm.DB.Close()
	return
}

func (u User) Login() (err error) {
	user := User{}
	orm.DB.Where("cid = ?", u.Cid).Find(&user) //先检查用户是否已注册
	if len(user.Cid) == 0 {
		err = errors.New(message.USER_NOT_EXIST_MSG)
		return
	} else if len(user.Address) == 0 {
		err = errors.New(message.USER_NOT_EXIST_MSG)
		return
	} else if user.Cid != u.Cid {
		err = errors.New(message.USER_NOT_EXIST_MSG)
		return
	}
	return
}

//获取所有的用户列表
func (user *User) Users() (users []User, err error) {
	if err = orm.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

//根据id获取用户信息
func (user *User) GetUserById() (u *User, err error) {
	us := make([]User, 10)
	//SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
	if err = orm.DB.Where(&User{ID: user.ID}).First(&us).Error; err != nil {
		return
	}
	return
}

//根据id修改用户信息
func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.DB.Select([]string{"id", "cid"}).First(&updateUser, id).Error; err != nil {
		return
	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.DB.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//根据id删除用户
func (user *User) Delete(id int64) (Result User, err error) {
	if err = orm.DB.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.DB.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}
