function resizeCanvas() {
  cvs.width = window.innerWidth;
  cvs.height = window.innerHeight * 0.7;
  window.setPlayerView(window.innerWidth, Math.floor(window.innerHeight * 0.7));
}
resizeCanvas();
window.onresize = resizeCanvas;

let count;
let renderFunction = drawPlayer;

async function loop(n) {
  n = n ? n : 0;
  playerCoord = await window.requestPlayerCoord(userInput);
  lineMap = await window.requestLine();
  drawFloor(lineMap);
  renderFunction(playerCoord, n);
  updateUI(playerCoord);
  setTimeout(
    (n) => {
      loop(n + 1);
    },
    10,
    n
  );
}


const physicDiv = document.getElementById("physic");
const mapDiv = document.getElementById("map");

document.addEventListener("DOMContentLoaded", () => {
  resizeCanvas();
  populateMenus()
});

setTimeout(loop, 500, 0);
