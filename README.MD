### Создать приложение веб-сервер (REST API) с функционалом по работе со структурой данных типа
*задание->задачи->трудозатраты, база данных любая.*

### Операции:

1. [x] Создать задание
2. [x] Создать задачи для задания
3. [x] Создать трудозатраты для задачи
4. [ ] Выбрать дерево задание->задачи->трудозатраты, отобразить на странице,
дерево строится средствами кода, а не базы данных.
Для каждого из элементов считается общее потраченное время и среднее время по всем подчиненным элементам.
Продемонстрировать при расчетах знания goroutines, mutex и каналов.
5. [ ] Удаление любого из элементов из отображенного списка с перерасчетом данных.
6. [ ] Переподчинение задач, трудозатрат другим родителям. Названия одного уровня уникальны

```json
{
  "task1": {
    "subTask1": {
      "cost1": 12,
      "cost2": 5,
      "cost3": 8
    }
  }
}
```

*Примечания: Уровень логирования приложения должен задаваться параметром командной строки.
По возможности не использовать глобальные переменные.
Структура приложения должна быть модульной и содержать модули модели, сервиса и конфигурирования.
Код должен быть покрыт комментариями.
ЖЕЛАТЕЛЬНО чтобы приложение компилировалось и запускалось в контейнере.*