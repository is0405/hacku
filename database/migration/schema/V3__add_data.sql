INSERT INTO `user` (name, mail, age, faculty, password, gender)
VALUES ("田中", "tnaka@ac.jp", 23, 1, "65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5", 2);

INSERT INTO `user` (name, mail, age, faculty, password, gender)
VALUES ("佐藤", "sato@ac.jp", 23, 5, "65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5", 1);

INSERT INTO `user` (name, mail, age, faculty, password, gender)
VALUES ("岸", "kishi@ac.jp", 20, 10, "65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5", 1);

INSERT INTO `recruitment` (title, contents, conditions, max_participation, reward, period, submit_id)
VALUES ("A実験被験者募集", "〇〇のような実験をやるので被験者を募集しています. ", "本大学に所属している大学生", 30, "時給1000円", "1時間", 1);

INSERT INTO `recruitment` (title, contents, conditions, max_participation, reward, period, submit_id)
VALUES ("B実験被験者募集", "〇〇のような実験をやるので被験者を募集しています. ", "本大学に所属している大学生", 3, "時給2000円", "5時間", 2);

INSERT INTO `participation` (recruitment_id, user_id)
VALUES (1, 2);

INSERT INTO `participation` (recruitment_id, user_id)
VALUES (1, 3);

INSERT INTO `participation` (recruitment_id, user_id)
VALUES (1, 4);

INSERT INTO `participation` (recruitment_id, user_id)
VALUES (2, 1);

INSERT INTO `participation` (recruitment_id, user_id)
VALUES (2, 3);

