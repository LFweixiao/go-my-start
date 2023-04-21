package model

type SysFile struct {
	Name     string `json:"name"`    //文件名
	Postfix  string `json:"postfix"` //后缀
	Year     string `json:"year"`
	Month    string `json:"month"`
	UUIDName string `json:"UUIDName"` //uudi 文件名
}
