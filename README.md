# Репозиторий для скриншотного тестирования страницы сайта Avito - Eco Impact

Тестируемый роут
https://www.avito.ru/avito-care/eco-impact

___
Инструкция по установке:
- Скачать и установить go версии `1.22.1`
- Скопировать репозиторий с github командой `git clone https://github.com/eziarange/AVT.git`
- Перейти в папку со скачанным репозиторием
- Выполнить команду `go mod vendor`
- Выполнить команду `go run .`
- Скриншоты появятся в папке `output` спустя 8+ минут. Время обусловлено тем, что между запросами стоит Sleep, чтобы автоматика не забанила тесты (а она сделала это уже 1 раз)
___
- Тест кейсы лежат в файле [TESTCASES.md](TESTCASES.md)
- Баг-репорты лежат в файле [BUGS.md](BUGS.md)
- Баг-репорты по первому заданию лежат в файле [BUGS_FIRST_TASK.md](BUGS_FIRST_TASK.md)