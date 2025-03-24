package models

// {
//     "responses": [
//         {
//             "id": "13781977501846641",
//             "code": 200
//         },
//         {
//             "id": "13781977501846641",
//             "code": 200
//         }
//     ]
// }

type Responses struct {
	Responses []*Response `json:"responses"`
}

type Response struct {
	Id   string `json:"id"`
	Code int    `json:"code"`
}
