# Конвертер из XML в JSON для сайта radio-t.com
Полное описание задачи можно находится [здесь](https://github.com/radio-t/radio-t-site/issues/23)

rtxml имеет 4 параметра:
* `input` - путь до xml файла для обработки
* `output` - путь до json файла, в который будет записан результат
* `podcast` - путь до json файла, полученного из api radio-t.com по запросу `GET /last/1000?categories=podcast` (все выпуски подкастов)
* `prep` - - путь до json файла, полученного из api radio-t.com по запросу `GET /last/1000?categories=prep` (все темы для выпусков)

Параметры вводятся в стиле Plan 9 (дефолтным для go):

`./rtxml -input ./IntenseDebate_clean.xml -output ./rt.json -podcast ./podcast.json -prep ./prep.json`
