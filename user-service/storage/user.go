package storage

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"user/config"
	pb "user/genuser"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	col *mongo.Collection
	cfg *config.Config
}

func NewUserRepo(ctx context.Context, cfg config.Config) (*UserRepo, error) {
	col, err := Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &UserRepo{col: col, cfg: &cfg}, nil
}

func (u *UserRepo) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	uu_id := uuid.New()
	uuidV5 := uuid.NewSHA1(uu_id, []byte(req.UserName))
	var res pb.UserResponse
	res.UserId = uuidV5.String()
	res.UserName = req.UserName
	res.Email = req.Email
	res.Password = req.Password
	_, err := u.col.InsertOne(ctx, res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *UserRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	filter := bson.D{{
		Key: "username", Value: req.UserName,
	}}
	users := []pb.UserRequest{}
	cursor, err := u.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.UserName == req.UserName && user.Password == req.Password {
			token, err := CreateToken(u.cfg.Jwt_Secret_Key, u.cfg.Expiration_Time)
			if err != nil {
				return nil, err
			}

			var res pb.LoginResponse
			res.Token = token

			start := time.Now()
			d, err := strconv.Atoi(u.cfg.Expiration_Time)
			if err != nil {
				return nil, err
			}
			dur := time.Duration(d)
			end := start.Add(dur * time.Hour)
			seconds := end.Sub(start)
			i := int(seconds.Seconds())
			s := strconv.Itoa(i)
			res.ExpiredAt = s + " seconds"

			return &res, nil
		}
	}
	return nil, errors.New("given data did not match with user name and password")
}

func CreateToken(secretKey string, expiration string) (string, error) {
	hours, err := strconv.Atoi(expiration)
	if err != nil {
		return "", fmt.Errorf("invalid expiration time format: %v", err)
	}

	exp := time.Now().Add(time.Duration(hours) * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp.Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
