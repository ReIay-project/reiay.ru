// Объект переводов
const translations = {
    en: {
        title: 'Relay Messenger',
        tagline: 'We work for you',
        description: 'Next-generation messenger',
        requestAccess: 'Request Access',
        enterEmail: 'Enter email'
    },
    ru: {
        title: 'Relay Мессенджер',
        tagline: 'Мы работаем для Вас',
        description: 'Мессенджер нового поколения',
        requestAccess: 'Запросить доступ',
        enterEmail: 'Введите email'
    }
};

// Функция для изменения языка
function changeLanguage(language) {
    document.getElementById('title').textContent = translations[language].title;
    document.getElementById('tagline').textContent = translations[language].tagline;
    document.getElementById('description').textContent = translations[language].description;
    document.getElementById('submitButton').textContent = translations[language].requestAccess;
    document.getElementById('email').setAttribute('placeholder', translations[language].enterEmail);
}

// Установка языка по умолчанию
document.addEventListener('DOMContentLoaded', () => {
    const userLang = navigator.language || 'ru';
    const defaultLang = userLang.startsWith('en') ? 'en' : 'ru';
    document.documentElement.lang = defaultLang;
    changeLanguage(defaultLang);

    // Кнопка для смены языка
    const languageToggle = document.getElementById('languageToggle');
    languageToggle.addEventListener('click', () => {
        const currentLang = document.documentElement.lang;
        const newLang = currentLang === 'ru' ? 'en' : 'ru';
        document.documentElement.lang = newLang;
        changeLanguage(newLang);
    });
});
