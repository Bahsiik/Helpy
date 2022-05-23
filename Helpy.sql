-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 23 mai 2022 à 22:19
-- Version du serveur : 8.0.27
-- Version de PHP : 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `forum`
--

-- --------------------------------------------------------

--
-- Structure de la table `post`
--

DROP TABLE IF EXISTS `post`;
CREATE TABLE IF NOT EXISTS `post` (
  `Post_id` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Content` varchar(1000) NOT NULL,
  `creation_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `reply_number` int DEFAULT '0',
  `Topic_id` int NOT NULL,
  `User_id` int NOT NULL,
  PRIMARY KEY (`Post_id`),
  KEY `Topic_id_fk` (`Topic_id`),
  KEY `User_id_fk` (`User_id`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `post`
--

INSERT INTO `post` (`Post_id`, `Title`, `Content`, `creation_date`, `reply_number`, `Topic_id`, `User_id`) VALUES
(66, 'Problème avec mon HTML', 'Bonjour les gars, pour le forum j\'ai un petit problème je n\'arrive pas a centrer ma div, on pourrait m\'aider svp ?', '2022-05-23 16:54:40', 2, 5, 32),
(67, 'Bug Wamp Serveur', 'En lançant mon serveur wamp je  me rend compte que j\'en ai que 2 d\'ouvert, je ne sais pas pourquoi ça me fait ça', '2022-05-23 17:28:44', 1, 5, 32),
(68, 'Javascript dans une instruction Twig', 'Bonjour,\r\nDans ma vue Twig, j\'ai à un moment besoin d\'ajouter une redirection vers une autre page avec un paramètre.\r\nCependant je voudrais que ce soit dynamique, donc que le paramètre soit une fonction qui retourne une valeur, comme ci-dessous:\r\n\r\nhref=\"{{path(\'app_compute\', {id : \'test()\'} )}}\"\r\nJ\'ai compris que cela ne marche pas parce qu\'il va envoyer le texte \'test()\' au lieu du résultat, mais je voudrais savoir comment faire pour avoir la valeur retournée par la fonction au lieu du texte.\r\n\r\nN.B: La fonction test() est une fonction Javascript, qui retourne une simple valeur.\r\n\r\nMerci d\'avance', '2022-05-23 17:31:14', 1, 5, 32),
(69, 'Aide bouton double action HTML', 'Bonjour,\r\n\r\nAlors voila , je cherche a savoir comment faire pour qu\'un bouton ait une double action , avec verification de la premiere.Pour faire simple , ce bouton permet de placer un pixel sur un canvas , mais je souhaite que pour faire cela les gens fasse un don sur Paypal de 10 centimes. Est ce que vous auriez une idee de comment faire ca , merci davance et bonne journée', '2022-05-23 17:32:15', 1, 5, 32),
(70, 'Menu non afficher à cause d\'une div', 'Bonjour, j\'ai un menu sur mobile sauf que le problème c\'est que quand j\'ouvre le menu, je ne vois que la div.', '2022-05-23 17:33:14', 0, 5, 32),
(71, 'Code HTML pour l\'affichage d\'un champ déjà saisi', 'Bonjour,\r\nsur mon site, j\'ai deux champ d\'un formulaire qui sont remplis sur une page (le champ N° de série et le champ email)\r\n\r\nhttps://data-perso.online/user/456/\r\n\r\nje souhaite afficher sur cette meme page uniquement le contenu des champs via html en dessous des éléments existants.\r\n\r\nJ\'ai tenté nombre de manipulations issues du net mais rien n\'y fait...\r\n\r\nSi vous aviez une piste d\'un code simple CE SERAIT MAGNIFIQUE !\r\nplus d\'une semaine mais pas de résultat, certainement du à mon niveau plus que faible dans ce domaine...\r\n\r\nBien à vous', '2022-05-23 17:34:21', 2, 5, 33),
(72, 'Je galère pour des trucs basiques', 'Hello !\r\nJ\'ai un mini projet que je galère a realiser parce que je suis naze en HTML haha.\r\nJ\'aimerais créer un bouton qui redirige vers une page QUE si l\'input est correct (comme une sorte de mot de passe).\r\nJ\'ai essayé plein d\'approche mais je galère comme pas possible avec les conditions.\r\nSi quelqu\'un pouvais m\'aider ça serait incroyable\r\nJ\'ai mis un petit schema de ce que j\'ai en tete.\r\n\r\nhttps://imgur.com/a/Vv005qc\r\n\r\nComme vous voyez c\'est ultra simple...\r\nJe vais pas vous mettre mes tentatives j\'en ai trop honte haha.\r\n\r\nMerci beaucoup pour votre attention :)', '2022-05-23 17:34:58', 5, 5, 33),
(73, 'Aide pour travail sur du javascript\r\n', 'Bonjour à tous,\r\nJe vous demande votre aide aujourd\'hui car je n\'ai pas réussi à répondre à une question sur du javascript d\'un travail que je dois rendre.\r\nIl s\'agit de la question numéro 4 dans laquelle il m\'est demandé de compléter la fonction createFigureBlock.\r\nVoici le lien de l\'exercice : https://www.fil.univ-lille.fr/~abdelkaf ... cript.html\r\n\r\nMerci d\'avance et une très bonne journée à vous.:(', '2022-05-23 17:35:56', 0, 5, 33),
(74, 'Questionnaire transformation digitale', 'Bonjour à tous,\r\nDans le cadre de mes études et comme beaucoup d\'étudiants, j\'ai le devoir de rédiger un mémoire de fin d\'année.\r\nCompte tenu de mon expérience et connaissance, j\'ai fait le choix de m\'intéresser au sujet de la transformation digitale, qui est pour moi un sujet passionnant. Cependant, j\'ai besoin de votre aide, de manière à partager les connaissances et expériences de tous à ce sujet.\r\nPouvez-vous svp prendre le temps de répondre à mon questionnaire ?\r\nPour ceux à qui ça intéresse, vos réponses me seront utiles dans la partie \"\' Analyse et expérience de la transformation digitale \".\r\nMerci à vous pour vos réponses et bonne continuation !7', '2022-05-23 17:51:04', 1, 1, 38),
(75, 'Se référencer sur google actualités', 'Bonjour à tous.\r\nJe voudrais savoir comment se référencer sur google actualités récentes? Merci :)', '2022-05-23 17:52:20', 0, 1, 38),
(76, 'Référencement Youtube', 'Bonjour à tous.\r\nJ\'ai trouvé sur ce site des spécialistes du référencement.\r\nSavez-vous comment créer une stratégie de référencement pour Youtube ?\r\nMerci', '2022-05-23 19:02:07', 3, 1, 38),
(77, 'Se référencer sur google actualités', 'Bonjour à tous.\r\nJe voudrais savoir comment se référencer sur google actualités récentes? Merci :)', '2022-05-23 19:07:01', 0, 1, 38),
(78, 'GoPro Son', 'Un logiciel pour améliorer le son enregistrer par une GoPro', '2022-05-23 19:12:50', 2, 2, 37),
(79, 'BtoB BtoC', 'On pourrait m\'expliquer le BtoB et le BtoC svp ?', '2022-05-23 19:14:20', 0, 1, 37),
(80, 'Qui a un modèle de flyer', 'Salut j\'ai besoin de flyer pour ma startup mais je trouve pas ou je pourrais en faire ', '2022-05-23 19:18:27', 1, 3, 37),
(81, 'Sketchup', 'Il faut une License pour SketchUp', '2022-05-23 19:26:44', 0, 4, 33),
(82, ' présence percutante', 'Des Tips pour avoir une présence impactante sur les réseaux ?', '2022-05-23 19:28:21', 0, 6, 33),
(83, 'Asus ou Canon', 'Quel caméra est la meilleure ?', '2022-05-23 19:29:33', 1, 2, 33),
(84, 'templates sites', 'Canva c\'est vraiment trop bien pour les templates mais j\'aimerais découvrir autre chose', '2022-05-23 19:30:20', 0, 3, 33),
(85, 'LoL', 'je voudrais créer un chroma de Garen mais je ne sais pas comment m\'y prendre', '2022-05-23 19:31:37', 0, 7, 33),
(86, 'Je voudrais améliorer mon pinceau', 'Je dessine beaucoup avec des HB mais du coup sa me limite dans mes rendu mais je sais pas sur quelle autre style de pinceaux prendre', '2022-05-23 19:33:13', 0, 8, 35),
(87, 'Laure me manque', 'Vous savez quand Laure revient ?', '2022-05-23 19:34:34', 3, 10, 35),
(88, 'Absence injustifié', 'Je ne peux pas justifier mes absences je ne comprend pas pourquoi ', '2022-05-23 19:35:25', 2, 10, 37),
(89, 'Multiprises', 'On manque de multiprise en A7 c\'est relou', '2022-05-23 19:37:03', 1, 10, 36),
(90, 'Hâte d\'être a la soirée des délègues', 'Je ramènerez mon jeu de loup garou on va bien s\'amuser', '2022-05-23 19:40:09', 1, 9, 39),
(91, 'Propreté des salles', 'Les salles de classes sont souvent sales a la fin des cours un petit rappel sur l\'hygiène ne ferait pas de mal', '2022-05-23 19:41:09', 0, 10, 39),
(92, 'Du coup l\'annexe c\'est ouvert ?', 'Jeudi j\'aimerais aller taffer a l\'annexe, c\'est ouvert ?', '2022-05-23 19:43:32', 1, 2, 39),
(93, 'Aide charte graphique', 'Si vous avez besoin d\'aide je suis la', '2022-05-23 19:44:42', 1, 3, 39),
(94, 'Qui utilise SketchUp encore', 'J\'ai vue que une personne utiliser encore SketchUp, faut passer a autre chose Monsieur', '2022-05-23 19:45:53', 0, 4, 39),
(95, 'Qui veut devenir mon modèle 3D', 'Je cherche un modèle 3D, Merci', '2022-05-23 19:46:46', 0, 7, 39),
(96, 'Besoin d\'un Logo Urgent', 'Bonjour je suis prêt a payer pour un logo pour après-demain ', '2022-05-23 19:47:57', 0, 8, 39),
(97, 'Le cours de Monsieur POETSH', 'Le cours de Monsieur POETSH sont avancer de 1 heure jeudi 25 mai', '2022-05-23 19:51:05', 0, 2, 40),
(98, 'License gratuite', 'Certaines licences sont gratuite sur l\'intra d\'Ynov', '2022-05-23 19:51:58', 0, 3, 40),
(99, 'Aide SketchUp', 'Je crois que Monsieur Jouan s\'y connait sur ce domaine la', '2022-05-23 19:52:54', 0, 4, 40),
(100, 'Télétravail', 'Monsieur MOUNIER ne pourra pas assurer les cours en présentiel toute la semaine , nous vous invitons a passer en distanciel', '2022-05-23 19:54:33', 0, 6, 40),
(101, 'Concours Modéle', 'Nous proposons de Créer un Concours de modélisation 3D au sein d\'Ynov Le 26 mai ', '2022-05-23 19:56:26', 0, 7, 40),
(102, 'Soirée annulé ', 'La soirée des délégués est annuler ', '2022-05-23 19:58:56', 0, 9, 40),
(103, 'Vos messages nous font plaisir', 'Demain gouter organisé a l\'admin pour les B1 infos', '2022-05-23 20:00:11', 0, 10, 40),
(104, 'Nouveau clip', 'Allez voir mon nouveau clip svp !!!!!!', '2022-05-23 20:03:35', 0, 2, 32),
(105, 'Dev\'XP', 'Cette start-up est juste incroyable si vous voulez vous améliorer n\'hésitez pas a posez votre candidature', '2022-05-23 20:06:21', 0, 3, 32),
(106, 'Ptdrrr Monsieur PIZZETA', 'Le prof a oublié de débrancher le projecteur on a vue tout les messages avec sa meuf !! hihi il est trop mignon avec elle', '2022-05-23 20:09:57', 0, 4, 32),
(107, 'Meuble IKEA', 'Les nouveaux meubles IKEA sont incroyable', '2022-05-23 20:12:21', 0, 4, 37),
(108, 'Stage Web manager', 'Des entreprises qui recrutent des B2 pour un stage de 6 semaine ?', '2022-05-23 20:15:33', 0, 6, 37),
(109, 'Post', 'Salut je test les posts', '2022-05-23 20:15:55', 0, 6, 37),
(110, 'Morgana ou Caitlyn ?', 'Vous préférez quelle design Morgana ou Caitlyn ?', '2022-05-23 20:16:55', 0, 7, 37),
(111, 'Ptdr la soirée', 'J\'en connais qui vont être dégouté pour la soirée', '2022-05-23 20:18:04', 0, 9, 37),
(112, 'Besoin d\'aide', 'Besoin d\'aide Urgemment la !!!!!', '2022-05-23 20:19:12', 0, 6, 35),
(113, 'Maya', 'C\'est moi ou maya beug ?', '2022-05-23 20:19:39', 0, 7, 35),
(114, 'Tournoi LOL', 'ON A GAGNER le TOURNOIS LOOOOOOL !!!!!!!!!', '2022-05-23 20:20:56', 0, 9, 35),
(115, 'BDE', 'Je suis trop heureuse que Le BDE organise des tournois comme celui de LOL', '2022-05-23 20:21:47', 1, 9, 35),
(116, 'Je suis largué en ce moment', 'je ne sais plus ou j\'en suis j\'aurais besoin d\'aide svp !!!!', '2022-05-23 20:22:44', 1, 8, 33),
(117, 'C\'est la CATA', 'Je n\'y arrive plus je pense a arrêter ', '2022-05-23 20:24:30', 1, 8, 33);

-- --------------------------------------------------------

--
-- Structure de la table `replies`
--

DROP TABLE IF EXISTS `replies`;
CREATE TABLE IF NOT EXISTS `replies` (
  `Reply_id` int NOT NULL AUTO_INCREMENT,
  `Content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `reply_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `Post_id` int DEFAULT '0',
  `ReplyTo_id` int DEFAULT '0',
  `User_id` int NOT NULL,
  `Deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Reply_id`),
  KEY `Post_id` (`Post_id`),
  KEY `User_id` (`User_id`),
  KEY `ReplyTo_id` (`ReplyTo_id`)
) ENGINE=InnoDB AUTO_INCREMENT=247 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `replies`
--

INSERT INTO `replies` (`Reply_id`, `Content`, `reply_date`, `Post_id`, `ReplyTo_id`, `User_id`, `Deleted`) VALUES
(156, 'Bonjour les gars, pour le forum j\'ai un petit problème je n\'arrive pas a centrer ma div, on pourrait m\'aider svp ?', '2022-05-23 16:54:40', 66, 0, 32, 0),
(157, 'En lançant mon serveur wamp je  me rend compte que j\'en ai que 2 d\'ouvert, je ne sais pas pourquoi ça me fait ça', '2022-05-23 17:28:44', 67, 0, 32, 0),
(158, 'Bonjour,\r\nDans ma vue Twig, j\'ai à un moment besoin d\'ajouter une redirection vers une autre page avec un paramètre.\r\nCependant je voudrais que ce soit dynamique, donc que le paramètre soit une fonction qui retourne une valeur, comme ci-dessous:\r\n\r\nhref=\"{{path(\'app_compute\', {id : \'test()\'} )}}\"\r\nJ\'ai compris que cela ne marche pas parce qu\'il va envoyer le texte \'test()\' au lieu du résultat, mais je voudrais savoir comment faire pour avoir la valeur retournée par la fonction au lieu du texte.\r\n\r\nN.B: La fonction test() est une fonction Javascript, qui retourne une simple valeur.\r\n\r\nMerci d\'avance', '2022-05-23 17:31:14', 68, 0, 32, 0),
(159, 'Bonjour,\r\n\r\nAlors voila , je cherche a savoir comment faire pour qu\'un bouton ait une double action , avec verification de la premiere.Pour faire simple , ce bouton permet de placer un pixel sur un canvas , mais je souhaite que pour faire cela les gens fasse un don sur Paypal de 10 centimes. Est ce que vous auriez une idee de comment faire ca , merci davance et bonne journée', '2022-05-23 17:32:15', 69, 0, 32, 0),
(160, 'Bonjour, j\'ai un menu sur mobile sauf que le problème c\'est que quand j\'ouvre le menu, je ne vois que la div.', '2022-05-23 17:33:14', 70, 0, 32, 0),
(161, 'Bonjour,\r\nsur mon site, j\'ai deux champ d\'un formulaire qui sont remplis sur une page (le champ N° de série et le champ email)\r\n\r\nhttps://data-perso.online/user/456/\r\n\r\nje souhaite afficher sur cette meme page uniquement le contenu des champs via html en dessous des éléments existants.\r\n\r\nJ\'ai tenté nombre de manipulations issues du net mais rien n\'y fait...\r\n\r\nSi vous aviez une piste d\'un code simple CE SERAIT MAGNIFIQUE !\r\nplus d\'une semaine mais pas de résultat, certainement du à mon niveau plus que faible dans ce domaine...\r\n\r\nBien à vous', '2022-05-23 17:34:21', 71, 0, 33, 0),
(162, 'Hello !\r\nJ\'ai un mini projet que je galère a realiser parce que je suis naze en HTML haha.\r\nJ\'aimerais créer un bouton qui redirige vers une page QUE si l\'input est correct (comme une sorte de mot de passe).\r\nJ\'ai essayé plein d\'approche mais je galère comme pas possible avec les conditions.\r\nSi quelqu\'un pouvais m\'aider ça serait incroyable\r\nJ\'ai mis un petit schema de ce que j\'ai en tete.\r\n\r\nhttps://imgur.com/a/Vv005qc\r\n\r\nComme vous voyez c\'est ultra simple...\r\nJe vais pas vous mettre mes tentatives j\'en ai trop honte haha.\r\n\r\nMerci beaucoup pour votre attention :)', '2022-05-23 17:34:58', 72, 0, 33, 0),
(163, 'Bonjour à tous,\r\nJe vous demande votre aide aujourd\'hui car je n\'ai pas réussi à répondre à une question sur du javascript d\'un travail que je dois rendre.\r\nIl s\'agit de la question numéro 4 dans laquelle il m\'est demandé de compléter la fonction createFigureBlock.\r\nVoici le lien de l\'exercice : https://www.fil.univ-lille.fr/~abdelkaf ... cript.html\r\n\r\nMerci d\'avance et une très bonne journée à vous.:(', '2022-05-23 17:35:56', 73, 0, 33, 0),
(164, 'Bonjour, montre moi ton code et je pourrais peut-être t\'aider ', '2022-05-23 17:36:40', 66, 156, 33, 0),
(165, 'Merci j\'ai pus régler mon problème, j\'avais juste pas link le bon CSS mdr :)', '2022-05-23 17:37:48', 66, 164, 32, 0),
(166, 'Pas sûr d\'avoir compris ce que tu cherche à faire : ici le numéro de série et l\'adresse e-mail ne sont pas dans un formulaire.\r\nQu\'entends-tu par \"uniquement le contenu des champs\", et c\'est quoi les \"éléments existants\" ?', '2022-05-23 17:38:55', 71, 161, 32, 0),
(167, 'Je ne comprends toujours pas ta problématique : si tu arrive à afficher du contenu dynamique sur cette page, pourquoi tu ne pourrais pas l\'afficher ailleurs (ici dans un mail) ?\r\nTu génère du html avec du contenu dynamique (ici le numéro et l\'adresse e-mail) exactement de la même façon que pour afficher ta page.', '2022-05-23 17:39:19', 71, 161, 32, 0),
(168, 'Hello,\r\nTu pourrais nous faire voir une de tes tentatives pour être sur de ce que tu as besoin ?\r\n', '2022-05-23 17:39:52', 72, 162, 32, 0),
(169, 'Il faudrait déjà bien fermer ton button <button>Envoyer</button> .\r\nEnsuite si tu veux faire de cette manière, moi de mon coté ça fonctionne .', '2022-05-23 17:40:59', 72, 162, 35, 0),
(170, 'Hello,\r\nTu pourrais nous faire voir une de tes tentatives pour être sur de ce que tu as besoin ?\r\n', '2022-05-23 17:41:13', 72, 162, 35, 0),
(171, 'J\'ai pas gardé beaucoup de mes essais. Si tu le souhaites vraiment je peux essayer de recreer ce que j\'ai fait mais ça ne servirait à rien si ce n\'est que de rendre le problème incomprehensible.\r\nVraiment c\'est ultra simple : une case d\'input, à coté (ou en dessous) un bouton. Quand t\'appuies sur le bouton, ca vérifie que ce qui ya d\'inscrit dans l\'input correspond au mot de passe, et si c\'est le cas ; ça te redirige vers une page. Pas besoin de css ni de mise en page.\r\n\r\n', '2022-05-23 17:42:05', 72, 168, 32, 0),
(172, 'Relance ton Wamp, ca devrait marcher.', '2022-05-23 17:43:35', 67, 157, 33, 0),
(173, 'Bonjour,\r\n\r\ntu peux passer l\'url de redirection dans un element data-url dans ta page et tu récupère cet élément en js pour faire ta redirection (le nom data-url est totalement arbitraire, tu choisis celui que tu veux)', '2022-05-23 17:44:55', 68, 158, 33, 0),
(174, 'Message supprimé par son créateur.', '2022-05-23 17:45:27', 69, 159, 33, 1),
(175, 'Message supprimé par son créateur.', '2022-05-23 17:45:49', 69, 159, 33, 1),
(176, 'Hello :)\r\n\r\nAlors deux solutions:\r\n\r\nLa première serait de passer par JS, quand on clique sur le bouton, coté JS tu fait une verification pour voir si le don a été fait, si oui ça place le pixel en JS, sinon ça fera le don en JS\r\n\r\nLa deuxième peut se faire en JS ou PHP, tu ne fait pas un mais deux boutons identiques, la vérification du don se fait avant de charger la page pour le PHP, ou bien au chargement de la page si tu passes par JS, si le don n\'a pas été fait, le bouton pour placer le pixel sera caché et il y aura un bouton pour faire le don, si le don a été fait ce sera l\'inverse (alors oui il y aura deux bouton, mais en aparence du point de vue de l\'utilisateur il n\'y en a que un)', '2022-05-23 17:46:07', 69, 159, 33, 0),
(177, 'Bonjour à tous,\r\nDans le cadre de mes études et comme beaucoup d\'étudiants, j\'ai le devoir de rédiger un mémoire de fin d\'année.\r\nCompte tenu de mon expérience et connaissance, j\'ai fait le choix de m\'intéresser au sujet de la transformation digitale, qui est pour moi un sujet passionnant. Cependant, j\'ai besoin de votre aide, de manière à partager les connaissances et expériences de tous à ce sujet.\r\nPouvez-vous svp prendre le temps de répondre à mon questionnaire ?\r\nPour ceux à qui ça intéresse, vos réponses me seront utiles dans la partie \"\' Analyse et expérience de la transformation digitale \".\r\nMerci à vous pour vos réponses et bonne continuation !7', '2022-05-23 17:51:04', 74, 0, 38, 0),
(178, 'Le lien Questionnaire : https://docs.google.com/forms/d/e/1FAIpQLSf0gw1nE5_mu7fhRreZ9_mMcm7BvYFS3xkhOlGWorNupBIpsw/viewform', '2022-05-23 17:51:33', 74, 177, 38, 0),
(179, 'Bonjour à tous.\r\nJe voudrais savoir comment se référencer sur google actualités récentes? Merci :)', '2022-05-23 17:52:20', 75, 0, 38, 0),
(180, 'Bonjour à tous.\r\nJ\'ai trouvé sur ce site des spécialistes du référencement.\r\nSavez-vous comment créer une stratégie de référencement pour Youtube ?\r\nMerci', '2022-05-23 19:02:07', 76, 0, 38, 0),
(181, 'Message supprimé par son créateur.', '2022-05-23 19:02:28', 76, 180, 38, 1),
(182, 'Message supprimé par son créateur.', '2022-05-23 19:02:54', 76, 180, 38, 1),
(183, 'Pour mon nouveau projet de chaine YouTube, j\'utilise Morningfame. Vous pouvez tester un mois en utilisant ce lien : https://morningfa.me/invite/zeciekya\r\nL\'idée est de commencer à créer des vidéos avec des mots clés très niché. Recherché mais pas utilisé par les concurrents.\r\nEnsuite au fur et à mesure que la chaine grossie, ont peut s\'attaquer à des requêtes (mots-clés) de plus en plus importantes.\r\nL\'outils moningfame vous permet de savoir sur quels mots clé vous pouvez faire votre vidéo en fonction des analytics de votre chaine. Ensuite, il vous donne un score d\'optimisation de votre titre, description et tag ;)', '2022-05-23 19:05:09', 76, 180, 39, 0),
(184, 'Bonjour.\r\nLe marketing de YouTube est souvent négligé par les spécialistes du marketing des médias sociaux. Certains considèrent YouTube comme un réseau de médias sociaux. D\'autres le considèrent plutôt comme une plateforme vidéo en ligne.\r\n\r\nLa mise en œuvre d\'une stratégie de référencement sur YouTube est une tâche plutôt complexe. Quoi qu\'il en soit, il existe d\'innombrables possibilités de marketing sur YouTube, surtout si votre public est présent sur la plateforme et que vos concurrents ne le sont pas.\r\n\r\nJe ne saurais l\'expliquer mais je vous conseille de consulter cet article : https://x-com-agency.net/referencement-youtube', '2022-05-23 19:05:26', 76, 180, 39, 0),
(185, 'Bonjour.\r\nLe marketing de YouTube est souvent négligé par les spécialistes du marketing des médias sociaux. Certains considèrent YouTube comme un réseau de médias sociaux. D\'autres le considèrent plutôt comme une plateforme vidéo en ligne.\r\n\r\nLa mise en œuvre d\'une stratégie de référencement sur YouTube est une tâche plutôt complexe. Quoi qu\'il en soit, il existe d\'innombrables possibilités de marketing sur YouTube, surtout si votre public est présent sur la plateforme et que vos concurrents ne le sont pas.\r\n\r\nJe ne saurais l\'expliquer mais je vous conseille de consulter cet article : https://x-com-agency.net/referencement-youtube', '2022-05-23 19:05:56', 76, 184, 38, 0),
(186, 'Bonjour à tous.\r\nJe voudrais savoir comment se référencer sur google actualités récentes? Merci :)', '2022-05-23 19:07:01', 75, 0, 38, 0),
(189, 'Un logiciel pour améliorer le son enregistrer par une GoPro', '2022-05-23 19:12:50', 78, 0, 37, 0),
(190, 'On pourrait m\'expliquer le BtoB et le BtoC svp ?', '2022-05-23 19:14:20', 79, 0, 37, 0),
(191, 'Salut j\'ai besoin de flyer pour ma startup mais je trouve pas ou je pourrais en faire ', '2022-05-23 19:18:27', 80, 0, 37, 0),
(192, 'Il faut une License pour SketchUp', '2022-05-23 19:26:44', 81, 0, 33, 0),
(193, 'Des Tips pour avoir une présence impactante sur les réseaux ?', '2022-05-23 19:28:21', 82, 0, 33, 0),
(194, 'Quel caméra est la meilleure ?', '2022-05-23 19:29:33', 83, 0, 33, 0),
(195, 'Canva c\'est vraiment trop bien pour les templates mais j\'aimerais découvrir autre chose', '2022-05-23 19:30:20', 84, 0, 33, 0),
(196, 'je voudrais créer un chroma de Garen mais je ne sais pas comment m\'y prendre', '2022-05-23 19:31:37', 85, 0, 33, 0),
(197, 'Je dessine beaucoup avec des HB mais du coup sa me limite dans mes rendu mais je sais pas sur quelle autre style de pinceaux prendre', '2022-05-23 19:33:13', 86, 0, 35, 0),
(198, 'Vous savez quand Laure revient ?', '2022-05-23 19:34:34', 87, 0, 35, 0),
(199, 'Je ne peux pas justifier mes absences je ne comprend pas pourquoi ', '2022-05-23 19:35:25', 88, 0, 37, 0),
(200, 'Je crois l\'année prochaine', '2022-05-23 19:35:42', 87, 198, 37, 0),
(201, 'On manque de multiprise en A7 c\'est relou', '2022-05-23 19:37:03', 89, 0, 36, 0),
(202, 'Perso je préfère Alexia ', '2022-05-23 19:37:26', 87, 198, 36, 0),
(203, 'Essaye d\'en parler a Alexia elle te réglera ca très très vite, c\'est la meilleure', '2022-05-23 19:38:05', 88, 199, 36, 0),
(204, 'Je ramènerez mon jeu de loup garou on va bien s\'amuser', '2022-05-23 19:40:09', 90, 0, 39, 0),
(205, 'Les salles de classes sont souvent sales a la fin des cours un petit rappel sur l\'hygiène ne ferait pas de mal', '2022-05-23 19:41:09', 91, 0, 39, 0),
(206, 'j\'y vais que pour prendre des bonbons perso', '2022-05-23 19:41:51', 87, 202, 39, 0),
(207, 'C\'est vrai que c\'est une crème', '2022-05-23 19:42:17', 88, 203, 39, 0),
(208, 'J\'en parlerais a l\'admin demain', '2022-05-23 19:42:36', 89, 201, 39, 0),
(209, 'Jeudi j\'aimerais aller taffer a l\'annexe, c\'est ouvert ?', '2022-05-23 19:43:32', 92, 0, 39, 0),
(210, 'Franchement aucune idée ', '2022-05-23 19:43:47', 78, 189, 39, 0),
(211, 'Celle de ton iphone mdrr', '2022-05-23 19:44:06', 83, 194, 39, 0),
(212, 'Si vous avez besoin d\'aide je suis la', '2022-05-23 19:44:42', 93, 0, 39, 0),
(213, 'J\'ai vue que une personne utiliser encore SketchUp, faut passer a autre chose Monsieur', '2022-05-23 19:45:53', 94, 0, 39, 0),
(214, 'Je cherche un modèle 3D, Merci', '2022-05-23 19:46:46', 95, 0, 39, 0),
(215, 'Bonjour je suis prêt a payer pour un logo pour après-demain ', '2022-05-23 19:47:57', 96, 0, 39, 0),
(216, 'Le cours de Monsieur POETSH sont avancer de 1 heure jeudi 25 mai', '2022-05-23 19:51:05', 97, 0, 40, 0),
(217, 'Certaines licences sont gratuite sur l\'intra d\'Ynov', '2022-05-23 19:51:58', 98, 0, 40, 0),
(218, 'Je crois que Monsieur Jouan s\'y connait sur ce domaine la', '2022-05-23 19:52:54', 99, 0, 40, 0),
(219, 'Monsieur MOUNIER ne pourra pas assurer les cours en présentiel toute la semaine , nous vous invitons a passer en distanciel', '2022-05-23 19:54:33', 100, 0, 40, 0),
(220, 'Nous proposons de Créer un Concours de modélisation 3D au sein d\'Ynov Le 26 mai ', '2022-05-23 19:56:26', 101, 0, 40, 0),
(221, 'La soirée des délégués est annuler ', '2022-05-23 19:58:56', 102, 0, 40, 0),
(222, 'Ne venez pas svp', '2022-05-23 19:59:18', 90, 204, 40, 0),
(223, 'Demain gouter organisé a l\'admin pour les B1 infos', '2022-05-23 20:00:11', 103, 0, 40, 0),
(224, 'Allez voir mon nouveau clip svp !!!!!!', '2022-05-23 20:03:35', 104, 0, 32, 0),
(225, 'Cette start-up est juste incroyable si vous voulez vous améliorer n\'hésitez pas a posez votre candidature', '2022-05-23 20:06:21', 105, 0, 32, 0),
(226, 'si c\'est pour dire ca évite de polluer le forum stp ', '2022-05-23 20:07:35', 78, 210, 32, 0),
(227, 'Ta qu\'a aller voir de toi même ', '2022-05-23 20:07:56', 92, 209, 32, 0),
(228, 'Franchement tu te gave frérot ', '2022-05-23 20:08:24', 93, 212, 32, 0),
(229, 'Demande a Jennifer Piambino', '2022-05-23 20:08:45', 80, 191, 32, 0),
(230, 'Le prof a oublié de débrancher le projecteur on a vue tout les messages avec sa meuf !! hihi il est trop mignon avec elle', '2022-05-23 20:09:57', 106, 0, 32, 0),
(231, 'Les nouveaux meubles IKEA sont incroyable', '2022-05-23 20:12:21', 107, 0, 37, 0),
(232, 'Des entreprises qui recrutent des B2 pour un stage de 6 semaine ?', '2022-05-23 20:15:33', 108, 0, 37, 0),
(233, 'Salut je test les posts', '2022-05-23 20:15:55', 109, 0, 37, 0),
(234, 'Vous préférez quelle design Morgana ou Caitlyn ?', '2022-05-23 20:16:55', 110, 0, 37, 0),
(235, 'J\'en connais qui vont être dégouté pour la soirée', '2022-05-23 20:18:04', 111, 0, 37, 0),
(236, 'Besoin d\'aide Urgemment la !!!!!', '2022-05-23 20:19:12', 112, 0, 35, 0),
(237, 'C\'est moi ou maya beug ?', '2022-05-23 20:19:39', 113, 0, 35, 0),
(238, 'ON A GAGNER le TOURNOIS LOOOOOOL !!!!!!!!!', '2022-05-23 20:20:56', 114, 0, 35, 0),
(239, 'Je suis trop heureuse que Le BDE organise des tournois comme celui de LOL', '2022-05-23 20:21:47', 115, 0, 35, 0),
(240, 'je ne sais plus ou j\'en suis j\'aurais besoin d\'aide svp !!!!', '2022-05-23 20:22:44', 116, 0, 33, 0),
(241, 'Je n\'y arrive plus je pense a arrêter ', '2022-05-23 20:24:30', 117, 0, 33, 0),
(242, 'Vous pouvez prendre RDV avec la psychologue de l\'école', '2022-05-23 20:25:36', 117, 241, 40, 0),
(243, 'n\'hésitez pas a en parler a votre référant filière ', '2022-05-23 20:26:01', 116, 240, 40, 0),
(244, 'Le meilleur BDE de tout les campus :)', '2022-05-23 20:26:33', 115, 239, 40, 0),
(245, 'Continuez comme ca', '2022-05-23 20:27:07', 72, 162, 40, 0);

-- --------------------------------------------------------

--
-- Structure de la table `session`
--

DROP TABLE IF EXISTS `session`;
CREATE TABLE IF NOT EXISTS `session` (
  `User_id` int NOT NULL,
  `Session_id` varchar(500) NOT NULL,
  `Creation` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `User_id_fk` (`User_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `session`
--

INSERT INTO `session` (`User_id`, `Session_id`, `Creation`) VALUES
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 16:17:15'),
(26, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-17 16:17:15'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 16:17:15'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 16:27:06'),
(26, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-17 16:27:37'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 21:36:02'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 21:39:12'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-17 21:40:50'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 13:04:26'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 14:22:37'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 14:29:27'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 15:42:02'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 20:27:37'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 20:30:38'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:25:14'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:52:26'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:53:48'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:57:33'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:58:17'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 21:58:58'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-18 22:11:15'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 09:12:53'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-19 09:13:32'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 09:48:42'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-19 09:50:18'),
(27, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-19 10:02:47'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 12:43:03'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 13:02:52'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 14:27:52'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-19 15:00:00'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-20 14:03:24'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-20 14:09:10'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-20 14:40:00'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-20 18:03:47'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-21 15:33:37'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-21 15:35:05'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-21 16:11:42'),
(26, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-21 16:12:50'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-21 17:29:05'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 16:33:44'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 16:56:35'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 17:20:20'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 17:47:01'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-22 17:47:04'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 17:52:00'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-22 17:53:37'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 18:04:09'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 18:05:28'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 18:21:48'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-22 18:23:03'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 19:02:32'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 21:18:35'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-22 21:40:59'),
(26, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-22 21:41:15'),
(27, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-22 21:41:57'),
(27, 'etHsbZRjxAwnwekrBEmfdzdcEkXBAkjQ', '2022-05-22 21:42:19'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 08:26:04'),
(27, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 08:26:14'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 10:47:09'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 12:17:57'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 13:19:34'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 14:05:26'),
(26, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 14:35:13'),
(27, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-23 14:35:33'),
(26, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 15:19:24'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 15:42:30'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 15:50:39'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 16:17:43'),
(27, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 16:27:21'),
(32, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 16:44:03'),
(33, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-23 16:45:34'),
(32, 'etHsbZRjxAwnwekrBEmfdzdcEkXBAkjQ', '2022-05-23 16:53:24'),
(33, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 17:33:43'),
(32, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 17:36:58'),
(35, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-23 17:40:31'),
(32, 'etHsbZRjxAwnwekrBEmfdzdcEkXBAkjQ', '2022-05-23 17:41:35'),
(33, 'ZLCtTMtTCoaNatyyiNKAReKJyiXJrscc', '2022-05-23 17:43:09'),
(38, 'tNswYNsGRussVmaozFZBsbOJiFQGZsnw', '2022-05-23 17:50:04'),
(38, 'TKSmVoiGLOpbUOpEdKupdOMeRVjaRzLN', '2022-05-23 19:00:53'),
(39, 'TXYeUCWKsXbGyRAOmBTvKSJfjzaLbtZs', '2022-05-23 19:04:34'),
(38, 'yMGeuDtRzQMDQiYCOhgHOvgSeycJPJHY', '2022-05-23 19:05:41'),
(37, 'NufNjJhhjUVRuSqfgqVMkPYVkURUpiFv', '2022-05-23 19:07:59'),
(33, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 19:23:53'),
(35, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 19:32:04'),
(37, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 19:34:48'),
(36, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-23 19:36:16'),
(39, 'etHsbZRjxAwnwekrBEmfdzdcEkXBAkjQ', '2022-05-23 19:38:18'),
(40, 'ZLCtTMtTCoaNatyyiNKAReKJyiXJrscc', '2022-05-23 19:49:33'),
(32, 'tNswYNsGRussVmaozFZBsbOJiFQGZsnw', '2022-05-23 20:02:58'),
(38, 'TKSmVoiGLOpbUOpEdKupdOMeRVjaRzLN', '2022-05-23 20:10:14'),
(37, 'TXYeUCWKsXbGyRAOmBTvKSJfjzaLbtZs', '2022-05-23 20:11:20'),
(35, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 20:18:22'),
(33, 'FpLSjFbcXoEFfRsWxPLDnJObCsNVlgTe', '2022-05-23 20:22:05'),
(40, 'MaPEZQleQYhYzRyWJjPjzpfRFEgmotaF', '2022-05-23 20:25:17'),
(33, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 21:12:36'),
(41, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 21:18:07'),
(33, 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2022-05-23 21:56:32');

-- --------------------------------------------------------

--
-- Structure de la table `topics`
--

DROP TABLE IF EXISTS `topics`;
CREATE TABLE IF NOT EXISTS `topics` (
  `Topic_id` int NOT NULL AUTO_INCREMENT,
  `Topic_name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Topic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `topics`
--

INSERT INTO `topics` (`Topic_id`, `Topic_name`) VALUES
(1, 'MarketCom'),
(2, 'Audiovisuel'),
(3, 'CreaDesign'),
(4, 'Architecture'),
(5, 'Info'),
(6, 'WebMgmt'),
(7, '3D'),
(8, '2D'),
(9, 'CampusLife'),
(10, 'Admin');

-- --------------------------------------------------------

--
-- Structure de la table `users`
--

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
  `User_id` int NOT NULL AUTO_INCREMENT,
  `Username` varchar(50) NOT NULL,
  `Admin` tinyint(1) DEFAULT '0',
  `Muted` tinyint(1) DEFAULT '0',
  `Email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Profil_Pic` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`User_id`),
  UNIQUE KEY `Username` (`Username`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `users`
--

INSERT INTO `users` (`User_id`, `Username`, `Admin`, `Muted`, `Email`, `Password`, `Profil_Pic`) VALUES
(32, 'MDavy13', 1, 0, 'davy.marthely@ynov.com', '$2a$10$I1FpRg2vm81WBo6ruHxYAuUQa32tRCmxmAwUyXjxG1L89gUDCZg4m', 1),
(33, 'MOlivier13', 1, 0, 'olivier.mistral@ynov.com', '$2a$10$xsY3730NkF/PvAcIdVShDubcP.sln4LFk9WyL4opSmv1uNsYOJcX.', 3),
(35, 'TMargot', 0, 0, 'margot.thomatis@ynov.com', '$2a$10$LG.wlo6HPOlWiGzwQ0HvPe1D7YrPV/2MaEpLPkbLtrRoMUGMWBnWS', 5),
(36, 'RPierre', 0, 1, 'pierre.roy@ynov.com', '$2a$10$xBCMzBzPsgUa7FLF3iic.Oq./5tKyL68.V00iDPC/8OXv6mTmBaWC', 6),
(37, 'PAntoine', 0, 0, 'antoine.pizzetta@ynov.com', '$2a$10$g3eHLSpa09iDJuTjbYm1QelX1kHHk.KEVc.EO7Iq1XzH2DhRadoYK', 2),
(38, 'MHind2', 0, 0, 'hind.attik@ynov.com', '$2a$10$WnyUwrrOIP1PKvCmOxrP6OEeDiW6V5JK0d.NJd6MSGhoaW8dkRZ.K', 4),
(39, 'MKheir13', 1, 0, 'kheireddine.mederreg@ynov.com', '$2a$10$5QkMV5iLZBOrqxBNznJLZeb1ia.rAv6BTFDbgkPoBSgV0UE1/MA7a', 5),
(40, 'Administration', 1, 0, 'administration@ynov.com', '$2a$10$.Ipw/jH1E36S8p1GpVTNNeMSax2XI7cVCnBXqogF0n5j3Ze6yQihe', 1);

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `post`
--
ALTER TABLE `post`
  ADD CONSTRAINT `post_ibfk_1` FOREIGN KEY (`Topic_id`) REFERENCES `topics` (`Topic_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `post_ibfk_2` FOREIGN KEY (`User_id`) REFERENCES `users` (`User_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Contraintes pour la table `replies`
--
ALTER TABLE `replies`
  ADD CONSTRAINT `replies_ibfk_1` FOREIGN KEY (`Post_id`) REFERENCES `post` (`Post_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `replies_ibfk_3` FOREIGN KEY (`User_id`) REFERENCES `users` (`User_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
