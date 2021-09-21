package main

import "testing"

// 版本一
//func TestHello(t *testing.T) {
//	got := Hello("World")
//	want := "Hello, World"
//
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}

// 版本二enecccbclvldjcbfnhgrjlfhhleeneuttrbfiinfirlk

func TestHello(t *testing.T) {
	// 子测试函数, 减少代码重复
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper()告诉测试套件这个方法是辅助函数。
		// 当测试失败的时候，报错的行号在函数调用处，而不是辅助函数处
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Maolin", "")
		want := "Hello, Maolin"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string to 'World", func(t *testing.T) {
		got := Hello("World", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
