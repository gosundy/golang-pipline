package main

import (
	"log"
	"net/http"
	"time"
)
type Pipline struct{
	middlewares []Middleware
}
type  Middleware func(handlerFunc http.HandlerFunc)http.HandlerFunc
func (pipline *Pipline)New(ms ...Middleware)*Pipline{
	pipline.middlewares=append(pipline.middlewares,ms...)
	return  pipline
}
func (pipline *Pipline)Pipe(ms ...Middleware)*Pipline  {
	pipline.middlewares=append(pipline.middlewares,ms...)
	return  pipline
}
func (pipline *Pipline)Process(h http.HandlerFunc)http.HandlerFunc{
	le:=len(pipline.middlewares)
	for i:=le-1;i>=0;i--{
		h=pipline.middlewares[i](h)
	}
	return h
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello,world!"))
	if f, ok := writer.(http.Flusher); ok {
		f.Flush()
	}
}
func main() {
	pipline:=&Pipline{}
	http.HandleFunc("/",pipline.New(Log).Process(Hello))
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}
func Log(handlerFunc http.HandlerFunc)http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		start:=time.Now()
		handlerFunc(writer,request)
		end:=time.Now()
		log.Println(end.Sub(start).String())
	}
}
