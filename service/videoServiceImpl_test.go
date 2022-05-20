package service

import (
	"fmt"
	"testing"
	"time"
)

func getVideoService() VideoService {
	var videoService VideoServiceImpl
	videoService.UserService = VideoSub{}
	videoService.LikeService = VideoSub{}
	videoService.CommentService = VideoSub{}
	return &videoService
}

func getVideoService2() VideoService {
	var userService UserServiceImpl
	var followService FollowServiceImp
	var videoService VideoServiceImpl
	var likeService LikeServiceImpl
	var commentService CommentServiceImpl

	userService.FollowService = &followService

	followService.UserService = &userService

	likeService.VideoService = &videoService

	commentService.UserService = &userService

	videoService.CommentService = &commentService
	videoService.LikeService = &likeService
	videoService.UserService = &userService

	return &videoService
}

func TestList(t *testing.T) {
	videoService := getVideoService2()
	list, err := videoService.List(999)
	if err != nil {
		return
	}
	for _, video := range list {
		fmt.Println(video)
	}

}

func TestGetVideo(t *testing.T) {
	videoService := getVideoService2()
	video, err := videoService.GetVideo(1, 2)
	if err != nil {
		return
	}
	fmt.Println(video)
}

func TestFeed(t *testing.T) {
	videoService := getVideoService2()
	feed, t2, err := videoService.Feed(time.Now(), 2)
	if err != nil {
		return
	}
	for _, video := range feed {
		fmt.Println(video)
	}
	fmt.Println(t2)
}
