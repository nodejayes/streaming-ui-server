import { ClientIdStore } from "./client-id";
import { Socket } from "./socket";
import { replaceElements } from "./incoming";
import { render } from "./rendering";

const CLIENT_ID_KEY = "clientId";
const IDENTITY_LOCATION = "/identity";
const WS_LOCATION = "ws://localhost:40000/ws";

(async function () {
  const clientStorage = new ClientIdStore(CLIENT_ID_KEY, IDENTITY_LOCATION);
  await clientStorage.ensureClientId();
  const socket = new Socket(WS_LOCATION, clientStorage);
  let api = await socket.openWebSocket();

  api.onClose = () => {
    socket.openWebSocket().then((a) => (api = a));
  };
  api.onMessage = async (msg) => {
    await replaceElements(JSON.parse(msg.data));
    if (api) {
      await render(api);
    }
  };
  await render(api);
})();
