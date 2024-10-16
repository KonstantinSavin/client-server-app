# Client server app
- [Описание](#description)
- [Использованные библиотеки](#libs)
- [Конфигурация](#config)
- [Установка](#management)

# Описание <a name="description"/>

1. Создана БД PostgreSQL. В БД создана, путем миграций, таблица из 4 полей (id, name, age, isconnected), заполненная произвольными данными в 6000 строк.
2. Разработано клиент-серверное приложение (по факту 2 приложения клиент и север). На форме клиента, при помощи консоли, клиент запрашивает данные. После получения запроса сервер обращается в БД и передаёт данные на клиент.
3. Клиент получает данные в одном потоке и передаёт для записи в файл в другой поток. Данные из потока в поток передаются построчно.
4. Все запросы пишутся в один и тот же файл.
5. Нажатие на кнопку не замораживает клиент. Есть возможность запустить еще один запрос при выполнении других запросов.


# Использованные пакеты <a name="libs"/>

- Логгирование: [sirupsen/logrus](https://github.com/sirupsen/logrus)
- Миграции: [pressly/goose](https://github.com/pressly/goose)
- Работа с HTTP: [gin-gonic/gin](https://github.com/gin-gonic/gin)

# Конфигурация <a name="config"/>

Конфигурация вынесена в файл .env

# Установка <a name="management"/>

1) Из дирректории server-app запускаем команду make для сборки контейнера с приложением,
затем команду make up для старта скрвера (при этом создается БД с помощью миграции).
2) Билдим main.go в дирректории client-app, запускаем, запрашиваем данные с сервера с помощью инструкций в консоли.

[Makefile](Makefile) 
