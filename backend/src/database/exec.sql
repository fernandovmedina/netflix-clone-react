--- @date   2024 October 28, 22:41
--- @author Fernando Vazquez

--- *************** INSERTS ***************

--- ADMIN TABLE
INSERT INTO ADMINS(NAME,LASTNAME,EMAIL,HASH_PASSWORD)VALUES("Fernando","Vazquez","fernandovazquez.dev@gmail.com","8524521cfe1b095d1ad9e506d28afa451f9e23cf2f32a0229d80e1be5afa6703");

--- *************** MOVIES ***************
--- /* BAYWATCH */
INSERT INTO ACTORS(ID_ACTOR,NAME)VALUES(1,"Dwayne Johnson"),(2,"Zac Efron"),(3,"Priyanka Chopra");
INSERT INTO CATEGORIES(ID_CATEGORY,NAME)VALUES(1,"Movies to laugh"),(2,"Action"),(3,"Adventure");
INSERT INTO DIRECTORS(ID_DIRECTOR,NAME)VALUES(1,"Seth Gordon");
INSERT INTO MOVIES(ID_MOVIE,NAME,BACKGROUND_URL,DURATION,YEAR_RELEASED,DESCRIPTION,VIDEO_URL)VALUES(1,"BAYWATCH","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/backgrounds/movies/baywatch/AAAABd_4ZYqUM_Kh3uNv4sQ8s3AW972ierLkULa-S0-HlZzMuqhBsm8nbcZBabJwwrJZRQKGkRFTjkmPStUuoEuH4sO5dfL9j0ukf4a9LfXLLV6fmIdlxI40hTVyDG87PPBLy1jJyv5IjV69dsPQRfYCgi-q46sWyV5Lghc.webp","1h 57min",2017,"Dangerous waters, high velocities nad movies rescues? Part of the every day routine of this life jackets full of action and comedy","/home/fernandovmedina/workspace/go/src/github.com/fernandovmedina/netflix-clone-react/backend/public/default_video.mp4");
INSERT INTO MOVIES_has_ACTORS(ID_MOVIE,ID_ACTOR)VALUES(1,1),(1,2),(1,3);
INSERT INTO MOVIES_has_CATEGORIES(ID_MOVIE,ID_CATEGORY)VALUES(1,1),(1,2),(1,3);
INSERT INTO MOVIES_has_DIRECTORS(ID_MOVIE,ID_DIRECTOR)VALUES(1,1);
--- /* INFIERNO */
--- /* NO MANCHES FRIDA */
--- /* PAUL */
--- /* SON COMO NINOS */

--- *************** SERIES ***************
--- /* BETTER CALL SAUL */
--- /* BREAKING BAD */
--- /* DAYBREAK */
--- /* PABLO ESCOBAR */
--- /* STRANGER THINGS */
