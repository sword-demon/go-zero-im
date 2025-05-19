package xerr

var codeText = map[int]string{
	ServerCommonError: "服务器异常，请稍好再试~",
	RequestParamError: "请求参数有误",
	DbError:           "数据库繁忙，稍后再尝试",
}

func ErrMsg(errCode int) string {
	if msg, ok := codeText[errCode]; ok {
		return msg
	}

	// 没找到对应的code就返回默认的错误信息
	return codeText[ServerCommonError]
}
