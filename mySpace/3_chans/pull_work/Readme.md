# 👷 ЗАДАЧА 5: "Пул воркеров"

## 📋 Условие:

* Есть слайс []int{1,2,3,4,5,6,7,8,9,10}
* Создай 3 воркера (горутин), которые:
* Читают числа из jobs канала
* Умножают на 2
* Пишут в results канал
* Заполни jobs, закрой его
* Прочти results и выведи результат

## 💡 Подсказка:

* Используй sync.WaitGroup для воркеров

* После завершения — закрой results