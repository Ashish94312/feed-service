package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	pb "feed-service/proto"
)

type Resolver struct {
	PostsClient pb.PostServiceClient
}
