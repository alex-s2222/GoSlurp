# ЗАДАЧА 2: "Параллельные рабочие"

📋 Условие:

* Создай 10 горутин, каждая должна:
* Вывести: Worker N started
* Подождать N * 100ms
* Вывести: Worker N done

Используй:
* sync.WaitGroup для ожидания завершения всех горутин

💡 Пример вывода:
```text
Worker 1 started
Worker 2 started
...
Worker 1 done
Worker 2 done
...
```