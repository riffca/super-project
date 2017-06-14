var express = require("express");
var exec = require("exec");
var app = express();

/* serves main page */
app.get("/build-client", function(req, res) {
  const spawn = require( 'child_process' ).spawn
  ls = spawn(/^win/.test(process.platform) ? 'npm.cmd' : 'npm', ['run',  'build']);
  ls.stdout.on( 'data', data => {
      console.log( `stdout: ${data}` );
  });
  ls.stderr.on( 'data', data => {
      console.log( `stderr: ${data}` );
  });
  ls.on( 'close', code => {
      console.log( `child process exited with code ${code}` );
      res.send(`stderr: ${code}` )
  });
});

var port = process.env.PORT || 5000;
 app.listen(port, function() {
  console.log("Listening on " + port);
 });

console.log(1)
