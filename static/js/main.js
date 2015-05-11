window.onload = function() {
    var ws = new WebSocket("ws://127.0.0.1:7777/game"),
        canvas = document.getElementById('canvas'),
        context = canvas.getContext('2d'),
        multiplier = 20;

    ws.onopen = function() {
        window.setInterval(function() {
            ws.send("state");
        }, 1000 / 25);
    };

    ws.onmessage = function(e) {
        context.clearRect(0, 0, 800, 800);
        var data = JSON.parse(e.data);


        data.forEach(function(entity) {
            var img = document.getElementById(entity.type),
                position = entity.position,
                x = position.x,
                y = position.y;

            context.drawImage(img, (x * multiplier) - 10, (y * multiplier) - 16);
        });
    }
};
