package db

import (
	"context"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	// 验证 TestMain 是否正确初始化了 testQueries
	if testQueries == nil {
		t.Fatal("TestMain failed to initialize testQueries - TestMain was not executed")
	}

	// 尝试使用 testQueries 执行一个简单的数据库操作
	// 创建一个测试账户
	account, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    "Test User2",
		Balance:  1000,
		Currency: "USD",
	})

	if err != nil {
		t.Fatalf("Failed to create account: %v", err)
	}

	if account.Owner != "Test User" {
		t.Errorf("Expected owner to be 'Test User', got '%s'", account.Owner)
	}

	if account.Balance != 1000 {
		t.Errorf("Expected balance to be 1000, got %d", account.Balance)
	}

	if account.Currency != "USD" {
		t.Errorf("Expected currency to be 'USD', got '%s'", account.Currency)
	}

	// 清理创建的账户
	// err = testQueries.DeleteAccount(context.Background(), account.ID)
	// if err != nil {
	// 	t.Errorf("Failed to delete account: %v", err)
	// }
}
