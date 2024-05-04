package otpmanager

type OtpManager interface {
	CreateOtp(string, string)
	VerifyOtp(string, string) (bool, error)
	DeleteOtp(string)
}
