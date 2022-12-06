package artisan

import (
	"github.com/aqie123/superAdminCore/artisan"
	"github.com/aqie123/superAdminCore/test/artisan/demo"
)

func Artisan() []artisan.Artisan {

	return []artisan.Artisan{
		new(demo.Demo),
	}
}
