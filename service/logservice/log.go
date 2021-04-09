package logservice

import (
	context "context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	grpc "google.golang.org/grpc"
)

func (c *logServiceClient) SendCountQuery(ctx context.Context, query bson.D, opts ...grpc.CallOption) (num *Num, err error) {
	in := &Query{}
	in.Query, err = json.Marshal(query)
	if err != nil {
		return nil, err
	}
	return c.CountQuery(ctx, in, opts...)
}
