package models

//	{
//		"messages":[
//			{
//				"msg": {
//					"product_id": "MLA19809670"
//				},
//				"id": "13781977501846641"
//			},
//				{
//				"msg": {
//					"product_id": "MLA27011547"
//				},
//				"id": "13781977501846641"
//			}
//		]
//		}
type Messages struct {
	Messages []*Message `json:"messages"`
}
type Message struct {
	Msg *Msg    `json:"msg"`
	Id  *string `json:"id"`
}

type Msg struct {
	ProductId *string `json:"product_id"`
}
