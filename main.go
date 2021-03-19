package main

import (
	"context"
	"log"
	"main/helpers"
	"main/service/logservice"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var col *mongo.Collection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MO2_MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(os.Getenv("MO2_DATABASE"))
	col = db.Collection(os.Getenv("LOG_COL"))
}

type LogModel struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	OperatorID        primitive.ObjectID `bson:"operator_id,omitempty"`
	Operation         int32              `bson:"operation,omitempty"`
	OperationTargetID primitive.ObjectID `bson:"operation_target_id,omitempty"`
	LogLevel          int32              `bson:"log_level,omitempty"`
	ExtraMessage      string             `bson:"extra_message,omitempty"`
	CreateTime        time.Time          `bson:"create_time,omitempty"`
	UpdateTime        time.Time          `bson:"update_time,omitempty"`
}
type server struct {
}

func (*server) Log(ctx context.Context, req *logservice.LogModel) (emp *logservice.Empty, err error) {
	model := LogModel{
		ID:                primitive.NewObjectID(),
		OperatorID:        helpers.BytesToMongoID(req.Operator),
		Operation:         req.Operation,
		OperationTargetID: helpers.BytesToMongoID(req.OperationTarget),
		LogLevel:          int32(req.LogLevel),
		ExtraMessage:      req.ExtraMessage,
		CreateTime:        time.Now(),
		UpdateTime:        time.Now(),
	}
	col.InsertOne(ctx, model)
	return
}
func model2Proto(log *LogModel) *logservice.LogModel {
	return &logservice.LogModel{
		Operator:        log.OperatorID[:],
		Operation:       log.Operation,
		OperationTarget: log.OperationTargetID[:],
		LogLevel:        logservice.LogModel_Level(log.LogLevel),
		ExtraMessage:    log.ExtraMessage,
		CreateTime:      log.CreateTime.UnixNano(),
		UpdateTime:      log.UpdateTime.UnixNano(),
	}
}
func (*server) Exist(ctx context.Context,
	req *logservice.ExtRequest) (model *logservice.LogModel, err error) {
	m := &LogModel{}
	err = col.FindOne(ctx, bson.M{
		"operator_id":         req.Operator,
		"operation":           req.Operation,
		"operation_target_id": req.OperationTarget}).Decode(m)
	model = model2Proto(m)

	return
}

func main() {
	lis, err := net.Listen("tcp", os.Getenv("LOG_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建 RPC 服务容器
	grpcServer := grpc.NewServer()
	logservice.RegisterLogServiceServer(grpcServer, &server{})
	log.Fatal(grpcServer.Serve(lis))
}
