CREATE TABLE Project
(
    id    serial PRIMARY KEY,
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
    projectId   INT NOT NULL,
    FOREIGN KEY (projectId) REFERENCES Project (id) ON DELETE CASCADE ON UPDATE CASCADE,
    authorId    INT NOT NULL,
    FOREIGN KEY (authorId) REFERENCES Author (id) ON DELETE CASCADE ON UPDATE CASCADE,
    assigneeId  INT NOT NULL,
    key         TEXT,
    summary     TEXT,
    description TEXT,
    type        TEXT,
    priority    TEXT,
    status      TEXT,
    createdTime TIMESTAMP WITHOUT TIME ZONE,
    closedTime  TIMESTAMP WITHOUT TIME ZONE,
    updatedTime TIMESTAMP WITHOUT TIME ZONE,
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

