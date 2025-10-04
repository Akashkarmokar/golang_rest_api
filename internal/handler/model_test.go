package handler_test

import (
	"testing"

	"github.com/Akashkarmokar/go_rest_api/internal/handler"
)

func TestNewsPostReqBody_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		req         handler.NewsPostReqBody
		expectedErr bool
	}{
		{
			name:        "author empty",
			req:         handler.NewsPostReqBody{},
			expectedErr: true,
		},
		{
			name: "title empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
			},
			expectedErr: true,
		},
		{
			name: "summary empty",
			req: handler.NewsPostReqBody{
				Author: "test-author",
				Title:  "test-title",
			},
			expectedErr: true,
		},
		{
			name: "created at empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "invalid",
			},
			expectedErr: true,
		},
		// {
		// 	name: "content empty",
		// 	req: handler.NewsPostReqBody{
		// 		Author:  "test-author",
		// 		Title:   "test-title",
		// 		Summary: "test-summary",
		// 	},
		// 	expectedErr: true,
		// },
		{
			name: "source empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2025-04-07T05:13:27+00:00",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2025-04-07T05:13:27+00:00",
				Source:    "https://www.goolge.com",
			},
			expectedErr: true,
		},
		{
			name: "tags empty",
			req: handler.NewsPostReqBody{
				Author:    "test-author",
				Title:     "test-title",
				Summary:   "test-summary",
				CreatedAt: "2025-04-07T05:13:27+00:00",
				Source:    "https://www.goolge.com",
				Tags:      []string{"test-tag"},
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()

			if tc.expectedErr && err == nil {
				t.Fatal("expected error but got nil")
			}
			if !tc.expectedErr && err != nil {
				t.Fatal("expected nil but got error")
			}
		})
	}
}
