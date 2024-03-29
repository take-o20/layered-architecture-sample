# Layered architecture sample with rest


## ADR
| HTTP REQUEST METHOD | URI | 処理 |
| :--- | :--- | :--- |
| GET | /api/v1/user | 全てのユーザーの情報を返す |
| GET | /api/v1/user/{user_id} | ユーザーの情報を返す |
| POST| /api/v1/user | ユーザーを作成する |
| PATCH or PUT| /api/v1/user/{user_id} | ユーザーの情報を更新する |
| DELETE| /api/v1/user | ユーザーを削除する |

### User

#### GET `/api/v1/user`
* request body
    * none
* response body
    ```json
    {
        "message": "got users",
        "users": [
            {
                "user_id": "$ID",
                "name": "$NAME",
                "email": "$EMAIL"
            },{
                "user_id": 2,
                "name": "$NAME",
                "email": "$EMAIL"
            },
        ]
    }
    ```
    * error case
        ```json
        {
            "message": "failed to got users"
        }
        ```


#### GET `/api/v1/user/{user_id}`
* uri example
    ```
    /api/v1/user/1
    ```
* request body
    * none
* response body
    ```json
        "message": "got user",
        "users": [
            {
                "user_id": "$ID",
                "name": "$NAME",
                "email": "$EMAIL"
            },
        ]
    ```
    * error case
        ```json
        {
            "message": "failed to got user"
        }
        ```


#### POST `/api/v1/user`

* request body
    ```json
    {
        "name": "$NAME",
        "email": "$EMAIL"
    }
    ```
* response
    ```json
    {
        "message": "created user",
        "users": [{
            "user_id": "$ID",
            "name": "$NAME",
            "email": "$EMAIL"
        }]
    }
    ```
    * error case
        ```json
        {
            "message": "failed to create user"
        }
        ```


#### PUT or PATCH `/api/v1/user/{user_id}`
* uri example
    ```
    /api/v1/user/1
    ```
* request body
    ```json
    {
        "name": "$NAME",
        "email": "$EMAIL"
    }
    ```
* response
    ```json
    {
        "message": "updated user",
        "users": [{
            "user_id": "$ID",
            "name": "new-name",
            "email": "new-email"
        }]
    }
    ```
    * error case
        ```json
        {
            "message": "failed to update user"
        }
        ```

#### DELETE
* uri
    ```
    /api/v1/user/{user_id}
    ```
* request body
    * none
* response
    ```json
    {
        "message": "deleted user",
        "users": [{
            "user_id": "$ID",
            "name": "$NAME",
            "email": "$EMAIL"
        }]
    }
    ```
    * error case
        ```json
        {
            "message": "failed to delete user"
        }
        ```






## 拡張案
* 現在の課題
    * (1) PATCH処理にて値が全く同じときにsqlが影響を与えた行が0になる
* Userテーブルに下記属性を追加
    * created_at
    * updated_at  // 課題(1)の解決策になる可能性あり

