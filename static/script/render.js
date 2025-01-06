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
  ctx.fill(floor)
  ctx.stroke(floor)
  // ctx.closePath();
}

const maxIdleSpeed = 30;
function drawPlayer(coord, n) {
  let p;
  if (coord.XSpeed < -maxIdleSpeed) {
    p = new PlayerLeft();
  } else if (coord.XSpeed > maxIdleSpeed) {
    p = new PlayerRight();
  } else {
    p = new PlayerIdle();
  }
  p.render(coord.X, coord.Y, n);
}

function drawBigPlayer(coord, n) {
  let p;
  if (coord.XSpeed < -maxIdleSpeed) {
    p = new BigPlayerLeft();
  } else if (coord.XSpeed > maxIdleSpeed) {
    p = new BigPlayerRight();
  } else {
    p = new BigPlayerIdle;
  }
  p.render(coord.X, coord.Y, n);
}

const xSpeed = document.getElementById("xspeed");
const ySpeed = document.getElementById("yspeed");
const xCoord = document.getElementById("xcoord");
const yCoord = document.getElementById("ycoord");
function updateUI(coord) {
  xSpeed.innerText = roundTo(coord.XSpeed, 2);
  ySpeed.innerText = roundTo(coord.YSpeed, 2);
  xCoord.innerText = roundTo(coord.XAbs, 2);
  yCoord.innerText = roundTo(coord.YAbs, 2);
}

function roundTo(x, n) {
  return Math.floor(x * 10 ** n) / 10 ** n;
}
