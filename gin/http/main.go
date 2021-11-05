package main

import "net/http"

func main()  {
	http.HandleFunc("/", func(re http.ResponseWriter, rs *http.Request) {
		re.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080",nil)

}
