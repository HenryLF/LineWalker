function resizeCanvas() {
  cvs.width = window.innerWidth;
  cvs.height = window.innerHeight * 0.7;
  window.setPlayerView(window.innerWidth, Math.floor(window.innerHeight * 0.7));
}
resizeCanvas();
window.onresize = resizeCanvas;


// let objCoord;
// let lineMap;
let renderFunction = drawPlayer;
async function loop(n) {
  n = n ? n : 0;
  objCoord = await window.requestObjectCoord(userInput);
  lineMap = await window.requestLine();
  drawFloor(lineMap);
  objCoord.forEach((obj, k) => {
    if (k == 0) {
      renderFunction(obj, n);
      updateUI(obj);
    } else {
      Ball.render(obj.ScreenCoord.X, obj.ScreenCoord.Y, obj.R);
    }
  });
  setTimeout(
    (n) => {
      loop(n + 1);
    },
    10,
    n
  );
}


document.addEventListener("DOMContentLoaded", () => {
  resizeCanvas();
  populateMenus();
  document.getElementById("soundtrack").play()
});

setTimeout(loop, 500, 0);
