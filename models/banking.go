package models

import (
	"context"
	"tahmid-saj/etl-elt-api/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type BankingAccount struct {
	UserId string `binding:"required" bson:"userId"`
	Email string `binding:"required" bson:"email"`
	Name string `binding:"required" bson:"name"`
	CurrentBalance float64 `binding:"required" bson:"currentBalance"`
	TotalIn float64 `binding:"required" bson:"totalIn"`
	TotalOut float64 `binding:"required" bson:"totalOut"`
	Transactions []Transaction `binding:"required" bson:"transactions"`
}

type Transaction struct {
	Amount float64 `binding:"required" bson:"amount"`
	Type string `binding:"required" bson:"type"`
	Reason string `bson:"reason"`
}

type BankingSummary struct {
	UserId string `binding:"required" bson:"userId"`
	Email string `binding:"required" bson:"email"`
	CurrentAllBalance float64 `binding:"required" bson:"currentAllBalance"`
	TotalAllBankingIn float64 `binding:"required" bson:"totalAllBankingIn"`
	TotalAllBankingOut float64 `binding:"required" bson:"totalAllBankingOut"`
}

func GetAllBankingAccounts() ([]BankingAccount, error) {
	collection := db.MongoDBClient.Database("test").Collection("bankingaccounts")

	filter := bson.D{}
	
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