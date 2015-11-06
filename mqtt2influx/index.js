var mqtt    = require('mqtt');
var client  = mqtt.connect('mqtt://infra.rzl');
var influx  = require('influx');
var config  = require('./config.json');

// Influx client configuration
var influx = influx({ 
  host : 'localhost',
  username : config.username,
  password : config.password,
  database : config.database
})

// Subscribe to all topics
client.subscribe('#');

// Handle new messages
client.on('message', function (topic, message) {
  message = JSON.parse(message);
  if (message.state !== undefined) {
    influx.writePoint(topic, message.state, {time: message.when}, function(err, response) {
      if (err) { console.log(err); }
    });
  } else {
    influx.writePoint(topic, message, {}, function(err, response) {
      if (err) { console.log(err); }
    });
  }
});
