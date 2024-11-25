package main

var ThreadData = []BaseThreadResponse{
	{
		Uuid:      "1234",
		Topic:     "This is sample1.",
		CreatedAt: "2006-01-02 15:04:05",
	},
	{
		Uuid:      "abcd",
		Topic:     "This is sample2.",
		CreatedAt: "2006-01-02 15:04:05",
	},
	{
		Uuid:      "5678",
		Topic:     "This is sample3.",
		CreatedAt: "2006-01-02 15:04:05",
	},
	{
		Uuid:      "efgh",
		Topic:     "This is sample4.",
		CreatedAt: "2006-01-02 15:04:05",
	},
	{
		Uuid:      "9101",
		Topic:     "This is sample5.",
		CreatedAt: "2006-01-02 15:04:05",
	},
	{
		Uuid:      "ijkl",
		Topic:     "This is sample6.",
		CreatedAt: "2006-01-02 15:04:05",
	},
}

var PostData = []BasePostResponse{
	{
		Uuid:       "1234post",
		Body:       "Test post1",
		ThreadUuid: "abcd",
		CreatedAt:  "2006-01-02 15:04:05",
	},
	{
		Uuid:       "abddkk",
		Body:       "Test post2",
		ThreadUuid: "ijkl",
		CreatedAt:  "2006-01-02 15:04:05",
	},
	{
		Uuid:       "lladfa",
		Body:       "Test post3",
		ThreadUuid: "abcd",
		CreatedAt:  "2006-01-02 15:04:05",
	},
	{
		Uuid:       "1233post",
		Body:       "Test post4",
		ThreadUuid: "9101",
		CreatedAt:  "2006-01-02 15:04:05",
	},
}
