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

## 4. Khai báo mảng
Có 3 cách. 
```go
cars := [3]string{"Toyota", "Mercedes", "BMW"}
cars := []string{"Toyota", "Mercedes", "BMW"}
cars := [...]string{"Toyota", "Mercedes", "BMW"}
```
Hai cách khai báo mảng 2 chiều

Không xác định số lượng phần tử trong mảng con
```go
langs := [][]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "RUST", "Crystal", "OCAML"}}
```

Xác định số lượng phần tử trong mảng con là 3 ! Kiểm tra lúc compile time.
```go
langs := [][3]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "RUST"}}
```

## 5. Defer
Đưa một lệnh vào một ngăn xếp đặc biệt. Trước khi hàm thoát, thì sẽ thực hiện tuần tự các lệnh trong ngăn xếp này theo cơ chế Last In First Out.

Xem ví dụ:
```go
func reverseLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	//fmt.Println(cars[0]) // Toyota
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}
```

## 6. Slice & Array

Array trong Go là static array. Số phần từ mảng không thay đổi sau khi khởi tạo.
Muốn thay đổi phần tử, thêm, xóa phải dùng Slice.
Phần tử trong mảng array có thể là value có thể là reference (con trỏ đến vùng nhớ) tùy thuộc vào kiểm phần tử mảng.

Chỉ mục truy xuất đến phần tử mảng luôn bắt đầu từ 0. Zero base tương tự như C, C++, Java

> Chú ý chỉ số sau ký tự `:` là exclusive chứ không phải là inclusive, có nghĩa là không bao gồm.
```go
fmt.Println(a[:2])        //Lấy 2 phần tử đầu tiên
fmt.Println(a[2:])        //Bỏ qua 2 phần tử đầu tiên
fmt.Println(a[len(a)-2:]) //Lấy 2 phần tử cuối cùng
fmt.Println(a[1:3])
```

## 7 Map: Kiểu lưu trữ key-value
Khai báo map sử dụng từ khóa `make` rất dài dòng
```go
dict := map(map[string]string)
```
Khai báo ngắn gọn. Recommended!
```go
dict := map[string]string{}
```
## 8. Remove item from slice
Có nhiều cách để xoá phần tử ra khỏi slice, chia thành 2 nhóm:

1. Ưu tiên tốc độ, không quan tâm đến thứ tự phần tử sau khi xoá
2. Ưu tiên giữ thứ tự phần tử

Cần có phương pháp đánh giá tốc độ các phương pháp này. Golang cung cấp kỹ thuật Benchmark.

## 8. Test và Benchmark

Golang cung cấp 2 kỹ thuật:
1. Unit Test kiểm thử logic chạy theo ý đồ của lập trình viên không
2. Benchmark kiểm thử tốc độ thực thi

Chúng ta có thể viết file *test.go nằm trong cùng package hoặc ra thư mục riêng.
Xem [remove_slice_bench_test.go](remove_slice_bench_test.go)
File test hay benchmark luôn phải kết thúc bằng `test.go` và `import "testing"`

Hàm benchmark luôn phải bắt đầu bằng `func Benchmark`

Hãy liên tục viết hàm benchmark để chọn ra cách tối ưu thực thi. Phong cách của Go là
- Đơn giản
- Ngắn gọn
- Chạy nhanh
- Tốn ít RAM

## 9. Truyền slice vào một hàm
Khi truyền vào một mạng/slice là truyền by reference

Khi một slice truyền vào một hàm, mặc dù nhìn có vẻ như là truyền bằng giá trị, nhưng thực chất là truyền bằng con trỏ tham chiếu đến cùng một mảng (array). Bất kỳ thay đổi nào lên slice cũng sẽ ảnh hưởng trực tiếp đến mảng.

**Qua đây chúng ta thấy quan hệ giữa mảng (array) và slice như sau:
- Mảng là nơi lưu trữ tất cả các phần tử thực sự trong bộ nhớ.
- Còn slice là cấu trúc để tham chiếu đến các phần tử mảng này. 
- Mảng thì tĩnh, không thể thay đổi kích thước. Nhưng với Slice thì có thể them, xóa, nghịch đảo phần tử.

## 10. Khai báo mảng sử dụng anonymous struct
Cách này vừa khai báo mảng chứa các struct. Không cần gì rõ tên struct mà chỉ cần các thuộc tính bên trong struct. Sau khi khai báo xong, thì khởi tạo dữ liệu luôn

