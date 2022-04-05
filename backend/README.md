起動時の処理
1.環境の確認
2.テンプレートの読み込み、キャッシュ


デプロイ
gcloud run deploy --source .

.envファイルと.env.devの作成
本番環境には.envを、develop環境には.env.devを書く
GITHUB_CLIENT_ID=
GITHUB_SECRET_ID=
APP_ENV=local
APP_HOST=localhost
DB_NAME=practices
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=""
DB_PORT=3306
DB_QUERY=parseTime=true

deploy時は.env.devを上げないため.envを読みに行く。



csv、zipのダウンロード、アップロード
inoueryo on 4/6/2022, 4:05:12 AM
8ac4f85  SoftHardTagBranchMore
ユーザー登録バックエンド
inoueryo on 4/6/2022, 4:04:39 AM
0eb1a66  SoftHardTagBranchMore
ユーザー登録
inoueryo on 4/6/2022, 4:03:04 AM
2064cb5  SoftHardTagBranchMore
その他フロント微調整
inoueryo on 4/6/2022, 4:02:40 AM
bff6132  SoftHardTagBranchMore
エラー関連など
inoueryo on 4/6/2022, 3:59:17 AM
5ca283f  SoftHardTagBranchMore
terraform
inoueryo on 4/6/2022, 3:57:12 AM
a03a7e4  SoftHardTagBranchMore
コマンド関連変更
