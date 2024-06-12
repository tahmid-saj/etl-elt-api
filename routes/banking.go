package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tahmid-saj/etl-elt-api/models"
)

// banking accounts
func getAllBankingAccounts(context *gin.Context) {
	bankingAccounts, err := models.GetAllBankingAccounts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch banking accounts"})
		return
	}

	context.JSON(http.StatusOK, bankingAccounts)
}

func getBankingAccountsByUser(context *gin.Context) {
	userId := context.Param("userid")
	email := context.Param("email")

	bankingAccounts, err := models.GetBankingAccountsByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch banking accounts"})
		return
	}

	context.JSON(http.StatusOK, bankingAccounts)
}

func createBankingAccount(context *gin.Context) {
	var bankingAccount models.BankingAccount
	err := context.ShouldBindJSON(&bankingAccount)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = bankingAccount.SaveBankingAccount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create banking account"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Banking account created", "bankingAccount": bankingAccount})
}

func updateBankingAccount(context *gin.Context) {

}

func deleteBankingAccount(context *gin.Context) {

}

// banking accounts summary
func getBankingAccountsSummary(context *gin.Context) {

}

func getBankingAccountSummary(context *gin.Context) {

}

func createBankingAccountSummary(context *gin.Context) {

}

func updateBankingAccountSummary(context *gin.Context) {

}

func deleteBankingAccountSummary(context *gin.Context) {

}