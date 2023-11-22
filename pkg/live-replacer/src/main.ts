import { Socket } from "./socket";
import { replaceElements } from "./incoming";
import { render } from "./rendering";
import { FilterStore } from "./filter";

const WS_LOCATION = "ws://localhost:40000/ws";
(window as any).Filters = FilterStore;

(async function () {
  const socket = new Socket(WS_LOCATION);
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
