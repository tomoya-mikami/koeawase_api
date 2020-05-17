# VoiceTask

## データの追加

使い方

```
${作成したバイナリ名} cli voice add ${保存する名前} ${./media以下においたファイル名}
```

結果

```
root@b21c36e8a2b1:/workdir# ./koeawase_api cli voice add myvoice myvoice.wav
add voice id:J66hym3b1dYWGRh1YF7E name:myvoice power spectrum frequency:41348
```

## 入稿したデータの類似度を計算する

使い方

```
${作成したバイナリ名} cli voice calclateSimilarity ${保存する名前} ${./media以下においたファイル名}
```

結果
```
root@b21c36e8a2b1:/workdir# ./koeawase_api cli voice calclateSimilarity mYAzzJUcOmHBUW3f7wqX lTVVPiwI0lNzoBuCFkCR
myvoiceさんとnansuさんの声の類似度は0.170049です
```

idはfirestoreに保存されているドキュメントのID 保存のCLIの実行結果をコピーして使う
