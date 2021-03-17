package weworklib

// #cgo LDFLAGS: -L${SRCDIR}/lib -lWeWorkFinanceSdk_C
// #cgo CFLAGS: -I ./lib/
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/wenzhenxi/gorsa"
)

type Client struct {
	ptr         *C.WeWorkFinanceSdk_t
	privateKeys map[uint32]string
}

/**
 * 初始化函数
 *
 * @param [in]  corpid      调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
 * @param [in]  secret      聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
 *
 */
func NewClient(corpId string, corpSecret string, privateKeys map[uint32]string) (*Client, error) {
	ptr := C.NewSdk()
	corpIdC := C.CString(corpId)
	corpSecretC := C.CString(corpSecret)
	defer func() {
		C.free(unsafe.Pointer(corpIdC))
		C.free(unsafe.Pointer(corpSecretC))
	}()
	retC := C.Init(ptr, corpIdC, corpSecretC)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	return &Client{
		ptr:         ptr,
		privateKeys: privateKeys,
	}, nil
}

func (this *Client) Free() {
	C.DestroySdk(this.ptr)
}

func (this *Client) GetChatList(seq uint64, limit uint64, proxy string, password string, timeout int) (messages []CommonMessage, err error) {
	chatDatas, err := this.GetChatData(seq, limit, proxy, password, timeout)
	if err != nil {
		return
	}
	messages = make([]CommonMessage, 0)
	for _, chatData := range chatDatas {
		message, decryptErr := this.DecryptData(chatData.PublickeyVer, chatData.EncryptRandomKey, chatData.EncryptChatMsg)
		if decryptErr != nil {
			err = fmt.Errorf("解析消息内容失败：%v", decryptErr)
			return
		}
		if v, exists := message.Content["sdkfileid"]; exists {
			mediaData, getMediaErr := this.GetMediaData("", v.(string), proxy, password, timeout)
			if getMediaErr != nil {
				err = fmt.Errorf("获取图片资源文件失败：%v", getMediaErr)
				return
			}
			message.Seq = chatData.Seq
			message.MediaData = mediaData.Data
		}
		messages = append(messages, *message)
	}
	return
}

/**
* 拉取聊天记录函数
*
*
* @param [in]  seq             从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
* @param [in]  limit           一次拉取的消息条数，最大值1000条，超过1000条会返回错误
* @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
* @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
* @param [in]  timeout         超时时间，单位秒
* @return chatDatas       返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：

{"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}
*/
func (this *Client) GetChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) ([]ChatData, error) {
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	chatSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeSlice(chatSlice)
	}()

	retC := C.GetChatData(this.ptr, C.ulonglong(seq), C.uint(limit), proxyC, passwdC, C.int(timeout), chatSlice)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	buf := this.GetContentFromSlice(chatSlice)
	var data ChatDataResponse
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	if data.IsError() {
		return nil, fmt.Errorf("get message from GetChatData failed: %s", data.ErrMsg)
	}
	return data.ChatDataList, nil
}

/**
* @brief 解析密文.企业微信自有解密内容
* @param [in]  encrypt_key, getchatdata返回的encrypt_random_key,使用企业自持对应版本秘钥RSA解密后的内容
* @param [in]  encrypt_msg, getchatdata返回的encrypt_chat_msg
* @return msg, 解密的消息明文
 */
func (this *Client) DecryptData(keyVersion uint32, encryptKey string, encryptMsg string) (message *CommonMessage, err error) {
	message = new(CommonMessage)
	decodeKey, err := gorsa.PriKeyDecrypt(encryptKey, this.privateKeys[keyVersion])
	if err != nil {
		return nil, fmt.Errorf("decode encryptKey(%s) failed: %v", encryptKey, err)
	}
	encryptKeyC := C.CString(string(decodeKey))
	encryptMsgC := C.CString(encryptMsg)
	msgSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(encryptKeyC))
		C.free(unsafe.Pointer(encryptMsgC))
		C.FreeSlice(msgSlice)
	}()
	retC := C.DecryptData(encryptKeyC, encryptMsgC, msgSlice)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	buf := this.GetContentFromSlice(msgSlice)
	err = json.Unmarshal(buf, message)
	var messageContentMap = make(map[string]interface{})
	err = json.Unmarshal(buf, &messageContentMap)
	if v, has := messageContentMap[string(message.MsgType)]; has {
		message.Content = v.(map[string]interface{})
	}
	return message, err
}

/**
 * 拉取媒体消息函数
 * Return值=0表示该API调用成功
 *
 *
 * @param [in]  sdkFileid       从GetChatData返回的聊天消息中，媒体消息包括的sdkfileid
 * @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
 * @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
 * @param [in]  indexbuf        媒体消息分片拉取，需要填入每次拉取的索引信息。首次不需要填写，默认拉取512k，后续每次调用只需要将上次调用返回的outindexbuf填入即可。
 * @param [in]  timeout         超时时间，单位秒
 * @return media_data      返回本次拉取的媒体数据.MediaData结构体.内容包括data(数据内容)/outindexbuf(下次索引)/is_finish(拉取完成标记)
 */

func (this *Client) GetMediaData(indexBuf string, sdkFileId string, proxy string, passwd string, timeout int) (*MediaData, error) {
	indexBufC := C.CString(indexBuf)
	sdkFileIdC := C.CString(sdkFileId)
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	mediaDataC := C.NewMediaData()
	defer func() {
		C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIdC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()

	retC := C.GetMediaData(this.ptr, indexBufC, sdkFileIdC, proxyC, passwdC, C.int(timeout), mediaDataC)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret)
	}
	return &MediaData{
		OutIndexBuf: C.GoString(C.GetOutIndexBuf(mediaDataC)),
		Data:        C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.GetDataLen(mediaDataC)),
		IsFinish:    int(C.IsMediaDataFinish(mediaDataC)) == 1,
	}, nil
}

func (this *Client) GetContentFromSlice(slice *C.struct_Slice_t) []byte {
	return C.GoBytes(unsafe.Pointer(C.GetContentFromSlice(slice)), C.GetSliceLen(slice))
}
