package appTypes

import "encoding/json"

type Storage int

const (
	Local Storage = iota
	Qiniu
)

func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Storage) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*s = ToStorage(str)
	return nil
}

// String 返回存储类型的字符串表示
func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	default:
		str = "未知存储"
	}
	return str
}

// ToStorage 将字符串转换为 Storage 类型
func ToStorage(str string) Storage {
	switch str {
	case "本地":
		return Local
	case "七牛云":
		return Qiniu
	default:
		return -1
	}
}
