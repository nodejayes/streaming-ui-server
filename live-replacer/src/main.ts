const CLIENT_ID_KEY = "clientId";
const IDENTITY_LOCATION = "/identity";
const WS_LOCATION = "ws://localhost:40000/ws";

interface Action<T> {
  type: string;
  payload: T;
}

function getClientId(): string | null {
  return localStorage.getItem(CLIENT_ID_KEY);
}

function setClientId(id: string) {
  localStorage.setItem(CLIENT_ID_KEY, id);
}

function ensureClientId(): Promise<void> {
  return new Promise((resolve, reject) => {
    const id = getClientId();
    if (!id) {
      fetch(IDENTITY_LOCATION)
        .then((response) => {
          return response.text();
        })
        .then((id) => {
          setClientId(id);
          resolve();
        })
        .catch((err) => reject(err));
    } else {
      resolve();
    }
  });
}

async function openWebSocket(): Promise<{
  send: (data: any) => void;
  onMessage: ((msg: MessageEvent<any>) => any) | null;
  onClose: ((e: CloseEvent) => any) | null;
}> {
  return new Promise((resolve) => {
    const ws = new WebSocket(`${WS_LOCATION}?clientId=${getClientId()}`);
    const api = {
      send: (data: any) => ws.send(JSON.stringify(data)),
      onMessage: (_ev: MessageEvent<any>) => {},
      onClose: (_ev: CloseEvent) => {},
    };
    ws.onopen = () => {
      console.info("open");
      resolve(api);
    };
    ws.onmessage = (ev) => api.onMessage(ev);
    ws.onclose = (ev) => api.onClose(ev);
  });
}

async function render() {
  // select all elements with listener
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
  await ensureClientId();
  let api = await openWebSocket();
  api.onClose = () => {
    console.info("close socket");
    openWebSocket().then((a) => (api = a));
  };
  api.onMessage = async (msg) => {
    console.info("onMessage", msg);
    await replaceElements(JSON.parse(msg.data));
    await render();
  };
  await render();
  setTimeout(() => {
    api.send({ type: "ping", payload: "hallo" });
  }, 1000);
  setInterval(() => {
    api.send({ type: "count increase", payload: 1 });
  }, 1000);
})();
