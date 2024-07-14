package helpers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/o1egl/paseto/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

const External = 7

func IsHasAccessTo(accessGroup []int, access int) bool {
	for _, a := range accessGroup {
		if a == access {
			return true
		}
	}
	return false
}

func CheckForAPIKey(toTestKey string) bool {
	if toTestKey == os.Getenv("API_KEY_FOR_MOBILE") || toTestKey == os.Getenv("API_KEY_FOR_WEB") {
		return true
	}
	return false
}

func Validate(s interface{}) error {

	validate := validator.New()

	err := validate.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		myErr := ""
		for _, err := range err.(validator.ValidationErrors) {
			myErr += " " + err.Field()
		}

		fmt.Println(err)
		return fmt.Errorf("Error in validating :" + myErr)
	}

	return nil
}

func ValidateToken(ctx context.Context, tokenString, key string, access int) (*paseto.JSONToken, error) {
	var receivedToken paseto.JSONToken
	var newFooter string

	fmt.Println("Decrypting Token")
	err := paseto.Decrypt(tokenString, []byte(key), &receivedToken, &newFooter)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Token %v", err)
	}

	fmt.Println("Validating Token")
	err = receivedToken.Validate(
		paseto.ValidAt(time.Now()),
		paseto.IssuedBy(os.Getenv("TOKEN_ISSUER")),
	)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid Token %v", err)
	}

	fmt.Println("Checking Access Group For Token")
	var accessGroup []int
	err = receivedToken.Get("accessGroup", &accessGroup)
	fmt.Println("Access Group  : ", accessGroup)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Token Is Not Valid %v", err)
	}

	if !IsHasAccessTo(accessGroup, access) {
		return nil, status.Errorf(codes.PermissionDenied, "Token Does Not Have Valid Permission To Access Resources %v", err)
	}

	fmt.Println("Checking Refresh Token ID Token")
	_, err = uuid.Parse(receivedToken.Jti)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument ", err)
	}

	fmt.Println("Checking User UUID Of Token")
	_, err = uuid.Parse(receivedToken.Audience)
	if err != nil {
		fmt.Println("Checking User UUID Of Token : ", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument ", err)
	}

	fmt.Println("Done With Token")
	return &receivedToken, nil
}

func ContextError(ctx context.Context) error {

	switch ctx.Err() {
	case context.Canceled:
		log.Println("Request Canceled")
		return status.Error(codes.DeadlineExceeded, "Request Canceled")
	case context.DeadlineExceeded:
		log.Println("DeadLine Exceeded")
		return status.Error(codes.DeadlineExceeded, "DeadLine Exceeded")
	default:
		return nil
	}
}

const AuthTokenExpiry = time.Minute * 5
const AuthTokenExpiryForUnAuthorized = time.Minute * 2

func GenerateAuthToken(userId string, userName string, refreshTokenId string, authorized bool, accessGroup []int) (string, error) {

	now := time.Now()
	var exp time.Time
	if authorized {
		exp = now.Add(AuthTokenExpiry)
	} else {
		exp = now.Add(AuthTokenExpiryForUnAuthorized)
	}
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   userId,
		Jti:        refreshTokenId,
		Subject:    userName,
		Issuer:     os.Getenv("TOKEN_ISSUER"),
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	jsonToken.Set("authorized", authorized)
	jsonToken.Set("accessGroup", accessGroup)
	footer := "Powered By AapanaVypar"

	// Encrypt data
	token, err := paseto.Encrypt([]byte(os.Getenv("AUTH_TOKEN_SECRETE")), jsonToken, footer)

	if err != nil {
		return "", err
	}
	return token, nil
}
