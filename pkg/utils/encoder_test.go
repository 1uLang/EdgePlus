// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package utils

import (
	"github.com/iwind/TeaGo/maps"
	"testing"
)

func TestEncodeMap(t *testing.T) {
	{
		t.Log(Encode([]byte("123")))
	}
	{
		s, err := EncodeMap(maps.Map{"a": 1})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(s)

		t.Log(Decode([]byte(s)))
	}
}
