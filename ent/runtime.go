// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/amaru0601/fluent/ent/chat"
	"github.com/amaru0601/fluent/ent/message"
	"github.com/amaru0601/fluent/ent/schema"
	"github.com/amaru0601/fluent/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chatFields := schema.Chat{}.Fields()
	_ = chatFields
	// chatDescDeleted is the schema descriptor for deleted field.
	chatDescDeleted := chatFields[1].Descriptor()
	// chat.DefaultDeleted holds the default value on creation for the deleted field.
	chat.DefaultDeleted = chatDescDeleted.Default.(bool)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescBody is the schema descriptor for body field.
	messageDescBody := messageFields[0].Descriptor()
	// message.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	message.BodyValidator = messageDescBody.Validators[0].(func(string) error)
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[1].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescPrivateKey is the schema descriptor for private_key field.
	userDescPrivateKey := userFields[3].Descriptor()
	// user.DefaultPrivateKey holds the default value on creation for the private_key field.
	user.DefaultPrivateKey = userDescPrivateKey.Default.(string)
}
