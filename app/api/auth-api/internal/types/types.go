// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}
