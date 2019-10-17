package main

import "testing"

func TestLoadConfig(t *testing.T) {
	t.Log(loadConfig())
}

func TestGetData(t *testing.T) {
	t.Log(getData())
}

func TestAddCer(t *testing.T) {
	t.Log(addCer(12, "lvcshu.info", "fullchain", "key"))
}

func TestAddService(t *testing.T) {
	if addService(":", "enable string", "disablestring") == false {
		t.Log("success")
	} else {
		t.Error(addService(":", "enable string", "disablestring"))
	}
}

func TestDelCer(t *testing.T) {
	t.Log(delCer(12))
}
