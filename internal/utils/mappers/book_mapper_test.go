package mappers

import (
	"library-management/internal/dto"
	"library-management/internal/models"
	"testing"
	"time"
)

func TestMapCreateRequestToBook(t *testing.T) {
	// Define a test case
	testCases := []struct {
		name     string
		input    dto.BookCreateRequest
		expected *models.Book
	}{
		{
			name: "Valid BookCreateRequest",
			input: dto.BookCreateRequest{
				Title:           "The Go Programming Language",
				Author:          "Alan A. A. Donovan",
				ISBN:            "978-0134190440",
				CopiesAvailable: 5,
				PublishedAt:     time.Date(2015, 11, 1, 0, 0, 0, 0, time.UTC),
			},
			expected: &models.Book{
				Title:           "The Go Programming Language",
				Author:          "Alan A. A. Donovan",
				ISBN:            "978-0134190440",
				CopiesAvailable: 5,
				PublishedAt:     time.Date(2015, 11, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			result := MapCreateRequestToBook(tc.input)

			// Compare the result with the expected output
			if result.Title != tc.expected.Title ||
				result.Author != tc.expected.Author ||
				result.ISBN != tc.expected.ISBN ||
				result.CopiesAvailable != tc.expected.CopiesAvailable ||
				!result.PublishedAt.Equal(tc.expected.PublishedAt) {
				t.Errorf("Test case %s failed: expected %+v, got %+v", tc.name, tc.expected, result)
			}
		})
	}
}
