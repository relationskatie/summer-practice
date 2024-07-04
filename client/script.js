document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('submitForm');
    const vacancyList = document.getElementById('vacancyList');

    const savedFormData = JSON.parse(localStorage.getItem('formData'));
    if (savedFormData) {
        document.getElementById('text').value = savedFormData.text;
        document.getElementById('salary').value = savedFormData.salary;
        document.getElementById('area').value = savedFormData.area;
    }

    form.addEventListener('submit', (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        
        const dataToSave = {
            text: formData.get('text'),
            salary: formData.get('salary'),
            area: formData.get('area')
        };
        localStorage.setItem('formData', JSON.stringify(dataToSave));

        fetch('http://localhost:8000/app/form', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataToSave)
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Проблемы с сетью ' + response.statusText);
            }
            return response.json();
        })
        .then(data => {
            console.log('Отправленные данные:', data);
        
            vacancyList.innerHTML = '';

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
});
