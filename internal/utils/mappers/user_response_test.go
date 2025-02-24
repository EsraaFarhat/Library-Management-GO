package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
	"testing"
)

func TestMapUserToResponse(t *testing.T) {
	// Define a test case
	testCases := []struct {
		name     string
		input    *models.User
		expected dto.UserResponse
	}{
		{
			name: "Valid User Model",
			input: &models.User{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Role:  "user",
			},
			expected: dto.UserResponse{
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Role:  "user",
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			result := MapUserToResponse(tc.input)

			// Compare the result with the expected output
			if result.ID != tc.expected.ID ||
				result.Name != tc.expected.Name ||
				result.Email != tc.expected.Email ||
				result.Role != tc.expected.Role ||
				!result.CreatedAt.Equal(tc.expected.CreatedAt) {
				t.Errorf("Test case %s failed: expected %+v, got %+v", tc.name, tc.expected, result)
			}
		})
	}
}
