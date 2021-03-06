package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

func Setup(){
	binding.Validator = new(defaultValidator)
}

type defaultValidator struct {
	once sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &defaultValidator{}

func(v *defaultValidator)ValidateStruct(obj interface{})error{
	if kindOfDate(obj) == reflect.Struct{
		v.lazyinit()
		if err := v.validate.Struct(obj);err!=nil{
			return error(err)
		}
	}
	return nil
}

func (v *defaultValidator)Engine() interface{}  {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator)lazyinit()  {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}

func kindOfDate(data interface{}) reflect.Kind{
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
