var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
var sassMiddleware = require('node-sass-middleware');
const cors = require('cors');

const header = require("./authorization/header");

var alarmRouter = require('./routes/alarm');

var app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(sassMiddleware({
  src: path.join(__dirname, 'public'),
  dest: path.join(__dirname, 'public'),
  indentedSyntax: true, // true = .sass and false = .scss
  sourceMap: true
}));
app.use(express.static(path.join(__dirname, 'public')));
app.use(cors());

app.use(header.authorize);

app.get('/apps/coffeetogo/api/v1/alarm/health', (req, res) => res.send({status: 200, message: 'Success'}));
app.use('/apps/coffeetogo/api/v1/alarm', alarmRouter);

module.exports = app;
