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
	var bankingAccount models.BankingAccount
	err := context.ShouldBindJSON(&bankingAccount)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	bankingAccounts, err := models.GetBankingAccountsByUser(bankingAccount.UserId, bankingAccount.Email)
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
func getAllBankingSummary(context *gin.Context) {
	bankingAccountsSummary, err := models.GetAllBankingSummary()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch banking accounts"})
	}

	context.JSON(http.StatusOK, bankingAccountsSummary)
}

func getBankingSummaryByUser(context *gin.Context) {
	userId := context.Param("userid")
	email := context.Param("email")

	bankingSummary, err := models.GetBankingSummaryByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch banking summary"})
		return
	}

	context.JSON(http.StatusOK, bankingSummary)
}

func createBankingSummary(context *gin.Context) {
	var bankingSummary models.BankingSummary
	err := context.ShouldBindJSON(&bankingSummary)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = bankingSummary.SaveBankingSummary()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create banking summary"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Banking summary created", "bankingSummary": bankingSummary})
}

func updateBankingSummary(context *gin.Context) {
	userId := context.Params.ByName("userid")
	email := context.Params.ByName("email")

	_, err := models.GetBankingSummaryByUser(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch banking summary"})
		return
	}

	var bankingSummary models.BankingSummary
	err = context.ShouldBindJSON(&bankingSummary)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	bankingSummary, err = bankingSummary.UpdateBankingSummary(userId, email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update banking summary"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated banking summary", "bankingSummary": bankingSummary})
}

func deleteBankingSummary(context *gin.Context) {
	var bankingSummary models.BankingSummary
	err := context.ShouldBindJSON(&bankingSummary)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	bankingSummary, err = models.GetBankingSummaryByUser(bankingSummary.UserId, bankingSummary.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch banking summary"})
		return
	}

	deletedBankingSummary, err := bankingSummary.DeleteBankingSummary()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete banking summary"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Banking summary successfully deleted", "bankingSummary": deletedBankingSummary})
}