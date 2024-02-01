package chat

import (
	"context"
	"github.com/qiniu/qmgo"
	pb_chat "pchat/pb/chat"
	"pchat/repository"
	"pchat/repository/bson"
	"pchat/utils"
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
	ChatId           bson.ObjectId   `bson:"chatId"`
}

func (*Message) CreateFromPb(ctx context.Context, pb *pb_chat.NewMessage) (Message, error) {
	message := Message{
		Sender: utils.GetUserIdAsObjectId(ctx),
	}
	err := utils.Copier().From(pb).To(&message)
	if err != nil {
		return message, err
	}
	if bson.IsObjectIdHex(pb.Id) {
		message.HasBeenEdited = true
	} else {
		message.Id = bson.NewObjectId()
	}
	condition := bson.M{
		"_id": message.Id,
	}
	setter := bson.M{
		"sender":           message.Sender,
		"replyTo":          message.ReplyTo,
		"hasBeenEdited":    message.HasBeenEdited,
		"threadId":         message.ThreadId,
		"isInThread":       message.IsInThread,
		"showThreadInChat": message.ShowThreadInChat,
		"type":             message.Type,
		"content":          message.Content,
		"fileUrl":          message.fileUrl,
		"mentionedUsers":   message.MentionedUsers,
	}
	if message.ReplyTo.IsZero() {
		delete(setter, "replyTo")
	}
	if message.ThreadId.IsZero() {
		delete(setter, "threadId")
	}
	updater := bson.M{
		"$set": setter,
		"$setOnInsert": bson.M{
			"createdAt": time.Now(),
			"isDeleted": false,
		},
	}
	change := qmgo.Change{
		Upsert:    true,
		ReturnNew: true,
		Update:    updater,
	}
	err = repository.FindAndApply(ctx, C_MESSAGE, condition, change, &message)
	return message, err
}
