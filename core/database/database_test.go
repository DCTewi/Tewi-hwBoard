package database_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dctewi/tewi-hwboard/core/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestInsert(t *testing.T) {
	database.Insert(database.TaskInfo{
		Subject:  "123",
		SubTitle: "sub",
		FileType: "docx",
		Start:    time.Now(),
		End:      time.Now(),
	})

	database.Insert(database.UserInfo{
		Email:     "123@qwe.com",
		StudentID: "20174215535",
	})

	database.Insert(database.UploadLog{
		Time:     time.Now(),
		SubmitTo: 1,
		Email:    "123@qwe.com",
		FileMD5:  "123214qarqwfser3ewqfdrewsdg43",
	})
}

func TestGetAll(t *testing.T) {
	tasks := database.GetAll("taskinfo")
	users := database.GetAll("userinfo")
	logs := database.GetAll("uploadlog")

	fmt.Println("--------")

	fmt.Println(tasks.Len())
	for i := tasks.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println(users.Len())
	for i := users.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println(logs.Len())
	for i := logs.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("--------")
}

func TestGetTaskBySubject(t *testing.T) {
	a, b := database.GetTaskBySubject("123"), database.GetTaskBySubject("567")

	fmt.Println("--------")

	fmt.Println(a.Len())
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println(b.Len())
	for i := b.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("--------")
}

func TestUserInfo(t *testing.T) {
	a, b := database.GetUserInfoBySID("20174215535"), database.GetUserInfoByEmail("123@qwe.com")

	fmt.Println("--------")

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println("--------")
}

func TestUpdate(t *testing.T) {
	fmt.Println("--------")

	a := database.GetUserInfoBySID("20174215535")

	a.Email = "hello@world.com"

	database.Update(a.ID, a)

	fmt.Println("--------")
}

func TestDelete(t *testing.T) {
	fmt.Println("--------")

	a := database.GetUserInfoBySID("20174215535")

	database.Delete(a)

	fmt.Println("--------")
}

func TestGetTaskByID(t *testing.T) {
	info := database.GetTaskByID(10)

	fmt.Println(info.End.Format("060102"))

	fmt.Println(info)
}

func TestGetUploadLogByEMail(t *testing.T) {
	infos := database.GetUploadLogByEMail("123@qwe.com")

	fmt.Println(infos)
}
