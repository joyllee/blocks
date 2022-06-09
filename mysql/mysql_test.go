package mysql

import (
	"fmt"
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
	"testing"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func TestInitMysql(t *testing.T) {
	logger.InitLogger()
	config.LoadConfig("../config/dev.yaml")
	err := InitMysql()
	defer DB.Close()
	if err != nil {
		t.Fatalf("init mysql failed: %v", err)
	}

	// 自动迁移
	DB.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "枯藤", "男", "篮球"}
	u2 := UserInfo{2, "topgoer.com", "女", "足球"}
	// 创建记录
	DB.Create(&u1)
	DB.Create(&u2)
	// 查询
	var u = new(UserInfo)
	DB.First(u)
	fmt.Printf("%#v\n", u)
	var uu UserInfo
	DB.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)
	// 更新
	DB.Model(&u).Update("hobby", "双色球")
	// 删除
	DB.Delete(&u)
}
