USE empresa_internet;

INSERT INTO Clients 
VALUES (1231, "Ivan", "Pineda", "1998-10-03", "Suba", "Bogotá"),
(1232, "Pepito", "Florez", "1998-11-30", "Castilla", "Bogotá"),
(1233, "Juan", "Perez", "1997-05-03", "Palmas", "Medellin"),
(1234, "Ivan", "Ruiz", "2000-10-03", "Poblado", "Medellin"),
(1235, "Luisa", "Pineda", "1999-12-28", "Suba", "Bogotá"),
(1236, "Tulio", "Garzon", "1998-10-03", "Bosa", "Bogotá"),
(1237, "Fernanda", "Perez", "1998-01-01", "Centenario", "Cali"),
(1238, "Julian", "Lijaca", "1997-12-20", "El peñon", "Cali"),
(1239, "Zoe", "Martinez", "1995-06-15", "Suba", "Bogotá"),
(1240, "Pedro", "Picapiedra", "2004-05-03", "Engativa", "Bogotá");

INSERT INTO InternetPlans
VALUES (1, 20, 30000, 0),
(2, 30, 40000, 0),
(3, 35, 50000, 5),
(4, 50, 65000, 10),
(5, 100, 90000, 15);

INSERT INTO Clients_has_InternetPlans
VALUES(1231, 1),
(1231, 5),
(1232, 3),
(1233, 4),
(1234, 1),
(1235, 2),
(1235, 3),
(1236, 5),
(1236, 2),
(1237, 5),
(1238, 3),
(1239, 4),
(1240, 2);
