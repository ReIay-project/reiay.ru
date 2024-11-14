const express = require('express');
const bodyParser = require('body-parser');
const { Pool } = require('pg');

const app = express();
const port = 5501; // Изменили порт на 5401

// Настройки подключения к базе данных PostgreSQL
const pool = new Pool({
    user: 'postgres',
    host: 'localhost',
    database: 'relay_db',
    password: 'zxsaqw1234Q!q', // ваш пароль
    port: 5432,
});

app.use(bodyParser.json());

// Маршрут для обработки запроса на доступ
app.post('/request-access', async (req, res) => {
    const email = req.body.email;

    // Проверка валидности email
    if (!email || !email.includes('@')) {
        return res.status(400).json({ message: 'Некорректный email' });
    }

    try {
        // Проверка, существует ли уже email
        const result = await pool.query('SELECT COUNT(*) FROM waiting_requests WHERE email = $1', [email]);
        if (parseInt(result.rows[0].count) > 0) {
            return res.status(200).json({ message: 'Этот email уже зарегистрирован.', redirect: '/request-sent' });
        }

        // Добавление email в таблицу
        await pool.query('INSERT INTO waiting_requests (email, approved) VALUES ($1, $2)', [email, false]);

        res.status(200).json({ message: 'Запрос успешно отправлен. Ожидайте подтверждения доступа.' });
    } catch (error) {
        console.error('Ошибка базы данных:', error);
        res.status(500).json({ message: 'Ошибка базы данных. Проверьте настройки.' });
    }
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});
app.get('/', (req, res) => {
    res.send('Добро пожаловать в Relay API!');
});
