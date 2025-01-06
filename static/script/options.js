const settingsTemp = document.getElementById("settings-temp");
const addObjectTemp = document.getElementById("add-object-temp");

const addObjectButton = document.getElementById("add-object-button");
const addObjectMenu = document.getElementById("add-object-menu");



let physicSettings = [
  ["Gravity", "G"],
  ["Air Friction", "AirFrictionCoeff"],
  ["Floor Friction", "FloorFrictionCoeff"],
  ["Slope Angle DX", "DX"],
  ["Player AccelerationX (Ground)", "LateralAcc"],
  ["Player AccelerationX (Air)", "LateralAirAcc"],
  ["Player AccelerationY (Ground)", "VerticalAcc"],
  ["Max Speed", "CapSpeed"],
  ["Slow Motion", "TimeSlow"],
];

let mapSettings = [
  ["Alpha", "A"],
  ["Beta", "B"],
  ["Scale X", "ScaleX"],
  ["ScaleY", "ScaleY"],
  ["Altitude", "Y0"],
];

let playerSettings = [
  ["X :", "X"],
  ["Y :", "Y"],
  ["Masse", "M"],
  ["Size", "R"],
];
let addObjectSettings = [
  ["X :", "X"],
  ["Y :", "Y"],
  ["Masse", "M"],
  ["Size", "R"],
];

function populateSetting(name, callBackString, callbackGet, callbackSet) {
  let t = settingsTemp.content.cloneNode(true);
  t.getElementById("name").innerText = name;
  let input = t.getElementById("input");
  t.getElementById("set").onclick = async () => {
    let val = parseFloat(input.value);
    let k;
    if (val) {
      k = await callbackSet(callBackString, val);
    } else {
      k = await callbackGet(callBackString);
    }
    input.value = k;
  };
  return t;
}


function populateMenus() {
  mapSettings.map((e) => {
    mapDiv.appendChild(populateSetting(...e, window.getMap, window.setMap));
  });
  physicSettings.map((e) => {
    physicDiv.appendChild(
      populateSetting(...e, window.getPhysic, window.setPhysic)
    );
  });
  for (let i of document.querySelectorAll("#set")) {
    if (i.onclick) {
      i.onclick();
    }
  }
  addObjectSettings.map((e) => {
    let t = addObjectTemp.content.cloneNode(true);
    t.getElementById("name").innerText = e[0];
    t.getElementById("input").value = 5;
    t.id = e[1];
    addObjectMenu.appendChild(t);
  });
}

const playerSelectButton = document.getElementById("playerselect-button");
const playerSelectMenu = document.getElementById("playerselect-menu");
playerSelectMenu.onclick = () => {
  playerSelectMenu.style.display = "none";
};

playerSelectButton.onclick = () => {
  playerSelectMenu.style.display = "flex";
};

const eggPlantButton = document.getElementById("eggplant");
eggPlantButton.onclick = () => {
  renderFunction = drawPlayer;
};
const bigCatButton = document.getElementById("bigcat");
bigCatButton.onclick = () => {
  renderFunction = drawBigPlayer;
};


