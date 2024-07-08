package main

import (
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
    cwd, _ := os.Getwd()
    logFile := filepath.Join(cwd, ".log")
    logger := logrus.New()
    logger.SetOutput(os.Stdout)
    file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
    if err != nil {
        logger.Fatal(err)
    }
    defer file.Close()
    logger.SetOutput(file)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

		http.HandleFunc("/Standart_calculator", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("For addition, use /sum?x=n&y=n\n" +
		"For subtraction, use /sub?x=n&y=n\n" +
		"For multiplication, use /mul/?x=n&y=n\n" +
		"For division, use /div/?x=n&y=n\n" +
		"For parsing, use /parsingCalculator?expr=expression"))
	})

    
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
    xStr := r.URL.Query().Get("x")
    yStr := r.URL.Query().Get("y")

    x, err := strconv.Atoi(xStr)
    if err != nil {
        http.Error(w, "Invalid x", http.StatusBadRequest)
        return
    }
    y, err := strconv.Atoi(yStr)
    if err != nil {
        http.Error(w, "Invalid y", http.StatusBadRequest)
        return
    }

    if x > 0 && y > 0 && x > math.MaxInt-y {
        logger.WithFields(logrus.Fields{
            "x": x,
            "y": y,
        }).Warning("Sum overflows int")
        w.Write([]byte("Error"))
    } else {
        sum := x + y
        sumStr := strconv.Itoa(sum)
        w.Write([]byte(sumStr))
    }
})


http.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
    xStr := r.URL.Query().Get("x")
    yStr := r.URL.Query().Get("y")

    x, err := strconv.Atoi(xStr)
    if err != nil {
        http.Error(w, "Invalid x", http.StatusBadRequest)
        return
    }
    y, err := strconv.Atoi(yStr)
    if err != nil {
        http.Error(w, "Invalid y", http.StatusBadRequest)
        return
    }

    if x > 0 && y > 0 && x > math.MaxInt-y {
        logger.WithFields(logrus.Fields{
            "x": x,
            "y": y,
        }).Warning("Subtraction overflows int")
        w.Write([]byte("Error"))
    } else {
        sub := x - y
        subStr := strconv.Itoa(sub)
        w.Write([]byte(subStr))
    }
})


http.HandleFunc("/mul", func(w http.ResponseWriter, r *http.Request) {
    xStr := r.URL.Query().Get("x")
    yStr := r.URL.Query().Get("y")

    x, err := strconv.Atoi(xStr)
    if err != nil {
        http.Error(w, "Invalid x", http.StatusBadRequest)
        return
    }
    y, err := strconv.Atoi(yStr)
    if err != nil {
        http.Error(w, "Invalid y", http.StatusBadRequest)
        return
    }

    if x > 0 && y > 0 && x  * y> math.MaxInt {
        logger.WithFields(logrus.Fields{
            "x": x,
            "y": y,
        }).Warning("Multiplication overflows int")
        w.Write([]byte("Error"))
    } else {
        mul := x * y
        mulStr := strconv.Itoa(mul)
        w.Write([]byte(mulStr))
    } 
})

http.HandleFunc("/div", func(w http.ResponseWriter, r *http.Request) {
    xStr := r.URL.Query().Get("x")
    yStr := r.URL.Query().Get("y")

    x, err := strconv.Atoi(xStr)
    if err != nil {
        http.Error(w, "Invalid x", http.StatusBadRequest)
        return
    }
    y, err := strconv.Atoi(yStr)
    if err != nil {
        http.Error(w, "Invalid y", http.StatusBadRequest)
        return
    }

    if x > 0 && y > 0 && x / y > math.MaxInt {
        logger.WithFields(logrus.Fields{
            "x": x,
            "y": y,
        }).Warning("Division overflows int")
        w.Write([]byte("Error"))
    } else if y != 0{
        div := x / y
        divStr := strconv.Itoa(div)
        w.Write([]byte(divStr))
    } else if y == 0 {
			w.Write([]byte("Division by zero"))
		}
})

   http.HandleFunc("/parsingCalculator", func(w http.ResponseWriter, r *http.Request) {
    expr := r.URL.Query().Get("expr")
    res, err := start(expr)
    if err != nil {
        http.Error(w, "Error evaluating expression", http.StatusInternalServerError)
        return
    }

    w.Write([]byte(res))
})

    port := "8080"
    logWithPort := logrus.WithFields(logrus.Fields{
        "port": port,
    })
    logWithPort.Info("Starting a web-server on port")
    logWithPort.Fatal(http.ListenAndServe(":"+port, nil))
}
