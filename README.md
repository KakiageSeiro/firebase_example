### GCPの操作

初期設定

```
# 認証情報を作成
$ gcloud auth login  --no-launch-browser
$ gcloud auth application-default login --no-launch-browser

# 接続確認とfunctionsが利用できる確認
$ gcloud functions list
```

デプロイ
※ただしgitlab ciからしか実行はできない

```
$ gcloud functions deploy hello --entry-point get_nayose_result --runtime go111 --trigger-http --entry-point HelloHTTP --project shinofara-233410
```