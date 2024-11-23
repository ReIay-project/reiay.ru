#!/bin/bash

# Папка с репозиторием
REPO_DIR="/var/www/reiay.ru"

# Переходим в директорию репозитория
cd $REPO_DIR

# Пытаемся обновить репозиторий
git fetch origin
git reset --hard origin/master
