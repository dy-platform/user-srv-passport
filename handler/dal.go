package handler

//import (
//	"github.com/jinzhu/gorm"
//	"github.com/sirupsen/logrus"
//)

//var (
//	DefaultDB
//)

//func init() {
//	db, err := gorm.Open("mysql", "test.db")
//	if err != nil {
//
//	}
//}

func signUpByMobile(mobile, password, email string, deviceID int32, appID int32) error {
	// 1. 查找所有该手机号相关的记录
	// 2. 检查是否已经注册，若注册，抛出异常
	// 3. 无记录创建一条记录， 并将状态设置为enable
	// 4. 若存在记录，则直接修改状态
	// 5. 更改用户的密码
	// 6. 写日志



	return nil

}
