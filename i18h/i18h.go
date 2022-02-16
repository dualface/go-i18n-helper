package i18h

import (
	"fmt"
	"strconv"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type (
	TranslatorFunc = func(msgId string, a ...interface{}) string
)

var (
	localizes = make(map[language.Tag]*i18n.Localizer)
	functions = make(map[language.Tag]TranslatorFunc)
)

// Load message file for specific languages
func Load(langs map[string]string, format string, unmarshalFunc i18n.UnmarshalFunc) error {
	for lang, filename := range langs {
		tag, err := language.All.Parse(lang)
		if err != nil {
			return err
		}

		b := i18n.NewBundle(tag)
		b.RegisterUnmarshalFunc(format, unmarshalFunc)
		_, err = b.LoadMessageFile(filename)
		if err != nil {
			return err
		}
		localizes[tag] = i18n.NewLocalizer(b, tag.String())
	}

	return nil
}

// T translate message to specific language
func T(lang string, msgId string, a ...interface{}) string {
	return Lang(lang)(msgId, a...)
}

// Lang get translate function for specific language
func Lang(lang string) TranslatorFunc {
	tag, err := language.All.Parse(lang)
	if err != nil {
		return defaultTranslator
	}
	l, ok := localizes[tag]
	if !ok {
		return defaultTranslator
	}
	f, ok := functions[tag]
	if !ok {
		f = createTranslator(l)
		functions[tag] = f
	}
	return f
}

//// local

func createTranslator(l *i18n.Localizer) TranslatorFunc {
	return func(msgId string, a ...interface{}) string {
		t := map[string]interface{}{}
		for i := range a {
			t["v"+strconv.Itoa(i+1)] = a[i]
		}
		r, err := l.Localize(&i18n.LocalizeConfig{MessageID: msgId, TemplateData: t})
		if err == nil {
			return r
		}
		return defaultTranslator(msgId, a...)
	}
}

func defaultTranslator(msg string, a ...interface{}) string {
	return fmt.Sprintf("#I18N MISS# %s %#v", msg, a)
}
