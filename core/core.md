### 1. "用户登录"

1. route definition

- Url: /user/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginResponse`

2. request definition



```golang
type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResponse struct {
	Token string `json:"token"`
}
```

### 2. "用户详情"

1. route definition

- Url: /user/info
- Method: GET
- Request: `InfoRequest`
- Response: `InfoResponse`

2. request definition



```golang
type InfoRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type InfoResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
```

