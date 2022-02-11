### API для создания, получения и редактирования пользователя. Программа спроектирована с примененим архитектурного паттерна clean architecture. Также это паттерн известен как паттерн портов и адаптеров 

## Точка входа в программу cmd/http/main.go

## Пример создания пользователя POST запросом localhost:8080/user

{
"firstname": "Boris",
"lastname": "Kramarovskiy",
"email": "weslastam01819@gmail.com",
"age": 23,
}

## Получение юзера GET localhost:8080/user/параметр id

## Обновление юзера PUT localhost:8080/user/параметр id

{
"firstname": "Elon ",
"lastname": "Musk",
"email": "I-love-space-@gmail.com",
"age": 50,
}
