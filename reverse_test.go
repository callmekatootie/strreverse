package strreverse

import "testing"

func TestReverse(t *testing.T) {
    original := []string{"Hello World!", "This is the best day of my life", "It's a beautiful world"}
    reverse := []string{"!dlroW olleH", "efil ym fo yad tseb eht si sihT", "dlrow lufituaeb a s'tI"}

    for i, s := range original {
        if r := Reverse(s); r != reverse[i] {
            t.Errorf("Reverse of %s is not %s, it is %s", s, r, reverse[i])
        }
    }
}
