package database

import (
	"container/list"
	"database/sql"
	"fmt"
	"time"

	"github.com/dctewi/tewi-hwboard/config"

	log "unknwon.dev/clog/v2"

	_ "github.com/mattn/go-sqlite3" // noerror
)

func checkerr(err error) {
	defer func() {
		if x := recover(); x != nil {
			log.Error("Database PANIC: " + fmt.Sprint(x))
		}
	}()

	if err != nil {
		panic(err)
	}
}

// TaskInfo table item
type TaskInfo struct {
	ID       int
	Subject  string
	SubTitle string
	FileType string
	Start    time.Time
	End      time.Time
	ClassNO  int
}

// UserInfo table item
type UserInfo struct {
	ID        int
	Email     string
	StudentID string
}

// UploadLog table item
type UploadLog struct {
	ID       int
	Time     time.Time
	SubmitTo int
	Email    string
	FileMD5  string
}

// TaskInfoSlice for sort type
type TaskInfoSlice []TaskInfo

// Len for sort
func (tasks TaskInfoSlice) Len() int {
	return len(tasks)
}

// Less for sort by id (Less)
func (tasks TaskInfoSlice) Less(i, j int) bool {
	return tasks[i].ID > tasks[j].ID
}

// Swap for sort
func (tasks TaskInfoSlice) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

func openDB() *sql.DB {
	db, err := sql.Open("sqlite3", config.Database.Path)
	checkerr(err)

	return db
}

// Insert item
func Insert(item interface{}) {
	db := openDB()
	defer db.Close()

	switch item.(type) {
	case TaskInfo:
		state, err := db.Prepare("INSERT INTO taskinfo(subject, subtitle, filetype, start, end, classno) VALUES(?, ?, ?, ?, ?, ?)")

		v := item.(TaskInfo)
		_, err = state.Exec(v.Subject, v.SubTitle, v.FileType, v.Start, v.End, v.ClassNO)
		checkerr(err)
	case UserInfo:
		state, err := db.Prepare("INSERT INTO userinfo(email, studentid) VALUES(?, ?)")

		v := item.(UserInfo)
		_, err = state.Exec(v.Email, v.StudentID)
		checkerr(err)
	case UploadLog:
		state, err := db.Prepare("INSERT INTO uploadlog(time, submitto, email, filemd5) VALUES(?, ?, ?, ?)")

		v := item.(UploadLog)
		_, err = state.Exec(v.Time, v.SubmitTo, v.Email, v.FileMD5)
		checkerr(err)
	default:
		panic("wrong type of interted item")
	}
}

// Delete item
func Delete(item interface{}) {
	db := openDB()
	defer db.Close()

	switch item.(type) {
	case TaskInfo:
		state, err := db.Prepare("SELECT * FROM taskinfo WHERE id=?")
		v := item.(TaskInfo)
		rows, err := state.Query(v.ID)
		checkerr(err)
		info := TaskInfo{}
		for rows.Next() {
			err := rows.Scan(&info.ID, &info.Subject, &info.SubTitle, &info.FileType, &info.Start, &info.End, &info.ClassNO)
			checkerr(err)
			break
		}
		rows.Close()
		if info == item.(TaskInfo) {
			state, err := db.Prepare("DELETE FROM taskinfo WHERE id=?")
			v := item.(TaskInfo)
			_, err = state.Exec(v.ID)
			checkerr(err)
		}
	case UserInfo:
		state, err := db.Prepare("SELECT * FROM userinfo WHERE id=?")
		v := item.(UserInfo)
		rows, err := state.Query(v.ID)
		checkerr(err)
		info := UserInfo{}
		for rows.Next() {
			err := rows.Scan(&info.ID, &info.Email, &info.StudentID)
			checkerr(err)
			break
		}
		rows.Close()
		if info == item.(UserInfo) {
			state, err := db.Prepare("DELETE FROM userinfo WHERE id=?")
			v := item.(UserInfo)
			_, err = state.Exec(v.ID)
			checkerr(err)
		}
	case UploadLog:
		state, err := db.Prepare("SELECT * FROM uploadlog WHERE id=?")
		v := item.(UploadLog)
		rows, err := state.Query(v.ID)
		checkerr(err)
		info := UploadLog{}
		for rows.Next() {
			err := rows.Scan(&info.ID, &info.Time, &info.SubmitTo, &info.Email, &info.FileMD5)
			checkerr(err)
			break
		}
		rows.Close()
		if info == item.(UploadLog) {
			state, err := db.Prepare("DELETE FROM uploadlog WHERE id=?")
			v := item.(UploadLog)
			_, err = state.Exec(v.ID)
			checkerr(err)
		}
	default:
		panic("wrong type of deleted item")
	}
}

// Update item
func Update(id int, item interface{}) {
	db := openDB()
	defer db.Close()

	switch item.(type) {
	case TaskInfo:
		state, err := db.Prepare("UPDATE taskinfo SET subject=?, subtitle=?, filetype=?, start=?, end=?, classno=? WHERE id=?")

		v := item.(TaskInfo)
		_, err = state.Exec(v.Subject, v.SubTitle, v.FileType, v.Start, v.End, v.ClassNO, id)
		checkerr(err)
	case UserInfo:
		state, err := db.Prepare("UPDATE userinfo SET email=?, studentid=? WHERE id=?")

		v := item.(UserInfo)
		_, err = state.Exec(v.Email, v.StudentID, id)
		checkerr(err)
	case UploadLog:
		state, err := db.Prepare("UPDATE uploadlog SET time=?, submitto=?, email=?, filemd5=? WHERE id=?")

		v := item.(UploadLog)
		_, err = state.Exec(v.Time, v.SubmitTo, v.Email, v.FileMD5, id)
		checkerr(err)
	default:
		panic("wrong type of updated item")
	}
}

// GetAll returns all items
func GetAll(table string) *list.List {
	db := openDB()
	defer db.Close()
	res := list.New()

	switch table {
	case "taskinfo":
		rows, err := db.Query("SELECT * FROM taskinfo")
		checkerr(err)
		defer rows.Close()

		for rows.Next() {
			info := TaskInfo{}
			err := rows.Scan(&info.ID, &info.Subject, &info.SubTitle, &info.FileType, &info.Start, &info.End, &info.ClassNO)
			checkerr(err)
			res.PushBack(info)
		}
	case "userinfo":
		rows, err := db.Query("SELECT * FROM userinfo")
		checkerr(err)
		defer rows.Close()

		for rows.Next() {
			info := UserInfo{}
			err := rows.Scan(&info.ID, &info.Email, &info.StudentID)
			checkerr(err)
			res.PushBack(info)
		}
	case "uploadlog":
		rows, err := db.Query("SELECT * FROM uploadlog")
		checkerr(err)
		defer rows.Close()

		for rows.Next() {
			info := UploadLog{}
			err := rows.Scan(&info.ID, &info.Time, &info.SubmitTo, &info.Email, &info.FileMD5)
			checkerr(err)
			res.PushBack(info)
		}
	default:
		panic("no such table")
	}

	return res
}

// GetTaskBySubject rt
func GetTaskBySubject(subject string) *list.List {
	db := openDB()
	defer db.Close()
	res := list.New()

	state, err := db.Prepare("SELECT * FROM taskinfo WHERE subject=?")
	checkerr(err)
	rows, err := state.Query(subject)
	checkerr(err)
	defer rows.Close()

	for rows.Next() {
		info := TaskInfo{}
		err := rows.Scan(&info.ID, &info.Subject, &info.SubTitle, &info.FileType, &info.Start, &info.End, &info.ClassNO)
		checkerr(err)
		res.PushBack(info)
	}

	return res
}

// GetTaskByID rt
func GetTaskByID(id int) *TaskInfo {
	db := openDB()
	defer db.Close()
	info := &TaskInfo{}

	state, err := db.Prepare("SELECT * FROM taskinfo WHERE id=?")
	checkerr(err)
	rows, err := state.Query(id)
	checkerr(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&info.ID, &info.Subject, &info.SubTitle, &info.FileType, &info.Start, &info.End, &info.ClassNO)
		checkerr(err)
		break
	}

	return info
}

// GetUserInfoByEmail rt
func GetUserInfoByEmail(email string) UserInfo {
	db := openDB()
	defer db.Close()

	state, err := db.Prepare("SELECT * FROM userinfo WHERE email=?")
	checkerr(err)
	rows, err := state.Query(email)
	checkerr(err)
	defer rows.Close()

	info := UserInfo{}
	for rows.Next() {
		checkerr(rows.Scan(&info.ID, &info.Email, &info.StudentID))
		break
	}
	return info
}

// GetUserInfoBySID rt
func GetUserInfoBySID(stuid string) UserInfo {
	db := openDB()
	defer db.Close()

	state, err := db.Prepare("SELECT * FROM userinfo WHERE studentid=?")
	checkerr(err)
	rows, err := state.Query(stuid)
	checkerr(err)
	defer rows.Close()

	info := UserInfo{}
	for rows.Next() {
		checkerr(rows.Scan(&info.ID, &info.Email, &info.StudentID))
		break
	}
	return info
}

// GetUploadLogByEMail rt
func GetUploadLogByEMail(email string) []UploadLog {
	db := openDB()
	defer db.Close()

	state, err := db.Prepare("SELECT * FROM uploadlog WHERE email=?")
	checkerr(err)
	rows, err := state.Query(email)
	checkerr(err)
	defer rows.Close()

	var res []UploadLog
	for rows.Next() {
		info := UploadLog{}
		checkerr(rows.Scan(&info.ID, &info.Time, &info.SubmitTo, &info.Email, &info.FileMD5))
		res = append(res, info)
	}

	return res
}
