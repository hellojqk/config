package service

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/model"
	"github.com/hellojqk/config/server/repository"
	"github.com/hellojqk/config/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//openssl genrsa -out private.key 2048
//openssl rsa -in private.key -pubout -out public.key

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	util.WaitInitFuncsAdd(initAuthKey)
}

func initAuthKey() error {
	privateKeyStr := viper.GetString("auth.privateKey")
	publicKeyStr := viper.GetString("auth.publicKey")
	if privateKeyStr == "" || publicKeyStr == "" {
		return errors.New("auth privateKey or publicKey not config")
	}
	var err error

	privateBts := []byte(privateKeyStr)
	publicBts := []byte(publicKeyStr)

	// fmt.Printf("privateKeyStr:%s\n", privateBts)
	// fmt.Printf("publicKeyStr:%s\n", publicBts)

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBts)
	if err != nil {
		return errors.WithMessage(err, "load privateKey err")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBts)
	if err != nil {
		return errors.WithMessage(err, "load publicKey err")
	}
	return nil
}

func userCollection() *mongo.Collection {
	return repository.DB.Collection("user")
}

// UserInsertOne .
func UserInsertOne(ctx context.Context, user entity.User) (err error) {
	user.SetCreator("")
	insertResult, err := repository.DB.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	if insertResult.InsertedID == nil {
		return errors.New("创建失败")
	}
	return nil
}

// UserLoginParam .
func UserLoginParam(ctx context.Context, userModel model.UserLoginParam) (jwtToken string, err error) {

	user := entity.User{}
	err = userCollection().FindOne(ctx, bson.M{"key": userModel.Key}).Decode(&user)
	if err != nil {
		return
	}

	var encryptPassword string
	encryptPassword, err = util.EncryptPassword(userModel.Password, userModel.Key)
	if err != nil {
		return
	}
	// println("encryptPassword", userModel.Key, userModel.Password, encryptPassword, user.Password)
	if encryptPassword == "" || encryptPassword != user.Password {
		return "", errors.New("incorrect key or password")
	}

	now := time.Now()

	claims := jwt.StandardClaims{
		Issuer:    "config1",
		Subject:   user.Key,
		IssuedAt:  jwt.NewTime(float64(now.Unix())),
		ExpiresAt: jwt.NewTime(float64(now.Add(20 * time.Hour).Unix())),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwtToken, err = token.SignedString(privateKey)
	if err != nil {
		err = errors.WithMessage(err, "sign token error")
		return
	}

	return
}

// UserTokenValid .
func UserTokenValid(ctx context.Context, jwtToken string) (userKey string, err error) {
	claim := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, &claim, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("token valid falied")
		return
	}
	userKey = claim.Subject
	return
}
