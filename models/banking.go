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

// banking accounts
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

	filter := bson.D{
		{Key: "userId", Value: userId}, 
		{Key: "email", Value: email},
	}

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

func (bankingAccount *BankingAccount) SaveBankingAccount() error {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_ACCOUNTS_COLLECTION"))
	newBankingAccount := BankingAccount{
		UserId: bankingAccount.UserId,
		Email: bankingAccount.Email,
		Name: bankingAccount.Name,
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

func (bankingAccount *BankingAccount) DeleteBankingAccount() (BankingAccount, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_ACCOUNTS_COLLECTION"))

	filter := bson.D{
		{Key: "userId", Value: bankingAccount.UserId}, 
		{Key: "email", Value: bankingAccount.Email},
		{Key: "name", Value: bankingAccount.Name},
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return *bankingAccount, nil
}

// banking summary
func GetAllBankingSummary() ([]BankingSummary, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_SUMMARY_COLLECTION"))

	filter := bson.D{{}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var bankingAccountsSummary []BankingSummary
	if err = cursor.All(context.TODO(), &bankingAccountsSummary); err != nil {
		panic(err)
	}

	return bankingAccountsSummary, nil
}

func GetBankingSummaryByUser(userId string, email string) (BankingSummary, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_SUMMARY_COLLECTION"))

	filter := bson.D{
		{Key: "userId", Value: userId},
		{Key: "email", Value: email},
	}
	
	var bankingSummary BankingSummary
	err := collection.FindOne(context.TODO(), filter).Decode(&bankingSummary)
	if err != nil {
		panic(err)
	}

	return bankingSummary, nil
}

func (bankingSummary *BankingSummary) SaveBankingSummary() error {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_SUMMARY_COLLECTION"))
	newBankingSummary := BankingSummary{
		UserId: bankingSummary.UserId,
		Email: bankingSummary.Email,
		CurrentAllBalance: 0.0,
		TotalAllBankingIn: 0.0,
		TotalAllBankingOut: 0.0,
	}

	_, err := collection.InsertOne(context.TODO(), newBankingSummary)
	if err != nil {
		panic(err)
	}

	return nil
}

func (bankingSummary *BankingSummary) UpdateBankingSummary(userId string, email string) (BankingSummary, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_SUMMARY_COLLECTION"))

	filter := bson.D{		
		{Key: "userid", Value: userId},
		{Key: "email", Value: email},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "userid", Value: bankingSummary.UserId},
			{Key: "email", Value: bankingSummary.Email},
			{Key: "currentAllBalance", Value: bankingSummary.CurrentAllBalance},
			{Key: "totalAllBankingIn", Value: bankingSummary.TotalAllBankingIn},
			{Key: "totalAllBankingOut", Value: bankingSummary.TotalAllBankingOut},
		},},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return *bankingSummary, nil
}

func (bankingSummary *BankingSummary) DeleteBankingSummary() (BankingSummary, error) {
	collection := db.MongoDBClient.Database(os.Getenv("MONGO_TEST_DB")).Collection(os.Getenv("MONGO_BANKING_SUMMARY_COLLECTION"))

	filter := bson.D{
		{Key: "userId", Value: bankingSummary.UserId}, 
		{Key: "email", Value: bankingSummary.Email},
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	return *bankingSummary, nil
}