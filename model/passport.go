package model


type UserPassport struct {
	UID       int64  `bson:"_id"`
	DeviceID  int64  `bson:"device_id"`
	Name      string `bson:"name"`
	Password  string `bson:"password"`
	Salt      string `bson:"salt"`
	Mobile    string `bson:"mobile"`
	Email     string `bson:"email"`
	WeChatID  string `bson:"wechat_id"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}


