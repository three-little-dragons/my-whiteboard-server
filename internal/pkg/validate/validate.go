package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ens "github.com/go-playground/validator/v10/translations/en"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"github.com/three-little-dragons/my-whiteboard-server/internal/pkg/com"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var translator *ut.UniversalTranslator
var defaultTr ut.Translator
var once sync.Once

func lazyInit() {
	once.Do(func() {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			// change tag name
			v.SetTagName("valid")
			// setup message translations
			registerTranslations(v)
			// custom field name displayed on error
			v.RegisterTagNameFunc(func(fld reflect.StructField) string {
				name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
				if name == "" {
					name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
				}
				// skip if tag key says it should be ignored
				if name == "-" {
					return ""
				}
				return name
			})
		}
	})
}

func registerTranslations(v *validator.Validate) {
	var err error
	var tr ut.Translator

	cn := zh.New()
	translator = ut.New(cn, cn, en.New())

	tr, _ = translator.GetTranslator("zh")
	err = zhs.RegisterDefaultTranslations(v, tr)
	if err != nil {
		panic(err)
	}

	defaultTr = tr

	tr, _ = translator.GetTranslator("en")
	err = ens.RegisterDefaultTranslations(v, tr)
	if err != nil {
		panic(err)
	}
}

func Struct[T any](c *gin.Context, obj *T) *T {
	if c.Request.Method == "GET" {
		return StructQuery(c, obj)
	}
	return StructBody(c, obj)
}

func StructBody[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBind, obj)
}

func StructQuery[T any](c *gin.Context, obj *T) *T {
	return bind(c, c.ShouldBindQuery, obj)
}

type Wrapper struct {
	arr []string
}

func (w *Wrapper) Len() int {
	return len(w.arr)
}
func (w *Wrapper) Less(i, j int) bool {
	return w.arr[i] < w.arr[j]
}

func getLanguagesSortedByPriority(c *gin.Context) []string {
	arr := strings.Split(c.Request.Header.Get("Accept-Language"), ",")
	m := make(map[float64][]string, len(arr))
	for _, item := range arr {
		var idx int

		var lang = item
		idx = strings.LastIndex(lang, ";")
		if idx > 0 {
			lang = item[0:idx]
		}

		var priority = 1.0
		idx = strings.LastIndex(item, "q=")
		if idx > 0 && idx < len(item) {
			if p, err := strconv.ParseFloat(item[idx:], 64); err == nil {
				priority = p
			}
		}

		m[priority] = append(m[priority], lang)
	}

	keys := make([]float64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	result := make([]string, 0, len(m))
	for _, k := range keys {
		for _, l := range m[k] {
			result = append(result, l)
		}
	}

	return result
}

func getTranslator(c *gin.Context) ut.Translator {
	var tr = defaultTr
	for _, locale := range getLanguagesSortedByPriority(c) {
		var found = false
		tr, found = translator.GetTranslator(locale)
		if found {
			break
		}
	}
	return tr
}

// bind automatically send errors when parameter errors occur
func bind[T any](c *gin.Context, bind func(obj interface{}) error, obj *T) *T {
	lazyInit()
	if err := bind(obj); err != nil {
		// if invalid
		validationErrors := err.(validator.ValidationErrors)

		var errors []string
		for _, e := range validationErrors {
			errors = append(errors, e.Translate(getTranslator(c)))
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, com.Response{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  strings.Join(errors, "; "),
		})
		return nil
	}
	// if valid
	return obj
}
