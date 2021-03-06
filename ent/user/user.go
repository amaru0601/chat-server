// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeChats holds the string denoting the chats edge name in mutations.
	EdgeChats = "chats"
	// Table holds the table name of the user in the database.
	Table = "users"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "user_messages"
	// ChatsTable is the table that holds the chats relation/edge. The primary key declared below.
	ChatsTable = "chat_members"
	// ChatsInverseTable is the table name for the Chat entity.
	// It exists in this package in order to avoid circular dependency with the "chat" package.
	ChatsInverseTable = "chats"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldCreatedAt,
	FieldPassword,
	FieldPrivateKey,
	FieldPublicKey,
}

var (
	// ChatsPrimaryKey and ChatsColumn2 are the table columns denoting the
	// primary key for the chats relation (M2M).
	ChatsPrimaryKey = []string{"chat_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultPrivateKey holds the default value on creation for the "private_key" field.
	DefaultPrivateKey string
)
