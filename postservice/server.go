package postservice

import (
	"context"
	"feed-service/data"
	pb "feed-service/proto"
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer
}

func (s *PostServiceServer) ListPostsByUser(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	userPosts := data.Posts[req.UserId]
	var protoPosts []*pb.Post
	for _, p := range userPosts {
		protoPosts = append(protoPosts, &pb.Post{
			Id:        p.ID,
			AuthorId:  p.AuthorID,
			Content:   p.Content,
			Timestamp: p.Timestamp.Unix(),
		})
	}
	return &pb.ListPostsResponse{Posts: protoPosts}, nil
}
