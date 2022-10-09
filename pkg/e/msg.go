package e

var MsgFlags = map[int]string{
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误,图片格式有问题",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
