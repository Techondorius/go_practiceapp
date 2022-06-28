# API Document

## ユーザー一覧取得 API

### リクエスト

```
GET /userManage/getUser
```

### レスポンス

#### 成功時

| param                 | type   | description |
| --------------------- | ------ | ----------- |
| detail[].ID           | number | ユーザーID    |
| detail[].FirstName    | string | 名前         |
| detail[].LastName     | string | 苗字         |
| detail[].Hourly_wage  | number | 時給         |

```javascript
{
    "detail": [
        {
            "ID": number,
            "FirstName": string,
            "LastName": string,
            "Hourly_wage": number
        },
        ...
    ]
}
```

## ユーザー作成 API

### リクエスト

```
PUT /userManage/newUser
```

| param       | type   | description   |
| ----------- | ------ | -----------   |
| FirstName   | string | 名前           |
| LastName    | string | 苗字           |
| Hourly_wage | number | 時給(Optional) |


```javascript
{
    "FirstName": string,
    "LastName": string
    "Hourly_wage": number
}
```

### レスポンス

#### 成功時

| param               | type   | description |
| ------------------- | ------ | ----------- |
| message             | string | Created     |
| detail[].ID         | number | ID          |
| detail[].FirstName  | string | 名前         |
| detail[].LastName   | string | 苗字         |
| detail[].Hourly_wage| string | 時給         |


```javascript
{
    "message": "Created"
    "detail": {
        "ID": 12,
        "FirstName": "Kyosuke",
        "LastName": "Fujita"
        "Hourly_wage": 10000,
    },
    ...
}
```

### 失敗時

#### Request bodyが不完全な時

400 Bad Request
