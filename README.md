### Soliton NK Wiki
Soliton NK のマニュアルの最新版は https://docs.example.com/ で公開しています。
また、このリポジトリをダウンロードすれば、オフラインでもマニュアルを参照することができます。

オフラインでマニュアルを参照する場合は、このリポジトリをドキュメントルートとしてウェブサーバを起ち上げ、ウェブブラウザからアクセスして下さい。

リポジトリにある server.go は、Go 言語で書かれた簡易ウェブサーバプログラムです。$ go run serve.go でサーバが起動します。ウェブブラウザから http://localhost:3001/ にアクセスして下さい。

マニュアルは Markdown で書かれており、MDwiki http://dynalon.github.io/mdwiki/ を使って HTML に変換しています。
