package main

import (
  "fmt"
  "os"
  "strconv"
  "net/http"
  "io"
)

func calculator (w http.ResponseWriter, req *http.Request){
  req.ParseForm()
  a:=req.FormValue("a")
  b:=req.FormValue("b")
  operator:=req.FormValue("operator")
  fmt.Fprintf(w, "b = %s\n", a)
  fmt.Fprintf(w, "a = %s\n", b)
  convA, err:=strconv.Atoi(a)
  convB, err:=strconv.Atoi(b)
  if err!=nil{
    fmt.Println(err)
    os.Exit(2)
  }
  var result int
  if operator=="add"{
    result=convA+convB
  }else if operator=="subtr"{
    result=convA-convB
  }else if operator=="mul"{
    result=convA*convB
  }else if operator=="div"{
    if convB!=0{
      result=convA/convB
    }else{
      io.WriteString(w, "You can't divide on 0")
    }
  }
  fmt.Fprintf(w, "Result: %d", result)
}
func main(){
  http.HandleFunc("/calc", calculator)
  http.ListenAndServe(":8080", nil)
}
