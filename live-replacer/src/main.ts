import { ClientIdStore } from "./client-id";
import { Socket } from "./socket";

const CLIENT_ID_KEY = "clientId";
const IDENTITY_LOCATION = "/identity";
const WS_LOCATION = "ws://localhost:40000/ws";
let API: WsApi | null = null;

const clickEvent: EventListenerOrEventListenerObject = (e: Event) => {
  if (!e.target) {
    console.warn("missing click event target", e);
    return;
  }
  const action = (e.target as any)?.getAttribute("lronclick");
  const payload = (e.target as any)?.getAttribute("lrclickpayload");
  API?.send({ type: action, payload: JSON.parse(payload) });
};

async function render(api: WsApi) {
  document.querySelectorAll("[lrOnClick]").forEach((el) => {
    el.removeEventListener("click", clickEvent);
    el.addEventListener("click", clickEvent);
  });
}

async function replaceElements(action: Action<string>) {
  if (action.type.startsWith("replaceHtml::")) {
    const selector = action.type.split("::")[1];
    const elements = document.querySelectorAll(selector);
    if (!elements) {
      return;
    }
    elements.forEach((element) => (element.innerHTML = action.payload));
  }
}

(async function () {
  const clientStorage = new ClientIdStore(CLIENT_ID_KEY, IDENTITY_LOCATION);
  await clientStorage.ensureClientId();
  const socket = new Socket(WS_LOCATION, clientStorage);
  API = await socket.openWebSocket();

  API.onClose = () => {
    socket.openWebSocket().then((a) => (API = a));
  };
  API.onMessage = async (msg) => {
    await replaceElements(JSON.parse(msg.data));
    if (API) {
      await render(API);
    }
  };
  await render(API);
  setTimeout(() => {
    if (API) {
      API.send({ type: "ping", payload: "hallo" });
    }
  }, 1000);
})();
