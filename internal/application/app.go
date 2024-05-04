package application

import otpmanager "auth-service/internal/ports/otpManager"

type ApplicationStruct struct {
	otpManager       otpmanager.OtpManager
	MessageQueueChan chan string
}

// constructor to the struct
func NewApplication(otpManager otpmanager.OtpManager) *ApplicationStruct {
	return &ApplicationStruct{
		otpManager:       otpManager,
		MessageQueueChan: make(chan string),
	}
}

func (app *ApplicationStruct) SignupWithPhoneNumber(number string) {
	// contacts otp-service to generate otp to the number provided using MQ/gRPC

	//check if user already exists
	// user exists -> true snd failure message -> User already exists
	//else contact otp-service to generate otp to the number provided using MQ/gRPC

	// if method gRPC => call OTP injection here
	// else terminate
}

/*
 * AddOtpToPhoneNumber : gets otp from OTP service
 *
 */
func (app *ApplicationStruct) AddOtpToPhoneNumber(otp string) {
	//
}

func (app *ApplicationStruct) VerifyAccount(number string) {
	// contacts otp-service to generate otp to the number provided using MQ/gRPC

}

func (app *ApplicationStruct) Login(number string) {
	// contacts otp-service to generate otp to the number provided using MQ/gRPC

}

func (app *ApplicationStruct) GetProfile(number string) {
	// contacts otp-service to generate otp to the number provided using MQ/gRPC

}

func (app *ApplicationStruct) ReturnMessageQueueChan() *chan string {
	return &app.MessageQueueChan
}
