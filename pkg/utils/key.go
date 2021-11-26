// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils

import timeutil "github.com/iwind/TeaGo/utils/time"

type Key struct {
	DayFrom      string   `json:"dayFrom"`
	DayTo        string   `json:"dayTo"`
	MacAddresses []string `json:"macAddresses"`
	Hostname     string   `json:"hostname"`
	Company      string   `json:"company"`
	Nodes        int      `json:"nodes"`
}

func (this *Key) IsValid() bool {
	return this.DayTo >= timeutil.Format("Y-m-d")
}
