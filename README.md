# 💬 Go RPC Chatroom

A simple chatroom built in **Go** using the **net/rpc** package for communication between a server and multiple clients.

---

## 🧠 Overview

This project demonstrates how to implement a minimal chatroom using **Remote Procedure Calls (RPC)** in Go.

- The **server** coordinates messages between clients and keeps the full chat history.
- Each **client** connects via RPC, sends messages, and retrieves the chat history from the server.
- Each client enters their **name** before chatting, and all messages are displayed with usernames.
- 

---
🎥 Demo Video

🎬 Watch the running demo here:
https://drive.google.com/file/d/1SJKkNGV2WZy_IryrVP7blwXuGjTiyc3o/view?usp=sharing

🧑‍💻 Author

Mahmoud Salem
GitHub Profile: https://github.com/mahmoudsalem253

🗂 Repository Link

🔗 https://github.com/mahmoudsalem253/chatroom

🧾 Documentation Summary

This chatroom project uses Go RPC to allow communication between multiple clients through a single coordinating server.
It demonstrates:

RPC method definition and invocation

Concurrency-safe data handling with sync.Mutex

Persistent in-memory message storage

Basic error handling for client-server disconnections
## 🚀 Features

✅ Multi-client chatroom via Go RPC  
✅ Thread-safe message storage on server  
✅ Each user enters a name before chatting  
✅ Server logs all messages in real-time  
✅ Clients can exit with `exit` or `Ctrl + C`  

---

## 🧩 Project Structure


---

## ⚙️ How It Works

### 🖥 Server
- Listens on TCP port `1234`
- Exposes two RPC methods:
  - `SendMessage(string, *[]string)` → stores message and returns updated history
  - `GetMessages(string, *[]string)` → returns the full chat history
- Displays all messages in its own terminal

### 💬 Client
- Connects to the server using `rpc.Dial`
- Prompts the user to enter their name
- Sends messages in the format:  

- Displays full chat history after each message

---

## 🧰 How to Run

1. **Start the server**  
 ```bash
 go run server.go
- Displays full chat history after each message

---

## 🧰 How to Run

1. **Start the server**  
 ```bash
 go run server.go
Chat server running on port 1234...
go run client.go

Enter your name: Mahmoud
Connected to chat server.
You: Hello everyone!

--- Chat History ---
Mahmoud: Hello everyone!
--------------------




