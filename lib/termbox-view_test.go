package thelm

import (
	"reflect"
	"testing"
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
			}
			if gotOffset := u.nextLineOffset(tt.args.start); gotOffset != tt.wantOffset {
				t.Errorf("UIView.nextLineOffset() = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}

func TestUIView_lineToByteOffset(t *testing.T) {
	twoline := []byte("a\nb\n")

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
	}
	type args struct {
		inputLine int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOffset int
	}{
		{"Empty buffer", fields{}, args{}, 0},
		{"One line", fields{buffer: []byte("abc")}, args{1}, 3},
		{"Two lines", fields{buffer: twoline}, args{1}, 2},
		{"Last line", fields{buffer: twoline}, args{100}, len(twoline)},
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
			}
			if gotOffset := u.lineToByteOffset(tt.args.inputLine); gotOffset != tt.wantOffset {
				t.Errorf("UIView.lineToByteOffset() = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}

func TestUIView_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name       string
		u          UIView
		args       args
		wantN      int
		wantErr    bool
		wantBuffer []byte
		wantLines  int
	}{
		{"Empty to empty", UIView{buffer: []byte{}}, args{}, 0, false, []byte{}, 0},
		{"Empty to some", UIView{buffer: []byte("a")}, args{}, 0, false, []byte("a"), 0},
		{"Some to some", UIView{buffer: []byte("a")}, args{[]byte("b")}, 1, false, []byte("ab"), 0},
		{"One line", UIView{buffer: []byte("a")}, args{[]byte("b\n")}, 2, false, []byte("ab\n"), 1},
		{"Many lines", UIView{buffer: []byte("a\nb\n"), lines: 2}, args{[]byte("c\n")}, 2, false, []byte("a\nb\nc\n"), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &tt.u
			gotN, err := u.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("UIView.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("UIView.Write() = %v, want %v", gotN, tt.wantN)
			}
			if !reflect.DeepEqual(u.buffer, tt.wantBuffer) {
				t.Errorf("UIView.buffer = %v, want %v", u.buffer, tt.wantBuffer)
			}
			if u.lines != tt.wantLines {
				t.Errorf("UIView.lines = %v, want %v", u.lines, tt.wantLines)
			}
		})
	}
}
