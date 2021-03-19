package logmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogModel struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	OperatorID             primitive.ObjectID `bson:"operator_id,omitempty"`
	Operation              int32              `bson:"operation,omitempty"`
	OperationTargetID      primitive.ObjectID `bson:"operation_target_id,omitempty"`
	LogLevel               int32              `bson:"log_level,omitempty"`
	ExtraMessage           string             `bson:"extra_message,omitempty"`
	CreateTime             time.Time          `bson:"create_time,omitempty"`
	UpdateTime             time.Time          `bson:"update_time,omitempty"`
	OperationTargetOwnerID primitive.ObjectID `bson:"operation_target_owner_id,omitempty"`
	Processed              bool               `bson:"processed,omitempty"`
}