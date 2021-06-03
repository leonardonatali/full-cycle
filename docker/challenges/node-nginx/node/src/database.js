const config = {
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE
}

const connect = async () => {

    if (global.connection && global.connection.state !== 'disconnected') {
        return global.connection
    }

    const mysql = require("mysql2/promise")
    const connection = await mysql.createConnection(`mysql://${config.user}:${config.password}@${config.host}:3306/${config.database}`)
    global.connection = connection
    return connection
}

module.exports = connect