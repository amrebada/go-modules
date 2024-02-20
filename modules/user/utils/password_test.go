package utils

import "testing"

func TestPasswordUtils(t *testing.T) {

	t.Run("HashPassword showed hash Password", func(t *testing.T) {
		password := "password"
		hashedPassword := HashPassword(password)
		if hashedPassword == password {
			t.Errorf("HashPassword() = %v, want %v", hashedPassword, password)
		}
	})

	t.Run("ComparePassword should success with correct password", func(t *testing.T) {
		password := "password"
		hashedPassword := HashPassword(password)
		if !ComparePassword(hashedPassword, password) {
			t.Errorf("ComparePassword() = %v, want %v", false, true)
		}
	})

	t.Run("ComparePassword should fail with wrong password", func(t *testing.T) {
		password := "password"
		hashedPassword := HashPassword(password)
		if ComparePassword(hashedPassword, "wrongpassword") {
			t.Errorf("ComparePassword() = %v, want %v", true, false)
		}
	})
}
