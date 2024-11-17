#!/bin/bash

# Переход в папку с проектом
cd /var/www/reiay.ru/relay-client || { 
    echo "Папка /var/www/reiay.ru/relay-client не найдена" >> /var/log/relay_client_update.log
    exit 1
}

# Проверка обновлений в главной ветке (master)
git fetch origin master

# Сравниваем текущую ревизию с удалённой.
# Если они отличаются, выполняем pull для обновления.
if [ "$(git rev-parse HEAD)" != "$(git rev-parse @{u})" ]; then
    git pull origin master
    # Логируем успешное обновление
    echo "Relay-client обновлён: $(date)" >> /var/log/relay_client_update.log
    # Отправка уведомления на почту
    echo "Relay-client успешно обновлён на сервере в $(date)" | mail -s "Relay-client обновлён" Aimeithi@yandex.ru
else
    # Если обновлений нет, записываем об этом в лог
    echo "Обновлений нет: $(date)" >> /var/log/relay_client_update.log
fi

