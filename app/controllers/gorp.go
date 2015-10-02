package controllers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/modules/db/app"
	r "github.com/revel/revel"
	"revel-bbs/app/models"
	"time"
)

var (
	Dbm *gorp.DbMap
)

const (
	DATE_FORMAT = "2006-01-02 15:04:05"
)

func InitDB() {
	r.INFO.Println("Start InitDB")
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTable(models.Message{}).SetKeys(true, "MessageId")
	setColumnSizes(t, map[string]int{
		"Name":    50,
		"Email":   50,
		"Title":   50,
		"Message": 200,
		"Created": 50,
	})

	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	messages := []*models.Message{
		&models.Message{0, "Scoot Marriott", "Tower@Pl.com", "Title Atlanta", "GA Message TEST", time.Now().Format(DATE_FORMAT)},
		&models.Message{0, "James W Hote", "Union@Square.com", "タイトル　Manhattan", "テスト　NY Message", time.Now().Format(DATE_FORMAT)},
		&models.Message{0, "Henry Rouge", "1315@16th.com", "TITLE @Washington", "てすと　メッセージ　DC Message", time.Now().Format(DATE_FORMAT)},
	}
	for _, message := range messages {
		if err := Dbm.Insert(message); err != nil {
			panic(err)
		}
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	r.INFO.Println("Start GorpController.Begin")
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	r.INFO.Println("Start GorpController.Commit")
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	r.INFO.Println("Start GorpController.Rollback")
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
