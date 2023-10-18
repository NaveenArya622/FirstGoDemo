package main

import (
	"FIRSTSERVER/ACCOUNTS"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signUp", func(writer http.ResponseWriter, request *http.Request) {
		methodType := request.Method
		switch methodType {
		case "POST":
			{
				var user ACCOUNTS.User
				err := json.NewDecoder(request.Body).Decode(&user)
				if err != nil {
					fmt.Printf("error occured while decoding")
				}
				code, body := user.AddUser()
				writer.WriteHeader(code)
				er := json.NewEncoder(writer).Encode(body)
				if er != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		default:
			{
				writer.WriteHeader(404)
				err := json.NewEncoder(writer).Encode(ACCOUNTS.Return{Message: "Invalid request"})
				if err != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		}
	})
	http.HandleFunc("/signIn", func(writer http.ResponseWriter, request *http.Request) {
		methodType := request.Method
		switch methodType {
		case "POST":
			{
				var user ACCOUNTS.User
				err := json.NewDecoder(request.Body).Decode(&user)
				if err != nil {
					fmt.Printf("error occured while decoding")
				}
				agent := request.Header.Get("User-Agent")
				code, body := user.CheckUser(agent)
				writer.WriteHeader(code)
				er := json.NewEncoder(writer).Encode(body)
				if er != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		default:
			{
				writer.WriteHeader(404)
				fmt.Fprintf(writer, "Rout not found!")
			}
		}
	})
	http.HandleFunc("/user/details", func(writer http.ResponseWriter, request *http.Request) {
		methodType := request.Method
		switch methodType {
		case "GET":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				writer.WriteHeader(code)
				er := json.NewEncoder(writer).Encode(body)
				if er != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		case "PUT":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					user = *body.Details
					err := json.NewDecoder(request.Body).Decode(&user)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body = user.UpdateUser()
				}
				writer.WriteHeader(code)
				er := json.NewEncoder(writer).Encode(body)
				if er != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		case "DELETE":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					user = *body.Details
					err := json.NewDecoder(request.Body).Decode(&user)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body = user.RemoveUser()
				}
				writer.WriteHeader(code)
				er := json.NewEncoder(writer).Encode(body)
				if er != nil {
					fmt.Printf("error occured while decoding")
				}
			}
		default:
			{
				writer.WriteHeader(404)
				fmt.Fprintf(writer, "Rout not found!")
			}
		}
	})
	http.HandleFunc("/user/tasks", func(writer http.ResponseWriter, request *http.Request) {
		methodType := request.Method
		switch methodType {
		case "POST":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					var task ACCOUNTS.Tasks
					user = *body.Details
					task.AssiniedTo = body.Details.Id
					err := json.NewDecoder(request.Body).Decode(&task)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body := task.CreateTask()
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				} else {
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				}
			}
		case "GET":
			{
				var user ACCOUNTS.User
				var task ACCOUNTS.Tasks
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					user = *body.Details
					task.AssiniedTo = body.Details.Id
					err := json.NewDecoder(request.Body).Decode(&task)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body := task.GetTasks()
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				} else {
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				}
			}
		case "PUT":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					var task ACCOUNTS.Tasks
					user = *body.Details
					task.AssiniedTo = body.Details.Id
					err := json.NewDecoder(request.Body).Decode(&task)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body := task.UpdateTask()
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				} else {
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				}
			}
		case "PATCH":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					var task ACCOUNTS.Tasks
					user = *body.Details
					task.AssiniedTo = body.Details.Id
					err := json.NewDecoder(request.Body).Decode(&task)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body := task.CompletedTask()
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				} else {
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				}
			}
		case "DELETE":
			{
				var user ACCOUNTS.User
				agent := request.Header.Get("User-Agent")
				token := request.Header.Get("Authorization")
				code, body := user.GetDetail(agent, token[7:])
				if code == 200 {
					var task ACCOUNTS.Tasks
					user = *body.Details
					task.AssiniedTo = body.Details.Id
					err := json.NewDecoder(request.Body).Decode(&task)
					if err != nil {
						fmt.Printf("error occured while decoding")
					}
					code, body := task.RemoveTask()
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				} else {
					writer.WriteHeader(code)
					er := json.NewEncoder(writer).Encode(body)
					if er != nil {
						fmt.Printf("error occured while decoding")
					}
				}
			}
		default:
			{
				writer.WriteHeader(404)
				fmt.Fprintf(writer, "Rout not found!")
			}
		}
	})
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
