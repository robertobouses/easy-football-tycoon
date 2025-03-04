# Project Structure

This document explains the organization of the project's folder and file structure.  
It helps developers understand how different parts of the project are organized and where to find specific components.

## Folder Overview

This project follows a **Layered Architecture** pattern, which is commonly used for separating different concerns in the system. The architecture helps to maintain a clean, modular codebase that can scale and be easily maintained. The main layers in this project are:

- **App Layer**: Contains the core business logic and functionality (found in the `app` folder).
- **API Layer**: Manages the HTTP routes and endpoints, exposing services to external consumers (found in the `http` and `api` folders).
- **Repository Layer**: Handles data access and persistence, interacting with the database (found in the `repository` folder).
- **Infrastructure Layer**: Contains configurations and migrations for database setup (found in the `migrations` folder).

This layered approach allows us to separate concerns, making it easier to manage, test, and extend each part of the application independently.


The project has the following main folders:
easy-football-tycoon
│   .env.config                       # Configuration file with environment variables for database and server setup
│   .gitignore                         # Specifies files and directories to be ignored by Git
│   dev.go                             # Development-related code or configurations
│   easy-football-tycoon.exe           # Executable file for the project
│   go.mod                             # Go module dependencies
│   go.sum                             # Go checksum file for module verification
│   main.go                            # Main entry point of the application
│   README.md                          # Project documentation and basic setup instructions
│   Taskfile.yml                       # Task automation file, possibly for deployment or builds
│
├───app
│   │   analytics.go                   # Analytics-related business logic
│   │   app_service.go                 # Main application service logic
│   │   auto_player_decline_by_injury.go  # Logic to automatically handle player injury decline
│   │   auto_player_development_by_age.go  # Automatically calculate player development based on age
│   │   auto_player_development_by_training.go  # Automatically calculate player development via training
│   │   calculate_ball_possession.go      # Calculates ball possession statistics for the game
│   │   calculate_player_stats.go         # Logic for calculating player statistics
│   │   calculate_player_stats_assists.go  # Calculates player assists statistics
│   │   calculate_player_stats_chances.go  # Calculates player chances statistics
│   │   calculate_player_stats_goals.go    # Calculates player goals statistics
│   │   calculate_rating.go               # Calculates player/team ratings
│   │   calculate_strategy.go             # Calculates strategy-related data
│   │   calculate_ticket_sales.go         # Calculates ticket sales
│   │   calculate_total_salaries.go       # Calculates total team salaries
│   │   calendary.go                      # Manages calendary-related operations
│   │   errors.go                         # Handles error-related operations
│   │   get_analytics.go                  # Fetches analytics data
│   │   get_calendary.go                  # Fetches calendary data
│   │   get_current_rival.go              # Fetches the current rival information
│   │   get_lineup.go                     # Fetches the team lineup
│   │   get_player_stats.go               # Fetches player statistics
│   │   get_resume.go                     # Fetches game resume information
│   │   get_rival.go                      # Fetches rival team data
│   │   get_strategy.go                   # Fetches strategy data
│   │   get_team.go                       # Fetches team data
│   │   get_team_staff.go                 # Fetches team staff data
│   │   lineup.go                         # Logic to manage lineup-related functions
│   │   match.go                          # Match-related logic and data
│   │   match_events.go                   # Logic for match events processing
│   │   player.go                         # Manages player-related operations
│   │   player_stats.go                   # Handles player stats-related functions
│   │   post_analtytics.go                # Posts analytics data
│   │   post_calendary.go                 # Posts calendary data
│   │   post_lineup.go                    # Posts lineup data
│   │   post_rival.go                     # Posts rival data
│   │   post_strategy.go                  # Posts strategy data
│   │   post_team.go                      # Posts team data
│   │   post_team_staff.go                # Posts team staff data
│   │   probabilistic_increment.go        # Handles probabilistic increments (e.g., player development)
│   │   process_injury.go                 # Processes player injuries
│   │   process_match_play.go             # Processes match play data
│   │   process_match_play_utils.go       # Utility functions for processing match play
│   │   process_match_simulation.go      # Simulates match play
│   │   process_player_sale.go            # Handles player sale operations
│   │   process_player_signing.go         # Handles player signing operations
│   │   process_staff_sale.go             # Handles staff sale operations
│   │   process_staff_signing.go          # Handles staff signing operations
│   │   rival.go                          # Rival-related logic
│   │   strategy.go                       # Strategy-related logic and functions
│
│   ├───signings
│   │       auto_player_generator.go      # Logic for auto-generating player signings
│   │       calculate_player_atributes.go # Calculates player attributes for signing
│   │       calculate_player_fee_and_salary.go # Calculates player signing fee and salary
│   │       get_signings.go               # Fetches player signing data
│   │       post_signings.go              # Posts player signing data
│   │       signings.go                   # General logic for player signings
│   │       signings_service.go           # Service functions related to player signings
│
│   └───staff
│           auto_staff_generator.go       # Logic for auto-generating staff signings
│           calculate_staff_atributes.go  # Calculates staff attributes for signing
│           calculate_staff_fee_and_salary.go # Calculates staff signing fee and salary
│           post_staff.go                 # Posts staff data
│           staff.go                      # General staff-related logic
│           staff_service.go              # Service functions related to staff management
│
├───docs
│       api.md                           # API documentation
│       environment.md                   # Documentation on environment variables
│       installation.md                  # Instructions for installing and setting up the project
│       project-structure.md             # Explanation of the project structure
│
├───http
│   │   server.go                        # Main HTTP server setup and configuration
│   ├───analytics
│   │       get_analytics.go              # Fetch analytics data via HTTP
│   │       handler.go                   # HTTP handler for analytics
│   │       post_analytics.go             # Post analytics data via HTTP
│   ├───calendary
│   │       get_calendary.go              # Get calendary-related data via HTTP
│   │       handler.go                   # HTTP handler for calendary
│   │       post_calendary.go             # Post calendary data via HTTP
│   ├───lineup
│   │       get_lineup.go                 # Get lineup data via HTTP
│   │       handler.go                   # HTTP handler for lineup
│   │       post_lineup.go                # Post lineup data via HTTP
│   ├───player_stats
│   │       get_player_stats.go           # Get player statistics via HTTP
│   │       handler.go                   # HTTP handler for player stats
│   ├───resume
│   │       get_resume.go                 # Get match/resume data via HTTP
│   │       handler.go                   # HTTP handler for resume data
│   │       post_match_decision.go        # Post match decision via HTTP
│   │       post_match_next_event.go      # Post next match event decision
│   │       post_player_sale_decision.go  # Post player sale decision via HTTP
│   │       post_player_signing_decision.go # Post player signing decision
│   │       post_staff_sale_decision.go   # Post staff sale decision via HTTP
│   │       post_staff_signing_decision.go # Post staff signing decision via HTTP
│   ├───rival
│   │       get_rival.go                  # Get rival data via HTTP
│   │       handler.go                   # HTTP handler for rival
│   │       post_rival.go                 # Post rival data via HTTP
│   ├───signings
│   │       get_signings.go               # Get signings data via HTTP
│   │       handler.go                   # HTTP handler for signings
│   │       post_auto_player_generator.go # Post auto-generated player signings via HTTP
│   │       post_signings.go              # Post player signings via HTTP
│   ├───staff
│   │       handler.go                   # HTTP handler for staff
│   │       post_auto_player_generator.go # Post auto-generated staff signings via HTTP
│   │       post_staff.go                 # Post staff data via HTTP
│   ├───strategy
│   │       get_strategy.go               # Get strategy data via HTTP
│   │       handler.go                   # HTTP handler for strategy
│   │       post_strategy.go              # Post strategy data via HTTP
│   ├───team
│   │       get_team.go                   # Get team data via HTTP
│   │       handler.go                   # HTTP handler for team data
│   │       post_team.go                  # Post team data via HTTP
│   └───team_staff
│           get_team_staff.go             # Get team staff data via HTTP
│           handler.go                   # HTTP handler for team staff
│           post_team_staff.go            # Post team staff data via HTTP
│
├───internal
│       postgres.go                       # Internal PostgreSQL database connection logic
│
├───migrations
│       000001_create_eft_schema.down.sql  # Schema creation (down migration)
│       000001_create_eft_schema.up.sql    # Schema creation (up migration)
│       000002_create_eft_team.down.sql    # Team table migration (down)
│       000002_create_eft_team.up.sql      # Team table migration (up)
│       000003_create_eft_lineup.down.sql  # Lineup table migration (down)
│       000003_create_eft_lineup.up.sql    # Lineup table migration (up)
│       000004_create_etf_rivals.sql.down.sql # Rival table migration (down)
│       000004_create_etf_rivals.sql.up.sql   # Rival table migration (up)
│       000005_create_etf_signings.sql.down.sql # Signings table migration (down)
│       000005_create_etf_signings.sql.up.sql   # Signings table migration (up)
│       000006_create_etf_calendary.sql.down.sql # Calendary table migration (down)
│       000006_create_etf_calendary.sql.up.sql   # Calendary table migration (up)
│       000007_create_etf_analytics.sql.down.sql # Analytics table migration (down)
│       000007_create_etf_analytics.sql.up.sql   # Analytics table migration (up)
│       000008_create_eft_bank.down.sql     # Bank table migration (down)
│       000008_create_eft_bank.up.sql       # Bank table migration (up)
│       000009_create_eft_staff.down.sql    # Staff table migration (down)
│       000009_create_eft_staff.up.sql      # Staff table migration (up)
│       000010_create_eft_team_staff.down.sql # Team staff table migration (down)
│       000010_create_eft_team_staff.up.sql   # Team staff table migration (up)
│       000011_create_eft_match.down.sql    # Match table migration (down)
│       000011_create_eft_match.up.sql      # Match table migration (up)
│       000012_create_eft_strategy.down.sql # Strategy table migration (down)
│       000012_create_eft_strategy.up.sql   # Strategy table migration (up)
│       000013_create_eft_player_stats.down.sql # Player stats table migration (down)
│       000013_create_eft_player_stats.up.sql   # Player stats table migration (up)
│
└───repository
    ├───analytics
    │   │   get_analytics.go         # Repository logic for getting analytics data
    │   │   post_analytics.go        # Repository logic for posting analytics data
    │   │   repository.go            # Main repository logic for analytics
    │   └───sql
    │       get_analytics.sql        # SQL to fetch analytics data
    │       post_analytics.sql       # SQL to insert or update analytics data
    │
    ├───bank
    │   │   get_balance.go           # Repository logic for getting bank balance
    │   │   post_transactions.go     # Repository logic for posting transactions
    │   │   repository.go            # Main repository logic for bank
    │   └───sql
    │       get_balance.sql          # SQL to fetch bank balance data
    │       post_transactions.sql    # SQL to insert or update transactions data
    │
    ├───calendary
    │   │   get_calendary.go         # Repository logic for getting calendary data
    │   │   post_calendary.go        # Repository logic for posting calendary data
    │   │   repository.go            # Main repository logic for calendary
    │   └───sql
    │       get_calendary.sql        # SQL to fetch calendary data
    │       post_calendary.sql       # SQL to insert or update calendary data
    │
    ├───lineup
    │   │   delete_player_from_lineup.go # Repository logic to delete player from lineup
    │   │   get_lineup.go            # Repository logic for getting lineup data
    │   │   player_exists_in_lineup.go  # Repository logic for checking player in lineup
    │   │   post_lineup.go           # Repository logic for posting lineup data
    │   │   repository.go            # Main repository logic for lineup
    │   └───sql
    │       delete_player_from_lineup.sql # SQL to delete player from lineup
    │       get_lineup.sql           # SQL to fetch lineup data
    │       player_exists_in_lineup.sql  # SQL to check player exists in lineup
    │       post_lineup.sql          # SQL to insert or update lineup data
    │
    ├───match
    │   │   get_matches.go           # Repository logic for getting match data
    │   │   post_match.go            # Repository logic for posting match data
    │   │   repository.go            # Main repository logic for match
    │   └───sql
    │       get_matches.sql          # SQL to fetch match data
    │       post_match.sql           # SQL to insert or update match data
    │
    ├───player_stats
    │   │   get_player_stats.go      # Repository logic for getting player stats data
    │   │   get_player_stats_by_id.go # Repository logic for getting player stats by ID
    │   │   repository.go            # Main repository logic for player stats
    │   │   update_player_stats.go   # Repository logic for updating player stats
    │   └───sql
    │       get_player_stats.sql     # SQL to fetch player stats data
    │       get_player_stats_by_id.sql # SQL to fetch player stats by ID
    │       update_player_stats.sql  # SQL to update player stats data
    │
    ├───rival
    │   │   get_rival.go             # Repository logic for getting rival data
    │   │   post_rival.go            # Repository logic for posting rival data
    │   │   repository.go            # Main repository logic for rival
    │   └───sql
    │       get_rival.sql            # SQL to fetch rival data
    │       post_rival.sql           # SQL to insert or update rival data
    │
    ├───signings
    │   │   delete_signing.go        # Repository logic for deleting signings
    │   │   get_signings.go          # Repository logic for getting signings data
    │   │   get_signings_random_by_analytics.go # Repository logic for random signings by analytics
    │   │   post_signings.go         # Repository logic for posting signings data
    │   │   repository.go            # Main repository logic for signings
    │   └───sql
    │       delete_signing.sql       # SQL to delete signings
    │       get_signings.sql         # SQL to fetch signings data
    │       get_signings_random_by_analytics.sql # SQL to get random signings by analytics
    │       post_signings.sql        # SQL to insert or update signings data
    │
    ├───staff
    │   │   delete_staff_signing.go  # Repository logic for deleting staff signings
    │   │   get_staff.go             # Repository logic for getting staff data
    │   │   get_staff_random_by_analytics.go # Repository logic for getting staff by analytics
    │   │   post_staff.go            # Repository logic for posting staff data
    │   │   repository.go            # Main repository logic for staff
    │   └───sql
    │       delete_staff_signing.sql # SQL to delete staff signing
    │       get_staff.sql            # SQL to fetch staff data
    │       get_staff_random_by_analytics.sql # SQL to get staff by analytics
    │       post_staff.sql           # SQL to insert or update staff data
    │
    ├───strategy
    │   │   get_strategy.go         # Repository logic for getting strategy data
    │   │   post_strategy.go        # Repository logic for posting strategy data
    │   │   repository.go           # Main repository logic for strategy
    │   └───sql
    │       get_strategy.sql        # SQL to fetch strategy data
    │       post_strategy.sql       # SQL to insert or update strategy data
    │
    ├───team
    │   │   delete_player_from_team.go  # Repository logic to delete player from team
    │   │   get_player_by_playerid.go  # Repository logic for getting player by player ID
    │   │   get_team.go               # Repository logic for getting team data
    │   │   post_team.go              # Repository logic for posting team data
    │   │   repository.go            # Main repository logic for team
    │   │   update_player_data.go     # Repository logic for updating player data
    │   │   update_player_lined_status.go # Repository logic for updating player's lined status
    │   └───sql
    │       delete_player_from_team.sql  # SQL to delete player from team
    │       get_player_by_playerid.sql  # SQL to fetch player by player ID
    │       get_team.sql               # SQL to fetch team data
    │       post_team.sql              # SQL to insert or update team data
    │       update_player_data.sql     # SQL to update player data
    │       update_player_lined_status.sql # SQL to update player's lined status
    │
    └───team_staff
        │   delete_team_staff.go      # Repository logic for deleting team staff
        │   get_team_staff.go         # Repository logic for getting team staff data
        │   post_team_staff.go        # Repository logic for posting team staff data
        │   repository.go             # Main repository logic for team staff
        └───sql
            delete_team_staff.sql     # SQL to delete team staff data
            get_team_staff.sql        # SQL to fetch team staff data
            post_team_staff.sql       # SQL to insert or update team staff data


## Key Files

- `.env.config`: Configuration file for environment variables (to be renamed `.env`).
- `README.md`: Project overview and documentation.
- `docker-compose.yml`: Docker configuration file for setting up the project environment.
- `.gitignore`: Defines which files and directories should be ignored by Git (e.g., `.env`).

## Detailed Folder Breakdown

### `app/signings/`
This folder contains the business logic related to managing player signings.  
You’ll find the core services that handle player data, contract details, and validation.

### `http/analytics/`
This folder includes the API routes and controllers for retrieving team analytics data, such as player performance and match statistics.

### `repository/analytics/sql/`
This folder contains SQL files specifically for fetching or manipulating analytics data within the database.

## Conclusion

This project structure is designed to separate different concerns (application logic, API routes, database interactions) to keep the codebase modular and maintainable.  
If you have any questions or need clarification, feel free to reach out to the team or refer to the documentation.
