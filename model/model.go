package model

import (
	"time"
)

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	TrueID       string    `db:"true_id"`
	School       string    `db:"school"`
	Password     string    `db:"password"`
	Email        string    `db:"email"`
	Salt         string    `db:"salt"`
	SessionToken string    `db:"session_token"`
	Gender       int       `db:"gender"`
	IsDeleted    bool      `db:"is_deleted"`
	LastLogin    time.Time `db:"last_login"`
}

type Domain struct {
	ID        int    `db:"id"`
	OwnerID   int    `db:"owner_id"`
	Name      string `db:"name"`
	Announce  string `db:"announce"`
	IsDeleted bool   `db:"is_deleted"`
}

type Problem struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Desc        string    `db:"description"`
	CreatorID   int       `db:"creator_id"`
	DomainID    int       `db:"domain_id"`
	InFmt       string    `db:"in_fmt"`
	OutFmt      string    `db:"out_fmt"`
	Other       string    `db:"other"`
	TimeLimit   int       `db:"time_limit"`
	MemoryLimit int       `db:"memory_limit"`
	Diff        int       `db:"diff"`
	Public      bool      `db:"public"`
	CreateTime  time.Time `db:"create_time"`
	UpdateTime  time.Time `db:"update_time"`
	IsDeleted   bool      `db:"is_deleted"`
}

type Homework struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Desc       string    `db:"description"`
	CreatorID  int       `db:"creator_id"`
	DomainID   int       `db:"domain_id"`
	StartTime  time.Time `db:"start_time"`
	EndTime    time.Time `db:"end_time"`
	Public     bool      `db:"public"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	IsDeleted  bool      `db:"is_deleted"`
}

type Contest struct {
	ID             int       `db:"id"`
	Title          string    `db:"title"`
	Desc           string    `db:"description"`
	Typee          string    `db:"type"`
	ParticipantNum int       `db:"participant_num"`
	CreatorID      int       `db:"creator_id"`
	DomainID       int       `db:"domain_id"`
	StartTime      time.Time `db:"start_time"`
	EndTime        time.Time `db:"end_time"`
	Public         bool      `db:"public"`
	CreateTime     time.Time `db:"create_time"`
	UpdateTime     time.Time `db:"update_time"`
	IsDeleted      bool      `db:"is_deleted"`
}

type Discussion struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	CommentNum int       `db:"comment_num"`
	CreatorID  int       `db:"creator_id"`
	DomainID   int       `db:"domain_id"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	IsDeleted  bool      `db:"is_deleted"`
}

type DiscussionComment struct {
	ID           int       `db:"id"`
	Content      string    `db:"content"`
	CreatorID    int       `db:"creator_id"`
	DiscussionID int       `db:"discussion_id"`
	FloorID      int       `db:"floor_id"`
	ReplyID      int       `db:"reply_id"`
	CreateTime   time.Time `db:"create_time"`
	IsDeleted    bool      `db:"is_deleted"`
}

type Role struct {
}

type Permission struct {
}

type RolePermission struct {
}
