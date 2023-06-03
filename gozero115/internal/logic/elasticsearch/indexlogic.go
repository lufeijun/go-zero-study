package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var esClient *elasticsearch.Client

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = funcs.ResponseInit()

	var r map[string]interface{}
	// res, err := esClient.Info()
	// if err != nil {
	// 	fmt.Println("Error getting response: %s", err)
	// 	return
	// }

	// json.NewDecoder(res.Body).Decode(&r)
	// fmt.Println("=======")
	// fmt.Println(res.StatusCode)
	// fmt.Println("=======")
	// defer res.Body.Close()
	// resp.Data = r
	// return

	// 创建
	// req := esapi.CreateRequest{
	// 	Index:      "my_user_log",
	// 	DocumentID: "2",
	// 	Body: strings.NewReader(`{
	// 		"id" : 1,
	// 		"user_id" : 1,
	// 		"openid" : "ooxxooxxooxxooxxoo",
	// 		"action" : "测试",
	// 		"time" : 1685769738,
	// 		"data_id": 11,
	// 		"data_unboxing_id" : 12,
	// 		"data_source" : "测试机"
	// 	}`),
	// }

	// req := esapi.UpdateRequest{
	// 	Index:      "my_user_log",
	// 	DocumentID: "2",
	// 	Body: strings.NewReader(`{
	// 		"doc" : {
	// 			"id" : 2,
	// 			"user_id" : 2,
	// 			"openid" : "ooxxooxxooxxooxxoo",
	// 			"action" : "测试",
	// 			"time" : 1685769738,
	// 			"data_id": 11,
	// 			"data_unboxing_id" : 12,
	// 			"data_source" : "测试机"
	// 		}
	// 	}`),
	// }

	data := make(map[string]interface{})
	msg := make(map[string]interface{})
	msg["id"] = 2
	msg["user_id"] = 2
	msg["action"] = "测试" + time.Now().Format("2006-01-02 15:04:05")
	msg["data_source"] = "测试机" + time.Now().Format("2006-01-02 15:04:05")
	data["doc"] = msg
	fmt.Println(data)
	str, _ := json.Marshal(data)
	req := esapi.UpdateRequest{
		Index:      "my_user_log",
		DocumentID: "2",
		Body:       strings.NewReader(string(str)),
	}

	res, err := req.Do(context.Background(), esClient)

	if err != nil {
		fmt.Println("出错了，这个你麻麻滴错误是", err)
	}

	json.NewDecoder(res.Body).Decode(&r)

	resp.Data = r

	return
}

func init() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.0.86:9200",
			"http://192.168.0.87:9200",
			"http://192.168.0.89:9200",
		},
	}
	var err error
	esClient, err = elasticsearch.NewClient(cfg)

	fmt.Println("es version：", elasticsearch.Version)

	if err != nil {
		fmt.Println("Error creating the client: %s", err)
	}
}
