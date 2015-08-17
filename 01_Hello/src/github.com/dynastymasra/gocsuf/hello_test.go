package stringutil

import (
  "testing"
  "github.com/dynastymasra/gocsuf/stringutil"
)

type Words struct {
  in, want string
}

func TestReverse(t *testing.T) {
  testCase := []Words{
    {"Hello World!", "!dlroW olleH"}, {"Hello Golang", "gnaloG olleH"}, {"", ""},
  }
  for _, c := range testCase {
    got := stringutil.Reverse(c.in)
    if got != c.want {
      t.Errorf("stringutil.Reverse(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}
