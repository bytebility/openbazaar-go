package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	notif "github.com/OpenBazaar/openbazaar-go/api/notifications"
	"strconv"
	"strings"
	"sync"
	"time"
)

type NotficationsDB struct {
	db   *sql.DB
	lock sync.RWMutex
}

func (n *NotficationsDB) Put(notification notif.Data, notifType string, timestamp time.Time) error {
	ser, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	n.lock.Lock()
	defer n.lock.Unlock()
	tx, _ := n.db.Begin()
	stmt, _ := tx.Prepare("insert into notifications(serializedNotification, type, timestamp, read) values(?,?,?,?)")

	defer stmt.Close()
	_, err = stmt.Exec(string(ser), strings.ToLower(notifType), int(timestamp.Unix()), 0)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (n *NotficationsDB) GetAll(offsetId int, limit int, typeFilter string) []notif.Notification {
	var ret []notif.Notification

	n.lock.RLock()
	defer n.lock.RUnlock()

	var stm string
	var filter string
	var args []interface{}
	if offsetId > 0 {
		args = append(args, offsetId)
		if typeFilter != "" {
			filter = " and type=?"
			args = append(args, strings.ToLower(typeFilter))
		}
		stm = "select rowid, serializedNotification, timestamp, read from notifications where rowid<?" + filter + " order by rowid desc limit " + strconv.Itoa(limit) + " ;"
	} else {
		if typeFilter != "" {
			filter = " where type=?"
			args = append(args, strings.ToLower(typeFilter))
		}
		stm = "select rowid, serializedNotification, timestamp, read from notifications" + filter + " order by timestamp desc limit " + strconv.Itoa(limit) + ";"
	}
	rows, err := n.db.Query(stm, args...)
	if err != nil {
		return ret
	}
	for rows.Next() {
		var notifId int
		var data []byte
		var timestampInt int
		var readInt int
		if err := rows.Scan(&notifId, &data, &timestampInt, &readInt); err != nil {
			fmt.Println(err)
			continue
		}
		var read bool
		if readInt == 1 {
			read = true
		}
		timestamp := time.Unix(int64(timestampInt), 0)
		var ni interface{}
		err := json.Unmarshal(data, &ni)
		if err != nil {
			fmt.Println(err)
			continue
		}
		n := notif.Notification{
			ID:        notifId,
			Data:      ni,
			Timestamp: timestamp,
			Read:      read,
		}
		ret = append(ret, n)
	}
	return ret
}

func (n *NotficationsDB) MarkAsRead(notifID int) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	tx, err := n.db.Begin()
	if err != nil {
		return err
	}
	stmt, _ := tx.Prepare("update notifications set read=1 where rowid=?")

	defer stmt.Close()
	_, err = stmt.Exec(notifID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (n *NotficationsDB) MarkAllAsRead() error {
	n.lock.Lock()
	defer n.lock.Unlock()
	_, err := n.db.Exec("update notifications set read=1")
	return err
}

func (n *NotficationsDB) Delete(notifID int) error {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.db.Exec("delete from notifications where rowid=?", notifID)
	return nil
}

func (n *NotficationsDB) GetUnreadCount() (int, error) {
	stm := "select Count(*) from notifications where read=0;"
	row := n.db.QueryRow(stm)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
