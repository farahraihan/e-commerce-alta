package factory

import (
	"TokoGadget/configs"
	dt_hnd "TokoGadget/internal/features/detail_transactions/handler"
	dt_qry "TokoGadget/internal/features/detail_transactions/repository"
	dt_srv "TokoGadget/internal/features/detail_transactions/services"
	t_hnd "TokoGadget/internal/features/transactions/handler"
	t_qry "TokoGadget/internal/features/transactions/repository"
	t_srv "TokoGadget/internal/features/transactions/services"
	u_hnd "TokoGadget/internal/features/users/handler"
	u_qry "TokoGadget/internal/features/users/repository"
	u_srv "TokoGadget/internal/features/users/services"

	// p_hnd "TokoGadget/internal/features/products/handler"
	p_qry "TokoGadget/internal/features/products/repository"
	// p_srv "TokoGadget/internal/features/products/services"
	"TokoGadget/internal/routes"
	"TokoGadget/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&u_qry.User{}, &p_qry.Product{}, &t_qry.Transaction{}, &dt_qry.DetailTransaction{})

	mu := utils.NewMidtransPayment("SB-Mid-server-cYM2or6TUkO8UHAjMzaWc7Zx")
	pu := utils.NewPasswordUtility()
	tu := utils.NewTokenUtility()

	um := u_qry.NewUserModel(db)
	us := u_srv.NewUserService(um, pu, tu)
	uc := u_hnd.NewUserController(us)

	// pm := u_qry.NewProductModel(db)
	// ps := u_srv.NewProductService(pm)
	// pc := u_hnd.NewProdyctController(ps)

	tq := t_qry.NewTransactionQuery(db)
	ts := t_srv.NewTransactionServices(tq, mu)
	th := t_hnd.NewTransactionHandler(ts, tu)

	dtq := dt_qry.NewDetailTransactionQuery(db)
	dts := dt_srv.NewDetailTransactionServices(dtq, tq)
	dth := dt_hnd.NewDetailTransactionHandler(dts, tu)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// Register

	// t.GET("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())

	routes.InitRoute(e, uc, th, dth)
}
