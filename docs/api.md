API Documentation
Overview
This document provides detailed information about the API for the Easy Football Tycoon project. The API allows interaction with various resources such as teams, players, staff, and match data. It uses REST principles and returns data in JSON format.

Base URL
https://api.easy-football-tycoon.com/v1/


It is important to follow this order when applying the endpoints, since we need to have certain data in our database:

Endpoints

1. POST / calendary
Creates the calendar by which the simulation will be developed. This is a function that generates dynamism to change the order of events.

URL: /calendary/create
Method: POST



2. POST / team
Manually create a player that belongs to the user's team

URL: /team/player
Method: POST

{
    "firstname": "Gonzalo",
    "lastname": "Gomez",
    "nationality": "es",
    "position": "Forward",
    "age": 39,
    "fee": 20,
    "salary": 30,
    "technique": 82,
    "mental": 99,
    "physique": 61,
    "familiarity": 50,
    "fitness": 51,
    "happiness": 52
}





3. POST / rival
Manually create rival teams

URL: /rival/team
Method: POST

  {
    "rivalname": "Everton",
    "technique": 460,
    "mental": 672,
    "physique": 650
  }



4. POST / signings
Manually create a player that will be stored in the db as a player available for signing

URL: /signings/person
Method: POST

 {
  "firstname": "Lucas",
  "lastname": "Paquetá",
  "nationality": "br",
  "position": "Midfielder",
  "age": 27,
  "fee": 70000,
  "salary": 9100500,
  "technique": 69,
  "mental": 92,
  "physique": 89,
  "injurydays": 0,
  "rarity": 69,
  "fitness": 83
}


5. POST / signings generator
Automatically creates a number that the user indicates of players that will be stored in the database as players available for signing

URL: /signings/auto
Method: POST

 {
  "firstname": "Lucas",
  "lastname": "Paquetá",
  "nationality": "br",
  "position": "Midfielder",
  "age": 27,
  "fee": 70000,
  "salary": 9100500,
  "technique": 69,
  "mental": 92,
  "physique": 89,
  "injurydays": 0,
  "rarity": 69,
  "fitness": 83
}




## Contact

If you have any questions or suggestions, feel free to contact me via my GitHub profile:

[robertobouses](https://github.com/robertobouses)
