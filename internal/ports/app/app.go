package app

type App interface {
	SignupWithPhoneNumber(string)
	AddOtpToPhoneNumber(string)
	VerifyAccount(string)
	Login(string)
	GetProfile(string)
	ReturnMessageQueueChan() *chan string
}
