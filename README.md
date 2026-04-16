
# taskmaster

Simple CLI app to learn Go


## Authors

- [@itsmeviiper](https://www.github.com/ItsMeViipeR)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Features

- Add task
- Remove task
- Edit task
- Mark task as done
## Run Locally

Clone the project

```bash
  git clone https://github.com/ItsMeViipeR/taskmaster.git
```

Go to the project directory

```bash
  cd taskmaster
```

Start the server

```bash
  go run .
```

## Prerequisites
- [Go](https://go.dev/doc/install) (version 1.25.8 or higher)

## Usage/Examples

```sh
taskmaster add  "Buy bread"
taskmaster add "Buy bread"
taskmaster remove "Buy bread" # or taskmaster remove <id>
taskmaster idof "Buy Bread" # returns id of all the instances of "Buy Bread" - here 0 and 1
taskmaster done "Buy bread" # or taskmaster done <id>
```
