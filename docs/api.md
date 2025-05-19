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





2. POST / analytics
Manually create initial values ​​for user's computer analytics

URL: /analytics/create
Method: POST

{
  "training": 40,
  "finances": 12,
  "scouting": 50,
  "physiotherapy": 36,
  "totalsalaries": 100,
  "totaltraining": 80,
  "totalfinances": 60,
  "totalscouting": 90,
  "totalphysiotherapy": 70,
  "trainingtaffcount": 0,
  "financestaffcount": 0,
  "scoutingstaffcount": 0,
  "physiotherapystaffcount": 0,
  "trust": 80,
  "stadiumcapacity": 30000
}





3. POST / strategy
Manually create the initial values ​​for the user's team strategy

URL: /strategy/create
Method: POST

{
  "formation": "4-4-2",
  "playing_style": "possession",
  "game_tempo": "balanced_tempo",
  "passing_style": "short",
  "defensive_positioning": "zonal_marking",
  "build_up_play": "play_from_back",
  "attack_focus": "wide_play",
  "key_player_usage": "reference_player"
}






4. POST / rival
Manually create rival teams

URL: /rival/team
Method: POST

  {
    "rivalname": "Everton",
    "technique": 460,
    "mental": 672,
    "physique": 650
  }






5. POST / team
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






6. POST / signings
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





7. POST / signings generator
Automatically creates a number that the user indicates of players that will be stored in the database as players available for signing

URL: /signings/auto
Method: POST

{
  "numberofplayertogenerate": 2
}





8. POST / staff
Manually create a player that will be stored in the db as a player available for signing

URL: /staff/create
Method: POST

{
  "firstname": "Nadia",
  "lastname": "Curadora",
  "nationality": "fr",
  "job": "physiotherapy",
  "age": 36,
  "fee": 50,
  "salary": 20,
  "training": 99,
  "finances": 90,
  "scouting": 70,
  "physiotherapy": 70,
  "rarity": 2,
   "knowledge": 24,
  "intelligence": 44
}






9. POST / staff generator
Automatically creates a number that the user indicates of employees that will be stored in the database as employees available for signing

URL: /staff/auto
Method: POST

{
  "numberofstafftogenerate": 3
}






10. POST / team staff
Manually create a employee that belongs to the user's team

URL: /team_staff/create
Method: POST

{
  "firstname": "Kolovin",
  "lastname": "Gurnavov",
  "nationality": "fr",
  "job": "trainer",
  "age": 40,
  "fee": 50,
  "salary": 20,
  "training": 99,
  "finances": 90,
  "scouting": 70,
  "physiotherapy": 70,
  "rarity": 26,
  "knowledge": 24,
  "intelligence": 44 
}





11. POST / lineup
Send a request to put a player in the lineup based on his id

URL: /lineup/player
Method: POST

  {
    "playerid": "d88b144e-92ed-46c6-8c67-7a909c481248"
  }






12. GET / calendary
Get the simulation calendar

URL: /calendary
Method: GET





  
13. GET / team
Get all the players that belong to the team

URL: /team
Method: GET




  

14. GET / rivals
Get all the rival teams of the DB

URL: /rival
Method: GET






15. GET / lineup
Get the players that are in the lineup

URL: /lineup
Method: GET






16. GET / player stats
Gets the statistics of the players on the user's team

URL: /stats
Method: GET






17. GET / signings
Get all the players available for signing

URL: /signings
Method: GET






18. GET / team staff
Get all employees that belong to the team

URL: /team_staff
Method: GET






19. GET / analytics
Gets the values ​​from the user's team analytics

URL: /analytics
Method: GET






20. GET / strategy
Gets the values ​​from the user's team strategy

URL: /strategy
Method: GET







21. GET / resume
continues the day, advances to the next event on the calendar. causes the simulation to run

URL: /resume
Method: GET







22. POST / player signing decision
Decision on signing a player to the user's team

URL: /resume/player-signing-decision
Method: POST

  {
    "accept": true
  }  







23. POST / player sale decision
Decision on sale of a player of the user's team

URL: /resume/player-sale-decision
Method: POST

  {
    "accept": true
  }  







24. POST / staff signing decision
Decision on signing a employee to the user's team

URL: /resume/staff-signing-decision
Method: POST

  {
    "accept": true
  }  







25. POST / staff sale decision
Decision on sale of a employee of the user's team

URL: /resume/staff-sale-decision
Method: POST

  {
    "accept": true
  }  







26. POST / match decision
decision on whether to play a match

URL: /resume/match-decision
Method: POST

{
    "decision": "play"
}
 






27. POST / next match event
continue to get the next event of the current match

URL: /resume/match-next-event
Method: POST






## Contact

If you have any questions or suggestions, feel free to contact me via my GitHub profile:

[robertobouses](https://github.com/robertobouses)
