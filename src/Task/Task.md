# 各タスクの使い方基本

## 準備

1. `make firestore-start` でfirestoreのエミュレータを動かす
2. `make run` でコンテナを起動(環境変数が必要なのでmakeファイル経由で起動すること)
3. `make bash` でコンテナに入る

コンテナに入ったあとは `go build`でバイナリを作成し, バイナリ経由でタスクを実行してください

## Task一覧

- [Voice](./VoiceTask.md)
