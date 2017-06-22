var express = require("express");
var app = express();

/* serves main page */
app.get("/build-client", function(req, res) {
  const spawn = require( 'child_process' ).spawn

  Promise.resolve().then(()=>{

    let dat


    ls = spawn(/^win/.test(process.platform) ? 'git.cmd' : 'git', ['pull']);
    ls.stdout.on( 'data', data => {
        dat = data
        console.log( `stdout: ${data}` );
    });
    ls.stderr.on( 'data', data => {
        dat = data
        console.log( `stderr: ${data}` );
    });
    ls.on( 'close', code => {
        console.log( `child process exited with code ${code}` );
        res.send(`stderr: ${code}` )
    });


  }).then(()=>{

    ls = spawn(/^win/.test(process.platform) ? 'npm.cmd' : 'npm', ['run',  'build']);
    ls.stdout.on( 'data', data => {
        dat = data
        console.log( `stdout: ${data}` );
    });
    ls.stderr.on( 'data', data => {
        dat = data
        console.log( `stderr: ${data}` );
    });
    ls.on( 'close', code => {
        console.log( `child process exited with code ${code}` );
        res.send(`stderr: ${code}` )
    });


  })




});

var port = process.env.PORT || 5000;
 app.listen(port, function() {
  console.log("Listening on " + port);
 });

console.log(1)
