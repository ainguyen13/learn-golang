## 1 package - module 
package main chứa func main là điểm bắt đầu chạy chương trình go. 
- Một ứng dụng Go (module go) có thể chứa nhiều package
- Mỗi package nằm ở một thư mục khác nhau
- Trong một package có thể chưa 1-nhiều file go. Mỗi file go hãy giữ dưới 200 dòng
- Một ứng dụng Go có thể import nhiều module ngoài. Tương tự mỗi module ngoài có thể chứa 1-nhiều package 

## 2 Hàm trả về error
Do Go không có try/catch nên phải luôn kiểm tra lỗi trả về từ mỗi func.
Quy ước tham số trả về cuối cùng luôn là error
Golang có cú pháp đặc biệt là if assignment gồm 2 vế:
1. Assignment: result, err := Sqrt(-1)
2. Condition: err != nil

Ví dụ: 
```go
if reult, err := Sqrt(-1); err != nil {
    fmt.Println(err,Error())
} else {
    fmt.Println(reult)
}
```

## 3. Debug trong Go và VSCode
Để debug được ứng dụng Golang chúng ta cần phải:
1. Tạo go module go mod init ten_module
2. Trong VSCode "Create launch JSON file"

Hãy tận dụng tối đa chức năng debug, đừng chỉ dùng lệnh in ra console.
Sau khi triển khai ứng dụng lên production. Hãy xóa hết tất cả các lệnh debug in ra console, vì những lệnh này gây chậm tốc độ

fmt.Println(variable)

Trong debug tool bar của VSCode có 5 nút:
- Continue: chạy tiếp
- Step Over: chạy qua đến dòng lệnh tiếp theo
- Step Into: đi sâu vào lệnh hiện tại vào hàm được gọi
- Step Out: chạy ra ngoài hàm hiện tại
- Restart: khởi động lại ứng dụng

Cần chú ý khung debug
1. Variables gòm 2 phần: Arguments (tham số truyền vào), Locals (các biến cục bộ trong hàm đang được debug)
2. Watch: lập trình viên có thể thêm các biểu thức để quan sát
3. BreakPoints: các điểm dừng do lập trình đặt để dừng chương trình
4. Call Stack: ngăn xếp các hàm chồng lên nhau khi chạy debug

Condition BreakPoint hữu ích khi chúng ta phải debug một vòng lặp nhiều lần. Hãy viết biểu thức boolean để tạm ngưng vòng lặp ở điểm chúng ta cần debug

```go
import "fmt"
func main() {
    for i:= 0; i < 100; i++ {
        fmt.Println(i)
    }
}
```