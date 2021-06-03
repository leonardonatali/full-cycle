const connect = require("./database");

const people = {
    create: async () => {
        // Creates people table if not exists
        const conn = await connect()
        await conn.query("CREATE TABLE IF NOT EXISTS `people` (`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY, `name` VARCHAR(255))")
    },
    insert: async () => {

        // Creates a random people name
        const name = Math.random().toString(36).substring(7)

        // Insert into database
        const conn = await connect()
        await conn.query(`INSERT INTO people (name) VALUES('${name}')`)
    },

    get: async () => {
        const conn = await connect()
        const [rows] = await conn.query("SELECT * FROM `people`")
        return rows
    }
}

module.exports.people = people