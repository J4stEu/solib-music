CREATE TABLE IF NOT EXISTS music (
    music_id INTEGER NOT NULL,
    album_types INTEGER[] NOT NULL,
    music_name text UNIQUE NOT NULL,
    music_info text,
    music_image text,
    file_path text UNIQUE NOT NULL,
    file_ext varchar(5) NOT NULL,
    CONSTRAINT music_pk PRIMARY KEY (music_id)
);

CREATE TABLE IF NOT EXISTS albums (
    album_id INTEGER NOT NULL,
    album_info text,
    album_image text,
    music_types INTEGER[] NOT NULL,
    CONSTRAINT album_pk PRIMARY KEY (album_id)
);

CREATE TABLE IF NOT EXISTS music_types (
    music_type_id INTEGER NOT NULL,
    music_type text NOT NULL,
    CONSTRAINT music_types_pk PRIMARY KEY (music_type_id)
);