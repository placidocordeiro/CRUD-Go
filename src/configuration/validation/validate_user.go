package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/placidocordeiro/CRUD-Go/src/configuration/errs"
)

var (
	Validate = validator.New();
	transl ut.Translator;
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New();
		unt := ut.New(en, en);
		transl, _ = unt.GetTranslator("en");
		en_translation.RegisterDefaultTranslations(val, transl);
	}
}

func ValidateUserError(validation_err error) *errs.Errs {
	var jsonError *json.UnmarshalTypeError;
	var jsonValidationErr validator.ValidationErrors;

	if errors.As(validation_err, &jsonError) {
		return errs.NewBadRequestError("Invalid field type");
	} else if errors.As(validation_err, &jsonValidationErr) {
		errorCauses := []errs.Causes{};

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := errs.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			};
			errorCauses = append(errorCauses, cause);
		}
		return errs.NewBadRequestValidationError("Invalid fields", errorCauses);
	} 
	return errs.NewBadRequestError("Error trying to convert fields");
}
