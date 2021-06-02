const mysql = require("mysql");

const MySQLConfig = {
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE
}

const conn = mysql.createConnection(MySQLConfig)

module.exports = conn