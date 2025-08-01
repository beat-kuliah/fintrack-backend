package api

import (
	"github.com/go-playground/validator/v10"
	"github/beat-kuliah/fintrack_backend/utils"
)

var currencyValidator validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsValidCurrency(currency)
	}
	return false
}
