package model

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	TrueID   string `db:"true_id"`
	School   string `db:"school"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Salt     string `db:"salt"`
	// SessionToken string    `db:"session_token"`
	Gender    int       `db:"gender"`
	IsDeleted bool      `db:"is_deleted"`
	LastLogin time.Time `db:"last_login"`
	Website   string    `db:"website"`
}

type Domain struct {
	ID         int       `db:"id"`
	OwnerID    int       `db:"owner_id"`
	Name       string    `db:"name"`
	Announce   string    `db:"announce"`
	Recommend  string    `db:"recommend"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	IsDeleted  bool      `db:"is_deleted"`
}

type Problem struct {
	ID          int       `db:"id"`
	Title       string    `db:"title"`
	Desc        string    `db:"description"`
	JudgeType   int       `db:"judge_type"`
	CreatorID   int       `db:"creator_id"`
	DomainID    int       `db:"domain_id"`
	InFmt       string    `db:"in_fmt"`
	OutFmt      string    `db:"out_fmt"`
	Other       string    `db:"other"`
	TestCases   string    `db:"test_cases"`
	SpecialCode string    `db:"special_code"`
	TimeLimit   int       `db:"time_limit"`
	MemoryLimit int       `db:"memory_limit"`
	Diff        int       `db:"diff"`
	Public      bool      `db:"public"`
	CreateTime  time.Time `db:"create_time"`
	UpdateTime  time.Time `db:"update_time"`
	IsDeleted   bool      `db:"is_deleted"`
	ACNum       int       `db:"ac_num"`
	SubmitNum   int       `db:"submit_num"`
}

type Homework struct {
	ID         int           `db:"id"`
	Title      string        `db:"title"`
	Desc       string        `db:"description"`
	ProblemIDs pq.Int64Array `db:"problem_ids"`
	CreatorID  int           `db:"creator_id"`
	DomainID   int           `db:"domain_id"`
	StartTime  time.Time     `db:"start_time"`
	EndTime    time.Time     `db:"end_time"`
	Public     bool          `db:"public"`
	CreateTime time.Time     `db:"create_time"`
	UpdateTime time.Time     `db:"update_time"`
	IsDeleted  bool          `db:"is_deleted"`
}

type Contest struct {
	ID         int           `db:"id"`
	Title      string        `db:"title"`
	Desc       string        `db:"description"`
	ProblemIDs pq.Int64Array `db:"problem_ids"`
	Typee      string        `db:"type"`
	CreatorID  int           `db:"creator_id"`
	DomainID   int           `db:"domain_id"`
	StartTime  time.Time     `db:"start_time"`
	EndTime    time.Time     `db:"end_time"`
	Public     bool          `db:"public"`
	CreateTime time.Time     `db:"create_time"`
	UpdateTime time.Time     `db:"update_time"`
	IsDeleted  bool          `db:"is_deleted"`
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
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Desc       string    `db:"description"`
	Permission int       `db:"permission"`
	DomainID   int       `db:"domain_id"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	IsDeleted  bool      `db:"is_deleted"`
}

type TestCase struct {
	Input    string `json:"input"`
	Expect   string `json:"expect"`
	IsSample bool   `json:"isSample"`
}

type Permission struct {
	Name string `db:"name" json:"name"`
	Bit  int    `db:"bit" json:"bit"`
}

type Submission struct {
	ID            int       `db:"id"`
	ProblemID     int       `db:"problem_id"`
	DomainID      int       `db:"domain_id"`
	FromType      string    `db:"from_type"`
	UserID        int       `db:"user_id"`
	Code          string    `db:"code"`
	SubmitTime    time.Time `db:"submit_time"`
	Status        int       `db:"status"`
	MaxMemory     int       `db:"max_memory"`
	MaxTime       int       `db:"max_time"`
	PassPercent   float32   `db:"pass_percent"`
	LastJudgeTime time.Time `db:"last_judge_time"`
	FromID        int       `db:"from_id"`
	Lang          string    `db:"lang"`
	Log           string    `db:"log"`
}

type Config struct {
	AddressList pq.StringArray `db:"address_list" json:"addressList"`
	Compilers   pq.StringArray `db:"compilers" json:"compilers"`
	Recommend   string         `db:"recommend" json:"recommend"`
	Announce    string         `db:"announce" json:"announce"`
}
