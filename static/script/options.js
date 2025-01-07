const settingsTemp = document.getElementById("settings-temp");
const addObjectTemp = document.getElementById("add-object-temp");

const physicDiv = document.getElementById("physic");
const mapDiv = document.getElementById("map");
const globalDiv = document.getElementById("global");

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
  ["Player AccelerationY (Air Down)", "VerticalAccDown"],
  ["Max Speed", "CapSpeed"],
  ["Slow Motion", "TimeSlow"],
  ["Max time step", "MaxTimeDelay"],
  ["Colision Energy Transfert", "ElasticColision"],
];

let globalSettings = [
  ["GlobalScaleX", "ScaleX"],
  ["GlobalScaleY", "ScaleY"],
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
  ["X :", "objX"],
  ["Y :", "objY"],
  ["Masse", "objM"],
  ["Size", "objR"],
];
let addObjectCurrentSettings = new Map([
  ["objX", 5],
  ["objY", 5],
  ["objM", 5],
  ["objR", 5],
]);

function populateSetting(
  name,
  callBackString,
  callbackGet,
  callbackSet,
  asInt
) {
  let t = settingsTemp.content.cloneNode(true);
  t.getElementById("name").innerText = name;
  let input = t.getElementById("input");
  t.getElementById("set").onclick = async () => {
    let val = asInt ? parseInt(input.value) : parseFloat(input.value);
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

  globalSettings.map((e) => {
    globalDiv.appendChild(
      populateSetting(...e, window.getGlobals, window.setGlobals, true)
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
    let input = t.getElementById("input");
    input.value = addObjectCurrentSettings.get(e[1]);
    input.onchange = () => {
      addObjectCurrentSettings.set(e[1], parseFloat(input.value));
    };
    addObjectMenu.appendChild(t);
  });

  addObjectButton.onclick = () => {
    let x = addObjectCurrentSettings.get("objX");
    let y = addObjectCurrentSettings.get("objY");
    let m = addObjectCurrentSettings.get("objM");
    let r = addObjectCurrentSettings.get("objR");
    if (x && y && r && m) {
      window.addObject(x, y, m, r);
    } else {
      window.addObject(0, 0, 5, 5);
    }
  };
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
