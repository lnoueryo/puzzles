package cloudfunctions

import (
    "context"
    "encoding/json"
    "log"

    "golang.org/x/oauth2/google"
    sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

// PubSubMessage は、Pub/Sub eventのペイロード
type PubSubMessage struct {
    Data []byte `json:"data"`
}
{
    "data": {
        "Instance": "production",
        "Project": "puzzles-345814",
        "Action": "start"
    }
}
type MessagePayload struct {
    Instance string // Cloud SQL のインスタンス名
    Project  string // プロジェクトID
    Action   string // 起動か停止か
}

// Pub/Sub message を受け取って処理
func ProcessPubSub(ctx context.Context, m PubSubMessage) error {
    var psData MessagePayload
    // PubSubMessageのjsonデータをMessagePayloadオブジェクトに変換する
    err := json.Unmarshal(m.Data, &psData)
    if err != nil {
        log.Println(err)
    }
    log.Printf("Request received for Cloud SQL instance %s action: %s, %s", psData.Action, psData.Instance, psData.Project)

    // デフォルトの認証情報を使って、http.Cllientを作成
    hc, err := google.DefaultClient(ctx, sqladmin.CloudPlatformScope)
    if err != nil {
        return err
    }

    // Cloud SQL service を作成
    service, err := sqladmin.New(hc)
    if err != nil {
        return err
    }

    // 起動と停止のアクションのリクエストを読み取る
    action := "UNDEFINED"
    switch psData.Action {
    case "start":
        action = "ALWAYS"
    case "stop":
        action = "NEVER"
    default:
        log.Fatal("No valid action provided.")
    }

    // 読み取ったアクションをCloud SQLに設定する準備をする
    rb := &sqladmin.DatabaseInstance{
        Settings: &sqladmin.Settings{
            ActivationPolicy: action,
        },
    }
    // CloudSQLインスタンス内のデータベースに関する情報を含むリソースを部分的に更新
    resp, err := service.Instances.Patch(psData.Project, psData.Instance, rb).Context(ctx).Do()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%#v\n", resp)
    return nil
} 