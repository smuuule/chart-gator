package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/smuuule/char-gator/internal/db"
	"github.com/thinkerou/favicon"
)

type Server struct {
	queries *db.Queries
}

func NewServer(queries *db.Queries) *Server {
	return &Server{queries: queries}
}

func (s *Server) SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(favicon.New("./favicon.ico"))
	_ = r.SetTrustedProxies(nil)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/songs", s.getSongs)
	r.POST("/songs", s.createSong)

	return r
}

type createSongReq struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
}

func (s *Server) getSongs(c *gin.Context) {
	songs, err := s.queries.ListSongs(c.Request.Context(), db.ListSongsParams{
		Limit:  50,
		Offset: 0,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list songs"})
		return
	}

	if songs == nil {
		songs = []db.Song{}
	}
	c.IndentedJSON(http.StatusOK, songs)
}

func (s *Server) createSong(c *gin.Context) {
	var req createSongReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	normTitle := strings.ToLower(strings.ReplaceAll(req.Title, " ", ""))
	normArtist := strings.ToLower(strings.ReplaceAll(req.Artist, " ", ""))

	song, err := s.queries.CreateSong(c.Request.Context(), db.CreateSongParams{
		Title:            req.Title,
		Artist:           req.Artist,
		NormalizedTitle:  normTitle,
		NormalizedArtist: normArtist,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create song: %v", err)})
		return
	}

	c.IndentedJSON(http.StatusCreated, song)
}
