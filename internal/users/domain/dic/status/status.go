package status

type Status int8

const (
	Enable  Status = iota + 1 // 正常
	Disable                   // 禁用
	Locked                    // 上锁
	Delete                    // 删除
)
