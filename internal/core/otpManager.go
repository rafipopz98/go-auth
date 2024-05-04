package core

import (
	"errors"
	"sync"
)

type OTPManager struct {
	mu               sync.RWMutex
	PhoneNumberToOtp map[string]string
}

func NewOTPManager() *OTPManager {
	return &OTPManager{
		PhoneNumberToOtp: make(map[string]string),
	}
}

func (otpManager *OTPManager) CreateOtp(phoneNumber, otp string) {
	otpManager.mu.Lock()
	defer otpManager.mu.Unlock()
	otpManager.PhoneNumberToOtp[phoneNumber] = otp
}

func (otpManager *OTPManager) VerifyOtp(phoneNumber, otp string) (bool, error) {
	storedOtp, ok := otpManager.PhoneNumberToOtp[phoneNumber]
	if !ok {
		return false, errors.New("phone number not found")
	}
	if storedOtp == otp {
		return true, nil
	}
	return false, errors.New("invalid otp")
}

func (otpManager *OTPManager) DeleteOtp(phoneNumber string) {
	otpManager.mu.Lock()
	defer otpManager.mu.Unlock()
	delete(otpManager.PhoneNumberToOtp, phoneNumber)
}
