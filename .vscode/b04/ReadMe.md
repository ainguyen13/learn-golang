# Chủ đề hôm nay về Interfance

Các bạn hay chạy code ở thư mục [demointerface](demointerface) trước.

Sau đó chúng sẽ xây dựng chức năng gửi email có nhiều biến thể khác nhau ở [mailer](mailer)

## Tuân thủ interface - implement interface
- Chỉ làm việc trên các prototype function mà đối tượng tuân thủ Interface 
- Không bận tâm đến kiểu thực sự của đối tượng
- Tính đa hình giúp lập trình viên quản lý được một nhóm đối tượng rất da dạng 
- Một đối tượng trong Golang có thể tuân thủ nhiều interface 
- Để một đối tượng gọi là `tuân thủ` một interface thì đối tượng đó phải implement tất cả các prototype function trong interface đó.
- Việc tuân thủ interface trong golang là `implicit` không cần khai báo, cứ tuân thủ đủ các prototype function của một interface là đạt.

```go
shapes := []Shape{rec, cir, tri}
for _, shape := range shapes {
  //Polymorphism
  fmt.Printf("%s perimeter = %.2f\n", shape, shape.Perimeter())
}
```

## Empty Interface `interface{}`
Kiểu Empty Interface chứa đối tượng thuộc bất kỳ kiểu nào bởi vì Empty Interface không chứa bất kỳ prototype function nào.

- Empty Interface là `interface{}`
- `interface{}` chấp nhận bất kỳ kiểu gì int, string, array, slice, struct, func
- Dùng `interface{}` khi nào? - Dùng khi không biết trước kiểu sẽ truyền vào 

## Kiểm tra kiểu Interface
Cú pháp kiểm tra kiểu bằng if assigment

```go
if f, ok := i.(string); ok { //Golang if assignment statement
  fmt.Println(f)
}
```

Kiểm tra kiểu bằng cú pháp `switch`
```go
for _, v := range arrayAnyType {
  switch v.(type) {
  case string:
    fmt.Printf("%v is string\n", v)
  case Person:
    fmt.Printf("%v is Person\n", v)
  case func(int, int) int:
    fmt.Printf("%v is func(int, int) int\n", v)
  }

}
```
> Gợi ý: khi lưu trữ thì bao dung (dùng kiểu `interface{}`) nhưng khi làm việc với từng đối tượng cụ thể, cần phải tìm hiểu rõ đối tượng đó là kiểu nào hoặc tuân thủ interface nào. Nếu dùng sai, lập tức sẽ gặp lỗi panic