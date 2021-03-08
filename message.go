package weworklib

type ActionType = string

const (
	SEND_ACTION   ActionType = "send"
	RECALL_ACTION ActionType = "recall"
	SWITCH_ACTION ActionType = "switch"
)

type MessageType = string

const (
	TEXT_MSG               MessageType = "text"
	IMG_MSG                MessageType = "image"
	REVOKE_MSG             MessageType = "revoke"
	AGREE_MSG              MessageType = "agree"
	DISAGREE_MSG           MessageType = "disagree"
	VOICE_MSG              MessageType = "voice"
	VIDEO_MSG              MessageType = "video"
	CARD_MSG               MessageType = "card"
	LOC_MSG                MessageType = "location"
	EMOTION_MSG            MessageType = "emotion"
	FILE_MSG               MessageType = "file"
	LINK_MSG               MessageType = "link"
	WEAPP_MSG              MessageType = "weapp"
	CHATRECORD_MSG         MessageType = "chatrecord"
	TODO_MSG               MessageType = "todo"
	VOTE_MSG               MessageType = "vote"
	COLLECT_MSG            MessageType = "collect"
	REDPACKET_MSG          MessageType = "redpacket"
	MEETING_MSG            MessageType = "meeting"
	DOC_MSG                MessageType = "docmsg"
	MARKDOWN_MSG           MessageType = "markdown"
	NEWS_MSG               MessageType = "news"
	CALENDAR_MSG           MessageType = "calendar"
	MIXED_MSG              MessageType = "mixed"
	MEETING_VOICE_CALL_MSG MessageType = "meeting_voice_call"
	VOIP_DOC_SHARE_MSG     MessageType = "voip_doc_share"
)

type EmotionType = uint32

const (
	GIF EmotionType = 1
	PNG EmotionType = 2
)

type ChatRecordType = string

const (
	TEXT_REC_MSG    ChatRecordType = "ChatRecordText"
	IMG_REC_MSG     ChatRecordType = "ChatRecordImage"
	VOICE_REC_MSG   ChatRecordType = "ChatRecordLocationVoice"
	VIDEO_REC_MSG   ChatRecordType = "ChatRecordVideo"
	CARD_REC_MSG    ChatRecordType = "ChatRecordCard"
	LOC_REC_MSG     ChatRecordType = "ChatRecordLocation"
	EMOTION_REC_MSG ChatRecordType = "ChatRecordEmotion"
	FILE_REC_MSG    ChatRecordType = "ChatRecordFile"
	LINK_REC_MSG    ChatRecordType = "ChatRecordLink"
	WEAPP_REC_MSG   ChatRecordType = "ChatRecordWeapp"
	MIXED_REC_MSG   ChatRecordType = "ChatRecordMixed"
)

type VoteType = uint32

const (
	LAUNCH_VOTE VoteType = 101
	JOIN_VOTE   VoteType = 102
)

type CollectDetailsType = string

const (
	TEXT   CollectDetailsType = "Text"
	NUMBER CollectDetailsType = "Number"
	DATE   CollectDetailsType = "Date"
	TIME   CollectDetailsType = "Time"
)

type RedpacketType = uint32

const (
	NORMAL_REDPACKET    RedpacketType = 1
	SPELL_REDPACKET     RedpacketType = 2
	INCENTIVE_REDPACKET RedpacketType = 3
)

type MeetingType = uint32

const (
	LAUNCH_MEETING MeetingType = 101
	HANDLE_MEETING MeetingType = 102
)

type MeetingStatus = uint32

const (
	JOIN_MEETING     MeetingStatus = 1
	REJECT_MEETING   MeetingStatus = 2
	PENDING_MEETING  MeetingStatus = 3
	NOINVITE_MEETING MeetingStatus = 4
	CANCELED_MEETING MeetingStatus = 5
	EXPIRED_MEETING  MeetingStatus = 6
	ABSENT_MEETING   MeetingStatus = 7
)

type Message interface {
	ID() string
}

type BaseMessage struct {
	MsgId   string      `json:"msgid,omitempty"`   // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	Action  ActionType  `json:"action,omitempty"`  // 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	From    string      `json:"from,omitempty"`    // 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	ToList  []string    `json:"tolist,omitempty"`  // 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	RoomId  string      `json:"roomid,omitempty"`  // 群聊消息的群id。如果是单聊则为空。
	MsgTime int64       `json:"msgtime,omitempty"` // 消息发送时间戳，utc时间，ms单位。
	MsgType MessageType `json:"msgtype,omitempty"` // 文本消息为：text。
}

func (this BaseMessage) ID() string {
	return this.MsgId
}

type CommonMessage struct {
	BaseMessage
	Seq       uint64
	MediaData []byte
	Content   map[string]interface{}
}

type TextMessage struct {
	BaseMessage
	Content string `json:"content,omitempty"` // 消息内容。
}

type ImageMessage struct {
	BaseMessage
	SdkFileId string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	Md5Sum    string `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
	FileSize  uint32 `json:"filesize,omitempty"`  // 图片资源的文件大小。
}

type RevokeMessage struct {
	BaseMessage
	PreMsgId string `json:"pre_msgid,omitempty"` // 标识撤回的原消息的msgid
}

type AggreeMessage struct {
	BaseMessage
	UserId    string `json:"userid,omitempty"`     // 同意/不同意协议者的userid，外部企业默认为external_userid。
	AgreeTime int64  `json:"agree_time,omitempty"` // 同意/不同意协议的时间，utc时间，ms单位。
}

type VoiceMessage struct {
	BaseMessage
	SdkFileId  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
	VoiceSize  uint32 `json:"voice_size,omitempty"`  // 语音消息大小。
	PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
	Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
}

type VideoMessage struct {
	BaseMessage
	SdkFileId  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
	FileSize   uint32 `json:"filesize,omitempty"`    // 图片资源的文件大小。
	PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
	Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
}

type CardMessage struct {
	BaseMessage
	CorpName string `json:"corpname,omitempty"` // 名片所有者所在的公司名称。
	UserId   string `json:"userid,omitempty"`   // 名片所有者的id，同一公司是userid，不同公司是external_userid
}

type LocationMessage struct {
	BaseMessage
	Lng     float64 `json:"longitude,omitempty"` // 经度，单位double
	Lat     float64 `json:"latitude,omitempty"`  // 纬度，单位double
	Address string  `json:"address,omitempty"`   // 地址信息
	Title   string  `json:"title,omitempty"`     // 位置信息的title。
	Zoom    uint32  `json:"zoom,omitempty"`      // 缩放比例。
}

type EmotionMessage struct {
	BaseMessage
	Type      EmotionType `json:"type,omitempty"`      // 表情类型，png或者gif.1表示gif 2表示png。
	Width     uint32      `json:"width,omitempty"`     // 表情图片宽度。
	Height    uint32      `json:"height,omitempty"`    // 表情图片高度。
	ImageSize uint32      `json:"imagesize,omitempty"` // 资源的文件大小。
	SdkFileId string      `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	Md5Sum    string      `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
}

type FileMessage struct {
	BaseMessage
	FileName  string `json:"filename,omitempty"`  // 文件名称。
	FileExt   string `json:"fileext,omitempty"`   // 文件类型后缀。
	SdkFileId string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	FileSize  uint32 `json:"filesize,omitempty"`  // 文件大小。
	Md5Sum    string `json:"md5sum,omitempty"`    // 资源的md5值，供进行校验。
}

type LinkMessage struct {
	BaseMessage
	Title    string `json:"title,omitempty"`       // 消息标题。
	Desc     string `json:"description,omitempty"` // 消息描述。
	LinkUrl  string `json:"link_url,omitempty"`    // 链接url地址
	ImageUrl string `json:"image_url,omitempty"`   // 链接图片url。
}

type WeappMessage struct {
	BaseMessage
	Title       string `json:"title,omitempty"`       // 消息标题。
	Desc        string `json:"description,omitempty"` // 消息描述。
	Username    string `json:"username,omitempty"`    // 用户名称。
	DisplayName string `json:"displayname,omitempty"` // 小程序名称
}

type ChatRecordMessage struct {
	BaseMessage
	Title string       `json:"title,omitempty"` // 聊天记录标题
	Item  []ChatRecord `json:"item,omitempty"`  // 消息记录内的消息内容，批量数据
}

type TodoMessage struct {
	BaseMessage
	VoteTitle string   `json:"votetitle,omitempty"` // 投票主题。
	VoteItem  []string `json:"voteitem,omitempty"`  // 投票选项，可能多个内容。
	VoteType  VoteType `json:"votetype,omitempty"`  // 投票类型.101发起投票、102参与投票。
	VoteId    string   `json:"voteid,omitempty"`    // 投票id，方便将参与投票消息与发起投票消息进行前后对照。
}

type CollectMessage struct {
	BaseMessage
	RoomName   string           `json:"room_name,omitempty"`   // 填表消息所在的群名称。
	Creator    string           `json:"creator,omitempty"`     // 创建者在群中的名字
	CreateTime string           `json:"create_time,omitempty"` // 创建的时间
	Details    []CollectDetails `json:"details,omitempty"`     // 表内容
}

type RedpacketMessage struct {
	BaseMessage
	Type        RedpacketType `json:"type,omitempty"`        // 红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包。
	Wish        string        `json:"wish,omitempty"`        // 红包祝福语
	TotalCnt    uint32        `json:"totalcnt,omitempty"`    // 红包总个数
	TotalAmount uint32        `json:"totalamount,omitempty"` // 红包总金额。单位为分。
}

type MeetingMessage struct {
	BaseMessage
	Topic       string        `json:"topic,omitempty"`       // 会议主题
	StartTime   int64         `json:"starttime,omitempty"`   // 会议开始时间。Utc时间
	EndTime     int64         `json:"endtime,omitempty"`     // 会议结束时间。Utc时间
	Address     string        `json:"address,omitempty"`     // 会议地址
	Remarks     string        `json:"remarks,omitempty"`     // 会议备注
	MeetingType MeetingType   `json:"meetingtype,omitempty"` // 会议消息类型。101发起会议邀请消息、102处理会议邀请消息
	MeetingId   uint64        `json:"meetingid,omitempty"`   // 会议id。方便将发起、处理消息进行对照
	Status      MeetingStatus `json:"status,omitempty"`      // 会议邀请处理状态。1 参加会议、2 拒绝会议、3 待定、4 未被邀请、5 会议已取消、6 会议已过期、7 不在房间内。
}

type DocMessage struct {
	BaseMessage
	Title      string `json:"title,omitempty"`       // 在线文档名称
	LinkUrl    string `json:"link_url,omitempty"`    // 在线文档链接
	DocCreator string `json:"doc_creator,omitempty"` // 在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid
}

type MarkdownMessage struct {
	BaseMessage
	Content string `json:"content,omitempty"` // markdown消息内容，目前为机器人发出的消息
}

type NewsMessage struct {
	BaseMessage
	Info struct {
		Item []News `json:"item,omitempty"` // 图文消息数组
	} `json:"info,omitempty"` // 图文消息的内容
}

type CalendarMessage struct {
	BaseMessage
	Title        string   `json:"title,omitempty"`        // 日程主题
	CreatorName  string   `json:"creatorname,omitempty"`  // 日程组织者
	AttendeeName []string `json:"attendeename,omitempty"` // 日程参与人。数组，内容为String类型
	StartTime    int64    `json:"starttime,omitempty"`    // 日程开始时间。Utc时间，单位秒
	EndTime      int64    `json:"endtime,omitempty"`      // 日程结束时间。Utc时间，单位秒
	Place        string   `json:"place,omitempty"`        // 日程地点
	Remarks      string   `json:"remarks,omitempty"`      // 日程备注
}

type MixedMessage struct {
	BaseMessage
	Mixed struct {
		Item []MixedMsg `json:"item,omitempty"`
	} `json:"mixed,omitempty"` // 消息内容。可包含图片、文字、表情等多种消息。Object类型
}

type MeetingVoiceCallMessage struct {
	BaseMessage
	VoiceId          string            `json:"voiceid,omitempty"`            // 音频id
	MeetingVoiceCall *MeetingVoiceCall `json:"meeting_voice_call,omitempty"` // 音频消息内容。包括结束时间、fileid，可能包括多个demofiledata、sharescreendata消息，demofiledata表示文档共享信息，sharescreendata表示屏幕共享信息。Object类型
}

type VoipDocShareMessage struct {
	BaseMessage
	VoipId       string        `json:"voipid,omitempty"`         // 音频id
	VoipDocShare *VoipDocShare `json:"voip_doc_share,omitempty"` // 共享文档消息内容。包括filename、md5sum、filesize、sdkfileid字段。Object类型
}

type SwitchMessage struct {
	MsgId  string `json:"msgid,omitempty"`  // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重
	Action string `json:"action,omitempty"` // 消息动作，切换企业为switch
	Time   int64  `json:"time,omitempty"`   // 消息发送时间戳，utc时间，ms单位。
	User   string `json:"user,omitempty"`   // 具体为切换企业的成员的userid。
}

func (this SwitchMessage) ID() string {
	return this.MsgId
}

type ChatRecord struct {
	Type         ChatRecordType `json:"type,omitempty"`          // 每条聊天记录的具体消息类型：ChatRecordText/ ChatRecordFile/ ChatRecordImage/ ChatRecordVideo/ ChatRecordLink/ ChatRecordLocation/ ChatRecordMixed ….
	Content      string         `json:"content,omitempty"`       // 消息内容。Json串，内容为对应类型的json
	MsgTime      int64          `json:"msgtime,omitempty"`       // 消息时间，utc时间，ms单位。
	FromChatroom bool           `json:"from_chatroom,omitempty"` // 是否来自群会话。
}

type CollectDetails struct {
	Id   uint64             `json:"id,omitempty"`   // 表项id
	Ques string             `json:"ques,omitempty"` // 表项名称
	Type CollectDetailsType `json:"type,omitempty"` // 表项类型，有Text(文本),Number(数字),Date(日期),Time(时间)
}

type News struct {
	Title  string `json:"title,omitempty"`       // 图文消息标题
	Desc   string `json:"description,omitempty"` // 图文消息描述
	Url    string `json:"url,omitempty"`         // 图文消息点击跳转地址
	PicUrl string `json:"picurl,omitempty"`      // 图文消息配图的url
}

type MixedMsg struct {
	Type    MessageType `json:"type,omitempty"`
	Content string      `json:"content,omitempty"`
}

type MeetingVoiceCall struct {
	EndTime         int64             `json:"endtime,omitempty"`         // 音频结束时间
	SdkFileId       string            `json:"sdkfileid,omitempty"`       // 音频媒体下载的id
	DemoFileData    []DemoFileData    `json:"demofiledata,omitempty"`    // 文档分享对象，Object类型
	ShareScreenData []ShareScreenData `json:"sharescreendata,omitempty"` // 屏幕共享对象，Object类型
}

type DemoFileData struct {
	FileName     string `json:"filename,omitempty"`     // 文档共享名称
	DemoOperator string `json:"demooperator,omitempty"` // 文档共享操作用户的id
	StartTime    int64  `json:"starttime,omitempty"`    // 文档共享开始时间
	EndTime      int64  `json:"endtime,omitempty"`      // 文档共享结束时间
}

type ShareScreenData struct {
	Share     string `json:"share,omitempty"`     // 屏幕共享用户的id
	StartTime int64  `json:"starttime,omitempty"` // 屏幕共享开始时间
	EndTime   int64  `json:"endtime,omitempty"`   // 屏幕共享结束时间
}

type VoipDocShare struct {
	FileName  string `json:"filename,omitempty"`  // 文档共享文件名称
	Md5Sum    string `json:"md5sum,omitempty"`    // 共享文件的md5值
	FileSize  uint64 `json:"filesize,omitempty"`  // 共享文件的大小
	SdkFileId string `json:"sdkfileid,omitempty"` // 共享文件的sdkfile，通过此字段进行媒体数据下载
}
