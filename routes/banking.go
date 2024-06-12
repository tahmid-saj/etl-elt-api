package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"tahmid-saj/etl-elt-api/models"
)

// banking accounts
func getBankingAccounts(context *gin.Context) {
	bankingAccounts, err := models.GetAllBankingAccounts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch banking accounts"})
		return
	}

	context.JSON(http.StatusOK, bankingAccounts)
}

func getBankingAccount(context *gin.Context) {

}

func createBankingAccount(context *gin.Context) {

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