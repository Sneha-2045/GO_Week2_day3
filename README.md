# Day 3 - WaitGroup and Synchronization in Go

## Objective

The objective of this task is to understand Go's concurrency model using Goroutines and synchronize their execution using the `sync.WaitGroup` package. The program demonstrates concurrent file processing where multiple files are read simultaneously without blocking the main thread.

---

## Concepts Covered

- Goroutines
- Concurrency in Go
- WaitGroup Synchronization
- File Handling
- Error Handling
- Concurrent Processing

---

## Project Structure

```
Day3/
│
├── main.go
├── file1.txt
├── file2.txt
├── file3.txt
└── README.md
```

---

## Program Description

The program processes multiple text files concurrently using Goroutines.

### Workflow

1. A list of file names is created.
2. A `WaitGroup` is initialized.
3. A separate Goroutine is launched for each file.
4. Each Goroutine:
   - Reads the file content.
   - Calculates the file size.
   - Displays the file name and size.
5. The main Goroutine waits until all file-processing Goroutines complete.
6. A completion message is displayed.

---

## Code Features

### Concurrent File Processing
Multiple files are processed simultaneously, improving efficiency compared to sequential execution.

### WaitGroup Synchronization
`sync.WaitGroup` ensures the main function waits until all Goroutines finish their tasks.

### Error Handling
If a file is missing or cannot be read, the program displays an appropriate error message.

---

## Sample Input Files

### file1.txt
```
Hello from file 1
```

### file2.txt
```
Hello from file 2
```

### file3.txt
```
Hello from file 3
```

---

## How to Run

### Step 1: Initialize Go Module

```bash
go mod init day3
```

### Step 2: Create Sample Files

```bash
echo "Hello from file 1" > file1.txt
echo "Hello from file 2" > file2.txt
echo "Hello from file 3" > file3.txt
```

### Step 3: Run the Program

```bash
go run main.go
```

---

## Sample Output

```
File: file1.txt | Size: 18 bytes
File: file2.txt | Size: 18 bytes
File: file3.txt | Size: 18 bytes
All files processed successfully!
```

> Note: Since the files are processed concurrently, the output order may vary.

---

## Learning Outcomes

After completing this task, you will be able to:

- Create and execute Goroutines.
- Synchronize concurrent operations using WaitGroups.
- Implement concurrent file-processing applications.
- Handle file operations in Go.
- Understand the basics of parallel execution and synchronization.

---

## Technologies Used

- Go Programming Language
- sync.WaitGroup
- os Package

---

