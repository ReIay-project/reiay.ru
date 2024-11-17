from flask import Flask, request, jsonify
from flask_mail import Mail, Message
from config import Config
from models import db, WaitingRequest
import logging
import re

app = Flask(__name__)
app.config.from_object(Config)
app.logger.setLevel(logging.INFO)
db.init_app(app)
mail = Mail(app)


def validate_email(email):
    if not email:
        return False
    pattern = r'^[^\s@]+@[^\s@]+\.[^\s@]+$'
    if re.match(pattern, email):
        return True
    else:
        return False


def save_request(email):
    request = WaitingRequest(email=email)
    db.session.add(request)
    db.session.commit()


def send_request_notification(email):
    msg = Message(
        'Запрос на доступ',
        recipients=[app.config['MAIL_NOTIFICATION_ADDRESS']],
        body=f'Пользователь с email {email} запросил доступ'
    )
    mail.send(msg)


@app.route('/request-access', methods=['POST'])
def request_access():
    try:
        content = request.json
        if not content:
            app.logger.info('request-access: ошибка валидации, пустое тело запроса')
            return '', 400
        email = content['email']
        if not validate_email(email):
            app.logger.info(f'request-access: ошибка валидации, некорректное значение email {email}')
            return jsonify({'error': 'invalid-email'}), 400
        save_request(email)
        app.logger.info(f'request-access: запрос сохранён в базу данных для email {email}')
        send_request_notification(email)
        app.logger.info(f'request-access: отправлено уведомление на почту для запроса от email {email}')
        return '', 200
    except Exception as e:
        message = e.message if hasattr(e, 'message') else ''
        app.logger.error(f'request-access: при выполнении запроса возникла ошибка {message}')
        return '', 500


if __name__ == "__main__":
    with app.app_context():
        db.create_all()
    app.run(host='0.0.0.0', port=80)
