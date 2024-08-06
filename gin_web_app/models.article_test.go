package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check that the length of the list of articles returned is the same as the length of the global variable holding the list of articles
	// same as the length of the global variable holding the list of articles

	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for idx, val := range alist {
		if val.Content != articleList[idx].Content || val.ID != articleList[idx].ID || val.Title != articleList[idx].Title {
			t.Fail()
			break
		}
	}
}
