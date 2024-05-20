To-do List API

Privia Security Siber Savaş Akademisi staj projesi. 

Go dili kullanılarak login ve CRUD işlemleri yapılabilen to-do list uygulaması. Projede HTTP framework olarak gin, kullanıcılara tanımlanacak tokenler için jwt-go, environment değişkenleri için
godotenv framework'leri kullanıldı. Proje MVC yapısına uygun olarak yapıldı. Kullanıcıları ve todo elementlerini kayıt altında tutabilmek için JSON dosyaları kullanıldı. Ön tanımlı kullanıcılardan biri
user diğeri admin olacak şekilde oluşturuldu.

Projeyi test etmek için Postman kullanıldı. Login işlemi sonrası verilen token'i Authorization Header'ına koyarak diğer endpointler test edilebilir. Aksi halde "Authorization header required" hatası ile karşılaşılabilir.

Admin tüm todo listelerini görme yetkisine sahip, user'lar yalnızca kendi oluşturdukları todo'lar üzerinde işlem yapabilmekte.

User todo listesini veya todo mesajını sildiğinde, JSON dosyasındaki kaydı silmek yerine "deleted_at" değişkeni tanımlanmakta ve silinen değerler kullanıcıya gösterilmemekte.

Kullanıcı, listedeki bir todo'yu tamamlamak için update route'unu kullanmalı. updateTodoMessage fonksiyonu todo message'ın içeriğini ve tamamlanma durumunu  değiştirmek için kullanılmakta.

CRUD işlemleri ve login işlemi için gerekli route'lar main.go dosyasında bulunmakta.

Projeyi başlatmak için go.mod ve go.sum dosyalarındaki dependency'leri indirdikten sonra CompileDaemon -command="./<directory_adı>"  komutunu kullanabilirsiniz.
