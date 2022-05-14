CREATE TABLE IF NOT EXISTS public.collaborators
(
    id serial PRIMARY KEY,
    name VARCHAR ( 250 ) NOT NULL,
    login VARCHAR ( 50 ) NOT NULL,
    occupation VARCHAR ( 50 ) NOT NULL
)


CREATE TABLE IF NOT EXISTS public.tables
(
    id serial PRIMARY KEY,
    places INT NOT NULL
)


CREATE TABLE IF NOT EXISTS public.leases
(
    id serial PRIMARY KEY,
    date DATE NOT NULL,
    hour TIME NOT NULL,
	collaborator_id INT NOT NULL,
    FOREIGN KEY (collaborator_id) REFERENCES collaborators (id)
)

CREATE TABLE IF NOT EXISTS public.offices
(
    id serial PRIMARY KEY,
    office_number INT NOT NULL,
    hour TIME NOT NULL,
    date DATE NOT NULL,
    table_id INT NOT NULL,
    collaborator_id INT NOT NULL,
    lease_id INT NOT NULL,
    FOREIGN KEY (lease_id) REFERENCES leases (id),
    FOREIGN KEY (collaborator_id) REFERENCES collaborators (id),
    FOREIGN KEY (table_id) REFERENCES tables (id) 
)