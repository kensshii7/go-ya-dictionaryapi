package main

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
)

// Структура DictionaryAPI описывает поле ApiKey и методы GetLangs, Lookup, doRequest
type DictionaryAPI struct {
	/*
		Поле ApiKey используется для взаимодействия с API
	*/
	ApiKey string
}

const (
	/* Опциональные флаги для поиска */

	/*
		Флаг FAMILY применяет family-search фильтер
		Слова "для взрослых" и слова с нецензурной лексикой будут исключены
	*/
	FLAGS_FAMILY = 0x0001

	/*
		Флаг MORPHO включает поиск слова по словоформе
	*/
	FLAGS_MORPHO = 0x0004

	/*
		Флаг POS_FILTER применяет family-search фильтер
		Слова "для взрослых" и слова с нецензурной лексикой будут исключены
	*/
	FLAGS_POS_FILTER = 0x0008
)

// Метод Init инициализирует DictionaryAPI
func Init(apiKey string) *DictionaryAPI {
	return &DictionaryAPI{
		ApiKey: apiKey,
	}
}

// Метод doRequest используется для создания запросов к API
func (api *DictionaryAPI) doRequest(url string) (*bytes.Buffer, error) {

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	switch req.StatusCode {
	case 401:
		err = errors.New("ERR_KEY_INVALID")
	case 402:
		err = errors.New("ERR_KEY_BLOCKED")
	case 403:
		err = errors.New("ERR_DAILY_REQ_LIMIT_EXCEEDED")
	case 413:
		err = errors.New("ERR_TEXT_TOO_LONG")
	case 501:
		err = errors.New("ERR_LANG_NOT_SUPPORTED")
	default:
		err = nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)

	return buf, err
}

/*
	Метод GetLangs используется для вывода списка направлений перевода , которые используются сервисом
	Например:
		en-ru, ru-ru
*/
func (api *DictionaryAPI) GetLangs() (*bytes.Buffer, error) {
	return api.doRequest("https://dictionary.yandex.net/api/v1/dicservice.json/getLangs?key=" + api.ApiKey)
}

/*
	Метод Lookup ищет слову/фразу в словаре
	Аргументы:
		lang - направление перевода (en-ru, ru-ru, see GetLangs)
		text - слово или фраза
		ui - язык пользовательского интерфейса
		flags - опциональные флаги
*/
func (api *DictionaryAPI) Lookup(text, lang, ui string, flags int) (*bytes.Buffer, error) {
	return api.doRequest("https://dictionary.yandex.net/api/v1/dicservice.json/lookup?key=" + api.ApiKey + "&lang=" + lang + "&text=" + text + "&ui=" + ui + "&flags=" + strconv.Itoa(flags))
}
