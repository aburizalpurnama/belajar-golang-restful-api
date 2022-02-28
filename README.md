
## 5. Buat Repository

- Harus mengimplementasi context
- Harus mengimol# Belajar Membpelmentasi trSaat menggunakan database transactional, hnsactionalt Aplikasi Resful API dengan golang

## 1. Tambahkan Dependency yang diperlukan

- Driver MySQL : https://github.com/go-sql-driver/mysql 
- HTTP Router : https://github.com/julienschmidt/httprouter 
- Validation : https://github.com/go-playground/validator

## 2. Membuat Dokumentasi OpenAPI 

- API Spec List Categories
- API Spec Get Category
- API 

### Litle Note

- Jika terdapat object yang berulang, lebih baik dibuatkan component

## 3. Membuat Security Schema

## 4. Membuat Database

## 5. Membuat Repository

## 6. Membuat Service

- Jumlah function service akan sama dengan jumlah endpoint API
Biasanya sebuah API hanya memanggil satu buah function service yang merupakan bussiness prosesnya

- Harus mengimplementasi context

- Saat menggunakan database transactional, harus mengimpelmentasi transactional

- Response tidak boleh berupa model domain, domain adalah merupakan representasi dari sebuah table, jika dibuat objek response maka akan mengekspose semua data yang ada di table termasuk data yang sesitif (password, dll)

- Transactional dilakukan di service layer, karena kemungkinan akan berinteraksi dengan berbagai repository sehingga jika salah satu terjadi error maka seluruh operasi pada semua repository akan dirollback

## 7. Tambahkan Validator

### Go-Playground Validator

- Instalation :

	go get github.com/go-playground/validator/v10
	
- Instruction :

	https://github.com/go-playground/validator

- Validator digunakan untuk menvalidasi input dari user (CreateRequest & UpdateRequest )

- Create dan Update saja yang ditambahkan validasi.

## 8. Membuat Controller

### Membuat Object Response

### Membuat Helper Function untuk memproses request body dan response body

## 9. Membuat instance function untuk repository, service, dan controller

## 10. Membuat helper untuk create db connection

## 11. Membuat Router

- Buat API mapping untuk tiap endpoint yang diperlukan sesuai API specs 

- Buat di main function

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.POST("/api/v1/categories", categoryController.Create)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

## 12. Buat HTTP Router

- Dibuat di main function, dibawah router

## 13. Tambahan-tambahan lain

- Buat object data pada respon menjadi lower case semua

	`json:"code`

- Connection harus diclose (FindById, FindAll)


## Notes

- Nama module 'domain-name'/'project-name'

	contoh : programmerzamannow/belajar-golang-restful-api
	
- Semua parameter dengan tipe struct harus dijadikan pointer, namun jika berupa interface tidak perlu
