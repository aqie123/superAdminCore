package queue

import (
	"github.com/aqie123/superAdminCore/queue/template"
	"github.com/aqie123/superAdminCore/task/email"
	"github.com/aqie123/superAdminCore/test/task/demo2"
	//namespace
)

var Queues = map[string]template.Task{
	"Demo2Task": &demo2.Demo2Task{Parameters: &demo2.Parameter{}},

	"email": &email.EmailTask{Parameters: &email.Parameter{}},

	//taskRegister
}
