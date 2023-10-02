package routes

import "github.com/gin-gonic/gin"

func RouterRegister(engine *gin.Engine) {
	LoginRouterRegistry(engine)
	GeneratorRouterRegistry(engine)
	UserRouterRegistry(engine)
	DeptRouterRegistry(engine)
	MenuRouterRegistry(engine)
	RoleRouterRegistry(engine)
	BalanceRouterRegistry(engine)
	MessageRouterRegistry(engine)
	RechargeRouterRegistry(engine)
	SensitiveWordRouterRegistry(engine)
	SessionRouterRegistry(engine)
	TokenConsumptionRouterRegistry(engine)
	WebUserRouterRegistry(engine)
	ConversationRouterRegistry(engine)
}
