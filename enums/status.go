package enums

const (
	ParamsError  = 1001
	DBError      = 1002
	SaveError    = 1003
	DeleteError  = 1004
	NotFound     = 1005
	Unauthorized = 401
	Forbidden    = 403
)

var resultCodeText = map[int]string{
	ParamsError:  "请求参数错误",
	DBError:      "数据库错误",
	SaveError:    "保存失败",
	DeleteError:  "删除失败",
	NotFound:     "数据没找到",
	Unauthorized: "请登录",
	Forbidden:    "你没有权限",
}

func StatusText(code int) (string, bool) {
	message, ok := resultCodeText[code]
	return message, ok
}

func ForceStatusText(code int) string {
	message, ok := resultCodeText[code]
	if !ok {
		return "未知错误"
	}

	return message
}
