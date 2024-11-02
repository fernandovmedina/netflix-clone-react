--- @date   2024 October 28, 22:41
--- @author Fernando Vazquez

--- *************** VARIABLES TO CHANGE ***************
SET @DEFAULT_VIDEO = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_video.mp4";

--- *************** VARIABLES TO NOT CHANGE (JUST CHANGE THE DIR "/home/.../github.com/fernandovmedina") ***************
SET @DEFAULT_ONE   = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_one.webp";
SET @DEFAULT_TWO   = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_two.webp";
SET @DEFAULT_THREE = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_three.webp";
SET @DEFAULT_FOUR  = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_four.webp";
SET @DEFAULT_FIVE  = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_five.webp";
SET @DEFAULT_SIX   = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_six.webp";
SET @DEFAULT_SEVEN = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_seven.webp";
SET @DEFAULT_EIGHT = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_eight.webp";
SET @DEFAULT_NINE  = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_nine.webp";
SET @DEFAULT_TEN   = "/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_ten.webp";

--- *************** INSERTS ***************

--- ADMIN TABLE
INSERT INTO ADMINS(NAME,LASTNAME,EMAIL,HASH_PASSWORD)VALUES("Fernando","Vazquez","fernandovazquez.dev@gmail.com","8524521cfe1b095d1ad9e506d28afa451f9e23cf2f32a0229d80e1be5afa6703");

--- *************** MOVIES ***************
--- /* BAYWATCH */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(1,"Dwayne Johnson"),(2,"Zac Efron"),(3,"Priyanka Chopra");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(1,"Movies to laugh"),(2,"Action"),(3,"Adventure");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(1,"Seth Gordon");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(1,"BAYWATCH","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/baywatch/AAAABd_4ZYqUM_Kh3uNv4sQ8s3AW972ierLkULa-S0-HlZzMuqhBsm8nbcZBabJwwrJZRQKGkRFTjkmPStUuoEuH4sO5dfL9j0ukf4a9LfXLLV6fmIdlxI40hTVyDG87PPBLy1jJyv5IjV69dsPQRfYCgi-q46sWyV5Lghc.webp","1h 57min",2017,"Dangerous waters, high velocities nad movies rescues? Part of the every day routine of this life jackets full of action and comedy",@DEFAULT_VIDEO);
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(1,1),(1,2),(1,3);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(1,1),(1,2),(1,3);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(1,1);
--- /* INFIERNO */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(4,"Damian Alcazar"),(5,"Joaquin Cosio"),(6,"Ernesto Gomez Cruz");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(4,"From Mexico"),(5,"Dramatic"),(6,"Action"),(7,"Adventure");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(2,"Luis Estrada");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(2,"El Infierno","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/infierno/AAAABWNtvjuUXXsrVi5BWcggC9ILlHDBd9BH6qXPL0YZq40etfoQBmsFhLHC9ZUCiahLOTwTCwKDqVGokcoeVQLeRbefiSrukQYADP0.webp","2h 28 min",2010,"Violent cartel members, corrupt members of the gobernment and a lot of deaths. The town were Benny was born has changed nad it would not take that much to change him.",@DEFAULT_VIDEO);
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(2,4),(2,5),(2,6);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(2,4),(2,5),(2,6),(2,7);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(2,2);
--- /* NO MANCHES FRIDA */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(7,"Omar Chaparro"),(8,"Martha Higareda"),(9,"Monica Dionne");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(8,"Comedy");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(3,"Nacho G. Velilla");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(3,"No Manches Frida","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/no-manches-frida/AAAABVZBGyqaWNKt2wSlSOwuFuN43C-3_-_O14p0haQ1frGvCJ1YJT-byVOEWQhhzZId4TARGqWe9VrVXFL5SjYPayGvbPSQjD6HvYw.webp","1h 54min",2016,"When he just comes out of the prison, an exconvict takes a job as a substitute teacher. In the same school where he hided a treasure.",@DEFAULT_VIDEO);
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(3,7),(3,8),(3,9);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(3,8);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(3,3);
--- /* PAUL */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(10,"Simon Pegg"),(11,"Nick Frost"),(12,"Seth Rogen");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(9,"Ingenious"),(10,"Comedy");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(4,"Greg Mottola");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(4,"PAUL","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/paul/AAAABdZntOqYW7kdCOvwCBapsGf14AUqIZOWVIj5_gFXwdW9nluO6nu8eKjpt-Z23i-J9Moc5qbVUp09nU1MJvf5_qy_gmRsn9tss6w.webp","1h 43min",2011,"In searching behind of the famous area 51, two science fiction fans meet an alien who tries to escape.",@DEFAULT_VIDEO);
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(4,10),(4,11),(4,12);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(4,9),(4,10);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(4,4);
--- /* SON COMO NINOS */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(13,"Adam Sandler"),(14,"Kevin James"),(15,"Chris Rock");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(11,"Comedy");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(5,"Dennis Dugan");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(5,"Son como ninos","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/son-como-ninos/AAAABR11cNBmm8eJR90-oos0_yeLqzmkVoykjK7V7sfQhAvVliqKKyZ_HiIhbbkAI5To6QSHFR-zFTY9-7muw_xfG_dLncTYjWnhupQ.webp","1h 41min",2013,"After coming back to his natal town, Lenny Feder and his friends of all his life learn that all the things are not the same as back then when they were children.",@DEFAULT_VIDEO);
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(5,13),(5,14),(5,15);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(5,11);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(5,5);

--- *************** SERIES ***************
--- /* BETTER CALL SAUL */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(16,"Bob Odenkirk"),(17,"Jonathan Banks"),(18,"Rhea Seehorn");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(12,"Dramatic series"),(13,"USA series");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(6,"Vince Gilligan"),(7,"Peter Gould");
INSERT INTO SERIES(ID_SERIE,NAME,NUMBER_OF_SEASONS,BACKGROUND_URL,YEAR_RELEASED,DESCRIPTION)VALUES(1,"Better Call Saul",6,"/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/series/better-call-saul/AAAABRyA7kGp-QnMqfHgjKvvh-BcnIjJ4ay4ZuSZQHdxkc45ETN68Ja38_Zxj42i5mIiFvgfph_7VZBmfaS1a48NUAb3Jr3Se_Sx6clSGJHgclVMpmfEAahvRTunXIciHN4sttZJ.jpg",2022,"The lawyer Jimmy McGill tries to forgets his past, but the old habits come out when a big oportunity comes out.");
INSERT INTO SERIES_has_ACTORS(ID_SERIE,ID_ACTOR)VALUES(1,16),(1,17),(1,18);
INSERT INTO SERIES_has_CATEGORIES(ID_SERIE,ID_CATEGORY)VALUES(1,12),(1,13);
INSERT INTO SERIES_has_DIRECTORS(ID_SERIE,ID_DIRECTOR)VALUES(1,6),(1,7);
INSERT INTO SEASONS(ID_SEASON,ID_SERIE,NUMBER_OF_SEASON)VALUES(1,1,1),(2,1,2);
INSERT INTO CHAPTERS(ID_CHAPTER,ID_SEASON,NAME,BACKGROUND_URL,DURATION,DESCRIPTION)VALUES
    (1,1,"Uno",@DEFAULT_ONE,"53 min","The lawyer Jimmy McGill tries to leaves his past, but the old habits come out with a big oportunity"),(2,1,"Mijo",@DEFAULT_TWO,"46 min","Jimmy is surrounded of dangerous criminals. His life will depend on how intellignet he can be"),(3,1,"Nacho",@DEFAULT_THREE,"46 min","The police caught his new and dangerours client, so jimmy must lye to the autorities without losing on trying"),(4,1,"Heroe",@DEFAULT_FOUR,"47 min","Jimmy is desesperate, but he gets creative to get new clients"),(5,1,"Pastorcito alpino",@DEFAULT_FIVE,"44 min","Jimmy test the waters in a new area of legal practice while doing his best to keep his brother out of harm's way"),(6,1,"Pasado",@DEFAULT_SIX,"42 min","Mike needs a favor and Jimmy does what he can but there is no way to calm the demons of the past"),(7,1,"Bingo",@DEFAULT_SEVEN,"47 min","As jimmy's new clients put him at a personal and professional crossroads, he enlists the help of a secret ally"),(8,1,"Rico",@DEFAULT_EIGHT,"47 min","As a lawyer for a company with dubious operations, Jimmy sets out to collect evidence and in the proccess raise Chuck"),(9,1,"Pimiento",@DEFAULT_NINE,"47 min","Jimmy's biggest case forces him to swallow his pride and seek help in the most unexpected places, meanwhile Mike returns to the job market"),(10,1,"Marco",@DEFAULT_TEN,"49 min","Jimmy visits his old town and meets with an old henchman. It wont be long before they get back to their old ways"),
    (11,2,"Gran cambio",@DEFAULT_ONE,"47 min","Unwanted Jimmy's pleasant life of freedom affect his relationship with Kim and Mike walks away from a reckless client"),(12,2,"La tarta de luna llena",@DEFAULT_TWO,"47 min","As Mike juggles controlling a volatile situation, Jimmy's old instincts surface when they try to clean up an eccentric client"),(13,2,"Amarillo",@DEFAULT_THREE,"42 min","When his attempts to get more clients for SandPiper arouse suspicion Jimmy changes tactics and finds a new usse for his astral gifts"),(14,2,"Guantes fuera",@DEFAULT_FOUR,"43 min","Thinking about the problems that could bring him in the future, Mike evaluates a lucrative job offer on his side, Jimmy faces the consequences of his campaign"),(15,2,"Rebecca",@DEFAULT_FIVE,"46 min","As Kim struggles to regain the respect of the office, an irritating associate strains Jimmy's patience, Chuck meets his past"),(16,2,"Esa famosa cancion",@DEFAULT_SIX,"46 min","Jimmy's find comfort in a familiar place while Kim weighs a life-changing proposal, an alarming threat pushes Mike to the edge"),(17,2,"Inflable",@DEFAULT_SEVEN,"42 min","Jimmy finds a creative solution to his problems at Davis & Main, Kim anxiously waits for a juicy offer to come to fruition"),(18,2,"Fifi",@DEFAULT_EIGHT,"48 min","Aware of the risk involved, Mike tries to maintain control of the problematic situation hre feared, Jimmy takes advantage of a surprising opportunity");
--- /* DAYBREAK */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(19,"Colin Ford"),(20,"Alyvia Alyn Lind"),(21,"Sophie Simnett");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(14,"Teenagers series"),(15,"Comedy"),(16,"Action");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(8,"Aron Eli Coleite"),(9,"Brad Peyton");
INSERT INTO SERIES(ID_SERIE,NAME,NUMBER_OF_SEASONS,BACKGROUND_URL,YEAR_RELEASED,DESCRIPTION)VALUES(2,"Daybreak",1,"/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/series/daybreak/AAAABRzs4cU9st72HmR5BXb4giLQHMjk4T8ffu9-m_zTW3B3_10A4CV5GaCno5ffV5piLhkptQb5GpzMoQ2RUY7Z7E-v9hVtxbIl06qPMtfZIQFJmPdSPWN2UftEM77GgjrRiut9.jpg",2019,"Los Angeles is devastated for the apocalipsis. It sounds terrible, but it is not that so. In fact, for Josh Wheeler, a normal student, the situacion got better.");
INSERT INTO SERIES_has_ACTORS(ID_SERIE,ID_ACTOR)VALUES(2,19),(2,20),(2,21);
INSERT INTO SERIES_has_CATEGORIES(ID_SERIE,ID_CATEGORY)VALUES(2,14),(2,15),(2,16);
INSERT INTO SERIES_has_DIRECTORS(ID_SERIE,ID_DIRECTOR)VALUES(2,8),(2,9);
INSERT INTO SEASONS(ID_SEASON,ID_SERIE,NUMBER_OF_SEASON)VALUES(3,2,1);
INSERT INTO CHAPTERS(ID_CHAPTER,ID_SEASON,NAME,BACKGROUND_URL,DURATION,DESCRIPTION)VALUES
    (19,3,"Josh contra el apocalipsis: Parte 1",@DEFAULT_ONE,"45 min","Los angeles is devastated for the apocalipsis, it sounds terrible but it is not"),
    (20,3,"Atrapatontos",@DEFAULT_TWO,"48 min","Josh, Angelica and Wezley hide on the mall that has not been touched"),
    (21,3,"La reina del slime de Glendale",@DEFAULT_THREE,"45 min","What hipes up Angelica? What is she looking for?"),
    (22,3,"MMMMMM-HMMMMMM",@DEFAULT_FOUR,"50 min","Josh promises to take revenge on a very important death for him"),
    (23,3,"Baile 2.0",@DEFAULT_FIVE,"48 min","Wezley plans a dance for everybody to hipe up"),
    (24,3,"7371745",@DEFAULT_SIX,"45 min","Turbo and Mona Lisa try to find the one who cheated on the team"),
    (25,3,"Canta tu vida",@DEFAULT_SEVEN,"46 min","Josh takes out Wesley, Angelica and Eli"),
    (26,3,"Amor a domicilio",@DEFAULT_EIGHT,"38 min","Before the apocalipsis, Josh and Sam didnt go to the school to enjoy a romantic day"),
    (27,3,"Josh contra el apocalipsis: Parte 2",@DEFAULT_NINE,"49 min","Josh and Eli come back to the school for diferent reasons"),
    (28,3,"CATAPLUUUUMMMMM!",@DEFAULT_TEN,"49 min","Will they beat the Principal Burr?");
--- /* BREAKING BAD */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(22,"Bryan Cranston"),(23,"Aaron Paul"),(24,"Anna Gunn");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(17,"Dramatic series"),(18,"Series about crimes"),(19,"Usa series");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(10,"Vince Gilligan");
INSERT INTO SERIES(ID_SERIE,NAME,NUMBER_OF_SEASONS,BACKGROUND_URL,YEAR_RELEASED,DESCRIPTION)VALUES(3,"Breaking Bad",2,"/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/public/backgrounds/series/breaking-bad/AAAABafEkwgV-ddaqb-IzfVJPCI7Kc4kSddq7NjUsnrOePu5t6IicZvd_lkxEjYiQkQ2sIpsVBC7Gk653tPkWAy9X9ef9dg3q7LDLpg.webp",2013,"A highschool chemical professor starts selling drugs to keep his family up");
INSERT INTO SERIES_has_ACTORS(ID_SERIE,ID_ACTOR)VALUES(3,22),(3,23),(3,24);
INSERT INTO SERIES_has_CATEGORIES(ID_SERIE,ID_CATEGORY)VALUES(3,17),(3,18),(3,19);
INSERT INTO SERIES_has_DIRECTORS(ID_SERIE,ID_DIRECTOR)VALUES(3,10);
INSERT INTO SEASONS(ID_SEASON,ID_SERIE,NUMBER_OF_SEASON)VALUES(4,3,1),(5,3,2);
INSERT INTO CHAPTERS(ID_CHAPTER,ID_SEASON,NAME,BACKGROUND_URL,DURATION,DESCRIPTION)VALUES
    (29,4,"Principio del fin",@DEFAULT_ONE,"58 min","A highschool chemical profesor starts selling drugs to keep his family up."),(30,4,"El gato en la bolsa",@DEFAULT_TWO,"48 min","After failure of their first grugs bussiness, Walt and Jesse must deseaper two bodies."),(31,4,"Y la bolsa en el rio",@DEFAULT_THREE,"48 min","While Walt cleans the disaster of their first bussiness, Walt thinks on saying Skyler everything he is doing."),(32,4,"Cancer",@DEFAULT_FOUR,"48 min","On being forced on telling the truth to his family, Walt fronts the problem of cancer with his family"),(33,4,"Materia gris",@DEFAULT_FIVE,"48 min","Skylar organices an intervention to convince Walter to accept his job expartner about the cancer treatment."),(34,4,"Un loco punado de nada",@DEFAULT_SIX,"48 min","The secondary effects and treatment cost grow up, Walt tells to Jesse they nedd to find the main supplier of the town"),(35,4,"Acuerdo no violento",@DEFAULT_SEVEN,"47 min","After jesse being about to die, Walt accepts to presents Tuco their drugs, while Skylar suspects that her sister stole a diamond necklace."),
    (36,5,"Siete con treinta y siete",@DEFAULT_ONE,"47 min","When they prepare new plans for their final bussiness drugs, Walt and Jesse start to get worried about Tuco killing them"),(37,5,"A las brasas",@DEFAULT_TWO,"48 min","While the DEA starts looking for Tuco, Skyler goes to Hank for him to help her go and find Walt"),(38,5,"Picadura de una abje muerta",@DEFAULT_THREE,"47 min","Walt tries to get back Skyler, and Jesse has no home after his parents took him out"),(39,5,"Abajo",@DEFAULT_FOUR,"47 min","Walt fights for trying to save his relationship with Skyler");
--- /* PABLO ESCOBAR */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(25,"Andres Parra"),(26,"Cecilia Nava");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(20,"Dramatic series"),(21,"Telenovels about crime"),(22,"From Colombia");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(11,"Juana Uribe"),(12,"Camilo Cano");
INSERT INTO SERIES(ID_SERIE,NAME,NUMBER_OF_SEASONS,BACKGROUND_URL,YEAR_RELEASED,DESCRIPTION)VALUES(4,"Pablo Escobar",1,"/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/series/pablo-escobar/AAAABQl4yXkB-DAKmy6mco3vS49Bz_B2p5cuFImUgxQCqEu14KTB9fkwvnho4oeqVbC_s7ZkF4McGFVeRVDPn50EtTcFxZIcrOblN2U.webp",2012,"Pablo demonstrated his skills for bussiness since he was a child. As an adult, he works for a man who introduces him to narcotrafic");
INSERT INTO SERIES_has_ACTORS(ID_SERIE,ID_ACTOR)VALUES(4,25),(4,26);
INSERT INTO SERIES_has_CATEGORIES(ID_SERIE,ID_CATEGORY)VALUES(4,20),(4,21),(4,22);
INSERT INTO SERIES_has_DIRECTORS(ID_SERIE,ID_DIRECTOR)VALUES(4,11),(4,12);
INSERT INTO SEASONS(ID_SEASON,ID_SERIE,NUMBER_OF_SEASON)VALUES(6,4,1);
INSERT INTO CHAPTERS(ID_CHAPTER,ID_SEASON,NAME,BACKGROUND_URL,DURATION,DESCRIPTION)VALUES
    (40,6,"Episode 1",@DEFAULT_ONE,"45 min","Pablo showed his skill for bussiness since he was a child"),
    (41,6,"Episode 2",@DEFAULT_TWO,"44 min","A pablos's neighboor watches a bank rob and recognizez the thieves"),
    (42,6,"Episode 3",@DEFAULT_THREE,"43 min","Pablo and Gonzalo are in jail for drug dealing. Fabio is forced to accept the relationship between his sister and Pablo"),
    (43,6,"Episode 4",@DEFAULT_FOUR,"45 min","With the planes he bought, Pablo send drugs to miami"),
    (44,6,"Episode 5",@DEFAULT_FIVE,"44 min","Pablo build the Napoles building while he realizes about the extradiction paper on Colombia"),
    (45,6,"Episode 6",@DEFAULT_SIX,"44 min","The operatives to find Irma are not going as expected"),
    (46,6,"Episode 7",@DEFAULT_SEVEN,"42 min","Pablo realizes about the political strength on his country, and contemplates to get into it"),
    (47,6,"Episode 8",@DEFAULT_EIGHT,"43 min","Pablo starts his campaign, but some of his partners disagree with him"),
    (48,6,"Episode 9",@DEFAULT_NINE,"44 min","The elections start and Pablo and Galan wait for the results, but Pablo is thinking on his next step"),
    (49,6,"Episode 10",@DEFAULT_TEN,"45 min","Pablo takes his charge. Rodrigo Lara is the new justice minister and uses his job to declare a war against drug dealing");
--- /* STRANGER THINGS */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(27,"Winona Ryder"),(28,"David Harbour"),(29,"Millie Bobby Brown");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(23,"Dramatic series"),(24,"Sci-fi series"),(25,"Teenagers series");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(13,"The Duffer Brothers");
INSERT INTO SERIES(ID_SERIE,NAME,NUMBER_OF_SEASONS,BACKGROUND_URL,YEAR_RELEASED,DESCRIPTION)VALUES(5,"Stranger Things",2,"/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/series/stranger-things/AAAABaJhbUQjS8mk9pa79jI-qjHrCibDw39Wfb2vd4E3gDAfRRy2fnPoNrqTV_6Fd6BEpM1ke1KD1DI0bYjAcRCo5mf8org11n6VkhbkahQqrbNj5sqGnh2J43f61TLqnvtYQkc4.jpg",2022,"Will watches something horrible when he comes back home, Near down there of his house, there is a gobernamental lab.");
INSERT INTO SERIES_has_ACTORS(ID_SERIE,ID_ACTOR)VALUES(5,27),(5,28),(5,29);
INSERT INTO SERIES_has_CATEGORIES(ID_SERIE,ID_CATEGORY)VALUES(5,23),(5,24),(5,25);
INSERT INTO SERIES_has_DIRECTORS(ID_SERIE,ID_DIRECTOR)VALUES(5,13);
INSERT INTO SEASONS(ID_SEASON,ID_SERIE,NUMBER_OF_SEASON)VALUES(7,5,1),(8,5,2);
INSERT INTO CHAPTERS(ID_CHAPTER,ID_SEASON,NAME,BACKGROUND_URL,DURATION,DESCRIPTION)VALUES
    (50,7,"Chapter one",@DEFAULT_ONE,"49 min","Will watches something horrible while he comes back home. Near there, a gobernamental laboratory hides a terrible secret on their deeps."),(51,7,"Chapter one",@DEFAULT_TWO,"56 min","Lucas, Mike and Dustin try to talk about the girl that they found."),(52,7,"Chapter one",@DEFAULT_THREE,"52 min","Nancy, each time more worried, looks for Barb and finds out what Jonathan is trying to do."),(53,7,"Chapter one",@DEFAULT_FOUR,"51 min","Joyce does not accept that Will is dead and tries to communicate with him."),(54,7,"Chapter one",@DEFAULT_FIVE,"53 min","Hopper brokes into the laboratory. Nancy and Jonathan fight against what took Will to another dimension"),(55,7,"Chapter one",@DEFAULT_SIX,"47 min","Jonathan looks for Nancy in the darkness and Steve does the same."),(56,7,"Chapter one",@DEFAULT_SEVEN,"42 min","Eleven tries to get to Willy, and Lucas tells'em about something terrible that's coming"),(57,7,"Chapter one",@DEFAULT_EIGHT,"55 min","Dr. Brenner stops Hoper and Joyce for an interrogatory."),
    (58,8,"Chapter one",@DEFAULT_ONE,"48 min","About some days of halloween, the guys discover that someone beat their score in some videogame, and Hopper investigates an inusual event on a pumpkin patch."),
    (59,8,"Chapter two",@DEFAULT_TWO,"56 min","Will watches something terrible on halloween's night. By the other side Mike does not lose hpe about eleven being near of them."),
    (60,8,"Chapter three",@DEFAULT_THREE,"51 min","Dustin get a new weird pet. Eleven gets desesperate with her situation, and Bob talks with Will about facing his fears"),
    (61,8,"Chapter four",@DEFAULT_FOUR,"46 min","Will's health gets worse and Joyce realizes about how big the problem is"),
    (62,8,"Chapter five",@DEFAULT_FIVE,"58 min","Nancy and Jonathan get a new partner who has his own theory about this."),
    (63,8,"Chapter six",@DEFAULT_SIX,"52 min","The bond between Will and the monster intensifies while everyone aks themselves how to stop it"),
    (64,8,"Chapter seven",@DEFAULT_SEVEN,"46 min","After a lot of vitions, Eleven leaves the town"),
    (65,8,"Chapter eight",@DEFAULT_EIGHT,"48 min","The situation on the laboratory activates a security protocol where nobody can scape of it."),
    (66,8,"Chapter nine",@DEFAULT_NINE,"62 min","Eleven knows the solution is on her hands and must finish what she started");