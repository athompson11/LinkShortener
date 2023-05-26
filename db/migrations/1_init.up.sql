CREATE TABLE IF NOT EXISTS links (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    linkid TEXT NOT NULL UNIQUE,
    link TEXT NOT NULL,
    creationtime DATETIME NOT NULL
);

CREATE INDEX indx_links_on_linkid ON links(linkid);
