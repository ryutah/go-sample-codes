package main

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSampleFoo(t *testing.T) {
	fooCtrl := gomock.NewController(t)
	defer fooCtrl.Finish()

	mockObj := NewMockFoo(fooCtrl)
	mockObj.EXPECT().Bar("In SampleFoo").Return("hoge fuga")
	if got := SampleFoo(mockObj); got != "hoge fuga" {
		t.Errorf("want %v, got %v", "hoge fuga", got)
	}
}
