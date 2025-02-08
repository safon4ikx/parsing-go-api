package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const city = "Курск"

// Структура - текущая погода из ключа json c массивом - current_condition
type СurrentСondition struct {
	FeelsLikeC string `json:"FeelsLikeC"` // по ощущению - ключ
	TempC      string `json:"temp_C"`     // градус по замеру - ключ
}

// Структура - Ответ от сервиса погоды
type WeatherResponse struct {
	СurrentСondition []СurrentСondition `json:"current_condition"` //Текущая погода - ключ
}

func createNewWeatcher() {
	/**

	"current_condition": [
	    {
	      "key1": "random text" //string
	      "key2": true, // boolean
	      "key3": 4000, //int64
		  "lang_ru": [{value:""}], // struct & map[string]string & interface{}
	*/
	jsonData, err := json.Marshal(&WeatherResponse{
		[]СurrentСondition{
			{
				FeelsLikeC: "-6",
				TempC:      "0",
			},
			{
				FeelsLikeC: "-6",
				TempC:      "0",
			},
			{
				FeelsLikeC: "-6",
				TempC:      "0",
			},
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jsonData)

	fmt.Println(string(jsonData))
}

func main() {
	urlTemplate := fmt.Sprintf("https://ru.wttr.in/%v?format=j2", city)

	//HTTP
	resp, err := http.Get(urlTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(body))

	/*Происходит преобразование из json в структуру */

	//Объявляем переменную и присваиваем структру WeatherResponse
	var wather_resp WeatherResponse

	/*Используем пакет json и метод Unmarshal для преобразования нашего байт-код строкового json из ответа сервер в body(тела-ответа-на запрос)
	Далее записываем преобразованные данные по ссылке на переменную с помощь &название_переменой(wather_resp) */
	err = json.Unmarshal(body, &wather_resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*Возвращаем преобразованный ответ из переменной по струкуре WeatherResponse -> CurrrentCondition[срез объекта 0 для яп-пк а для человека = 1] -> TempC, FeelsLikeC */

	//RESULT
	fmt.Printf("Погода сейчас: %+v*(%v)", wather_resp.СurrentСondition[0].TempC, wather_resp.СurrentСondition[0].FeelsLikeC)

	// fmt.Println(urlTemplate)
}
