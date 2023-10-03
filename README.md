# goerrorsplus

## Description

goerrorsplus is a simple error wrapper for go. It provides a simple way to wrap errors with additional information such as error, custom error message, go version, function name, arguments, and stack trace. It also provides a simple way to print the error as normal or verbose also keep the  err obj for further processing.

## Usage

### Installation

    ```bash
    go get github.com/fengdotdev/goerrorsplus
    ```

### Example

    ```go
        package main
        import(
        "fmt"
        "github.com/fengdotdev/goerrorsplus/e"
        "joepackage"
        )

        func main()

        // joe.division is a fictional function that returns an error if the second argument is zero
        result, err :=  joe.division(1,0)
        if err != nil {
        // you provide the err, custom message, and additional info such as tags, function name, trace and arguments (u can leave nil or empty if you don't want to provide them)
        errp := e.E(err,"div failed",[]string{"math-err", "joe-package"},division,1,0)
        fmt.Println(errp.V()) // prints: 2023/10/03 04:55:44 2023-10-03 04:55:44 go1.21.1 error: recover: div by zero msg: div failed fn: main.division trace: caller: G:/fengdotdev/github/goerrorsplus/example/main.go line: 33  args: [%!s(int=1) %!s(int=0)]
        }
    ```

### Output

    ```bash
    go run example/main.go
    2023/10/03 15:46:13 2023-10-03 15:46:13 go1.21.1 error: recover: div by zero msg: div failed fn: main.division trace: caller: G:/fengdotdev/github/goerrorsplus/example/main.go line: 33  args: [%!s(int=1) %!s(int=0)]
    ```

## License

under the MIT License.

## Author

    ALEJANDRO VALDERAS FUENTES (FENG) feng.dev <feng@vitruvio.cl>

