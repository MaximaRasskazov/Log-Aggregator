# Log Aggregator

**Сервис для сбора, обработки и анализа логов в реальном времени.**  
*Цель:* учебный проект по освоению высоконагруженных систем на Go.

## 🛠 Технологии
- **Язык:** Go (1.22+)
- **Концепции:** конкурентность, работа с файловой системой, обработка потоков данных

## 📦 Установка
```bash
# Клонирование репозитория
git clone https://github.com/ваш-username/log-aggregator.git
cd log-aggregator

# Сборка проекта
go build -o log-aggregator

# Использование
./log-aggregator -dir <директория> -keywords <ключевые_слова>

# Примеры
./log-aggregator
./log-aggregator -dir /var/logs -keywords ERROR,WARN,FAILURE
./log-aggregator -dir . -keywords CRITICAL