-- Enable the UUID extension for generating UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the Company table
CREATE TABLE Company (
    ID UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    Name VARCHAR(15) UNIQUE NOT NULL,
    Description TEXT,
    Amount_Of_Employees INT NOT NULL,
    Registered BOOLEAN NOT NULL,
    Company_Type VARCHAR(20) CHECK (Company_Type IN ('Corporation', 'Non Profit', 'Cooperative', 'Sole Proprietorship')) NOT NULL
);

-- Create an index on the Name column for faster lookups
CREATE INDEX idx_company_name ON Company (Name);

-- Insert Test Company 1
INSERT INTO Company (ID, Name, Description, Amount_Of_Employees, Registered, Company_Type)
VALUES
    ('a4e9f5c0-1c05-4d58-8a9c-5d9db33ec50f', 'Company A', 'Description for Company A', 100, true, 'Corporation');

-- Insert Test Company 2
INSERT INTO Company (ID, Name, Description, Amount_Of_Employees, Registered, Company_Type)
VALUES
    ('6c34a0e1-ebc8-4a9a-810f-6a7f0b65a27c', 'Company B', NULL, 50, false, 'Non Profit');

-- Insert Test Company 3
INSERT INTO Company (ID, Name, Description, Amount_Of_Employees, Registered, Company_Type)
VALUES
    ('ad5a1e53-1e6e-4d78-975a-72a0baf487a2', 'Company C', 'Description for Company C', 200, true, 'Cooperative');
