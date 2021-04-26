# naive-grpc
Naive grpc-server

ВНИМАНИЕ ! ЭТО КОД ДЕМОНСТРИРУЕТ КАК создать и запустить GRPC сервер
ЭТО УЧЕБНЫЙ ПРИМЕР, НЕ ИСПОЛЬЗУЙТЕ ТАКУЮ ОРГАНИЗАЦИЮ КОДА В СВОИХ РЕШЕНИЯХ
Вам нужен protoc, мой protoc --version:  
libprotoc 3.15.6, если другой protoc, могут быть отличия

Находясь в папке pb  
```protoc --go_out=. *.proto```  
Cгенерируется s.pb.go  
  
```protoc --go-grpc_out=. *.proto```  
Cгенерируется s_grpc.pb.go  

Обратите внимание, что какие публичные методы и структуры находится в файле s.pb.go и что s_grpc.pb.go    

Запустите сервер  
```go run main.go```  

Evans (https://github.com/ktr0731/evans.git)  
В другом терминале наберите  
```evans -p 8081 -r```  
Откроется интерактивный режим grpc клиента, наберите  
```
show package
...
show service
```  
Изучите вывод, проанализируйте взаимосвязь с кодом и proto файлом  
Вызовите сначала Hello, a затем UnimplementedMethod:  
```
api.SampleServer@127.0.0.1:8081> call Hello
name (TYPE_STRING) => Alex
{
  "name": "Hello Alex"
}
---
api.SampleServer@127.0.0.1:8081> call UnimplemetedMethod
name (TYPE_STRING) => alex
command call: rpc error: code = Unimplemented desc = method UnimplemetedMethod not implemented
```   
Проанализируйте причину разницы в поведении сервера,  
как "починить"  UnimplementedMethod?
