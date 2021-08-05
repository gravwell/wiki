＃Gravwellベータプログラム

ご挨拶

Gravwell Betaの早期アクセスグループに所属しているため、このメッセージを読んでいます。皆様のご参加に心より感謝申し上げますとともに、フィードバックやバグレポートをお待ちしております。

バグ報告やフィードバックは[beta@gravwell.io](mailto:beta@gravwell.io)までお願いします。

ありがとうございます！

## 現在の "ベータ版の状態"

Gravwell 4.2.0のリリースに向けて準備を進めています。このリリースの主な新機能は「データエクスプローラー」で、ポイント＆クリックで簡単にデータを操作できるようになりました。

### 希望するテスト

このスプリントのテスト希望（優先度順）

* データエクスプローラー - できるだけ多くの異なるデータソースでそれを試してみてください。フィルタを追加したり、フィルタを削除したり、他のクエリモードにピボットしたりします。
* クエリスタジオ - クエリ編集ボックスが使いやすく堅牢であることを確認し、多くのクエリとタグの間を移動し、他のレンダラーで新しいフォーマットをチェックします。
* タグごとのアクセラレータ - いくつかの [タグごとのアクセラレータ定義](#!configuration/accelerators.md#Accelerating_Specific_Tags) を追加します。

## インストールとアップグレード

このビルドが利用可能になり、テストできるようになったことを大変うれしく思います。新しいubuntuリポジトリとDockerイメージを作成しました。StableからBetaへの切り替えは、aptソースリポジトリ（または最初からインストールする場合はクイックスタート手順）を変更することで実行されます。

### アップグレード：

`etc/apt/sources.list.d/gravwell.list` ファイルを編集し、`https://update.gravwell.io/debian/`の代わりに`https://update.gravwell.io/debianbeta/`を使用します。次に `apt update` と `apt upgrade` を実行すると、新しいリリースになります。

### ゼロからのインストール：

```
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debianbeta/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install gravwell
```

### Docker

Dockerイメージは[gravwell/beta](https://hub.docker.com/r/gravwell/beta)にあります。ドキュメント内のどのDockerコマンドについても、`gravwell/gravwell`を`gravwell/beta`に変更すれば、"難なく動く "はずです。

## ありがとうございます

Gravwellがもたらす新しい機能にとても興奮しています。ベータプログラムに興味を持っていただき、ご参加いただきありがとうございました。あなたなしではできませんでした。

フィードバックやバグレポート、特に新しいツールを使って作ったクールなものを見せてください!
