package mysql

import (
	"testing"
)

type account struct {
	Id        int    `json:"id" db:"id"`
	UserName  string `json:"user_name" db:"user_name"`
	PassWord  string `json:"password" db:"password"`
	IsDeleted int    `json:"is_deleted" db:"is_deleted"`
}

func TestInit(t *testing.T) {
	err := Init(Config{
		DataSource: DataSource{
			Driver:   "mysql",
			Username: "root",
			Password: "123456",
			Protocol: "tcp",
			Address:  "127.0.0.1",
			Port:     "3306",
			Dbname:   "test",
			Params:   "charset=utf8mb4&sql_notes=false&sql_notes=false&timeout=90s&collation=utf8mb4_general_ci&parseTime=True&loc=Local",
		},
		MaxIdleConns: 100,
		MaxOpenConns: 1000,
		KeepAlive:    3600,
	})
	if err != nil {
		t.Fatalf("query faied, error:[%v]", err.Error())
	}
	rows, err := DB().Query("select * from account")
	if err != nil {
		t.Fatalf("query faied, error:[%v]", err.Error())
	}

	for rows.Next() {
		var id, is_deleted []uint8
		var user_name, password string
		err := rows.Scan(&id, &is_deleted, &user_name, &password)
		if err != nil {
			t.Fatalf("query faied, error:[%v]", err.Error())
		}
		t.Log(id, is_deleted, user_name, password)
	}
	rows.Close()

	var account []account
	err = DB().Select(&account, "SELECT * FROM account WHERE is_deleted = 0")
	if err != nil {
		t.Fatalf("query faied, error:[%v]", err.Error())
	}
	t.Log(account)

	_, err = DB().Exec(`update account set is_deleted = ? where is_deleted = 0`, 1)
	if err != nil {
		t.Fatalf("query faied, error:[%v]", err.Error())
	}
}
