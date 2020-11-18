package main

import (
	"github.com/hlib-go/hunpay/ocwap"
	"testing"
	"time"
)

var pemPriKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDCZi6OTtT6hk/V
RcQoburI0HlyynhlrKrYV2h10Ls398DMMU43ePJNGeetrRUeue4sQidWc47lEf7b
H/2+o5ues4+AJD5f34ezCP3VBD4emdaO+5KiYDfcQ0GL3F2+SgbM7dNpViG5T5SH
Sbi3ilCeVEPXMrNgkygfgU2UNOiepo2C5ApG/g2dhjVp/A5CkCZS5aa5pep0ImJ2
IEdJlZFmPwm6nkQnSJ0/s/hloGa1jEt82mJ6stj0wA/8ck+esCFkeifItyVY7cD/
4OwxH2+0LmgJ8q/C/gb1PMHtfdD4+0D/IpUzJXZyJsvcc9wLbO4J8ivE8Dt5Iv5/
ofOtC+m7AgMBAAECggEBAJNLizEGqDdet3S4lQWx7THVTIBey1T2vMiJijviuUiR
78KIBWTgvm8PFs0wnRUX+lAMm/PUQUxuEzYDCmd9XfekxDFqxNwt6YsNYA8cVNko
5TqXgaaI0yqQx0Rq55i8TGTQOuTQf2MC2l6bzFs9cRJWdMTExMeDGN6uQZLvd+Zk
pc/QXhXuj/HG5p11aIHrVDQyG2AN88dn9wqVguOMOBkBZjDAQdF5SYTl9am9HOuH
Zq37qqsHJNaY+t/7V5RnNo3bJC4Yai9BKQnvMy5yaYLfzGzvxo+0DE9Mdyt7rULF
70NPlOw+admUCVNb6lvK544NGCYJtYkBP/YE1fvyJGECgYEA+2a0ex1UWJNnc9WC
fsmiHZDuOMCv7D/ygdYrTOU6Bodrs/FIFrQJZhXab626fFAPUszmJDLc0mv2BJSz
w8taNzAL6pbLspEXCwv0f6UIiCjQgHi7fu39OI69Lt78gKYp+Um65JepZnwrD9Jh
LdJAbZjr6fTgCMwhUyLojsljy2sCgYEAxfSKQ2SN8ji+/kRgNN7LqFgv51bb6wIy
Xs4DG0R563HZTdNobLKV0LsFumWilUh/ddze+VrdjuZ2klOXQ7qjJxwlL/RtuRdZ
H+EwDS9h84tsA2RiLHKQPnIASToYM8vqgVA7VklH0+P9399MM0+EDbLhgjPckGLo
2LJ1f3dgvvECgYEAw2sXeefXi66xKPJbqLCVisQA2S620T74BAL1z6UTkMWta8dv
UO6Lq8Xq8QqrPjyBXMyXTKYYzpxLNU5d2iF1NBbt/GFRX7G2psiZOquPUT/gqyyu
GkFmp5MU7Z05y8reL8AnPc/CRz7Xvkm7boHTwR2wrEDD8TKz6Mrm2S8kmpcCgYEA
vovBe1WTfRE6Z66RnNLI3ubkVZ66WeDnc3KCcwDbCtOwBMX2woq0wxMDVIT4lxIN
/vn4d7YLhr44bGmiNUO2QLNK6Ho0E/Jxi8pLYqW1d6VA95LtHTO5vSInPFV7boBe
3tLICyrGxSO1AIYE528nAbiqcZZSPXm4AL7ncycKLUECgYEAi9/E8bBLlOil+N+b
n3qHhCCSxPpjPln0RJsiy2LO1qf08nDW44dE4Ua0HYldoeCzSbqcw8wszBzu9HLV
gfh5Ku1srRHt/1uiI8PPwZL5GiWFTr01qbXKUyMDpauQ41VyDjWyixOub3uAqONu
Kb9w9k6wXTjk147jImhBinG4DL0=
-----END RSA PRIVATE KEY-----`

func TestRsaWithSha256Sign(t *testing.T) {
	sign, err := ocwap.RsaWithSha256Sign("123", pemPriKey)
	if err != nil {
		t.Error(err)
	}
	t.Log("sign:" + sign)
	//value=123 sign=ulrSoxOLyPzsDpVk6QCgUFKLVVQpLAWyob3cdcyWiwSh0I2QuoD/nIJyFuXpNUdR5R/OLTqTwOdFos0UmQ5TxZucm49zLHJyj3B3Toc8lT7tiJICyQW80KNkpTKUEeRmzwOz5hhI0tMnQBZhV95oslps4iOdDb0DiDXYeI641UCPTZA1CkwGcYIKA1nwMAa7R8ZdKS5fPgoJrzAjrwC/NNmVQniykthClUDsIf+TJMrUK1KnRb56LHGwbGw5Uyswk4PaQSwG4IGIjslfqtyxEO1UrQjEKkqIpI+X+wKj+TlSoTHM0WFAmPYYZBPdE3BqMN58GdFib7/gCRd5AXRwCw==
	//value=123 sign=ulrSoxOLyPzsDpVk6QCgUFKLVVQpLAWyob3cdcyWiwSh0I2QuoD/nIJyFuXpNUdR5R/OLTqTwOdFos0UmQ5TxZucm49zLHJyj3B3Toc8lT7tiJICyQW80KNkpTKUEeRmzwOz5hhI0tMnQBZhV95oslps4iOdDb0DiDXYeI641UCPTZA1CkwGcYIKA1nwMAa7R8ZdKS5fPgoJrzAjrwC/NNmVQniykthClUDsIf+TJMrUK1KnRb56LHGwbGw5Uyswk4PaQSwG4IGIjslfqtyxEO1UrQjEKkqIpI+X+wKj+TlSoTHM0WFAmPYYZBPdE3BqMN58GdFib7/gCRd5AXRwCw==
	//value=123 sign=ulrSoxOLyPzsDpVk6QCgUFKLVVQpLAWyob3cdcyWiwSh0I2QuoD/nIJyFuXpNUdR5R/OLTqTwOdFos0UmQ5TxZucm49zLHJyj3B3Toc8lT7tiJICyQW80KNkpTKUEeRmzwOz5hhI0tMnQBZhV95oslps4iOdDb0DiDXYeI641UCPTZA1CkwGcYIKA1nwMAa7R8ZdKS5fPgoJrzAjrwC/NNmVQniykthClUDsIf+TJMrUK1KnRb56LHGwbGw5Uyswk4PaQSwG4IGIjslfqtyxEO1UrQjEKkqIpI+X+wKj+TlSoTHM0WFAmPYYZBPdE3BqMN58GdFib7/gCRd5AXRwCw==
}

func TestRsaWithSha256Sign2(t *testing.T) {
	beg := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		_, err := ocwap.RsaWithSha256Sign("123", pemPriKey)
		if err != nil {
			t.Error(err)
		}
	}

	end := time.Now().UnixNano()
	t.Log((end - beg) / 1e6)
	//1000次，1357毫秒
}
