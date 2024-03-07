package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/swarajp18/blogging-platform/proto"
	"google.golang.org/grpc"
)

var (
	port  = flag.Int("port", 50051, "gRPC server port")
	blogs []*pb.BlogData
)

type blogServer struct {
	pb.UnimplementedBlogServer
}

func init() {
	blog_1 := &pb.BlogData{PostID: "1", Title: "Spiderman", Content: "Spiderman is a good man", PublicationDate: "2009-11-10 23:00:00 +0000 UTC m=+0.000000001", Tags: []string{"spiderman", "blog", "mostliked"}}
	blog_2 := &pb.BlogData{PostID: "2", Title: "Ironman", Content: "Ironman is a bad man", PublicationDate: "2019-04-10 23:00:00 +0000 UTC m=+0.000000001", Tags: []string{"ironman", "blog", "worstliked"}}
	blogs = append(blogs, blog_1)
	blogs = append(blogs, blog_2)
}

func (s *blogServer) GetBlog(ctx context.Context, in *pb.ID) (*pb.BlogData, error) {
	log.Printf("Received GET BLOG request: %v", in)
	result := &pb.BlogData{}

	for _, blog := range blogs {
		if blog.GetPostID() == in.GetValue() {
			result = blog
			break
		}
	}

	return result, nil
}

func (s *blogServer) GetBlogs(in *pb.Empty, stream pb.Blog_GetBlogsServer) error {
	log.Printf("Received GET ALL BLOGS request")
	for _, blog := range blogs {
		if err := stream.Send(blog); err != nil {
			return err
		}
	}
	return nil
}

func (s *blogServer) CreateBlog(ctx context.Context, in *pb.BlogData) (*pb.BlogData, error) {
	log.Printf("Received CREATE BLOG request: %v", in)
	in.PostID = strconv.Itoa(rand.Intn(100000000))
	blogs = append(blogs, in)
	return in, nil
}

func (s *blogServer) UpdateBlog(ctx context.Context, in *pb.BlogData) (*pb.UpdateStatus, error) {
	log.Printf("Received UPDATE BLOG request: %v", in)

	result := &pb.Status{}
	for index, blog := range blogs {
		if blog.GetPostID() == in.GetPostID() {
			blogs = append(blogs[:index], blogs[index+1:]...)
			in.PostID = blog.GetPostID()
			blogs = append(blogs, in)
			result.Value = 1
			break
		}
	}
	updateStatus := &pb.UpdateStatus{
		Status:   result,
		BlogData: in,
	}

	return updateStatus, nil
}

func (s *blogServer) DeleteBlog(ctx context.Context, in *pb.ID) (*pb.Status, error) {
	log.Printf("Received DELETE BLOG request: %v", in)
	result := pb.Status{}
	for index, blog := range blogs {
		if blog.GetPostID() == in.GetValue() {
			blogs = append(blogs[:index], blogs[index+1:]...)
			result.Value = 1
			break
		}
	}

	return &result, nil
}

func main() {
	fmt.Println("gRPC server running ...")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBlogServer(s, &blogServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
