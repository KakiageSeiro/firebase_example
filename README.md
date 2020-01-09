### GCPの操作

初期設定

```
$ npm install -g firebase-tools
$ firebase init
```

デプロイ
※ただしgitlab ciからしか実行はできない

```
$ gcloud functions deploy NAME --runtime RUNTIME TRIGGER [FLAGS...]
```