package model

import "testing"

func TestBuild(t *testing.T) {
	b := NewSwagger()
	b.Info(NewInfo().Title("test")).Host("test.host.com")
	t.Logf("%#v", b.Build())
}
