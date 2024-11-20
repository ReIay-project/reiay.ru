import os

class Config:
    SQLALCHEMY_DATABASE_URI = os.environ.get('SQLALCHEMY_DATABASE_URI', '')
    SQLALCHEMY_TRACK_MODIFICATIONS = os.getenv('SQLALCHEMY_TRACK_MODIFICATIONS', 'False').lower() in ('true')
    MAIL_SERVER = os.environ.get('MAIL_SERVER', '')
    MAIL_PORT = os.environ.get('MAIL_PORT', 465)
    MAIL_USE_TLS = os.getenv('MAIL_USE_TLS', 'False').lower() in ('true')
    MAIL_USE_SSL = os.getenv('MAIL_USE_SSL', 'False').lower() in ('true')
    MAIL_USERNAME = os.environ.get('MAIL_USERNAME', '')
    MAIL_PASSWORD = os.environ.get('MAIL_PASSWORD', '')
    MAIL_DEFAULT_SENDER = os.environ.get('MAIL_DEFAULT_SENDER', '')
    MAIL_NOTIFICATION_ADDRESS = os.environ.get('MAIL_NOTIFICATION_ADDRESS', '')
