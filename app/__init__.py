from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)

# Конфигурация для подключения к базе данных PostgreSQL
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://your_user:your_password@localhost/your_db_name'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

from app import routes
