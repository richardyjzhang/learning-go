package dbops

import (
	"strconv"
	"testing"
	"time"
)

var tempVid string

func clearTables() {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDelUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("jack", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("jack")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDelUser(t *testing.T) {
	err := DeleteUser("jack", "123")
	if err != nil {
		t.Errorf("Error of DelUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("jack")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Error: DelUser Failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDelVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	res, err := AddNewVideo(1, "AV")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}

	tempVid = res.ID
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo")
	}
}

func testDelVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DelVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	res, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
	if res != nil {
		t.Errorf("Error of DelVideoInfo Failed")
	}
}

func TestCommentWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComments", testListComments)
}

func testAddComment(t *testing.T) {
	vid := "AV"
	aid := 1
	content := "I like this AV"

	err := AddNewComment(vid, aid, content)

	if err != nil {
		t.Errorf("Error of Add Comment: %s", err)
	}
}

func testListComments(t *testing.T) {
	vid := "AV"
	from := 0
	to, _ := strconv.Atoi(strconv.FormatInt(
		time.Now().UnixNano()/1e9, 10))
	_, err := ListComments(vid, from, to)

	if err != nil {
		t.Errorf("Error of List Comment: %s", err)
	}
}
