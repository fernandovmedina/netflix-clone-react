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
--- /* PABLO ESCOBAR */
--- /* STRANGER THINGS */
