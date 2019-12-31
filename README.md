# ihh - I have a headache

今後数時間で気圧低下による頭痛が起きるかどうか絵文字で教えてくれるプログラムです。

## 使い方

1. `git clone`して、`go install`してください。
2. 環境変数に、`owmCityID`と`owmAppID`を指定してください。
  * `owmCityID`は、OpenWeatherMapにおけるCityID。（東京は`1850147`）
  * `owmAppID`は、OpenWeatherMapのAPI Keyです。
3. `ihh`コマンドで以下の結果が得られます。
  * 😀👍：気圧維持or上昇
  * 🤢👎：気圧減少

## Exit Status

|status code|意味|
|:--|:--|
|3|環境変数が定義されていない|
|4|現在の天気の取得するGETでエラー|
|5|取得情報から気圧を見る過程でエラー|
|6|天気の予報を取得するGETでエラー|
|7|取得情報から気圧を見る過程でエラー|