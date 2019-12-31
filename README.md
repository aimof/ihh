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