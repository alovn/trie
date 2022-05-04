package trie

import (
	"testing"
)

func Test_trie_Find(t *testing.T) {
	tr := New()
	tr.Add("abc")
	tr.Add("abcd")
	tr.Add("hello")

	tests := []struct {
		name string
		args string
		want bool
	}{
		{name: "find a", args: "a", want: false},
		{name: "find ab", args: "ab", want: false},
		{name: "find abc", args: "abc", want: true},
		{name: "find abcd", args: "abcd", want: true},
		{name: "find abcde", args: "abcde", want: false},
		{name: "find he", args: "he", want: false},
		{name: "find hello", args: "hello", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tr.Find(tt.args); got != tt.want {
				t.Errorf("trie.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trie_Remove(t *testing.T) {
	tr := New()
	tr.Add("ab")
	tr.Add("abc")
	tr.Add("abcd")
	tr.Add("hello")

	tr.Remove("ab")
	tr.Remove("abcd")
	tr.Remove("hello")

	tests := []struct {
		name string
		args string
		want bool
	}{
		{name: "find ab", args: "ab", want: false},
		{name: "find abc", args: "abc", want: true},
		{name: "find abcd", args: "abcd", want: false},
		{name: "find hello", args: "hello", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tr.Find(tt.args); got != tt.want {
				t.Errorf("trie.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
