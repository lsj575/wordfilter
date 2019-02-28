package models

type CheckResult struct {
	Ret int `json:"ret"`         // 1 存在敏感词 0 不存在
	Result string `json:"result"`
	Find []string `json:"find"`
}
