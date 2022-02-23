package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Traking-work/traking-backend.git/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorizationRepo struct {
	db *mongo.Collection
}

func NewAuthorizationRepo(db *mongo.Database) *AuthorizationRepo {
	return &AuthorizationRepo{db: db.Collection(usersCollection)}
}

func (r *AuthorizationRepo) GetUser(ctx context.Context, username, password string) (domain.UserLogin, error) {
	var user domain.UserLogin

	if err := r.db.FindOne(ctx, bson.M{"username": username, "password": password}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserLogin{}, domain.ErrUserNotFound
		}
		return domain.UserLogin{}, err

	}

	return user, nil
}

func (r *AuthorizationRepo) SetSession(ctx context.Context, userID primitive.ObjectID, session domain.Session) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$set": bson.M{
			"session": session,
		}})

	return err
}

func (r *AuthorizationRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.UserLogin, error) {
	var user domain.UserLogin
	if err := r.db.FindOne(ctx, bson.M{
		"session.refreshToken": refreshToken,
		"session.expiresAt":    bson.M{"$gt": time.Now()},
	}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.UserLogin{}, domain.ErrUserNotFound
		}
		return domain.UserLogin{}, err

	}

	return user, nil
}

func (r *AuthorizationRepo) RemoveRefreshToken(ctx context.Context, refreshToken string) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"session.refreshToken": refreshToken}, bson.M{
		"$set": bson.M{
			"session.refreshToken": "",
		}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.ErrUserNotFound
		}
		return err
	}

	return nil
}
