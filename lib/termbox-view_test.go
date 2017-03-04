package thelm

import (
	"sync"
	"testing"
)

var (
	twoline = []byte("first\nsecond\n")
)

func TestUIView_nextLineOffset(t *testing.T) {

	type fields struct {
		buffer       []byte
		readpos      int
		lines        int
		sizeX        int
		sizeY        int
		offsetX      int
		offsetY      int
		highlightY   int
		statusLine   string
		inputLine    string
		inputCursorX int
		mutex        sync.Mutex
	}
	type args struct {
		start int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOffset int
	}{
		{"Empty buffer", fields{}, args{}, 0},

		{"One line = length of buffer", fields{buffer: []byte("abc")}, args{}, 3},

		{"Two lines", fields{buffer: []byte("a\nb\n")}, args{}, 2},
		{"Start in the middle",
			fields{buffer: []byte("aa\nbb\ncc\ndd")},
			args{4}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIView{
				buffer:       tt.fields.buffer,
				readpos:      tt.fields.readpos,
				lines:        tt.fields.lines,
				sizeX:        tt.fields.sizeX,
				sizeY:        tt.fields.sizeY,
				offsetX:      tt.fields.offsetX,
				offsetY:      tt.fields.offsetY,
				highlightY:   tt.fields.highlightY,
				statusLine:   tt.fields.statusLine,
				inputLine:    tt.fields.inputLine,
				inputCursorX: tt.fields.inputCursorX,
				mutex:        tt.fields.mutex,
			}
			if gotOffset := u.nextLineOffset(tt.args.start); gotOffset != tt.wantOffset {
				t.Errorf("UIView.nextLineOffset() = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}
