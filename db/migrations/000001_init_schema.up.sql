CREATE TABLE sources (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    last_scraped TIMESTAMP WITH TIME ZONE
);

CREATE TABLE songs (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    normalized_title TEXT NOT NULL,
    normalized_artist TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(normalized_title, normalized_artist)
);

CREATE TABLE charts (
    id BIGSERIAL PRIMARY KEY,
    song_id BIGINT NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    charter TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    source_id BIGINT NOT NULL REFERENCES sources(id) ON DELETE RESTRICT,
    download_url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
