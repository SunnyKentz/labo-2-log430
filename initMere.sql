-- Create transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    magasin VARCHAR(100) NOT NULL,
    caisse VARCHAR(100) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('VENTE', 'RETOUR')),
    produit_ids VARCHAR(100) NOT NULL,
    montant DECIMAL(10,2) NOT NULL,
    deja_retourne BOOLEAN NOT NULL DEFAULT FALSE
);

-- Create employe table
CREATE TABLE IF NOT EXISTS employes (
    id SERIAL PRIMARY KEY,
    nom VARCHAR(100) NOT NULL,
    role VARCHAR(100) NOT NULL
);
-- Insert sample employees
INSERT INTO employes (nom, role) VALUES
('Alice', 'commis'),
('Bob', 'manager'),
('Claire', 'commis'),
('David', 'commis'),
('Eva', 'manager');