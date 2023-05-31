### 3.2 Đọc từ bàn phím
Có 3 loại module: 
1. Module chuẩn của Google
2. Public module chia sẻ cho mọi user
3. Private module chỉ cho phép những nhân viên hoặc thành viên giới hạn được xem

Đã phát hành module là chia sẻ toàn bộ mã nguồn
Để debug thì cần phải tạo module bằng lệnh `go mod init`

```
Build Error: go build -o D:\go-workspace\.vscode\__debug_bin.exe -gcflags all=-N -l .
go: go.mod file not found in current directory or any parent directory; see 'go help modules' (exit status 1)
```

Tạo module 
```
$ go mod init lab01
$ go mod tidy
```

`go mod tidy` dọn dẹp lại các module 

### 3.3 Khai báo biến
Có mấy cách:

Khai báo + Gán trong 1 lệnh. Recommended 
```go
height := 1.6
```

Khai báo và Gán tách thành 2 lệnh. Not Recommended
```go 
var height float64
height = 1.6
```

Tên hàm bắt đầu bằng chữ thường chỉ có phạm vi hoạt động trong nội bộ package - private 

Tên hàm bắt đầu bằng chữ hoa có phạm vi hoạt động (scope) trong và ngoài package - public


----------------------------------------------------------------
Trong Go, `bufio` là một thư viện chuẩn được sử dụng để xử lý đầu vào và đầu ra văn bản (text input/output). Thư viện này cung cấp các công cụ cho phép đọc và ghi các chuỗi văn bản, đọc dữ liệu từ các luồng vào chuẩn (os.Stdin) và từ các tệp, và ghi dữ liệu vào các luồng xuất chuẩn (os.Stdout) và các tệp.

Một số tính năng của thư viện bufio bao gồm:

Đọc và ghi các chuỗi văn bản đơn giản từ các luồng vào và luồng xuất chuẩn.

Cung cấp các công cụ cho phép đọc và ghi các dòng và các phần văn bản từ các luồng vào.

Cung cấp đệm để giảm thiểu số lần đọc hoặc ghi truy cập vào các tệp hoặc các luồng.

Cung cấp hỗ trợ cho các hoạt động đọc và ghi dữ liệu trực tiếp từ và đến các tệp.

Các tính năng này giúp cho việc xử lý đầu vào và đầu ra văn bản trở nên đơn giản và hiệu quả hơn trong Go. Thư viện bufio thường được sử dụng trong các chương trình đọc và xử lý dữ liệu đầu vào từ bàn phím hoặc các tệp.
-----------------------------------------------------------------
