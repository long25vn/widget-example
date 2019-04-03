# Widgets

Dữ liệu hiển trị lưu trong file `data.json`.

Mỗi layout có `Type` là loại template, và `Code` là nội dung sẽ được render theo layout đã chọn, thứ tự code trong mảng tương đương với vị trí widget trong layout

```
{
   "Type": 2,
   "Code": [
      "<h2>This is H tag</h2>",
      "<p>This is p tag</p>"
   ]
},
```




### Một số layout có sẵn

![](https://i.ibb.co/Rvkd9mh/Untitled.jpg)


### Kết quả

![](https://i.ibb.co/P48RTgR/Hello.jpg)

### Chạy ví dụ:

```
go run main.go
```

Truy cập [localhost:8080](http://localhost:8080)