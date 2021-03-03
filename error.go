package weworklib

import (
	"fmt"
)

var errorMap = map[int]string{
	10000: "请求参数错误  检查Init接口corpid、secret参数；检查GetChatData接口limit参数是否未填或大于1000；检查GetMediaData接口sdkfileid是否为空，indexbuf是否正常",
	10001: "网络请求错误  检查是否网络有异常、波动；检查使用代理的情况下代理参数是否设置正确的用户名与密码",
	10002: "数据解析失败	建议重试请求。若仍失败，可以反馈给企业微信进行查询，请提供sdk接口参数与调用时间点等信息",
	10003: "系统调用失败	GetMediaData调用失败，建议重试请求。若仍失败，可以反馈给企业微信进行查询，请提供sdk接口参数与调用时间点等信息",
	10005: "fileid错误	检查在GetMediaData接口传入的sdkfileid是否正确",
	10006: "解密失败	请检查是否先进行base64decode再进行rsa私钥解密，再进行DecryptMsg调用",
	10007: "已废弃	目前不会返回此错误码",
	10008: "DecryptMsg错误	建议重试请求。若仍失败，可以反馈给企业微信进行查询，请提供sdk接口参数与调用时间点等信息",
	10009: "ip非法	请检查sdk访问外网的ip是否与管理端设置的可信ip匹配，若不匹配会返回此错误码",
	10010: "请求的数据过期	用户欲拉取的数据已过期，仅支持近3天内的数据拉取",
	10011: "ssl证书错误	使用openssl版本sdk，校验ssl证书失败",
}

type Error struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func (this Error) Error() string {
	return fmt.Sprintf("%d:%s", this.ErrCode, this.ErrMsg)
}

func NewSDKErr(code int) Error {
	return Error{ErrCode: code, ErrMsg: errorMap[code]}
}
