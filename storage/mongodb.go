package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"jacq/model"
	"log"
	"time"
)

// mongoDB collections
const (
	userCollection = "users"
)

var _ DataStore = &MongoStore{}

type MongoStore struct {
	mongoClient  *mongo.Client
	databaseName string
}

func New(connectURI, databaseName string) (DataStore, *mongo.Client, error) {
	//Connecting to MongoDB...
	log.Println("Connecting to MongoDB...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectURI))
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, nil, err
	}
	//Connected to MongoDB successfully ..
	log.Println("Connected to mongoDB Successfully")
	return &MongoStore{mongoClient: client, databaseName: databaseName}, client, nil
}

func (repo *MongoStore) CreateUser(data *model.User) (*model.User, error) {
	_, err := repo.col(userCollection).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *MongoStore) UpdateUser(ID string, data *model.User) (*model.User, error) {
	oldUser, err := repo.GetUserByID(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": ID}
	_, err = repo.col(userCollection).ReplaceOne(context.TODO(), filter, repo.buildUserPayload(data, oldUser))
	if err != nil {
		return nil, err
	}
	return repo.GetUserByID(ID)
}

func (repo *MongoStore) GetUserByEmail(email string) (*model.User, error) {
	filter := bson.M{"email": email}
	var user *model.User
	if err := repo.col(userCollection).FindOne(context.Background(), filter, nil).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MongoStore) VerifyNumber(data *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (repo *MongoStore) VerifyIdentity(data *model.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *MongoStore) SecureTransaction(data *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (repo *MongoStore) GetUserByID(ID string) (*model.User, error) {
	filter := bson.M{"id": ID}
	var user = &model.User{}
	err := repo.col(userCollection).FindOne(context.Background(), filter, nil).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *MongoStore) col(collectionName string) *mongo.Collection {
	return repo.mongoClient.Database(repo.databaseName).Collection(collectionName)
}

func (repo *MongoStore) buildUserPayload(newUser, oldUser *model.User) *model.User {
	if newUser == nil {
		return oldUser
	}
	if newUser.FirstName != "" {
		oldUser.FirstName = newUser.FirstName
	}
	if newUser.LastName != "" {
		oldUser.LastName = newUser.LastName
	}
	if newUser.DateOfBirth != "" {
		oldUser.DateOfBirth = newUser.DateOfBirth
	}
	if newUser.Gender != "" {
		oldUser.Gender = newUser.Gender
	}
	if newUser.Country != "" {
		oldUser.Country = newUser.Country
	}

	if newUser.State != "" {
		oldUser.State = newUser.State
	}

	if newUser.HomeAddress != "" {
		oldUser.HomeAddress = newUser.HomeAddress
	}

	if newUser.TransactionCode != 0 {
		oldUser.TransactionCode = newUser.TransactionCode
	}

	if newUser.Password != "" {
		oldUser.Password = newUser.Password
	}

	if newUser.ConfirmPassword != "" {
		oldUser.ConfirmPassword = newUser.ConfirmPassword

	}

	if newUser.DateOfBirth != "" {
		oldUser.DateOfBirth = newUser.DateOfBirth
	}

	if newUser.IDType != "" {
		oldUser.IDType = newUser.IDType
	}

	if newUser.Document != "" {
		oldUser.Document = newUser.Document
	}

	return oldUser
}
