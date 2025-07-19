package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golodash/galidator"
	"net/http"
	"strings"
)

var Currencies = map[string]string{
	"USD": "USD",
	"NGN": "NGN",
	"IDR": "IDR",
}

func IsValidCurrency(currency string) bool {
	if _, ok := Currencies[currency]; ok {
		return true
	}
	return false
}

func GetActiveUser(c *gin.Context) (int64, error) {
	value, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to access resource"})
		return 0, fmt.Errorf("Error occured")
	}

	userId, ok := value.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Encountered an issue"})
		return 0, fmt.Errorf("Error occured")
	}
	return userId, nil
}

func HandlerError(err error, c *gin.Context, gValid galidator.Validator) interface{} {
	var ve validator.ValidationErrors

	if c.Request.ContentLength == 0 {
		return "Provide body"
	}

	if e, ok := err.(*json.UnmarshalTypeError); ok {
		if e.Field == "" {
			return "Provide a json body"
		}
		msg := fmt.Sprintf("Invalid for field: '%s'. Expected a value of type '%s'", e.Field, e.Type)
		return msg
	}
	if errors.As(err, &ve) {
		// Format custom: bisa array atau map
		out := map[string]string{}
		for _, fe := range ve {
			out[fe.Field()] = customErrorMessage(fe)
		}
		fmt.Println(out)
		return out
	}

	return gValid.DecryptErrors(err)
}

func customErrorMessage(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("required %s", field)
	case "email":
		return fmt.Sprintf("email %s", field)
	default:
		return fmt.Sprintf("all %s", field)
	}
}
