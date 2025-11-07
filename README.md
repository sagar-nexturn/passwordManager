# Password Manager CLI

## Setup

1. Clone Project
2. Set environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME, MASTER_KEY).
3. Initialize table by own on render or by uncommenting the corresponding codes in main file.
4. Run the project.

## Example

```bash

# run(to know about commands -h flag)
go run main.go -h
go run main.go add -h
go run main.go get -h
go run main.go update -h
go run main.go delete -h
```

## Sample Run Examples
> The following are the example of commands.

- **Add Command**
---
-help

  ![Add Command](sampleRun/add_help.png)   

-main

  ![Add Command](sampleRun/add.png)

---

- **Get Command**
---
-help

![Get Command](sampleRun/get_help.png)

-by id

![Get Command](sampleRun/get_byId.png)

-by name

![Get Command](sampleRun/get_byName.png)

-all

![Get Command](sampleRun/get_all.png)

---

- **Update Command**
---
-help

![Update Command](sampleRun/update_help.png)

-main

![Update Command](sampleRun/update.png)

---

- **Delete Command**
---
-help

![Delete Command](sampleRun/delete_help.png)

-by id

![Delete Command](sampleRun/delete_byId.png)

-by name

![Delete Command](sampleRun/delete_byName.png)

---

