package base

import (
	"github.com/gostarer/domain/infra"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	validator9 "gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator9.Validate
var translator ut.Translator

func Validate() *validator9.Validate {
	Check(validate)
	return validate
}

func Transtate() ut.Translator {
	Check(translator)
	return translator
}

type ValidatorStarter struct {
	infra.BaseGoStarer
}

func (v *ValidatorStarter) Init(ctx infra.GoStarerContext) {
	validate = validator9.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("Not found translator: zh")
	}

}

func ValidateStruct(s interface{}) (err error) {
	err = Validate().Struct(s)
	if err != nil {
		_, ok := err.(*validator9.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(validator9.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(Transtate()))
			}
		}
		return err
	}
	return nil
}
