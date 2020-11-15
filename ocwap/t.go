package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"sync"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func uniqueId() string {
	b := make([]byte, 48)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func UUIDRand() string {
	var b = make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

func main() {
	var m sync.Map

	go ran("n1", m)
	go ran("n2", m)
	go ran("n3", m)
	go ran("n4", m)
	go ran("n5", m)
	go ran("n6", m)
	go ran("n7", m)
	go ran("n8", m)
	go ran("n9", m)
	go ran("n10", m)

	for i := 0; i < 10000000; i++ {
		a := UUIDRand()
		if _, ok := m.Load(a); !ok {
			m.Store(a, i)
		} else {
			fmt.Println("出现重复了: i = ", strconv.Itoa(i))
			break
		}
	}
	fmt.Println("完成...")
}

func ran(name string, m sync.Map) {
	for i := 0; i < 1000000; i++ {
		a := UUIDRand()
		if _, ok := m.Load(a); !ok {
			m.Store(a, i)
		} else {
			fmt.Println(name, "出现重复了: i = ", strconv.Itoa(i))
			break
		}
	}
	fmt.Println(name, " 完成...")
}
