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
  ["Floor Reaction", "FloorReactionCoeff"],
  ["Hard Ground", "GroundHardness"],
  ["Slope Angle DX", "DX"],
  ["Player AccelerationX (Ground)", "LateralAcc"],
  ["Player AccelerationX (Air)", "LateralAirAcc"],
  ["Player AccelerationY (Up,Ground)", "VerticalAcc"],
  ["Player AccelerationY (Down)", "VerticalAccDown"],
  ["Max Speed", "CapSpeed"],
  ["Slow Motion", "TimeSlow"],
  ["Max time step", "MaxTimeDelay"],
  ["Colision Energy Transfert", "ElasticColision"],
];

let globalSettings = [
  ["GlobalScaleX", "ScaleX"],
  ["GlobalScaleY", "ScaleY"],
];
let globalPlayerSettings = [
  ["Player Size", "R"],
  ["Player Mass", "M"],
  ["Player X", "X"],
  ["Player Y", "Y"],
];

let mapSettings = [
  ["Alpha", "A"],
  ["Beta", "B"],
  ["Scale X", "ScaleX"],
  ["ScaleY", "ScaleY"],
  ["Altitude", "Y0"],
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
let scale;
function populateSetting(
  name,
  callBackString,
  callbackGet,
  callbackSet,
  className
) {
  let t = settingsTemp.content.cloneNode(true);
  t.getElementById("name").innerText = name;
  let input = t.getElementById("input");
  if (className) {
    input.className += className;
    console.log(input.className)
  }
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

  globalSettings.map((e) => {
    globalDiv.appendChild(
      populateSetting(...e, window.getGlobals, window.setGlobals,e[1])
    );
  });

  globalPlayerSettings.map((e) => {
    globalDiv.appendChild(
      populateSetting(...e, window.getPlayer, window.setPlayer)
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

  addObjectButton.onclick = async () => {
    let x = addObjectCurrentSettings.get("objX");
    let y = addObjectCurrentSettings.get("objY");
    let m = addObjectCurrentSettings.get("objM");
    let r = addObjectCurrentSettings.get("objR");
    if (x && y && r && m) {
      window.addObject(x, y, m, r*playerSize/10);
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

    