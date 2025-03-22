package postgres

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("testing...")
}

func TestConnect(t *testing.T) {
	_, err := Connect(
		WithAuth("postgres", "123456"),
		WithDatabase("zhenjian"),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success!!!")
}
