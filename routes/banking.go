package routes

import (
	"net/http"
	"tahmid-saj/etl-elt-api/models"
	"github.com/gin-gonic/gin"
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

func deleteBankingAccount(context *gin.Context) {
	userId := context.Param("userid")
	email := context.Param("email")
	
	var bankingAccount models.BankingAccount
	err := context.ShouldBindJSON(&bankingAccount)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	bankingAccounts, err := models.GetBankingAccountsByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch banking accounts"})
		return
	}

	if len(bankingAccounts) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No banking account exists"})
		return
	}

	deletedBankingAccount, err := bankingAccount.DeleteBankingAccount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete banking account"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Banking account successfully deleted", "bankingAccount": deletedBankingAccount})
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