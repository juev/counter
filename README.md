# counter

ПО счетчик посещений веб-страниц.

Состоит из клиента и сервера.

## Клиент

Используется для управления серверной частью и запросов данных в сервер.

Управление:

 1. Добавление нового домена в счетчик
 2. Удаление домена из счетчика

Запрос данных:

 1. Получение данных по посещаемости определенного домена

## Сервер

Используется для подсчета посещений страницы и проведения агрегирования данных.

Счетчик висит на определенном дипазоне адресов вида:

```plain
/count/12345/
```

Число это номер присвоенный определенному доменну с помощью клиента. Если домен
не заведен, то есть проводиться обращение к незарегистрированному номеру, то
проводиться фиксация обращения на номер `1`, который является общим для всех
запросов, которые не зарегистрированы.

(?) возможно использование не порядковых номеров, а определенных uid?

Используются GET-запросы на указанные адреса для проведения подсчета.

Управление сервером осуществляется через GRPC:

```plain
/add
/remove
/stat
```

Хранение данных проводиться в redis для использования встроенной дедубликации.
Формат данных предстоит продумать.
