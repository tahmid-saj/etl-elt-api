package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// expenses
	server.GET("/expenses/expenses")
	server.GET("/expenses/expenses/:userid/:email")
	server.POST("/expenses/expenses")
	server.PUT("/expenses/expenses/:userid/:email")
	server.DELETE("/expenses/expenses/:userid/:email")

	server.GET("/expenses/expenses-summary")
	server.GET("/expenses/expenses-summary/:userid/:email")
	server.POST("/expenses/expenses-summary")
	server.PUT("/expenses/expenses-summary/:userid/:email")
	server.DELETE("/expenses/expenses-summary/:userid/:email")

	// banking
	server.GET("/banking/banking-accounts", getAllBankingAccounts)
	server.GET("/banking/banking-accounts/:userid/:email", getBankingAccountsByUser)
	server.POST("/banking/banking-accounts", createBankingAccount)
	server.PUT("/banking/banking-accounts/:userid/:email")
	server.DELETE("/banking/banking-accounts/:userid/:email")

	server.GET("/banking/banking-summary")
	server.GET("/banking/banking-summary/:userid/:email")
	server.POST("/banking/banking-summary")
	server.PUT("/banking/banking-summary/:userid/:email")
	server.DELETE("/banking/banking-summary/:userid/:email")

	// investments
	server.GET("/investments/investments")
	server.GET("/investments/investments/:userid/:email")
	server.POST("/investments/investments")
	server.PUT("/investments/investments/:userid/:email")
	server.DELETE("/investments/investments/:userid/:email")

	server.GET("/investments/investments-summary")
	server.GET("/investments/investments-summary/:userid/:email")
	server.POST("/investments/investments-summary")
	server.PUT("/investments/investments-summary/:userid/:email")
	server.DELETE("/investments/investments-summary/:userid/:email")

	// savings
	server.GET("/savings/savings-accounts")
	server.GET("/savings/savings-accounts/:userid/:email")
	server.POST("/savings/savings-accounts")
	server.PUT("/savings/savings-accounts/:userid/:email")
	server.DELETE("/savings/savings-accounts/:userid/:email")

	server.GET("/savings/savings-summary")
	server.GET("/savings/savings-summary/:userid/:email")
	server.POST("/savings/savings-summary")
	server.PUT("/savings/savings-summary/:userid/:email")
	server.DELETE("/savings/savings-summary/:userid/:email")

	// nutrition tracked days
	server.GET("/nutrition-tracked-days/nutrition-tracked-days")
	server.GET("/nutrition-tracked-days/nutrition-tracked-days/:userid/:email")
	server.POST("/nutrition-tracked-days/nutrition-tracked-days")
	server.PUT("/nutrition-tracked-days/nutrition-tracked-days/:userid/:email")
	server.DELETE("/nutrition-tracked-days/nutrition-tracked-days/:userid/:email")

	server.GET("/nutrition-tracked-days/nutrition-tracked-days-summary")
	server.GET("/nutrition-tracked-days/nutrition-tracked-days-summary/:userid/:email")
	server.POST("/nutrition-tracked-days/nutrition-tracked-days-summary")
	server.PUT("/nutrition-tracked-days/nutrition-tracked-days-summary/:userid/:email")
	server.DELETE("/nutrition-tracked-days/nutrition-tracked-days-summary/:userid/:email")

	// tracked calories burned
	server.GET("/tracked-calories-burned/tracked-calories-burned")
	server.GET("/tracked-calories-burned/tracked-calories-burned/:userid/:email")
	server.POST("/tracked-calories-burned/tracked-calories-burned")
	server.PUT("/tracked-calories-burned/tracked-calories-burned/:userid/:email")
	server.DELETE("/tracked-calories-burned/tracked-calories-burned/:userid/:email")

	// btc_forecast.daily prediction
	server.GET("/btc-forecast/daily-prediction")
	server.GET("/btc-forecast/daily-prediction/:currentDate")
	server.POST("/btc-forecast/daily-prediction")
	server.PUT("/btc-forecast/daily-prediction/:currentDate")
	server.DELETE("/btc-forecast/daily-prediction/:currentDate")

	// btc_forecast.two_week_prediction
	server.GET("/btc-forecast/two-week-prediction")
	server.GET("/btc-forecast/two-week-prediction/:currentDate")
	server.POST("/btc-forecast/two-week-prediction")
	server.PUT("/btc-forecast/two-week-prediction/:currentDate")
	server.DELETE("/btc-forecast/two-week-prediction/:currentDate")
}