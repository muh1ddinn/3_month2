package service

import (
	models "cars_with_sql/api/model"
	"cars_with_sql/config"
	"cars_with_sql/pkg"
	"cars_with_sql/pkg/jwt"
	"cars_with_sql/pkg/logger"
	"cars_with_sql/pkg/password"
	"cars_with_sql/pkg/smtp"
	"cars_with_sql/storage"
	"context"
	"fmt"
)

type authService struct {
	storage storage.IStorage
	log     logger.ILogger
	redis   storage.IREdisStorage
}

func NewAuthService(storage storage.IStorage, log logger.ILogger) authService {
	return authService{
		storage: storage,
		log:     log,
	}
}

func (a authService) CustomerLogin(ctx context.Context, loginRequest models.CustomerLoginRequest) (models.CustomerLoginResponse, error) {
	fmt.Println(" loginRequest.Login: ", loginRequest.Login)
	customer, err := a.storage.Customer().GetByLogin(ctx, loginRequest.Login)
	if err != nil {
		a.log.Error("error while getting customer credentials by login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	if err = password.CompareHashAndPassword(customer.Password, loginRequest.Password); err != nil {
		a.log.Error("error while comparing password", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	m := make(map[interface{}]interface{})

	m["user_id"] = customer.Id
	m["user_role"] = config.CUSTOMER_ROLE

	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		a.log.Error("error while generating tokens for customer login", logger.Error(err))
		return models.CustomerLoginResponse{}, err
	}

	return models.CustomerLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a authService) CustomerRegister(ctx context.Context, logRequest models.CustomerRegisterRequest) error {
	fmt.Println("loginRequest.Maillllllll:", logRequest.Mail)

	// Check if the customer already exists in the database
	customer, err := a.storage.Customer().Checklogin(ctx, logRequest.Mail)
	if err != nil {
		// Handle the error, e.g., customer not found or database error
		fmt.Println("Error while checking login:", err)
		return err
	}

	// Check if the email exists in the database
	if customer.Mail != logRequest.Mail {
		fmt.Println("You have already registered. Please login.")
		// Handle the case where the customer already exists
		return fmt.Errorf("customer with email %s already exists", logRequest.Mail)
	} else {
		// Generate OTP and send it to the user
		otpCode := pkg.GenerateOTP()
		msg := fmt.Sprintf("Your otp code is: %v, for registering RENT_CAR. Don't give it to anyone", otpCode)

		// err := a.redis.SetX(ctx, logRequest.Mail, otpCode, time.Minute*2)
		// if err != nil {
		// 	a.log.Error("Error while setting OTP code in Redis for customer registration: ", logger.Error(err))
		// 	return err
		// }

		err = smtp.Sendmail(logRequest.Mail, msg)
		if err != nil {
			a.log.Error("Error while sending OTP code to customer for registration: ", logger.Error(err))
			return err
		}
	}

	return nil
}
