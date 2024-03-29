package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string

	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)

	fmt.Printf("%v c.ShouldBind \n", err)
	if err != nil {
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for _, verr := range verrs {
			errs = append(errs, &ValidError{
				Key:     verr.Field(),
				Message: verr.Error(),
			})
		}

		return false, errs
	}

	return true, nil
}

func BindAndValidUri(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBindUri(v)

	if err != nil {
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for _, verr := range verrs {
			errs = append(errs, &ValidError{
				Key:     verr.Field(),
				Message: verr.Error(),
			})
		}

		return false, errs
	}

	return true, nil
}
