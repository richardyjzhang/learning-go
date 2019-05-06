package dbops

import (
	"database/sql"
	"log"
	"time"

	"github.com/richardyjzhang/learning-go/api/defs"
	"github.com/richardyjzhang/learning-go/api/utils"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		log.Printf("Add User Error: %s", err)
		return err
	}

	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("Get User Error: %s", err)
		return "", err
	}

	var pwd string
	err = stmt.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmt.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmt, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("Delete User Error: %s", err)
		return err
	}

	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// Create UUID
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmt, err := dbConn.Prepare(`INSERT INTO video_info 
		(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{ID: vid, AuthorID: aid, Name: name, DisplayCtime: ctime}

	defer stmt.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmt, err := dbConn.Prepare(`SELECT author_id, name, display_ctime 
		FROM video_info WHERE id = ?`)
	if err != nil {
		log.Printf("Get Video error: %s", err)
		return nil, err
	}

	var aid int
	var displayCTime string
	var name string

	err = stmt.QueryRow(vid).Scan(&aid, &name, &displayCTime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmt.Close()

	res := &defs.VideoInfo{ID: vid, AuthorID: aid, Name: name, DisplayCtime: displayCTime}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmt, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("Delete Video Error: %s", err)
		return err
	}

	_, err = stmt.Exec(vid)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func AddNewComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmt, err := dbConn.Prepare(`INSERT INTO comments 
		(id, video_id, author_id, content) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmt.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmt, err := dbConn.Prepare(`
		SELECT comments.id, users.login_name, comments.content
		FROM comments 
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? 
		AND comments.time > FROM_UNIXTIME(?)
		AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment
	if err != nil {
		log.Printf("Error List Comments: %s", err)
		return res, err
	}

	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{ID: id, VideoID: vid, AuthorID: name, Content: content}
		res = append(res, c)
	}

	defer stmt.Close()

	return res, nil
}
