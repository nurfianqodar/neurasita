package global

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTrans "github.com/go-playground/validator/v10/translations/id"
)

var Validate *validator.Validate
var Translator ut.Translator

func init() {
	Validate = validator.New()

	idLoc := id.New()
	enLoc := en.New()

	uni := ut.New(enLoc, idLoc, enLoc)
	Translator, _ = uni.GetTranslator("id")
	idTrans.RegisterDefaultTranslations(Validate, Translator)
}
