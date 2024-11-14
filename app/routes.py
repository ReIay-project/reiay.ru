from app import app, db
from flask import request, jsonify
import smtplib
from email.mime.text import MIMEText
from app.models import WaitingRequest

# Настройки SMTP для Yandex
SMTP_SERVER = "smtp.yandex.ru"
SMTP_PORT = 587
EMAIL_ADDRESS = "Aimeithi@yandex.ru"
EMAIL_PASSWORD = "your_app_password"   # Пароль приложения Yandex

@app.route('/request-access', methods=['POST'])
def request_access():
    # Получаем email из формы
    email = request.form.get('email')

    # Проверка на пустое поле
    if not email:
        return jsonify({'error': 'Поле email обязательно'}), 400

    # Запись в базу данных PostgreSQL
    new_request = WaitingRequest(email=email)
    db.session.add(new_request)
    db.session.commit()

    # Формируем текст письма
    subject = "Запрос на доступ"
    body = f"Пользователь с email {email} запросил доступ."

    # Создаем сообщение
    msg = MIMEText(body)
    msg['Subject'] = subject
    msg['From'] = EMAIL_ADDRESS
    msg['To'] = EMAIL_ADDRESS

    try:
        # Отправка письма через SMTP
        with smtplib.SMTP(SMTP_SERVER, SMTP_PORT) as server:
            server.starttls()
            server.login(EMAIL_ADDRESS, EMAIL_PASSWORD)
            server.sendmail(EMAIL_ADDRESS, EMAIL_ADDRESS, msg.as_string())

        return jsonify({'message': 'Запрос успешно отправлен!'}), 200
    except Exception as e:
        return jsonify({'error': str(e)}), 500
