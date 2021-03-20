package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/Monkey-Mouse/mo2log/helpers"
	"github.com/Monkey-Mouse/mo2log/logmodel"
	"github.com/Monkey-Mouse/mo2log/service/logservice"

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
	col.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{Keys: bson.M{"processed": 1}},
		{Keys: bson.M{"operation_target_owner_id": 1}},
		{Keys: bson.M{"operator_id": 1}},
		{Keys: bson.M{"operation": 1}},
		{Keys: bson.M{"log_level": 1}},
		{Keys: bson.M{"create_time": -1}},
		{Keys: bson.M{"update_time": -1}},
		{Keys: bson.M{"operation_target_id": -1}},
	})
}

type server struct {
}

func (*server) Log(ctx context.Context, req *logservice.LogModel) (emp *logservice.Empty, err error) {
	model := logmodel.LogModel{
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
	emp = &logservice.Empty{}
	col.InsertOne(ctx, model)
	return
}
func model2Proto(log *logmodel.LogModel) *logservice.LogModel {
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
	m := &logmodel.LogModel{}
	err = col.FindOne(ctx, bson.M{
		"operator_id":         req.Operator,
		"operation":           req.Operation,
		"operation_target_id": req.OperationTarget}).Decode(m)
	model = model2Proto(m)

	return
}

func (*server) GetUserMsgs(ctx context.Context, req *logservice.ListRequest) (arr *logservice.LogArray, err error) {
	arr = &logservice.LogArray{}
	id := helpers.BytesToMongoID(req.UserId)
	cur, err := col.Find(ctx, bson.M{"operation_target_owner_id": id},
		options.Find().SetSkip(int64(req.Page)*int64(req.Pagesize)).SetLimit(int64(req.Pagesize)).SetSort(bson.M{"update_time": -1}))
	col.UpdateMany(ctx, bson.M{"operation_target_owner_id": id}, bson.D{{"$set", bson.M{"processed": true}}})
	if err != nil {
		return nil, err
	}
	log := &logmodel.LogModel{}
	logs := make([]*logservice.LogModel, 0, req.Pagesize)
	for {
		if !cur.Next(ctx) {
			break
		}
		err = cur.Decode(log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, model2Proto(log))
	}
	arr.Logs = logs
	return
}
func (*server) GetUserNewMsgNum(ctx context.Context, uid *logservice.UserID) (num *logservice.Num, err error) {
	num = &logservice.Num{}
	num.Num, err = col.CountDocuments(ctx, bson.M{
		"operation_target_owner_id": helpers.BytesToMongoID(uid.UserId),
		"processed":                 false})
	return
}
func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("LOG_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建 RPC 服务容器
	grpcServer := grpc.NewServer()
	logservice.RegisterLogServiceServer(grpcServer, &server{})
	log.Fatal(grpcServer.Serve(lis))
}
