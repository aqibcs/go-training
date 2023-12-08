CREATE TABLE branch (
    BranchID INTEGER PRIMARY KEY,
    BranchName VARCHAR(100),
    Capacity INTEGER DEFAULT 200
);

CREATE TABLE students (
    StudentID INTEGER PRIMARY KEY,
    FirstName VARCHAR(100) NOT NULL,
    LastName VARCHAR(100),
    BranchID INTEGER,
    Email VARCHAR(100) NOT NULL,
    CONSTRAINT fk_BranchID FOREIGN KEY (BranchID) REFERENCES branch(BranchID)
);
