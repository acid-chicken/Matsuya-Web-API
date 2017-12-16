import express from 'express';
import path from 'path';
import logger from 'morgan';
import cookieParser from 'cookie-parser';
import bodyParser from 'body-parser';

import index from './routes/index';

//APIv1読み込み
import v1random from './api/v1/random';

//APIv2読み込み
import v2random from './api/v2/random';

const app = express();

// GraphQL
import graphqlHTTP from 'express-graphql';
import schema from './api/graphql/schema';

app.use('/graphql', graphqlHTTP({
  schema: schema.Schema,
  rootValue: schema.Root,
  graphiql: true, // プロダクション用だと消したほうがいいかも
}));


// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// uncomment after placing your favicon in /public
//app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')));
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({
  extended: false
}));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', index);

//APIv1実装
app.use('/v1/random', v1random);

//APIv2実装
app.use('/v2/random', v2random);

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  res.sendStatus(404);
});

// error handler
app.use(function (err, req, res) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
