# GoShell – A Custom Shell in Golang  

GoShell is a custom shell implemented in **Go**, designed to explore how shells interact with the operating system. It supports **built-in commands**, **external commands**, and **piping**.

I built this project to understand the underlying concepts of process management, system calls, and command execution in Unix-like environments.  

![shell](https://github.com/user-attachments/assets/4b9ca19d-281a-4f46-96e8-b400016b15cf)


---

## Features  
- **Built-in Commands**: `cd`, `exit`, and more.  
- **External Commands**: Run standard commands like `ls`, `cat`, etc.  
- **Command Piping**: Connect commands using `|` just like in Bash.  
- **Lightweight Prompt**: Simple interface focusing on functionality.  

---

## 🛠️ How It Works  
- Reads user input.  
- Parses commands and detects pipes.  
- Executes built-ins directly.  
- Spawns external processes using Go’s `os/exec` package.  
- Connects commands via `io.Pipe()` for piping functionality.
- 
<img width="1541" height="800" alt="Screenshot from 2025-08-25 03-17-47" src="https://github.com/user-attachments/assets/db0e0284-1b6a-4c56-a86a-fe094fb3709f" />


This helped me understand **how shells interpret and execute commands** at a low level.  

---

