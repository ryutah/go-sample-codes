package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMockReader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockReader := NewMockMyReader(ctrl)
	extStr := "test"
	mockReader.EXPECT().Read(gomock.Any()).Do(func(p []byte) {
		r := strings.NewReader(extStr)
		r.Read(p)
	}).Return(len(extStr), io.EOF).MinTimes(1)

	got, err := ioutil.ReadAll(mockReader)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != extStr {
		t.Errorf("want : %v, got : %v", extStr, string(got))
	}
}

func TestMockWrite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWriter := NewMockMyWriter(ctrl)
	var w bytes.Buffer
	extStr := "test"
	mockWriter.EXPECT().Write(gomock.Any()).Do(func(p []byte) {
		w.Write(p)
	}).Return(len(extStr), nil).MinTimes(1)

	buf := bufio.NewWriter(mockWriter)
	if _, err := buf.WriteString(extStr); err != nil {
		t.Fatal(err)
	}
	buf.Flush()

	if got := w.String(); got != extStr {
		t.Errorf("want : %v, got : %v", extStr, got)
	}
}
