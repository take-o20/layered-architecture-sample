# Layered architecture sample with rest


## ADR

### User

URI
```
/api/v1/user
```

#### GET
* url
    ```
    /api/v1/user?user_id
    ```
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
                "user_id": "$ID",
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


#### GET with parameter
* parameter
    ```
    /api/v1/user?user_id
    ```
* request body
    * none
* response body
    ```json
        "message": "get all users",
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


#### POST

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
        "message": "create user",
        "users": [{
            "name": "name",
            "email": "email"
        }]
    }
    ```
    * error case
        ```json
        {
            "message": "failed to create user"
        }
        ```


#### PUT or PATCH
* parameter
    ```
    /api/v1/user?user_id
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
        "message": "update user",
        "users": []{
            "id": "id",
            "name": "name",
            "email": "email"
        }
    }
    ```
    * error case
        ```json
        {
            "message": "failed to update user"
        }
        ```

#### DELETE
* parameter
    ```
    /api/v1/user?user_id
    ```
* request body
    * none
* response
    ```json
    {
        "message": "delete user",
        "users": []{
            "id": "id",
            "name": "name",
            "email": "email"
        }
    }
    ```
    * error case
        ```json
        {
            "message": "failed to delete user"
        }
        ```





