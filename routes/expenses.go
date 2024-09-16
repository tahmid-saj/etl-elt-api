package routes

import (
	"net/http"
	"tahmid-saj/etl-elt-api/models"
	"github.com/gin-gonic/gin"
)

// expenses
func getExpenses(context *gin.Context) {
	expenses, err := models.GetAllExpenses()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch expenses"})
		return
	}

	context.JSON(http.StatusOK, expenses)
}

func getExpensesByUser(context *gin.Context) {
	userId := context.Param("userid")
	email := context.Param("email")

	expenses, err := models.GetExpensesByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch expenses"})
	}

	context.JSON(http.StatusOK, expenses)
}

func createExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = expense.SaveExpense()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create expense"})
	}
}

func updateExpense(context *gin.Context) {
	userId := context.Params.ByName("userid")
	email := context.Params.ByName("email")

	_, err := models.GetExpensesByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch expenses"})
		return
	}

	var expense models.Expense
	err = context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated expense", "expense": expense})
}

func deleteExpense(context *gin.Context) {
	var expense models.Expense
	err := context.ShouldBindJSON(&expense)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	expenses, err := models.GetExpensesByUser(expense.UserId, expense.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch expenses"})
		return
	}

	if len(expenses) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No expense exists"})
		return
	}

	deletedExpense, err := expense.deleteExpense()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete expense"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Expense successfully deleted", "expense": expense})
}

// expenses summary
func getExpensesSummary(context *gin.Context) {

}

func getExpenseSummary(context *gin.Context) {

}

func createExpensesSummary(context *gin.Context) {

}

func updateExpensesSummary(context *gin.Context) {

}

func deleteExpensesSummary(context *gin.Context) {

}