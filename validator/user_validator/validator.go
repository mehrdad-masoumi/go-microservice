package user_validator

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"mlm/dto"
	"mlm/entity"
	"mlm/pkg/error_msg"
	"mlm/pkg/richerror"
	"regexp"
)

type Repository interface {
	FindNodeByReferral(referral string) (entity.Node, error)
	IsEmailUnique(email string) (bool, error)
}

type Validator struct {
	repo Repository
}

func NewUserValidator(repo Repository) Validator {
	return Validator{repo: repo}
}

func (v Validator) Validate(req dto.UserCreateRequest) (map[string]string, error) {

	const target = "node_validator.ValidateRegisterRequest"

	err := validation.ValidateStruct(&req,

		validation.Field(&req.Email,
			validation.Required,
			is.Email,
		),

		validation.Field(&req.Username,
			validation.Required,
			validation.Length(3, 20),
		),

		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`)),
		),

		validation.Field(&req.ConfirmPassword,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`)),
		),
	)
	if err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors,
			richerror.New(target).
				WithKind(richerror.KindInvalid).
				WithMeta(map[string]interface{}{"req": req})

	}
	return nil, nil
}

func (v Validator) checkEmailUniqueness(value interface{}) error {

	email := value.(string)
	if isUnique, err := v.repo.IsEmailUnique(email); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(error_msg.ErrorMsgEmailIsNotUnique)
		}
	}

	return nil
}
