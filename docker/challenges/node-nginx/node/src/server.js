const express = require("express")
const conn = require("./database");

// Creates people table if not exists
let sql = "CREATE TABLE IF NOT EXISTS `people` (`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY, `name` VARCHAR(255))";
conn.query(sql);

// Creates a random people name
const name = Math.random().toString(36).substring(7)

// Insert into database
sql = `INSERT INTO people (name) VALUES('${name}')`;
conn.query(sql)

sql = "SELECT * FROM `people`"
let peopleList = "";
conn.query(sql, (err, peoples) => {
    if (peoples) {
        peoples.forEach((people) => {
            peopleList = peopleList +
            `<strong>ID</strong>: ${people.id}
            <br>
            <strong>Name</strong>: ${people.name}
            <br>
            <br>
            `
        })
    }
})
conn.end()

const app = express()
const port = process.env.APP_PORT || 80

app.get('/', (req, res) => {
    // Retrieves all database records
    return res.send(
        `<h1>Full Cycle Rocks!</h1>
        <hr>
        People list:<br>
        ${peopleList}`
    )
})

app.listen(port, () => {
    console.log(`The app is listening at port ${port}`);
})
