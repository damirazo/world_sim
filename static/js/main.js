window.onload = function() {
    var multiplier = 20;
    var canvas = document.getElementById('canvas');
    var context = canvas.getContext('2d');

    var characterImage1 = document.getElementById("char1");
    var characterImage2 = document.getElementById("char2");
    //var targetImage = document.getElementById("chest");

    var ws = new WebSocket("ws://127.0.0.1:7777/game");

    window.setInterval(function() {
        ws.send("state");
    }, 1000 / 25);

    var prev = 1;

    ws.onmessage = function(e) {
        context.clearRect(0, 0, 800, 800);
        var data = JSON.parse(e.data);

        var img = (prev < 10) ? characterImage1 : characterImage2;

        data.forEach(function(playerParams) {
            var x = playerParams.position['X'];
            var y = playerParams.position['Y'];

            //var target = playerParams.Target;

            //context.drawImage(targetImage, (target.X * multiplier) - 10, (target.Y * multiplier) - 10);
            context.drawImage(img, (x * multiplier) - 10, (y * multiplier) - 16);
        });

        prev += 1;
        if (prev > 10) {
            prev = 1;
        }
    }
};