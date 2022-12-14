package controller

import (
	"github.com/aqie123/superAdminCore/contextPlus"
	"github.com/aqie123/superAdminCore/queue"
	"github.com/aqie123/superAdminCore/response"
	"github.com/aqie123/superAdminCore/task/email"
)

func Ping(c *contextPlus.Context) *response.Response {

	//e := errors.WithStack(ee.New("nice"))
	//
	//fmt.Println(fmt.Sprintf("%+v", e))

	return response.Resp().Api(1, "success", "ping")
}

func Test(c *contextPlus.Context) *response.Response {

	//logs.NewLogs().Error("123").Stdout()

	return response.Resp().Api(1, "success", "")
}

func Task(c *contextPlus.Context) *response.Response {

	queue.Dispatch(email.NewEmailTask("title", "904801074@qq.com", "content")).SetTryTime(10).Run()

	return response.Resp().Api(1, "success", "nice")

}
