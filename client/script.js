document.getElementById('submitForm').addEventListener('submit', (event) => {
    event.preventDefault();
    const formData = new FormData(event.target);
    fetch('http://localhost:8000/app/form', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            area: formData.get('area'),
            salary: formData.get('salary'),
            text: formData.get('text')
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Проблемы с сетью ' + response.statusText);
        }
        return response.json();
    })
    .then(data => {
        console.log('Отправленные данные:', data);
        fetch('http://localhost:8000/app/vacancies', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Проблемы с сетью ' + response.statusText);
            }
            return response.json();
        })
        .then(vacancies => {
            console.log('Полученные вакансии:', vacancies);
            const vacancyList = document.getElementById('vacancyList');
            vacancyList.innerHTML = '<h2>Вакансии</h2>';
            vacancies.forEach(vacancy => {
                const vacancyItem = document.createElement('div');
                vacancyItem.className = 'vacancy-item';
                vacancyItem.innerHTML = `
                    <h3>${vacancy.name}</h3>
                    <p>Зарплата: ${vacancy.salary.from}</p>
                    <p>Область: ${vacancy.area.name}</p>
                    <p>URL: <a href="${vacancy.url}" target="_blank">${vacancy.url}</a></p>
                `;
                vacancyList.appendChild(vacancyItem);
            });
        })
        .catch(error => {
            console.error('Проблемы с получением вакансий:', error);
        });
    })
    .catch(error => {
        console.error('Проблемы с отправкой данных:', error);
    });
});
