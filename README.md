# Конвертер из XML в JSON для сайта radio-t.com
Полное описание задачи находится [здесь](https://github.com/radio-t/radio-t-site/issues/23)

rtxml имеет 2 параметра:
* `input` - путь до xml файла для обработки
* `output` - путь до json файла, в который будет записан результат
* `podcast` - путь до json файла, полученного из api radio-t.com по запросу `GET /last/1000?categories=podcast` (все выпуски подкастов)
* `prep` - - путь до json файла, полученного из api radio-t.com по запросу `GET /last/1000?categories=prep` (все темы для выпусков)

Параметры вводятся в стиле Plan 9 (дефолтном для go), данные принимаются с stdin, выводятся на stdout (ошибки - в stderr):

`cat ./IntenseDebate_clean.xml | ./rtxml -podcast ./podcast.json -prep ./prep.json > rt.json`
