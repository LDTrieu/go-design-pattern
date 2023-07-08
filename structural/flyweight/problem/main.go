package main

type ChatMessage struct {
	Content      string
	SenderName   string
	SenderAvatar []byte
}

func main() {

	// fmt.Println([]ChatMessage{
	// 	{
	// 		Content:      "hi",
	// 		SenderName:   "Peter",
	// 		SenderAvatar: make([]byte, 1024*300), // 300kb
	// 	},
	// 	{
	// 		Content:      "oh here you are",
	// 		SenderName:   "Mary",
	// 		SenderAvatar: make([]byte, 1024*400), // 400kb
	// 	},
	// 	{
	// 		Content:      "how are you doing?",
	// 		SenderName:   "Peter",
	// 		SenderAvatar: make([]byte, 1024*300), // 300kb
	// 	},
	// 	{
	// 		Content:      "I'm doing well?",
	// 		SenderName:   "Mary",
	// 		SenderAvatar: make([]byte, 1024*400), // 400kb
	// 	},
	// })

}
