const CLIENT_ID_KEY = "clientId";
const IDENTITY_LOCATION = "/identity";
const WS_LOCATION = "ws://localhost:40000/ws";

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
      onMessage: ws.onmessage,
      onClose: ws.onclose,
    };
    ws.onopen = () => {
      console.info("open");
      resolve(api);
    };
    ws.onmessage = console.info;
    ws.onclose = console.info;
  });
}

async function render() {
  // select all elements with listener
}

async function replaceElements() {}

(async function () {
  await ensureClientId();
  let api = await openWebSocket();
  api.onClose = () => {
    console.info("close socket");
    openWebSocket().then((a) => (api = a));
  };
  api.onMessage = async (msg) => {
    console.info(msg);
    await replaceElements();
    await render();
  };
  await render();
  setTimeout(() => {
    api.send({ type: "ping", payload: "hallo" });
  }, 1000);
})();
