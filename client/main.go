package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/swarajp18/blogging-platform/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type Blog struct {
	Title           string   `json:"Title"`
	Content         string   `json:"Content"`
	PublicationDate string   `json:"PublicationDate"`
	Tags            []string `json:"Tags"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewBlogClient(conn)

	r := gin.Default()

	r.GET("/blogs", func(ctx *gin.Context) {
		request := &pb.Empty{}
		stream, err := client.GetBlogs(ctx, request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		for {
			row, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v.GetBlogs(_) = _, %v", client, err)
			}
			log.Printf("BlogData: %v", row)

			ctx.JSON(http.StatusOK, gin.H{
				"blogs": row,
			})
		}
	})

	r.GET("/blog/:id", func(ctx *gin.Context) {
		request := &pb.ID{Value: ctx.Param("id")}
		result, err := client.GetBlog(ctx, request)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"blog": result,
		})
	})

	r.POST("/blog", func(ctx *gin.Context) {
		var blog Blog
		err := ctx.ShouldBind(&blog)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		data := &pb.BlogData{
			PostID:          "",
			Title:           blog.Title,
			Content:         blog.Content,
			PublicationDate: blog.PublicationDate,
			Tags:            blog.Tags,
		}

		result, err := client.CreateBlog(ctx, data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"blog": result,
		})
	})

	r.PUT("/blog/:id", func(ctx *gin.Context) {
		request := &pb.ID{Value: ctx.Param("id")}
		var blog Blog
		err := ctx.ShouldBind(&blog)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		data := &pb.BlogData{
			PostID:          request.Value,
			Title:           blog.Title,
			Content:         blog.Content,
			PublicationDate: blog.PublicationDate,
			Tags:            blog.Tags,
		}

		result, err := client.UpdateBlog(ctx, data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if result.Status.Value == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"blog": result.BlogData,
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error updating blog",
			})
			return
		}
	})

	r.DELETE("/blog/:id", func(ctx *gin.Context) {
		request := &pb.ID{Value: ctx.Param("id")}
		result, err := client.DeleteBlog(ctx, request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if result.Value == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Blog deleted successfully",
			})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error deleting blog",
			})
			return
		}
	})

	r.Run(":5001")
}
