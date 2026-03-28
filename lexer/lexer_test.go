package lexer

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New("abc")
	if l.ch != 'a' {
		t.Errorf("New: ch = %q, want 'a'", l.ch)
	}
	if l.pos != 0 {
		t.Errorf("New: pos = %d, want 0", l.pos)
	}
	if l.readPos != 1 {
		t.Errorf("New: readPos = %d, want 1", l.readPos)
	}
}

func TestNew_Empty(t *testing.T) {
	l := New("")
	if l.ch != 0 {
		t.Errorf("New empty: ch = %q, want 0", l.ch)
	}
}

func TestReadChar(t *testing.T) {
	l := New("ab")
	// New already called readChar once, ch == 'a'
	l.readChar()
	if l.ch != 'b' {
		t.Errorf("readChar: ch = %q, want 'b'", l.ch)
	}
	l.readChar()
	if l.ch != 0 {
		t.Errorf("readChar at EOF: ch = %q, want 0", l.ch)
	}
}

func TestReadChar_Multibyte(t *testing.T) {
	l := New("あ")
	if l.ch != 'あ' {
		t.Errorf("readChar multibyte: ch = %q, want 'あ'", l.ch)
	}
}

func TestPeekChar(t *testing.T) {
	l := New("ab")
	// ch == 'a', peek should return 'b
	if got := l.peekChar(); got != 'b' {
		t.Errorf("peekChar: got %q, want 'b'", got)
	}
	// pos should not advance
	if l.ch != 'a' {
		t.Errorf("peekChar must not advance: ch = %q, want 'a'", l.ch)
	}
}

func TestPeekChar_AtEOF(t *testing.T) {
	l := New("a")
	// ch == 'a', no next char
	if got := l.peekChar(); got != 0 {
		t.Errorf("peekChar at EOF: got %q, want 0", got)
	}
}

func TestSkipWhitespace(t *testing.T) {
	tests := []struct {
		input  string
		wantCh rune
	}{
		{"   x", 'x'},
		{"\t\tx", 'x'},
		{"\n\nx", 'x'},
		{"\r\nx", 'x'},
		{"x", 'x'},
	}
	for _, tt := range tests {
		l := New(tt.input)
		l.skipWhitespace()
		if l.ch != tt.wantCh {
			t.Errorf("skipWhitespace(%q): ch = %q, want %q", tt.input, l.ch, tt.wantCh)
		}
	}
}

func TestSkipWhitespace_OnlySpaces(t *testing.T) {
	l := New("   ")
	l.skipWhitespace()
	if l.ch != 0 {
		t.Errorf("skipWhitespace all spaces: ch = %q, want 0", l.ch)
	}
}

func TestSkipLineComment(t *testing.T) {
	// skipLineComment は "//" の直後から呼ばれる想定
	// 改行まで読み飛ばし、改行の次の文字に進む
	l := New("this is a comment\nnext")
	l.skipLineComment()
	if l.ch != 'n' {
		t.Errorf("skipLineComment: ch = %q, want 'n'", l.ch)
	}
}

func TestSkipLineComment_NoNewline(t *testing.T) {
	// 改行なしでEOFに達した場合
	l := New("comment without newline")
	l.skipLineComment()
	if l.ch != 0 {
		t.Errorf("skipLineComment no newline: ch = %q, want 0", l.ch)
	}
}
