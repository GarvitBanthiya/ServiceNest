package service

import (
	"errors"
	"fmt"
	"serviceNest/errs"
	"serviceNest/interfaces"
	"serviceNest/model"
	"serviceNest/repository"
	"serviceNest/util"
)

type UserService struct {
	otpRepo  *repository.OtpRepository
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository, otpRepo *repository.OtpRepository) interfaces.UserService {
	return &UserService{userRepo: userRepo,
		otpRepo: otpRepo}
}

// View User

func (s *UserService) ViewProfileByID(userID string) (*model.User, error) {
	//if err := s.userRepo.EnsureConnection(); err != nil {
	//	return nil, err
	//}

	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("could not find user: %v", err)
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *UserService) UpdateUser(userID string, newEmail, newPassword, newAddress, newPhone *string) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("could not find user: %v", err)
	}

	// Update email
	if newEmail != nil {
		if err := util.ValidateEmail(*newEmail); err != nil {
			return err
		}
		existingUser, err := s.userRepo.GetUserByEmail(*newEmail)
		if err == nil && existingUser.ID != userID {
			return errors.New("email already in use by another user")
		}
		user.Email = *newEmail
	}

	// Update password
	if newPassword != nil {
		if err := util.ValidatePassword(*newPassword); err != nil {
			return err
		}
		user.Password = *newPassword
	}

	// Update contact
	if newPhone != nil {
		if err := util.ValidatePhoneNumber(*newPhone); err != nil {
			return err
		}
		user.Contact = *newPhone
	}
	// Update address
	if newAddress != nil {
		user.Address = *newAddress
	}

	// Save the updated user back to the repository_test
	if err := s.userRepo.UpdateUser(user); err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}

	return nil
}

func (s *UserService) CheckUserExists(email string) (*model.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("could not find user: %v", err)
	}
	return user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	user.ID = util.GenerateUUID()
	err := s.userRepo.SaveUser(user)
	if err != nil {
		return fmt.Errorf("could not save user: %v", err)
	}
	return nil
}

func (s *UserService) ForgetPasword(email string, answer string, updatedPassword string) error {
	securityAnswer, err := s.userRepo.GetSecurityAnswerByEmail(email)
	if err != nil {
		return err
	}
	if *securityAnswer != answer {
		return fmt.Errorf(errs.IncorrectSecurityAnswer)
	}

	err = s.userRepo.UpdatePassword(email, updatedPassword)
	if err != nil {
		return err
	}

	return nil

}

func (s *UserService) GenerateOtp(email string) error {
	_, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	otp, err := s.otpRepo.GenerateOTP()
	if err != nil {
		return err
	}
	s.otpRepo.SaveOTP(email, otp)

	return util.SendOTPEmail(email, otp)
}

func (s *UserService) VerifyAndUpdatePassword(email, otp string, password string) error {
	valid := s.otpRepo.ValidateOTP(email, otp)
	if valid == false {
		return errors.New("Invalid Otp")
	}
	err := s.userRepo.UpdatePassword(email, password)
	if err != nil {
		return err
	}

	return nil
}
