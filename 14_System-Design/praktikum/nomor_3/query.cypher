// Contoh berkas Cypher untuk Neo4j

// Data Contoh (pengguna)
CREATE (:User {name: "John"})
CREATE (:User {name: "Alice"})
CREATE (:User {name: "Bob"})

// Ambil data dari pengguna
MATCH (u:User)
RETURN u