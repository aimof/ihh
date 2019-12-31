package main

import "testing"

import "strings"

func TestCurrent(t *testing.T) {
	const j = `{"coord":{"lon":139.69,"lat":35.69},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02n"}],"base":"stations","main":{"temp":279.14,"feels_like":266.67,"temp_min":277.59,"temp_max":280.37,"pressure":1015,"humidity":41},"visibility":10000,"wind":{"speed":13.9,"deg":340,"gust":19.5},"clouds":{"all":20},"dt":1577790300,"sys":{"type":1,"id":8074,"country":"JP","sunrise":1577742635,"sunset":1577777810},"timezone":32400,"id":1850147,"name":"Tokyo","cod":200}`
	n, err := current(strings.NewReader(j))
	if err != nil {
		t.Error()
	}
	if n != 1015 {
		t.Error()
	}
}

func TestForecast(t *testing.T) {
	const j = `{"cod":"200","message":0,"cnt":40,"list":[{"dt":1577793600,"main":{"temp":277.4,"feels_like":268.75,"temp_min":277.4,"temp_max":280.02,"pressure":1019,"sea_level":1019,"grnd_level":1010,"humidity":32,"temp_kf":-2.62},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":33},"wind":{"speed":7.89,"deg":352},"sys":{"pod":"n"},"dt_txt":"2019-12-31 12:00:00"},{"dt":1577804400,"main":{"temp":276.68,"feels_like":269.58,"temp_min":276.68,"temp_max":278.64,"pressure":1021,"sea_level":1021,"grnd_level":1013,"humidity":32,"temp_kf":-1.96},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":1},"wind":{"speed":5.61,"deg":7},"sys":{"pod":"n"},"dt_txt":"2019-12-31 15:00:00"},{"dt":1577815200,"main":{"temp":276.12,"feels_like":270.09,"temp_min":276.12,"temp_max":277.43,"pressure":1022,"sea_level":1022,"grnd_level":1014,"humidity":32,"temp_kf":-1.31},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":1},"wind":{"speed":4.04,"deg":25},"sys":{"pod":"n"},"dt_txt":"2019-12-31 18:00:00"},{"dt":1577826000,"main":{"temp":276.58,"feels_like":270.89,"temp_min":276.58,"temp_max":277.23,"pressure":1023,"sea_level":1023,"grnd_level":1016,"humidity":31,"temp_kf":-0.65},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04n"}],"clouds":{"all":99},"wind":{"speed":3.56,"deg":32},"sys":{"pod":"n"},"dt_txt":"2019-12-31 21:00:00"},{"dt":1577836800,"main":{"temp":277.58,"feels_like":272.67,"temp_min":277.58,"temp_max":277.58,"pressure":1024,"sea_level":1024,"grnd_level":1017,"humidity":32,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":{"all":100},"wind":{"speed":2.56,"deg":29},"sys":{"pod":"d"},"dt_txt":"2020-01-01 00:00:00"},{"dt":1577847600,"main":{"temp":280.27,"feels_like":276.67,"temp_min":280.27,"temp_max":280.27,"pressure":1022,"sea_level":1022,"grnd_level":1015,"humidity":33,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03d"}],"clouds":{"all":31},"wind":{"speed":1,"deg":146},"sys":{"pod":"d"},"dt_txt":"2020-01-01 03:00:00"},{"dt":1577858400,"main":{"temp":282.76,"feels_like":278.61,"temp_min":282.76,"temp_max":282.76,"pressure":1020,"sea_level":1020,"grnd_level":1013,"humidity":52,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":15},"wind":{"speed":3.14,"deg":163},"sys":{"pod":"d"},"dt_txt":"2020-01-01 06:00:00"},{"dt":1577869200,"main":{"temp":281.55,"feels_like":277.64,"temp_min":281.55,"temp_max":281.55,"pressure":1021,"sea_level":1021,"grnd_level":1014,"humidity":50,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":2.46,"deg":46},"sys":{"pod":"n"},"dt_txt":"2020-01-01 09:00:00"},{"dt":1577880000,"main":{"temp":280.9,"feels_like":276.25,"temp_min":280.9,"temp_max":280.9,"pressure":1021,"sea_level":1021,"grnd_level":1014,"humidity":43,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":3.07,"deg":358},"sys":{"pod":"n"},"dt_txt":"2020-01-01 12:00:00"},{"dt":1577890800,"main":{"temp":280.06,"feels_like":276.16,"temp_min":280.06,"temp_max":280.06,"pressure":1021,"sea_level":1021,"grnd_level":1014,"humidity":47,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":2.06,"deg":349},"sys":{"pod":"n"},"dt_txt":"2020-01-01 15:00:00"},{"dt":1577901600,"main":{"temp":279.35,"feels_like":275.76,"temp_min":279.35,"temp_max":279.35,"pressure":1021,"sea_level":1021,"grnd_level":1013,"humidity":41,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.24,"deg":348},"sys":{"pod":"n"},"dt_txt":"2020-01-01 18:00:00"},{"dt":1577912400,"main":{"temp":278.6,"feels_like":274.01,"temp_min":278.6,"temp_max":278.6,"pressure":1022,"sea_level":1022,"grnd_level":1015,"humidity":37,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":2.41,"deg":14},"sys":{"pod":"n"},"dt_txt":"2020-01-01 21:00:00"},{"dt":1577923200,"main":{"temp":279.65,"feels_like":274.84,"temp_min":279.65,"temp_max":279.65,"pressure":1023,"sea_level":1023,"grnd_level":1016,"humidity":36,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":2.8,"deg":13},"sys":{"pod":"d"},"dt_txt":"2020-01-02 00:00:00"},{"dt":1577934000,"main":{"temp":282.67,"feels_like":278.34,"temp_min":282.67,"temp_max":282.67,"pressure":1022,"sea_level":1022,"grnd_level":1016,"humidity":31,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":2.2,"deg":33},"sys":{"pod":"d"},"dt_txt":"2020-01-02 03:00:00"},{"dt":1577944800,"main":{"temp":283.8,"feels_like":280.38,"temp_min":283.8,"temp_max":283.8,"pressure":1021,"sea_level":1021,"grnd_level":1015,"humidity":35,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":1.29,"deg":93},"sys":{"pod":"d"},"dt_txt":"2020-01-02 06:00:00"},{"dt":1577955600,"main":{"temp":283.25,"feels_like":279.99,"temp_min":283.25,"temp_max":283.25,"pressure":1023,"sea_level":1023,"grnd_level":1016,"humidity":37,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":1},"wind":{"speed":1.1,"deg":149},"sys":{"pod":"n"},"dt_txt":"2020-01-02 09:00:00"},{"dt":1577966400,"main":{"temp":281.99,"feels_like":277.61,"temp_min":281.99,"temp_max":281.99,"pressure":1023,"sea_level":1023,"grnd_level":1016,"humidity":39,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":2},"wind":{"speed":2.63,"deg":348},"sys":{"pod":"n"},"dt_txt":"2020-01-02 12:00:00"},{"dt":1577977200,"main":{"temp":280.19,"feels_like":274.83,"temp_min":280.19,"temp_max":280.19,"pressure":1023,"sea_level":1023,"grnd_level":1016,"humidity":42,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":30},"wind":{"speed":3.93,"deg":11},"sys":{"pod":"n"},"dt_txt":"2020-01-02 15:00:00"},{"dt":1577988000,"main":{"temp":279.12,"feels_like":273.91,"temp_min":279.12,"temp_max":279.12,"pressure":1021,"sea_level":1021,"grnd_level":1014,"humidity":47,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":39},"wind":{"speed":3.8,"deg":8},"sys":{"pod":"n"},"dt_txt":"2020-01-02 18:00:00"},{"dt":1577998800,"main":{"temp":278.48,"feels_like":273.16,"temp_min":278.48,"temp_max":278.48,"pressure":1021,"sea_level":1021,"grnd_level":1014,"humidity":47,"temp_kf":0},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"clouds":{"all":52},"wind":{"speed":3.86,"deg":12},"sys":{"pod":"n"},"dt_txt":"2020-01-02 21:00:00"},{"dt":1578009600,"main":{"temp":279.2,"feels_like":274.25,"temp_min":279.2,"temp_max":279.2,"pressure":1020,"sea_level":1020,"grnd_level":1014,"humidity":43,"temp_kf":0},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"clouds":{"all":72},"wind":{"speed":3.25,"deg":357},"sys":{"pod":"d"},"dt_txt":"2020-01-03 00:00:00"},{"dt":1578020400,"main":{"temp":282.31,"feels_like":278.07,"temp_min":282.31,"temp_max":282.31,"pressure":1017,"sea_level":1017,"grnd_level":1011,"humidity":36,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":18},"wind":{"speed":2.31,"deg":19},"sys":{"pod":"d"},"dt_txt":"2020-01-03 03:00:00"},{"dt":1578031200,"main":{"temp":283.5,"feels_like":280.69,"temp_min":283.5,"temp_max":283.5,"pressure":1016,"sea_level":1016,"grnd_level":1009,"humidity":40,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":9},"wind":{"speed":0.66,"deg":133},"sys":{"pod":"d"},"dt_txt":"2020-01-03 06:00:00"},{"dt":1578042000,"main":{"temp":283.08,"feels_like":280.16,"temp_min":283.08,"temp_max":283.08,"pressure":1017,"sea_level":1017,"grnd_level":1009,"humidity":42,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":0.87,"deg":311},"sys":{"pod":"n"},"dt_txt":"2020-01-03 09:00:00"},{"dt":1578052800,"main":{"temp":281.78,"feels_like":278.42,"temp_min":281.78,"temp_max":281.78,"pressure":1017,"sea_level":1017,"grnd_level":1009,"humidity":48,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.61,"deg":310},"sys":{"pod":"n"},"dt_txt":"2020-01-03 12:00:00"},{"dt":1578063600,"main":{"temp":281.02,"feels_like":278,"temp_min":281.02,"temp_max":281.02,"pressure":1016,"sea_level":1016,"grnd_level":1009,"humidity":54,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":1.31,"deg":312},"sys":{"pod":"n"},"dt_txt":"2020-01-03 15:00:00"},{"dt":1578074400,"main":{"temp":280.04,"feels_like":275.69,"temp_min":280.04,"temp_max":280.04,"pressure":1016,"sea_level":1016,"grnd_level":1009,"humidity":48,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":0},"wind":{"speed":2.74,"deg":10},"sys":{"pod":"n"},"dt_txt":"2020-01-03 18:00:00"},{"dt":1578085200,"main":{"temp":279.2,"feels_like":274.34,"temp_min":279.2,"temp_max":279.2,"pressure":1017,"sea_level":1017,"grnd_level":1010,"humidity":51,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":29},"wind":{"speed":3.48,"deg":8},"sys":{"pod":"n"},"dt_txt":"2020-01-03 21:00:00"},{"dt":1578096000,"main":{"temp":280.14,"feels_like":275.34,"temp_min":280.14,"temp_max":280.14,"pressure":1017,"sea_level":1017,"grnd_level":1011,"humidity":48,"temp_kf":0},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02d"}],"clouds":{"all":15},"wind":{"speed":3.4,"deg":22},"sys":{"pod":"d"},"dt_txt":"2020-01-04 00:00:00"},{"dt":1578106800,"main":{"temp":282.18,"feels_like":277.64,"temp_min":282.18,"temp_max":282.18,"pressure":1015,"sea_level":1015,"grnd_level":1009,"humidity":42,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":{"all":93},"wind":{"speed":3.04,"deg":27},"sys":{"pod":"d"},"dt_txt":"2020-01-04 03:00:00"},{"dt":1578117600,"main":{"temp":282.95,"feels_like":279.44,"temp_min":282.95,"temp_max":282.95,"pressure":1014,"sea_level":1014,"grnd_level":1007,"humidity":39,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04d"}],"clouds":{"all":97},"wind":{"speed":1.52,"deg":60},"sys":{"pod":"d"},"dt_txt":"2020-01-04 06:00:00"},{"dt":1578128400,"main":{"temp":282.95,"feels_like":279.46,"temp_min":282.95,"temp_max":282.95,"pressure":1014,"sea_level":1014,"grnd_level":1007,"humidity":41,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04n"}],"clouds":{"all":100},"wind":{"speed":1.61,"deg":82},"sys":{"pod":"n"},"dt_txt":"2020-01-04 09:00:00"},{"dt":1578139200,"main":{"temp":282.5,"feels_like":279.55,"temp_min":282.5,"temp_max":282.5,"pressure":1013,"sea_level":1013,"grnd_level":1006,"humidity":59,"temp_kf":0},"weather":[{"id":804,"main":"Clouds","description":"overcast clouds","icon":"04n"}],"clouds":{"all":99},"wind":{"speed":1.76,"deg":35},"sys":{"pod":"n"},"dt_txt":"2020-01-04 12:00:00"},{"dt":1578150000,"main":{"temp":280.86,"feels_like":276.49,"temp_min":280.86,"temp_max":280.86,"pressure":1012,"sea_level":1012,"grnd_level":1006,"humidity":61,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":49},"wind":{"speed":3.55,"deg":13},"sys":{"pod":"n"},"dt_txt":"2020-01-04 15:00:00"},{"dt":1578160800,"main":{"temp":279.54,"feels_like":273.63,"temp_min":279.54,"temp_max":279.54,"pressure":1013,"sea_level":1013,"grnd_level":1007,"humidity":73,"temp_kf":0},"weather":[{"id":802,"main":"Clouds","description":"scattered clouds","icon":"03n"}],"clouds":{"all":34},"wind":{"speed":6.03,"deg":6},"sys":{"pod":"n"},"dt_txt":"2020-01-04 18:00:00"},{"dt":1578171600,"main":{"temp":279.16,"feels_like":274.11,"temp_min":279.16,"temp_max":279.16,"pressure":1015,"sea_level":1015,"grnd_level":1008,"humidity":55,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01n"}],"clouds":{"all":9},"wind":{"speed":3.93,"deg":354},"sys":{"pod":"n"},"dt_txt":"2020-01-04 21:00:00"},{"dt":1578182400,"main":{"temp":281.91,"feels_like":274.1,"temp_min":281.91,"temp_max":281.91,"pressure":1018,"sea_level":1018,"grnd_level":1011,"humidity":38,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":4},"wind":{"speed":7.46,"deg":356},"sys":{"pod":"d"},"dt_txt":"2020-01-05 00:00:00"},{"dt":1578193200,"main":{"temp":283.65,"feels_like":276.09,"temp_min":283.65,"temp_max":283.65,"pressure":1018,"sea_level":1018,"grnd_level":1011,"humidity":30,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":0},"wind":{"speed":6.88,"deg":19},"sys":{"pod":"d"},"dt_txt":"2020-01-05 03:00:00"},{"dt":1578204000,"main":{"temp":284.34,"feels_like":278.88,"temp_min":284.34,"temp_max":284.34,"pressure":1019,"sea_level":1019,"grnd_level":1012,"humidity":29,"temp_kf":0},"weather":[{"id":800,"main":"Clear","description":"clear sky","icon":"01d"}],"clouds":{"all":3},"wind":{"speed":3.9,"deg":32},"sys":{"pod":"d"},"dt_txt":"2020-01-05 06:00:00"},{"dt":1578214800,"main":{"temp":283.17,"feels_like":278.01,"temp_min":283.17,"temp_max":283.17,"pressure":1022,"sea_level":1022,"grnd_level":1014,"humidity":30,"temp_kf":0},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"clouds":{"all":84},"wind":{"speed":3.4,"deg":20},"sys":{"pod":"n"},"dt_txt":"2020-01-05 09:00:00"}],"city":{"id":1850147,"name":"Tokyo","coord":{"lat":35.6895,"lon":139.6917},"country":"JP","timezone":32400,"sunrise":1577742634,"sunset":1577777810}}`
	n, err := forecast(strings.NewReader(j))
	if err != nil {
		t.Error()
	}
	if n != 1021 {
		t.Error()
	}
}