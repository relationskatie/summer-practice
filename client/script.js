document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('submitForm');
    const vacancyTableBody = document.getElementById('vacancyTableBody');

    const savedFormData = JSON.parse(localStorage.getItem('formData'));
    if (savedFormData) {
        document.getElementById('text').value = savedFormData.text;
        document.getElementById('salary').value = savedFormData.salary;
        document.getElementById('area').value = savedFormData.area;
        if (savedFormData.experience) {
            document.querySelector(`input[name="experience"][value="${savedFormData.experience}"]`).checked = true; 
        }
    }

    function getExperienceString(experienceValue) {
        switch (experienceValue) {
            case 'noExperience':
                return 'Нет опыта';
            case 'between1And3':
                return 'От 1 года до 3 лет';
            case 'between3And6':
                return 'От 3 до 6 лет';
            case 'moreThan6':
                return 'Более 6 лет';
            default:
                return '';
        }
    }

    function renderVacancies(vacancies) {
        vacancies.forEach(vacancy => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${vacancy.name}</td>
                <td>${vacancy.salary.from}</td>
                <td>${vacancy.area.name}</td>
                <td><a href="${vacancy.alternate_url}" target="_blank">${vacancy.alternate_url}</a></td>
            `;
            vacancyTableBody.appendChild(row);
        });
    }

    function loadVacancies() {
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
            renderVacancies(vacancies);
        })
        .catch(error => {
            console.error('Проблемы с получением вакансий:', error);
        });
    }

    form.addEventListener('submit', (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);

        const dataToSave = {
            text: formData.get('text'),
            salary: formData.get('salary'),
            area: formData.get('area'),
            experience: formData.get('experience')
        };
        localStorage.setItem('formData', JSON.stringify(dataToSave));

        const experienceString = getExperienceString(formData.get('experience'));
        dataToSave.experience = experienceString;

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
            vacancyTableBody.innerHTML = '';
            loadVacancies();
        })
        .catch(error => {
            console.error('Проблемы с отправкой данных:', error);
        });
    });

});
