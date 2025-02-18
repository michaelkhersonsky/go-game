const canvas = document.getElementById("gameCanvas");
const ctx = canvas.getContext("2d");

canvas.width = 800;
canvas.height = 500;

// Paddle properties
const paddleWidth = 100, paddleHeight = 10;
let paddleX = (canvas.width - paddleWidth) / 2;
const paddleSpeed = 10;

// Ball properties
let ballX = canvas.width / 2, ballY = canvas.height / 2;
const ballRadius = 10;
let ballSpeedX = 4, ballSpeedY = 4;

// Controls
let rightPressed = false, leftPressed = false;

document.addEventListener("keydown", (e) => {
    if (e.key === "ArrowRight") rightPressed = true;
    if (e.key === "ArrowLeft") leftPressed = true;
});

document.addEventListener("keyup", (e) => {
    if (e.key === "ArrowRight") rightPressed = false;
    if (e.key === "ArrowLeft") leftPressed = false;
});

function update() {
    // Move paddle
    if (rightPressed && paddleX + paddleWidth < canvas.width) paddleX += paddleSpeed;
    if (leftPressed && paddleX > 0) paddleX -= paddleSpeed;

    // Move ball
    ballX += ballSpeedX;
    ballY += ballSpeedY;

    // Ball collision with walls
    if (ballX - ballRadius < 0 || ballX + ballRadius > canvas.width) ballSpeedX *= -1;
    if (ballY - ballRadius < 0) ballSpeedY *= -1;

    // Ball collision with paddle
    if (ballY + ballRadius > canvas.height - paddleHeight && 
        ballX > paddleX && ballX < paddleX + paddleWidth) {
        ballSpeedY *= -1; // Reverse direction
    } else if (ballY + ballRadius > canvas.height) {
        resetGame(); // Reset if ball misses paddle
    }
}

function resetGame() {
    ballX = canvas.width / 2;
    ballY = canvas.height / 2;
    ballSpeedX = 4;
    ballSpeedY = 4;
}

function draw() {
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Draw paddle
    ctx.fillStyle = "blue";
    ctx.fillRect(paddleX, canvas.height - paddleHeight, paddleWidth, paddleHeight);

    // Draw ball
    ctx.beginPath();
    ctx.arc(ballX, ballY, ballRadius, 0, Math.PI * 2);
    ctx.fillStyle = "red";
    ctx.fill();
    ctx.closePath();
}

function gameLoop() {
    update();
    draw();
    requestAnimationFrame(gameLoop);
}

gameLoop();

