package enums

const (
    ParamsError = 1001
    DBError     = 1002
    SaveError   = 1003
    DeleteError = 1004
)

var resultCodeText = map[int]string{
    ParamsError: "请求参数错误",
    DBError:     "数据库错误",
    SaveError:   "保存失败",
    DeleteError: "删除失败",
}

func StatusText(code int) (string, bool) {
    message, ok := resultCodeText[code]
    return message, ok
}
