SELECT Actores.Nombre 
FROM Actores
JOIN ActoresSeries ON Actores.ActorID = ActoresSeries.ActorId
LEFT JOIN Series ON ActoresSeries.SerieID = Series.SeriesID
WHERE Series.Nombre = "Big Sister"

SELECT COUNT(EpisodioID)
FROM Directores
JOIN Episodios ON Directores.DirectorID = Episodios.DirectorID
GROUP BY Directores.Nombre
ORDER BY COUNT(EpisodioID)