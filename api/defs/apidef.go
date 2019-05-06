package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	ID           string
	AuthorID     int
	Name         string
	DisplayCtime string
}
