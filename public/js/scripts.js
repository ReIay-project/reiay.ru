// Объект переводов
const translations = {
    en: {
        title: 'Relay Messenger',
        tagline: 'We work for you',
        description: 'Next-generation messenger',
        requestAccess: 'Request Access',
        enterEmail: 'Enter email',
        oauthLogin: 'Login with OAuth'
    },
    ru: {
        title: 'Relay Мессенджер',
        tagline: 'Мы работаем для Вас',
        description: 'Мессенджер нового поколения',
        requestAccess: 'Запросить доступ',
        enterEmail: 'Введите email',
        oauthLogin: 'Войти через OAuth'
    }
};

// Функция для изменения языка
function changeLanguage(language) {
    document.getElementById('title').textContent = translations[language].title;
    document.getElementById('tagline').textContent = translations[language].tagline;
    document.getElementById('description').textContent = translations[language].description;
    document.getElementById('submitButton').textContent = translations[language].requestAccess;
    document.getElementById('email').setAttribute('placeholder', translations[language].enterEmail);
    document.getElementById('oauthLoginButton').textContent = translations[language].oauthLogin;
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

    // Кнопка для входа через OAuth
    document.getElementById('oauthLoginButton').addEventListener('click', () => {
        window.location.href = 'https://oauth.provider.com/auth?client_id=your_client_id';
    });
});

// Плавная анимация при прокрутке
window.addEventListener('scroll', () => {
    const elements = document.querySelectorAll('.element-on-scroll');
    elements.forEach((el) => {
        const position = el.getBoundingClientRect().top;
        if (position < window.innerHeight) {
            el.classList.add('visible');
        }
    });
});

document.addEventListener('DOMContentLoaded', () => {
    const userLang = localStorage.getItem('language') || (navigator.language.startsWith('en') ? 'en' : 'ru');
    document.documentElement.lang = userLang;
    changeLanguage(userLang);

    const languageToggle = document.getElementById('languageToggle');
    languageToggle.addEventListener('click', () => {
        const currentLang = document.documentElement.lang;
        const newLang = currentLang === 'ru' ? 'en' : 'ru';
        document.documentElement.lang = newLang;
        localStorage.setItem('language', newLang);
        changeLanguage(newLang);
    });

    // Кнопка для входа через OAuth
    document.getElementById('oauthLoginButton').addEventListener('click', () => {
        window.location.href = 'https://oauth.provider.com/auth?client_id=your_client_id';
    });
});

document.addEventListener("DOMContentLoaded", () => {
    const lazyImages = document.querySelectorAll('img.lazy-load');
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const img = entry.target;
                img.src = img.dataset.src;
                img.classList.remove('lazy-load');
                observer.unobserve(img);
            }
        });
    });

    lazyImages.forEach(img => {
        observer.observe(img);
    });
});

document.getElementById('accessForm').addEventListener('submit', (e) => {
    e.preventDefault();
    const email = document.getElementById('email').value;
    const errorMessage = document.getElementById('errorMessage');
    if (!email) {
        errorMessage.textContent = 'Поле email не может быть пустым';
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
        errorMessage.textContent = 'Неверный формат email';
    } else {
        errorMessage.textContent = '';
        // Отправить запрос...
    }
});

if ('Notification' in window) {
    Notification.requestPermission().then(permission => {
        if (permission === 'granted') {
            new Notification('Добро пожаловать в Relay!');
        }
    });
}
