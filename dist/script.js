document.getElementById('nameForm').addEventListener('submit', function (event) {
    event.preventDefault();
    const nameInput = document.getElementById('nameInput');
    const responseContainer = document.getElementById('responseContainer');
    // Очищаем контейнер с предыдущим ответом
    responseContainer.innerHTML = '';
    // Получаем имя из поля ввода
    const name = nameInput.value;
    // Отправляем GET-запрос на сервер
    fetch('/?name=' + encodeURIComponent(name))
        .then(function (response) {
        return response.text();
    })
        .then(function (data) {
        // Отображаем полученный ответ в контейнере
        responseContainer.innerHTML = data;
    });
});
