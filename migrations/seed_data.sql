-- 2. Insertar analytics
INSERT INTO eft.analytics (
  training, finances, scouting, physiotherapy,
  totalsalaries, totaltraining, totalfinances, totalscouting, totalphysiotherapy,
  trainingstaffcount, financestaffcount, scoutingstaffcount, physiotherapystaffcount,
  trust, stadiumcapacity
) VALUES (
  40, 12, 50, 36,
  100, 80, 60, 90, 70,
  0, 0, 0, 0,
  80, 30000
);

-- 3. Insertar strategy
INSERT INTO eft.strategy (
  formation, playing_style, game_tempo, passing_style,
  defensive_positioning, build_up_play, attack_focus, key_player_usage
) VALUES (
  '4-4-2', 'possession', 'balanced_tempo', 'short',
  'zonal_marking', 'play_from_back', 'wide_play', 'reference_player'
);

-- 4. Insertar rival
INSERT INTO eft.rivals (rivalname, technique, mental, physique) VALUES
  ('Manchester City',        900, 890, 870),
  ('Arsenal',                880, 870, 860),
  ('Liverpool',              860, 850, 840),
  ('Aston Villa',            820, 810, 800),
  ('Tottenham Hotspur',      810, 800, 790),
  ('Chelsea',                790, 770, 780),
  ('Newcastle United',       770, 760, 760),
  ('Manchester United',      750, 740, 750),
  ('West Ham United',        720, 710, 730),
  ('Crystal Palace',         700, 700, 710),
  ('Brighton & Hove Albion', 730, 720, 700),
  ('AFC Bournemouth',        690, 680, 690),
  ('Fulham',                 680, 670, 680),
  ('Wolverhampton Wanderers',670, 660, 670),
  ('Everton',                660, 650, 660),
  ('Brentford',              650, 640, 650),
  ('Nottingham Forest',      620, 610, 630),
  ('Luton Town',             600, 590, 620),
  ('Burnley',                590, 580, 610),
  ('Sheffield United',       570, 560, 600);

-- 5. Insertar player del equipo
INSERT INTO eft.team (
  firstname, lastname, nationality, position, age, fee, salary,
  technique, mental, physique, familiarity, fitness, happiness
) VALUES
  ('Alfred', 'Smith', 'eng', 'goalkeeper', 28, 15, 25, 75, 80, 85, 60, 70, 75),
  ('Lucas', 'Martinez', 'arg', 'defender', 26, 18, 22, 78, 82, 80, 65, 68, 70),
  ('John', 'Brown', 'eng', 'defender', 30, 12, 20, 74, 77, 82, 70, 69, 72),
  ('Carlos', 'Lopez', 'esp', 'defender', 24, 20, 28, 80, 79, 78, 68, 72, 74),
  ('Marco', 'Rossi', 'ita', 'defender', 27, 17, 23, 77, 76, 79, 64, 67, 70),
  ('David', 'Wilson', 'eng', 'midfielder', 25, 22, 30, 85, 88, 70, 75, 80, 78),
  ('Pierre', 'Dupont', 'fra', 'midfielder', 29, 20, 28, 83, 85, 72, 70, 75, 77),
  ('Hans', 'Schmidt', 'ger', 'midfielder', 31, 18, 26, 80, 82, 68, 73, 74, 76),
  ('Liam', 'Johnson', 'eng', 'midfielder', 23, 15, 20, 78, 80, 66, 70, 72, 73),
  ('Sergio', 'Gomez', 'esp', 'forward', 24, 25, 35, 88, 85, 70, 75, 78, 80),
  ('Oliver', 'Clark', 'eng', 'forward', 27, 23, 33, 85, 83, 72, 70, 75, 78),
  ('Ivan', 'Petrov', 'rus', 'forward', 29, 20, 30, 83, 81, 68, 72, 74, 75),
  ('Mateo', 'Silva', 'bra', 'forward', 26, 22, 31, 84, 82, 69, 74, 76, 77),
  ('Ethan', 'White', 'eng', 'goalkeeper', 32, 14, 21, 74, 78, 83, 67, 70, 72),
  ('Nicolas', 'Moreau', 'fra', 'midfielder', 28, 19, 27, 81, 84, 70, 71, 73, 75),
  ('Diego', 'Fernandez', 'arg', 'defender', 25, 18, 24, 79, 81, 77, 69, 70, 73);

-- 6. Insertar jugador para signings
INSERT INTO eft.signings (
  firstname, lastname, nationality, position, age, fee, salary,
  technique, mental, physique, injurydays, rarity, fitness
) VALUES
  ('Harry', 'Kane', 'eng', 'forward', 30, 85000, 10500000, 88, 90, 85, 5, 80, 87),
  ('Kevin', 'De Bruyne', 'bel', 'midfielder', 32, 90000, 11500000, 90, 93, 78, 3, 85, 90),
  ('Virgil', 'van Dijk', 'ned', 'defender', 31, 75000, 9800000, 85, 88, 90, 2, 78, 88),
  ('Alisson', 'Becker', 'bra', 'goalkeeper', 30, 60000, 9200000, 80, 85, 88, 0, 75, 86),
  ('Raheem', 'Sterling', 'eng', 'forward', 28, 70000, 9400000, 84, 87, 75, 4, 70, 85),
  ('Trent', 'Alexander-Arnold', 'eng', 'defender', 25, 68000, 8900000, 83, 86, 80, 1, 72, 84),
  ('Bruno', 'Fernandes', 'por', 'midfielder', 29, 72000, 9700000, 87, 89, 74, 2, 77, 88),
  ('Marc-André', 'ter Stegen', 'ger', 'goalkeeper', 30, 58000, 9000000, 79, 84, 85, 0, 74, 85),
  ('Son', 'Heung-min', 'kor', 'forward', 30, 73000, 9500000, 85, 88, 77, 3, 75, 87),
  ('Joshua', 'Kimmich', 'ger', 'midfielder', 27, 71000, 9300000, 86, 90, 79, 1, 76, 86);


-- 7. Insertar staff manual
INSERT INTO eft.staff (
  firstname, lastname, nationality, job, age, fee, salary,
  training, finances, scouting, physiotherapy,
  rarity, knowledge, intelligence
) VALUES
  ('Lucas', 'Martinez', 'es', 'trainer', 40, 60, 35, 95, 50, 55, 45, 3, 50, 70),
  ('Emma', 'Johnson', 'uk', 'financial', 45, 55, 40, 40, 98, 60, 50, 2, 70, 75),
  ('Oliver', 'Smith', 'us', 'scout', 38, 45, 30, 50, 55, 95, 40, 2, 60, 65),
  ('Sophia', 'Brown', 'de', 'physiotherapist', 35, 50, 25, 40, 50, 45, 95, 3, 55, 60),
  ('Mateo', 'Garcia', 'ar', 'trainer', 42, 65, 37, 90, 48, 50, 40, 3, 52, 68),
  ('Isabella', 'Lopez', 'it', 'financial', 39, 58, 42, 45, 96, 52, 55, 2, 68, 72),
  ('Ethan', 'Williams', 'ca', 'scout', 37, 47, 33, 52, 53, 92, 38, 2, 62, 67),
  ('Mia', 'Davis', 'au', 'physiotherapist', 36, 52, 28, 38, 48, 46, 90, 3, 57, 61),
  ('Liam', 'Miller', 'nl', 'trainer', 41, 63, 36, 88, 52, 48, 42, 3, 54, 69),
  ('Charlotte', 'Wilson', 'se', 'financial', 43, 57, 41, 44, 97, 50, 53, 2, 69, 74);


-- 8. Insertar team staff
INSERT INTO eft.team_staff (
  firstname, lastname, nationality, job, age, fee, salary,
  training, finances, scouting, physiotherapy,
  rarity, knowledge, intelligence
) VALUES
  ('Andrés', 'Ruiz', 'es', 'trainer', 45, 55, 30,
   95, 50, 60, 55,
   20, 30, 40),
   
  ('Julia', 'Martínez', 'uk', 'financial', 42, 52, 28,
   40, 98, 50, 45,
   22, 33, 38),

  ('Peter', 'Smith', 'us', 'scout', 37, 48, 25,
   45, 52, 92, 40,
   18, 28, 35),
   
  ('Laura', 'González', 'de', 'scout', 34, 46, 24,
   42, 50, 89, 38,
   19, 29, 37),

  ('Marta', 'López', 'it', 'physiotherapist', 39, 50, 27,
   38, 44, 48, 90,
   21, 31, 39);
