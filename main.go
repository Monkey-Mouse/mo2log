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
type server struct {
}

func (*server) Log(ctx context.Context, req *logservice.LogModel) (emp *logservice.Empty, err error) {
	model := LogModel{
		ID:                     primitive.NewObjectID(),
		OperatorID:             helpers.BytesToMongoID(req.Operator),
		Operation:              req.Operation,
		OperationTargetID:      helpers.BytesToMongoID(req.OperationTarget),
		LogLevel:               int32(req.LogLevel),
		ExtraMessage:           req.ExtraMessage,
		CreateTime:             time.Now(),
		UpdateTime:             time.Now(),
		OperationTargetOwnerID: helpers.BytesToMongoID(req.OperationTargetOwner),
		Processed:              false,
	}
	col.InsertOne(ctx, model)
	return
}
func model2Proto(log *LogModel) *logservice.LogModel {
	return &logservice.LogModel{
		Operator:             log.OperatorID[:],
		Operation:            log.Operation,
		OperationTarget:      log.OperationTargetID[:],
		LogLevel:             logservice.LogModel_Level(log.LogLevel),
		ExtraMessage:         log.ExtraMessage,
		OperationTargetOwner: log.OperationTargetOwnerID[:],
		CreateTime:           log.CreateTime.UnixNano(),
		UpdateTime:           log.UpdateTime.UnixNano(),
		Processed:            log.Processed,
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

func (*server) GetUserMsgs(ctx context.Context, req *logservice.ListRequest) (arr *logservice.LogArray, err error) {
	cur, err := col.Find(ctx, bson.M{"operation_target_owner_id": req.UserId},
		options.Find().SetSkip(int64(req.Page)*int64(req.Pagesize)).SetLimit(int64(req.Pagesize)).SetSort(bson.M{"update_time": -1}))
	col.UpdateMany(ctx, bson.M{"operation_target_owner_id": req.UserId}, bson.D{{"$set", bson.M{"processed": true}}})
	if err != nil {
		return nil, err
	}
	log := &LogModel{}
	logs := make([]*logservice.LogModel, 0, req.Pagesize)
	for {
		err = cur.Decode(log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, model2Proto(log))
		if !cur.TryNext(ctx) {
			break
		}
	}
	arr.Logs = logs
	return
}
func (*server) GetUserNewMsgNum(ctx context.Context, uid *logservice.UserID) (num *logservice.Num, err error) {
	num.Num, err = col.CountDocuments(ctx, bson.M{"operation_target_owner_id": uid.UserId})
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
