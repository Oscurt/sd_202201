const express = require('express');
const session = require('express-session');
const redis = require('redis');
const connectRedis = require('connect-redis');
const cors = require('cors');
var bodyParser = require('body-parser');

const app = express();

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const RedisStore = connectRedis(session);

const redisClient = redis.createClient({
    legacyMode: true,
    host: process.env.REDIS_HOST,   
    port: 6379
});

redisClient.connect().catch(console.error);

app.use(session({
    store: new RedisStore({ client: redisClient }),
    secret: 'mysecret',
    resave: false,
    saveUninitialized: false
}));

app.get("/", (req, res) => {
    const sess = req.session;
    if (sess.username && sess.password) {
        if (sess.username) {
            res.write(`<h1>Welcome ${sess.username} </h1><br>`)
            res.write(
                `<h3>This is the Home page</h3>`
            );
            res.send('<a href=' + '/logout' + '>Click here to log out</a >')
        }
    } else {
        res.sendFile(__dirname + "/login.html")
    }
});

app.post("/login", (req, res) => {
    const sess = req.session;
    const { username, password } = req.body
    sess.username = username
    sess.password = password
    res.send("success")
});

app.get("/logout", (req, res) => {
    req.session.destroy(err => {
        if (err) {
            return console.log(err);
        }
        res.redirect("/")
    });
});

app.listen(3000, () => {
    console.log("Server started at port 3000");
});