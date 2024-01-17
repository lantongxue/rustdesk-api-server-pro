package test

import (
	"rustdesk-api-server-pro/util"
	"testing"

	"github.com/golang-module/carbon/v2"
)

func TestTimeDiff(t *testing.T) {
	s := carbon.Now().DiffInSeconds(carbon.Parse("2024-01-02 15:44:30"))
	t.Logf("%d", s)
}

func TestRemoveArrayElement(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	n := util.RemoveElement(arr, "a")
	t.Log(n)
}
