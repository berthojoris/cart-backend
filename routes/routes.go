package routes

import (
	"github.com/berthojoris/cart-backend/app/repositories/impl"
	impl2 "github.com/berthojoris/cart-backend/app/services/impl"
	"github.com/berthojoris/cart-backend/app/utils"
	"github.com/berthojoris/cart-backend/app/web/controllers"
	"github.com/berthojoris/cart-backend/app/web/middleware"
	"github.com/berthojoris/cart-backend/bootstrap"
	"github.com/berthojoris/cart-backend/config"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Route struct {
	Config      *config.Configuration
	CorsHandler context.Handler
}

func NewRoute(config *config.Configuration) *Route {
	return &Route{
		Config: config,
		CorsHandler: cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			AllowedHeaders:   []string{"*"},
		}),
	}
}

func (r *Route) Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", controllers.GetHomeHandler)

	// repositories
	anouncementRequestRepository := impl.NewAnouncementRepositoryImpl()

	// services
	anouncementService := impl2.NewAnouncementServiceImpl(anouncementRequestRepository)

	v1 := b.Party("/v1", r.CorsHandler).AllowMethods(iris.MethodOptions)
	{
		announcements := v1.Party("/announcements")
		{
			anouncementController := controllers.NewAnouncementController(r.Config.Database.DB, anouncementService)
			announcements.Get("/", anouncementController.GetIndexHandler)
			announcements.Post("/", anouncementController.CreateAnnouncementHandler)
			announcements.Get("/{id:uint}/delete", anouncementController.DeleteAnnouncementHandler)
		}
	}
}
