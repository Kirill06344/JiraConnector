CREATE TABLE Project
(
    id    serial PRIMARY KEY,
    key TEXT,
    title TEXT
);

CREATE TABLE Author
(
    id   serial PRIMARY KEY,
    name TEXT
);

CREATE TABLE Issue
(
    id          serial PRIMARY KEY,
    project_id   INT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES Project (id) ON DELETE CASCADE ON UPDATE CASCADE,
    author_id    INT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,
    assignee_id  INT NOT NULL,
    FOREIGN KEY (assignee_id) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,
    key         TEXT,
    summary     TEXT,
    description TEXT,
    type        TEXT,
    priority    TEXT,
    status      TEXT,
    created_time TIMESTAMP WITHOUT TIME ZONE,
    closed_time  TIMESTAMP WITHOUT TIME ZONE,
    updated_time TIMESTAMP WITHOUT TIME ZONE,
    timeSpent   INT
);

CREATE TABLE StatusChanges
(
    id         serial primary key,
    issueId    INT NOT NULL,
    FOREIGN KEY (issueId) REFERENCES Issue (id) ON DELETE CASCADE ON UPDATE CASCADE,
    authorId   INT NOT NULL,
    FOREIGN KEY (authorId) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,
    changeTime TIMESTAMP WITHOUT TIME ZONE,
    fromStatus TEXT,
    toStatus   TEXT
);

DROP ROLE IF EXISTS pguser;
CREATE USER pguser WITH ENCRYPTED PASSWORD 'pgpwd';

GRANT ALL ON ALL TABLES IN SCHEMA public TO pguser;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO pguser;



