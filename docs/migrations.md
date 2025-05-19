One way to create data for the db:

📦 Database Setup Instructions
This guide explains how to set up the database for the Easy Football Tycoon project, including schema migrations, seed data, and pending manual inserts via API.

🔧 1. Run Schema Migrations
This step creates all tables and schemas defined for the project.

bash
Copiar
Editar
task migrate-up
This will execute all migration files inside the migrations/ directory and apply the database structure to the database defined in the .env.config file.

🌱 2. Run Seed Data Migration
This step inserts initial data into the database (sample teams, players, strategies, etc.).

bash
Copiar
Editar
task seed-db
This runs the seed_data.sql file and populates most of the tables.

⚠️ Note: The calendary table is intentionally excluded from this step due to a NOT NULL constraint on the dayType field. These records must be inserted through the API to ensure proper validation and logic.

📬 3. Insert Calendar Data via API
After running the migrations, use the POST /calendary endpoint to add day entries.

Example:

POST /calendary
Content-Type: application/json

{
  "id": 4,
  "dayType": "Training"
}
Repeat this for each calendar day required.
