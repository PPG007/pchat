package model

import (
	"pchat/repository/bson"
	"time"
)

const (
	C_MESSAGE = "message"

	MESSAGE_TYPE_FILE = "file"
	MESSAGE_TYPE_TEXT = "text"
)

var (
	CMessage = &Message{}
)

type Message struct {
	Id               bson.ObjectId   `bson:"_id"`
	CreatedAt        time.Time       `bson:"createdAt"`
	IsDeleted        bool            `bson:"isDeleted"`
	Sender           bson.ObjectId   `bson:"sender"`
	ReplyTo          bson.ObjectId   `bson:"replyTo,omitempty"`
	HasBeenEdited    bool            `bson:"hasBeenEdited"`
	ThreadId         bson.ObjectId   `bson:"threadId,omitempty"`
	IsInThread       bool            `bson:"isInThread"`
	ShowThreadInChat bool            `bson:"showThreadInChat"`
	ResponseEmojis   []string        `bson:"responseEmojis"`
	Type             string          `bson:"type"`
	Content          string          `bson:"content"`
	fileUrl          string          `bson:"fileUrl"`
	MentionedUsers   []bson.ObjectId `bson:"mentionedUsers,omitempty"`
}
