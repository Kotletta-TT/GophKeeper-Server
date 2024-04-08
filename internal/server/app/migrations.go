package app

// Создание таблицы для хранения пользователей
const TableUsers = `CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT,
    password TEXT
);`

// Создание таблицы для хранения секретных карточек
const TableSecretCards = ` CREATE TABLE IF NOT EXISTS secret_cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    name TEXT,
    url TEXT,
    login TEXT,
    password TEXT,
    text TEXT,
    update_time TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);`

// Создание таблицы для хранения файлов, связанных с секретными карточками
const TableFileSecretCards = `CREATE TABLE IF NOT EXISTS file_secret_cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id UUID,
    file BYTEA,
    FOREIGN KEY (card_id) REFERENCES secret_cards(id)
);`

// Создание таблицы для хранения метаданных, связанных с секретными карточками
const TableMetaSecretCards = `CREATE TABLE IF NOT EXISTS meta_secret_cards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    card_id UUID,
    key TEXT,
    value TEXT,
    FOREIGN KEY (card_id) REFERENCES secret_cards(id)
);`
