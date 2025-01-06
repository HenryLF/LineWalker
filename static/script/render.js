const cvs = document.getElementById("canvas");
const ctx = cvs.getContext("2d");

class Entity {
  spriteSheet = new Image();
  spriteWidth = 0;
  delayN = 20;
  maxN = 1;
  render(x, y, N) {
    N = Math.round(N / this.delayN);
    let sx = this.spriteWidth * (N % this.maxN);
    ctx.drawImage(
      this.spriteSheet,
      sx,
      0,
      this.spriteWidth,
      this.spriteSheet.height,
      x - this.spriteSheet.height / 2,
      y - this.spriteWidth / 2,
      this.spriteWidth,
      this.spriteSheet.height
    );
  }
}

let PlayerIdle_img = new Image();
PlayerIdle_img.src = "./assets/playerIdle.png";
class PlayerIdle extends Entity {
  spriteWidth = 48;
  maxN = 4;
  spriteSheet = PlayerIdle_img;
}

let PlayerLeft_img = new Image();
PlayerLeft_img.src = "./assets/playerLeft.png";
class PlayerLeft extends Entity {
  spriteWidth = 75;
  maxN = 6;
  spriteSheet = PlayerLeft_img;
}

let PlayerRight_img = new Image();
PlayerRight_img.src = "./assets/playerRight.png";
class PlayerRight extends PlayerLeft {
  spriteSheet = PlayerRight_img;
}

let BigPlayerIdle_img = new Image();
BigPlayerIdle_img.src = "./assets/bigPlayerIdle.png";
class BigPlayerIdle extends Entity {
  spriteWidth = 197;
  maxN = 8;
  spriteSheet = BigPlayerIdle_img;
}

let BigPlayerLeft_img = new Image();
BigPlayerLeft_img.src = "./assets/bigPlayerLeft.png";
class BigPlayerLeft extends Entity {
  spriteWidth = 193;
  maxN = 9;
  spriteSheet = BigPlayerLeft_img;
}

let BigPlayerRight_img = new Image();
BigPlayerRight_img.src = "./assets/bigPlayerRight.png";
class BigPlayerRight extends BigPlayerLeft {
  spriteSheet = BigPlayerRight_img;
}

const Ball = {
  color: "#0000FF",
  render(x, y, r) {
    ctx.save();
    ctx.fillStyle = this.color;
    let ball = new Path2D();
    ball.arc(x, y, r, 0, Math.PI * 2);
    ball.closePath();
    ctx.fill(ball);
    ctx.restore();
  },
};

let Deresolve = 10;
ctx.strokeStyle = "#000000";
function drawFloor(map) {
  ctx.fillStyle = "green";
  ctx.clearRect(0, 0, cvs.width, cvs.height);
  let floor = new Path2D();
  floor.moveTo(-Deresolve, map[-Deresolve]);
  for (let x = -Deresolve; x < cvs.width; x += Deresolve) {
    floor.lineTo(x, map[x]);
  }
  floor.lineTo(cvs.width + Deresolve, cvs.height + Deresolve);
  floor.lineTo(-Deresolve, cvs.height + Deresolve);
  floor.closePath();
  ctx.fill(floor);
  ctx.stroke(floor);
  // ctx.closePath();
}

const maxIdleSpeed = 30;
function drawPlayer(obj, n) {
  let p;
  if (obj.Speed.X < -maxIdleSpeed) {
    p = new PlayerLeft();
  } else if (obj.Speed.X > maxIdleSpeed) {
    p = new PlayerRight();
  } else {
    p = new PlayerIdle();
  }
  p.render(obj.ScreenCoord.X, obj.ScreenCoord.Y, n);
}

function drawBigPlayer(obj, n) {
  let p;
  if (obj.Speed.X < -maxIdleSpeed) {
    p = new BigPlayerLeft();
  } else if (obj.Speed.X > maxIdleSpeed) {
    p = new BigPlayerRight();
  } else {
    p = new BigPlayerIdle();
  }
  p.render(obj.ScreenCoord.X, obj.ScreenCoord.Y, n);
}

const xSpeed = document.getElementById("xspeed");
const ySpeed = document.getElementById("yspeed");
const xCoord = document.getElementById("xcoord");
const yCoord = document.getElementById("ycoord");
function updateUI(obj) {
  xSpeed.innerText = roundTo(obj.Speed.X, 2);
  ySpeed.innerText = roundTo(obj.Speed.Y, 2);
  xCoord.innerText = roundTo(obj.Coord.X, 0);
  yCoord.innerText = roundTo(obj.Coord.Y, 0);
}

function roundTo(x, n) {
  return Math.floor(x * 10 ** n) / 10 ** n;
}
