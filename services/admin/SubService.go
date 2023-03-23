package admin

import (
	"gin/utils"
)

func Assembly(params map[string]any) (bool, any) {
	length := len(params)
	if length <= 0 {
		return false, "参数错误!"
	}
	key := utils.GetMapFirstKey(params)
	switch key {
	case "role":
		return RoleSelect()
	case "rule":
		return RuleSelect(params[key].(string))
	case "dictionary":
		return DictionarySelect()
	default:
		return false, "组件未注册!"
	}
}
