package models

import (
	"context"
	"os"
	"tahmid-saj/etl-elt-api/db/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

type BankingAccount struct {
	UserId string `json:"userId" binding:"required" bson:"userId"`
	Email string `json:"email" binding:"required" bson:"email"`
	Name string `json:"name" binding:"required" bson:"name"`
	CurrentBalance float64 `json:"currentBalance" bson:"currentBalance"`
	TotalIn float64 `json:"totalIn" bson:"totalIn"`
	TotalOut float64 `json:"totalOut" bson:"totalOut"`
	Transactions []Transaction `json:"transactions" bson:"transactions"`
}

type Transaction struct {
	Amount float64 `json:"amount" bson:"amount"`
	Type string `json:"type" bson:"type"`
	Reason string `json:"reason" bson:"reason"`
}

type BankingSummary struct {
	UserId string `json:"userId" binding:"required" bson:"userId"`
	Email string `json:"email" binding:"required" bson:"email"`
	CurrentAllBalance float64 `json:"currentAllBalance" bson:"currentAllBalance"`
	TotalAllBankingIn float64 `json:"totalAllBankingIn" bson:"totalAllBankingIn"`
	TotalAllBankingOut float64 `json:"totalAllBankingOut" bson:"totalAllBankingOut"`
}

func GetAllBankingAccounts() ([]BankingAccount, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_ACCOUNTS_COLLECTION"))

	filter := bson.D{{}}
	
	// Retrieves documents that match the query filter
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	
	// Unpacks the cursor into a slice
	var bankingAccounts []BankingAccount
	if err = cursor.All(context.TODO(), &bankingAccounts); err != nil {
		panic(err)
	}

	return bankingAccounts, nil
}

func GetBankingAccountsByUser(userId string, email string) ([]BankingAccount, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_ACCOUNTS_COLLECTION"))

	filter := bson.D{{"userId", userId}, {"email", email}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var bankingAccounts []BankingAccount
	if err = cursor.All(context.TODO(), &bankingAccounts); err != nil {
		panic(err)
	}

	return bankingAccounts, nil
}

func (ba *BankingAccount) SaveBankingAccount() error {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_ACCOUNTS_COLLECTION"))
	newBankingAccount := BankingAccount{
		UserId: ba.UserId,
		Email: ba.Email,
		Name: ba.Name,
		CurrentBalance: 0.0,
		TotalIn: 0.0,
		TotalOut: 0.0,
		Transactions: []Transaction{},
	}

	_, err := collection.InsertOne(context.TODO(), newBankingAccount)
	if err != nil {
		panic(err)
	}

	return nil
}