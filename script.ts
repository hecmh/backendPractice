function sendLog(message: string): void {

    fetch('/log', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ message: message })
    })
      .then(response => {
        if (response.ok) {
          console.log('Лог успешно отправлен на сервер');
        } else {
          console.error('Ошибка при отправке лога на сервер');
        }
      })
      .catch(error => {
        console.error('Ошибка при отправке лога на сервер:', error);
      });
  }
  

  const logButton = document.getElementById('logButton');
  logButton.addEventListener('click', (event) => {
    event.preventDefault(); 
    const messageInput = document.getElementById('messageInput') as HTMLInputElement;
    const message = messageInput.value;
    sendLog(message);
  });
  