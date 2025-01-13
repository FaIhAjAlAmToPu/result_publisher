package database

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func ConnectToCluster() gocqlx.Session {
	cluster := gocql.NewCluster("localhost:9042", "localhost:9042", "localhost:9042")

	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"} // Default credentials
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Keyspace = "active_results"
	cluster.Consistency = gocql.Quorum

	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		panic("Connection to cluster failed")
	}
	return session
}

// CREATE KEYSPACE IF NOT EXISTS active_results
// WITH replication = {'class': 'NetworkTopologyStrategy', 'datacenter1': 3}
// AND durable_writes = true;

func CreateKeyspace(session gocqlx.Session) {
	err := session.Query(`
        CREATE KEYSPACE IF NOT EXISTS active_results 
        WITH replication = {'class': 'NetworkTopologyStrategy', 'datacenter1': 3} 
        AND durable_writes = true;`, nil).Exec()
	if err != nil {
		panic(fmt.Sprintf("Failed to create keyspace: %v", err))
	}
}

func CreateTables(session gocqlx.Session) {
	err := session.ExecStmt(`
        CREATE TABLE IF NOT EXISTS active_results.exam (
            id UUID PRIMARY KEY,
            name TEXT,
			separator_name TEXT,
			held_by TEXT,
   		 	format MAP<TEXT, double>,
            publishing_date TIMESTAMP,
			end_time TIMESTAMP
        ) WITH compaction = {'class': 'LeveledCompactionStrategy'};
    `)
	if err != nil {
		panic("Failed to create exam table: " + err.Error())
	}

	// Create Result Table
	err = session.ExecStmt(`
        CREATE TABLE IF NOT EXISTS active_results.result (
    		exam_id UUID,
    		group_id TEXT,
    		student_id TEXT,
    		student_name TEXT,
    		scores list<double>,
    		PRIMARY KEY ((exam_id, group_id), student_id)
		) WITH compaction = {'class': 'LeveledCompactionStrategy'};

    `)
	if err != nil {
		panic("Failed to create result table: " + err.Error())
	}
}

// type Song struct {
// 	Id         string
// 	Title      string
// 	Artist     string
// 	Album      string
// 	Created_at time.Time
// }

// func insert_into_song() {
// 	song := Song{}

// 	q := session.Query(
// 		`INSERT INTO media_player.playlist (id,title,artist,album,created_at) VALUES (now(),?,?,?,?)`,
// 		[]string{":title", ":artist", ":album", ":created_at"}).
// 		BindMap(map[string]interface{}{
// 			":title":      song.Title,
// 			":artist":     song.Artist,
// 			":album":      song.Album,
// 			":created_at": time.Now(),
// 		})

// 	err := q.Exec()
// 	if err != nil {
// 		panic("error in exec query to insert a song in playlist %w", err)
// 	}
// }

// func read_song() {
// 	func (s Song) String() string {
// 	return fmt.Sprintf("Id: %s\nTitle: %s\nArtist: %s\nAlbum: %s\nCreated At: %s\n", s.Id, s.Title, s.Artist, s.Album, s.Created_at)
// }

// song := Song{}

// q := session.Query("SELECT * FROM media_player.playlist", nil)

// if err := q.SelectRelease(&song); err != nil {
//     panic("error in exec query to list playlists: %w", err)
// }

// println(song)
// }

// func update_song() {
// q := session.Query(
//     `UPDATE media_player.playlist SET
//         id = :id,
//         title = :title,
//         artist = :artist,
//         album = :album,
//         created_ad = :created_at
//         WHERE id = :id`,
//     []string{":id", ":title", ":artist", ":album", ":created_at"}).
//     BindMap(map[string]interface{} {
//         ":id":         "40450211-42cc-11ee-b14c-3da98b5024c0",
//         ":title":      "CPFMGD",
//         ":artist":     "Canhassi",
//         ":album":      "Canhas desu",
//         ":created_at": time.Now(),
//     })

// err := q.Exec(); if err != nil {
//     panic("error in exec update query")
// }
// }

// func delete_song() {

// q := session.Query(`DELETE FROM media_player.playlist WHERE id = ?`,
//     []string{":id"}).
//     BindMap(map[string]interface{} {
//         ":id": songToDelete.Id,
//     })

// err := q.Exec(); if err != nil {
//     return fmt.Errorf("error to exec delete query %w", err)
// }
// }
