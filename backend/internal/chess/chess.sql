CREATE TABLE IF NOT EXISTS game_state (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    fen TEXT NOT NULL,
    captures TEXT,
    white TEXT,
    black TEXT,
    white_draw INTEGER NOT NULL DEFAULT 0 CHECK (white_draw IN (0, 1)),
    black_draw INTEGER NOT NULL DEFAULT 0 CHECK (black_draw IN (0, 1)),
    last_move DATETIME,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS moves (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    move_number INTEGER NOT NULL,
    player TEXT NOT NULL,
    move_text TEXT NOT NULL,
    fen_before TEXT NOT NULL,
    fen_after TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);