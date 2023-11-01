import { ClientIdStore } from "./client-id";
import { Socket } from "./socket";

const CLIENT_ID_KEY = "clientId";
const IDENTITY_LOCATION = "/identity";
const WS_LOCATION = "ws://localhost:40000/ws";
let API: WsApi | null = null;

const sendMessage = (target: HTMLElement | null, eventName: string) => {
  if (!target) {
    console.warn(`no element on target event ${eventName}`);
    return;
  }
  const actionName = `lr${eventName}action`;
  const payloadName = `lr${eventName}payload`;
  const inputsName = `lr${eventName}inputs`;
  const action = target.getAttribute(actionName);
  if (!action) {
    console.warn(`missing ${actionName} on element`, target);
    return;
  }
  const elementId = target.getAttribute("id") ?? null;
  const payload = target.getAttribute(payloadName) ?? null;
  const inputSelectors: string | null = target.getAttribute(inputsName) ?? null;
  if (!payload && !inputSelectors) {
    API?.send({
      elementId,
      type: action,
      payload: null,
      inputs: {},
    });
    return;
  }
  if (!inputSelectors) {
    API?.send({
      elementId,
      type: action,
      payload: JSON.parse(payload ?? "null"),
      inputs: {},
    });
    return;
  }

  const selectors = inputSelectors.split("<=>");
  const inputData: { [key: string]: { [key: string]: string } } = {};
  for (const selector of selectors) {
    document.querySelectorAll(selector)?.forEach((el) => {
      el.querySelectorAll("input")?.forEach((input: HTMLInputElement) => {
        const inputName = input.getAttribute("name");
        if (!inputName) {
          return;
        }
        if (!inputData[selector]) {
          inputData[selector] = {};
        }
        inputData[selector][inputName] = input.value;
      });
    });
  }
  API?.send({
    elementId,
    type: action,
    payload: JSON.parse(payload ?? "null"),
    inputs: {},
  });
};

const clickEvent: EventListenerOrEventListenerObject = (e: Event) => {
  if (!e.target) {
    console.warn("missing click event target", e);
    return;
  }
  sendMessage(e.target as HTMLElement, "click");
};

async function render(api: WsApi) {
  document.querySelectorAll("[lrClickAction]").forEach((el) => {
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
    elements.forEach((element) => (element.innerHTML = action.payload ?? ""));
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
})();
