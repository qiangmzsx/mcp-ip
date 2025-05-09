package ip

import "testing"

func Test_IP2RegionService(t *testing.T) {
	service, err := NewIP2RegionService("../data/ip2region.xdb")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := service.Searcher.SearchByStr("123.129.23.29")
	xdbLocation := XDB2Location(result)
	if xdbLocation.Country != "中国" {
		t.Fatal("error")
	}
}

func Test_QQwryService(t *testing.T) {
	service, err := NewQQWRYService("../data/qqwry.dat")
	if err != nil {
		t.Error(err)
		return
	}
	result, err := service.Searcher.Find("123.129.23.29")
	xdbLocation := XDB2Location(result)
	if xdbLocation.Country != "中国" {
		t.Fatal("error")
	}
}
